package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// Encrypt executes `Encrypt` operation.
func (svc *KMS) Encrypt(ctx context.Context, r EncryptRequest) (*EncryptResult, error) {
	out, err := svc.RawEncrypt(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "Encrypt",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewEncryptResult(out), nil
}

// EncryptRequest has parameters for `Encrypt` operation.
type EncryptRequest struct {
	KeyID     string
	Plaintext []byte

	// optional
	EncryptionAlgorithm EncryptionAlgorithmSpec
	EncryptionContext   map[string]string
	GrantTokens         []string
}

func (r EncryptRequest) ToInput() *SDK.EncryptInput {
	in := &SDK.EncryptInput{}

	if r.KeyID != "" {
		in.KeyId = pointers.String(r.KeyID)
	}
	in.Plaintext = r.Plaintext

	in.EncryptionAlgorithm = SDK.EncryptionAlgorithmSpec(r.EncryptionAlgorithm)
	in.EncryptionContext = r.EncryptionContext
	in.GrantTokens = r.GrantTokens
	return in
}

type EncryptResult struct {
	CiphertextBlob      []byte
	EncryptionAlgorithm EncryptionAlgorithmSpec
	KeyID               string
}

func NewEncryptResult(output *SDK.EncryptResponse) *EncryptResult {
	r := &EncryptResult{}
	if output == nil {
		return r
	}

	r.CiphertextBlob = output.CiphertextBlob
	r.EncryptionAlgorithm = EncryptionAlgorithmSpec(output.EncryptionAlgorithm)
	if output.KeyId != nil {
		r.KeyID = *output.KeyId
	}
	return r
}
