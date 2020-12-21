package athena

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/athena"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// StopQueryExecution executes `StopQueryExecution` operation.
func (svc *Athena) StopQueryExecution(ctx context.Context, r StopQueryExecutionRequest) (*StopQueryExecutionResult, error) {
	out, err := svc.RawStopQueryExecution(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "StopQueryExecution",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewStopQueryExecutionResult(out), nil
}

// StopQueryExecutionRequest has parameters for `StopQueryExecution` operation.
type StopQueryExecutionRequest struct {
	QueryExecutionID string
}

func (r StopQueryExecutionRequest) ToInput() *SDK.StopQueryExecutionInput {
	in := &SDK.StopQueryExecutionInput{}

	if r.QueryExecutionID != "" {
		in.QueryExecutionId = pointers.String(r.QueryExecutionID)
	}
	return in
}

type StopQueryExecutionResult struct {
}

func NewStopQueryExecutionResult(o *SDK.StopQueryExecutionResponse) *StopQueryExecutionResult {
	result := &StopQueryExecutionResult{}
	return result
}
