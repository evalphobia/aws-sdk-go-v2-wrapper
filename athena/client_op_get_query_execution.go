package athena

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/athena"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetQueryExecution executes `GetQueryExecution` operation.
func (svc *Athena) GetQueryExecution(ctx context.Context, r GetQueryExecutionRequest) (*GetQueryExecutionResult, error) {
	out, err := svc.RawGetQueryExecution(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetQueryExecution",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetQueryExecutionResult(out), nil
}

// GetQueryExecutionRequest has parameters for `GetQueryExecution` operation.
type GetQueryExecutionRequest struct {
	QueryExecutionID string
}

func (r GetQueryExecutionRequest) ToInput() *SDK.GetQueryExecutionInput {
	in := &SDK.GetQueryExecutionInput{}

	if r.QueryExecutionID != "" {
		in.QueryExecutionId = pointers.String(r.QueryExecutionID)
	}
	return in
}

type GetQueryExecutionResult struct {
	Query                 string
	QueryExecutionContext QueryExecutionContext
	QueryExecutionID      string
	ResultConfiguration   ResultConfiguration
	StatementType         StatementType
	Statistics            QueryExecutionStatistics
	Status                QueryExecutionStatus
	WorkGroup             string
}

func NewGetQueryExecutionResult(o *SDK.GetQueryExecutionResponse) *GetQueryExecutionResult {
	r := &GetQueryExecutionResult{}
	if o == nil {
		return r
	}
	oo := o.GetQueryExecutionOutput
	if oo == nil {
		return r
	}
	v := oo.QueryExecution
	if v == nil {
		return r
	}

	if v.Query != nil {
		r.Query = *v.Query
	}
	if v.QueryExecutionId != nil {
		r.QueryExecutionID = *v.QueryExecutionId
	}
	if v.WorkGroup != nil {
		r.WorkGroup = *v.WorkGroup
	}

	r.QueryExecutionContext = NewQueryExecutionContext(v.QueryExecutionContext)
	r.ResultConfiguration = NewResultConfiguration(v.ResultConfiguration)
	r.Statistics = NewQueryExecutionStatistics(v.Statistics)
	r.Status = NewQueryExecutionStatus(v.Status)
	r.StatementType = StatementType(v.StatementType)
	return r
}
