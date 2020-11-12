package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetParameter executes `GetParameter` operation.
func (svc *SSM) GetParameter(ctx context.Context, r GetParameterRequest) (*GetParameterResult, error) {
	out, err := svc.RawGetParameter(ctx, r.ToInput())
	if err == nil {
		return NewGetParameterResult(out), err
	}

	if errors.GetAWSErrorCode(err) == "ParameterNotFound" {
		return &GetParameterResult{
			NotFound: true,
		}, nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "GetParameter",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// GetParameterRequest has parameters for `GetParameter` operation.
type GetParameterRequest struct {
	Name string

	// optional
	WithDecryption bool
}

func (r GetParameterRequest) ToInput() *SDK.GetParameterInput {
	in := &SDK.GetParameterInput{}
	if r.Name != "" {
		in.Name = pointers.String(r.Name)
	}
	if r.WithDecryption {
		in.WithDecryption = pointers.Bool(r.WithDecryption)
	}
	return in
}

type GetParameterResult struct {
	NotFound bool

	Parameter
}

func NewGetParameterResult(o *SDK.GetParameterResponse) *GetParameterResult {
	result := &GetParameterResult{}
	if o == nil {
		return result
	}
	result.Parameter = newParameter(o.Parameter)
	return result
}
