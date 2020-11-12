package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutParameter executes `PutParameter` operation.
func (svc *SSM) PutParameter(ctx context.Context, r PutParameterRequest) (*PutParameterResult, error) {
	out, err := svc.RawPutParameter(ctx, r.ToInput())
	if err == nil {
		return NewPutParameterResult(out), err
	}

	if errors.GetAWSErrorCode(err) == "ParameterAlreadyExists" {
		return &PutParameterResult{
			AlreadyExists: true,
		}, nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "PutParameter",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// PutParameterRequest has parameters for `PutParameter` operation.
type PutParameterRequest struct {
	Name  string
	Value string

	// optional
	AllowedPattern string
	DataType       string
	Description    string
	KeyID          string
	Overwrite      bool
	Policies       string
	Tags           []Tag
	Tier           ParameterTier
	Type           ParameterType
}

func (r PutParameterRequest) ToInput() *SDK.PutParameterInput {
	in := &SDK.PutParameterInput{}
	if r.Name != "" {
		in.Name = pointers.String(r.Name)
	}
	if r.Value != "" {
		in.Value = pointers.String(r.Value)
	}
	if r.AllowedPattern != "" {
		in.AllowedPattern = pointers.String(r.AllowedPattern)
	}
	if r.DataType != "" {
		in.DataType = pointers.String(r.DataType)
	}
	if r.Description != "" {
		in.Description = pointers.String(r.Description)
	}
	if r.KeyID != "" {
		in.KeyId = pointers.String(r.KeyID)
	}
	if r.Overwrite {
		in.Overwrite = pointers.Bool(r.Overwrite)
	}
	if r.Policies != "" {
		in.Policies = pointers.String(r.Policies)
	}

	if len(r.Tags) != 0 {
		list := make([]SDK.Tag, len(r.Tags))
		for i, v := range r.Tags {
			list[i] = v.ToSDK()
		}
		in.Tags = list
	}
	in.Tier = SDK.ParameterTier(r.Tier)
	in.Type = SDK.ParameterType(r.Type)
	return in
}

type PutParameterResult struct {
	AlreadyExists bool

	Tier    ParameterTier
	Version int64
}

func NewPutParameterResult(o *SDK.PutParameterResponse) *PutParameterResult {
	result := &PutParameterResult{}
	if o == nil {
		return result
	}

	result.Tier = ParameterTier(o.Tier)
	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}
