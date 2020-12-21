package athena

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/athena"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// StartQueryExecution executes `StartQueryExecution` operation.
func (svc *Athena) StartQueryExecution(ctx context.Context, r StartQueryExecutionRequest) (*StartQueryExecutionResult, error) {
	out, err := svc.RawStartQueryExecution(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "StartQueryExecution",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewStartQueryExecutionResult(out), nil
}

// StartQueryExecutionRequest has parameters for `StartQueryExecution` operation.
type StartQueryExecutionRequest struct {
	QueryString string

	// optional
	ClientRequestToken    string
	QueryExecutionContext QueryExecutionContext
	ResultConfiguration   ResultConfiguration
	WorkGroup             string
}

func (r StartQueryExecutionRequest) ToInput() *SDK.StartQueryExecutionInput {
	in := &SDK.StartQueryExecutionInput{}

	if r.QueryString != "" {
		in.QueryString = pointers.String(r.QueryString)
	}
	if r.WorkGroup != "" {
		in.WorkGroup = pointers.String(r.WorkGroup)
	}
	in.QueryExecutionContext = r.QueryExecutionContext.ToSDK()
	in.ResultConfiguration = r.ResultConfiguration.ToSDK()
	return in
}

type StartQueryExecutionResult struct {
	QueryExecutionID string
}

func NewStartQueryExecutionResult(o *SDK.StartQueryExecutionResponse) *StartQueryExecutionResult {
	result := &StartQueryExecutionResult{}
	if o == nil {
		return result
	}

	if o.QueryExecutionId != nil {
		result.QueryExecutionID = *o.QueryExecutionId
	}
	return result
}
