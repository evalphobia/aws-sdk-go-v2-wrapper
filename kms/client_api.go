package kms

import (
	"context"
	"encoding/base64"
)

func (svc *KMS) DeleteKey(ctx context.Context, key string, day ...int64) error {
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

func (svc *KMS) EncryptBytes(ctx context.Context, key string, plainData []byte) (encryptedData []byte, err error) {
	result, err := svc.Encrypt(ctx, EncryptRequest{
		KeyID:     key,
		Plaintext: plainData,
	})
	if err != nil {
		return nil, err
	}

	return result.CiphertextBlob, nil
}

func (svc *KMS) EncryptString(ctx context.Context, key string, plainText string) (base64Text string, err error) {
	encryptedData, err := svc.EncryptBytes(ctx, key, []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func (svc *KMS) DecryptBytes(ctx context.Context, encryptedData []byte) (plainData []byte, err error) {
	result, err := svc.Decrypt(ctx, DecryptRequest{
		CiphertextBlob: encryptedData,
	})
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}

func (svc *KMS) DecryptString(ctx context.Context, base64Text string) (plainText string, err error) {
	byt, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return "", err
	}

	plainData, err := svc.DecryptBytes(ctx, byt)
	if err != nil {
		return "", err
	}
	return string(plainData), nil
}

func (svc *KMS) ReEncryptBytes(ctx context.Context, destinationKey string, encryptedData []byte) (resultEncryptedData []byte, err error) {
	result, err := svc.ReEncrypt(ctx, ReEncryptRequest{
		DestinationKeyID: destinationKey,
		CiphertextBlob:   encryptedData,
	})
	if err != nil {
		return nil, err
	}

	return result.CiphertextBlob, nil
}

func (svc *KMS) ReEncryptString(ctx context.Context, destinationKey, base64Text string) (resultBase64Text string, err error) {
	byt, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return "", err
	}

	encryptedData, err := svc.ReEncryptBytes(ctx, destinationKey, byt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}
