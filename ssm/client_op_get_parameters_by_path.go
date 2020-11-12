package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetParametersByPath executes `GetParametersByPath` operation.
func (svc *SSM) GetParametersByPath(ctx context.Context, r GetParametersByPathRequest) (*GetParametersByPathResult, error) {
	out, err := svc.RawGetParametersByPath(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetParametersByPath",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetParametersByPathResult(out), err
}

// GetParametersByPathRequest has parameters for `GetParametersByPath` operation.
type GetParametersByPathRequest struct {
	Path string

	// optional
	MaxResults       int64
	NextToken        string
	ParameterFilters []ParameterStringFilter
	Recursive        bool
	WithDecryption   bool
}

func (r GetParametersByPathRequest) ToInput() *SDK.GetParametersByPathInput {
	in := &SDK.GetParametersByPathInput{}
	if r.Path != "" {
		in.Path = pointers.String(r.Path)
	}
	if r.MaxResults != 0 {
		in.MaxResults = pointers.Long64(r.MaxResults)
	}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	if len(r.ParameterFilters) != 0 {
		list := make([]SDK.ParameterStringFilter, len(r.ParameterFilters))
		for i, v := range r.ParameterFilters {
			list[i] = v.ToSDK()
		}
		in.ParameterFilters = list
	}
	if r.Recursive {
		in.Recursive = pointers.Bool(r.Recursive)
	}
	if r.WithDecryption {
		in.WithDecryption = pointers.Bool(r.WithDecryption)
	}
	return in
}

type GetParametersByPathResult struct {
	NextToken  string
	Parameters []Parameter
}

func NewGetParametersByPathResult(o *SDK.GetParametersByPathResponse) *GetParametersByPathResult {
	result := &GetParametersByPathResult{}
	if o == nil {
		return result
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}

	if len(o.Parameters) != 0 {
		list := make([]Parameter, len(o.Parameters))
		for i := range o.Parameters {
			list[i] = newParameter(&o.Parameters[i])
		}
		result.Parameters = list
	}
	return result
}
