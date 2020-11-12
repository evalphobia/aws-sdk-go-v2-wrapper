package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetParameters executes `GetParameters` operation.
func (svc *SSM) GetParameters(ctx context.Context, r GetParametersRequest) (*GetParametersResult, error) {
	out, err := svc.RawGetParameters(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetParameters",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetParametersResult(out), err
}

// GetParametersRequest has parameters for `GetParameters` operation.
type GetParametersRequest struct {
	Names []string

	// optional
	WithDecryption bool
}

func (r GetParametersRequest) ToInput() *SDK.GetParametersInput {
	in := &SDK.GetParametersInput{}
	in.Names = r.Names
	if r.WithDecryption {
		in.WithDecryption = pointers.Bool(r.WithDecryption)
	}
	return in
}

type GetParametersResult struct {
	InvalidParameters []string
	Parameters        []Parameter
}

func NewGetParametersResult(o *SDK.GetParametersResponse) *GetParametersResult {
	result := &GetParametersResult{}
	if o == nil {
		return result
	}

	result.InvalidParameters = o.InvalidParameters
	if len(o.Parameters) != 0 {
		list := make([]Parameter, len(o.Parameters))
		for i := range o.Parameters {
			list[i] = newParameter(&o.Parameters[i])
		}
		result.Parameters = list
	}
	return result
}
