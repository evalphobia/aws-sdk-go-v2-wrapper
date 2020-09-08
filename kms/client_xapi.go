package kms

import (
	"context"
	"encoding/base64"
)

// XDeleteKey deletes the key using 'ScheduleKeyDeletion'.
func (svc *KMS) XDeleteKey(ctx context.Context, key string, day ...int64) error {
	metaData, err := svc.DescribeKey(ctx, DescribeKeyRequest{
		KeyID: key,
	})
	if err != nil {
		return err
	}

	const defaultWindowDay = 30
	d := int64(defaultWindowDay)
	if len(day) != 0 {
		d = day[0]
	}

	_, err = svc.ScheduleKeyDeletion(ctx, ScheduleKeyDeletionRequest{
		KeyID:               metaData.KeyMetadata.KeyID,
		PendingWindowInDays: d,
	})
	return err
}

// XEncryptBytes encrypts the 'plainData' using the 'key' and returns 'encryptedData'.
func (svc *KMS) XEncryptBytes(ctx context.Context, key string, plainData []byte) (encryptedData []byte, err error) {
	result, err := svc.Encrypt(ctx, EncryptRequest{
		KeyID:     key,
		Plaintext: plainData,
	})
	if err != nil {
		return nil, err
	}

	return result.CiphertextBlob, nil
}

// XEncryptString encrypts the 'plainText' using the 'key' and returns encrypted 'base64Text'.
func (svc *KMS) XEncryptString(ctx context.Context, key string, plainText string) (base64Text string, err error) {
	encryptedData, err := svc.XEncryptBytes(ctx, key, []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// XDecryptBytes decrypts the 'encryptedData'.
func (svc *KMS) XDecryptBytes(ctx context.Context, encryptedData []byte) (plainData []byte, err error) {
	result, err := svc.Decrypt(ctx, DecryptRequest{
		CiphertextBlob: encryptedData,
	})
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}

// XDecryptString decrypts the 'base64Text'.
func (svc *KMS) XDecryptString(ctx context.Context, base64Text string) (plainText string, err error) {
	byt, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return "", err
	}

	plainData, err := svc.XDecryptBytes(ctx, byt)
	if err != nil {
		return "", err
	}
	return string(plainData), nil
}

// XReEncryptBytes re-encrypts the 'encryptedData' using 'destinationKey'.
func (svc *KMS) XReEncryptBytes(ctx context.Context, destinationKey string, encryptedData []byte) (resultEncryptedData []byte, err error) {
	result, err := svc.ReEncrypt(ctx, ReEncryptRequest{
		DestinationKeyID: destinationKey,
		CiphertextBlob:   encryptedData,
	})
	if err != nil {
		return nil, err
	}

	return result.CiphertextBlob, nil
}

// XReEncryptString re-encrypts the 'base64Text' using 'destinationKey'.
func (svc *KMS) XReEncryptString(ctx context.Context, destinationKey, base64Text string) (resultBase64Text string, err error) {
	byt, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return "", err
	}

	encryptedData, err := svc.XReEncryptBytes(ctx, destinationKey, byt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}
