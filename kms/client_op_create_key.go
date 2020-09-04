package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateKey executes `CreateKey` operation.
func (svc *KMS) CreateKey(ctx context.Context, r CreateKeyRequest) (*CreateKeyResult, error) {
	out, err := svc.RawCreateKey(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateKey",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateKeyResult(out), nil
}

// CreateKeyRequest has parameters for `CreateKey` operation.
type CreateKeyRequest struct {
	BypassPolicyLockoutSafetyCheck bool
	CustomKeyStoreID               string
	CustomerMasterKeySpec          CustomerMasterKeySpec
	Description                    string
	KeyUsage                       KeyUsageType
	Origin                         OriginType
	Policy                         string
	Tags                           []Tag
}

func (r CreateKeyRequest) ToInput() *SDK.CreateKeyInput {
	in := &SDK.CreateKeyInput{}
	if r.BypassPolicyLockoutSafetyCheck {
		in.BypassPolicyLockoutSafetyCheck = pointers.Bool(r.BypassPolicyLockoutSafetyCheck)
	}
	if r.CustomKeyStoreID != "" {
		in.CustomKeyStoreId = pointers.String(r.CustomKeyStoreID)
	}

	in.CustomerMasterKeySpec = SDK.CustomerMasterKeySpec(r.CustomerMasterKeySpec)

	if r.Description != "" {
		in.Description = pointers.String(r.Description)
	}

	in.KeyUsage = SDK.KeyUsageType(r.KeyUsage)
	in.Origin = SDK.OriginType(r.Origin)

	if r.Policy != "" {
		in.Policy = pointers.String(r.Policy)
	}

	if len(r.Tags) != 0 {
		list := make([]SDK.Tag, len(r.Tags))
		for i, v := range r.Tags {
			list[i] = v.ToSDK()
		}
		in.Tags = list
	}
	return in
}

type CreateKeyResult struct {
	KeyMetadata KeyMetadata
}

func NewCreateKeyResult(output *SDK.CreateKeyResponse) *CreateKeyResult {
	r := &CreateKeyResult{}
	if output == nil {
		return r
	}

	r.KeyMetadata = newKeyMetadata(output.KeyMetadata)
	return r
}
