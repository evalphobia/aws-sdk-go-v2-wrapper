package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DescribeKey executes `DescribeKey` operation.
func (svc *KMS) DescribeKey(ctx context.Context, r DescribeKeyRequest) (*DescribeKeyResult, error) {
	out, err := svc.RawDescribeKey(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DescribeKey",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDescribeKeyResult(out), nil
}

// DescribeKeyRequest has parameters for `DescribeKey` operation.
type DescribeKeyRequest struct {
	KeyID string

	// optional
	GrantTokens []string
}

func (r DescribeKeyRequest) ToInput() *SDK.DescribeKeyInput {
	in := &SDK.DescribeKeyInput{}

	if r.KeyID != "" {
		in.KeyId = pointers.String(r.KeyID)
	}
	in.GrantTokens = r.GrantTokens
	return in
}

type DescribeKeyResult struct {
	KeyMetadata KeyMetadata
}

func NewDescribeKeyResult(output *SDK.DescribeKeyResponse) *DescribeKeyResult {
	r := &DescribeKeyResult{}
	if output == nil {
		return r
	}

	r.KeyMetadata = newKeyMetadata(output.KeyMetadata)
	return r
}
