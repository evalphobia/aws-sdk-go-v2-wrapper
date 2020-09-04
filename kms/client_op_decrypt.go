package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// Decrypt executes `Decrypt` operation.
func (svc *KMS) Decrypt(ctx context.Context, r DecryptRequest) (*DecryptResult, error) {
	out, err := svc.RawDecrypt(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "Decrypt",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDecryptResult(out), nil
}

// DecryptRequest has parameters for `Decrypt` operation.
type DecryptRequest struct {
	CiphertextBlob []byte

	// optional
	EncryptionAlgorithm EncryptionAlgorithmSpec
	EncryptionContext   map[string]string
	GrantTokens         []string
	KeyID               string
}

func (r DecryptRequest) ToInput() *SDK.DecryptInput {
	in := &SDK.DecryptInput{}

	in.CiphertextBlob = r.CiphertextBlob
	in.EncryptionAlgorithm = SDK.EncryptionAlgorithmSpec(r.EncryptionAlgorithm)
	in.EncryptionContext = r.EncryptionContext
	in.GrantTokens = r.GrantTokens

	if r.KeyID != "" {
		in.KeyId = pointers.String(r.KeyID)
	}
	return in
}

type DecryptResult struct {
	EncryptionAlgorithm EncryptionAlgorithmSpec
	KeyID               string
	Plaintext           []byte
}

func NewDecryptResult(output *SDK.DecryptResponse) *DecryptResult {
	r := &DecryptResult{}
	if output == nil {
		return r
	}

	r.EncryptionAlgorithm = EncryptionAlgorithmSpec(output.EncryptionAlgorithm)
	if output.KeyId != nil {
		r.KeyID = *output.KeyId
	}
	r.Plaintext = output.Plaintext
	return r
}
