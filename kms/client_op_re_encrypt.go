package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ReEncrypt executes `ReEncrypt` operation.
func (svc *KMS) ReEncrypt(ctx context.Context, r ReEncryptRequest) (*ReEncryptResult, error) {
	out, err := svc.RawReEncrypt(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ReEncrypt",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewReEncryptResult(out), nil
}

// ReEncryptRequest has parameters for `ReEncrypt` operation.
type ReEncryptRequest struct {
	CiphertextBlob   []byte
	DestinationKeyID string

	// optional
	DestinationEncryptionAlgorithm EncryptionAlgorithmSpec
	DestinationEncryptionContext   map[string]string
	GrantTokens                    []string
	SourceEncryptionAlgorithm      EncryptionAlgorithmSpec
	SourceEncryptionContext        map[string]string
	SourceKeyID                    string
}

func (r ReEncryptRequest) ToInput() *SDK.ReEncryptInput {
	in := &SDK.ReEncryptInput{}

	in.CiphertextBlob = r.CiphertextBlob
	if r.DestinationKeyID != "" {
		in.DestinationKeyId = pointers.String(r.DestinationKeyID)
	}

	in.DestinationEncryptionAlgorithm = SDK.EncryptionAlgorithmSpec(r.DestinationEncryptionAlgorithm)
	in.DestinationEncryptionContext = r.DestinationEncryptionContext
	in.GrantTokens = r.GrantTokens
	in.SourceEncryptionAlgorithm = SDK.EncryptionAlgorithmSpec(r.SourceEncryptionAlgorithm)
	in.SourceEncryptionContext = r.SourceEncryptionContext
	if r.SourceKeyID != "" {
		in.SourceKeyId = pointers.String(r.SourceKeyID)
	}
	return in
}

type ReEncryptResult struct {
	CiphertextBlob                 []byte
	DestinationEncryptionAlgorithm EncryptionAlgorithmSpec
	KeyID                          string
	SourceEncryptionAlgorithm      EncryptionAlgorithmSpec
	SourceKeyID                    string
}

func NewReEncryptResult(output *SDK.ReEncryptResponse) *ReEncryptResult {
	r := &ReEncryptResult{}
	if output == nil {
		return r
	}

	r.CiphertextBlob = output.CiphertextBlob
	r.DestinationEncryptionAlgorithm = EncryptionAlgorithmSpec(output.DestinationEncryptionAlgorithm)

	if output.KeyId != nil {
		r.KeyID = *output.KeyId
	}

	r.SourceEncryptionAlgorithm = EncryptionAlgorithmSpec(output.SourceEncryptionAlgorithm)

	if output.SourceKeyId != nil {
		r.SourceKeyID = *output.SourceKeyId
	}
	return r
}
