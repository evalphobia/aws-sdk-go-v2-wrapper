package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetParameterHistory executes `GetParameterHistory` operation.
func (svc *SSM) GetParameterHistory(ctx context.Context, r GetParameterHistoryRequest) (*GetParameterHistoryResult, error) {
	out, err := svc.RawGetParameterHistory(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetParameterHistory",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetParameterHistoryResult(out), err
}

// GetParameterHistoryRequest has parameters for `GetParameterHistory` operation.
type GetParameterHistoryRequest struct {
	Name string

	// optional
	MaxResults     int64
	NextToken      string
	WithDecryption bool
}

func (r GetParameterHistoryRequest) ToInput() *SDK.GetParameterHistoryInput {
	in := &SDK.GetParameterHistoryInput{}
	if r.Name != "" {
		in.Name = pointers.String(r.Name)
	}
	if r.MaxResults != 0 {
		in.MaxResults = pointers.Long64(r.MaxResults)
	}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	if r.WithDecryption {
		in.WithDecryption = pointers.Bool(r.WithDecryption)
	}
	return in
}

type GetParameterHistoryResult struct {
	NextToken  string
	Parameters []ParameterHistory
}

func NewGetParameterHistoryResult(o *SDK.GetParameterHistoryResponse) *GetParameterHistoryResult {
	result := &GetParameterHistoryResult{}
	if o == nil {
		return result
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}
	if len(o.Parameters) != 0 {
		list := make([]ParameterHistory, len(o.Parameters))
		for i := range o.Parameters {
			list[i] = newParameterHistory(&o.Parameters[i])
		}
		result.Parameters = list
	}
	return result
}
