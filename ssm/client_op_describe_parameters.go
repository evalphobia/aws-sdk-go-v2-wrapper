package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DescribeParameters executes `DescribeParameters` operation.
func (svc *SSM) DescribeParameters(ctx context.Context, r DescribeParametersRequest) (*DescribeParametersResult, error) {
	out, err := svc.RawDescribeParameters(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DescribeParameters",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDescribeParametersResult(out), err
}

// DescribeParametersRequest has parameters for `DescribeParameters` operation.
type DescribeParametersRequest struct {
	MaxResults       int64
	NextToken        string
	ParameterFilters []ParameterStringFilter
}

func (r DescribeParametersRequest) ToInput() *SDK.DescribeParametersInput {
	in := &SDK.DescribeParametersInput{}
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
	return in
}

type DescribeParametersResult struct {
	NextToken  string
	Parameters []ParameterMetadata
}

func NewDescribeParametersResult(o *SDK.DescribeParametersResponse) *DescribeParametersResult {
	result := &DescribeParametersResult{}
	if o == nil {
		return result
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}
	if len(o.Parameters) != 0 {
		list := make([]ParameterMetadata, len(o.Parameters))
		for i := range o.Parameters {
			list[i] = newParameterMetadata(o.Parameters[i])
		}
		result.Parameters = list
	}
	return result
}
