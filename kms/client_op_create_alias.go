package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateAlias executes `CreateAlias` operation.
func (svc *KMS) CreateAlias(ctx context.Context, r CreateAliasRequest) error {
	_, err := svc.RawCreateAlias(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateAlias",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// CreateAliasRequest has parameters for `CreateAlias` operation.
type CreateAliasRequest struct {
	AliasName   string
	TargetKeyID string
}

func (r CreateAliasRequest) ToInput() *SDK.CreateAliasInput {
	in := &SDK.CreateAliasInput{}
	if r.AliasName != "" {
		in.AliasName = pointers.String(r.AliasName)
	}
	if r.TargetKeyID != "" {
		in.TargetKeyId = pointers.String(r.TargetKeyID)
	}
	return in
}
