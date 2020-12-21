package athena

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/athena"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetQueryResults executes `GetQueryResults` operation.
func (svc *Athena) GetQueryResults(ctx context.Context, r GetQueryResultsRequest) (*GetQueryResultsResult, error) {
	out, err := svc.RawGetQueryResults(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetQueryResults",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetQueryResultsResult(out), nil
}

// GetQueryResultsRequest has parameters for `GetQueryResults` operation.
type GetQueryResultsRequest struct {
	QueryExecutionID string

	// optional
	MaxResults int64
	NextToken  string
}

func (r GetQueryResultsRequest) ToInput() *SDK.GetQueryResultsInput {
	in := &SDK.GetQueryResultsInput{}

	if r.QueryExecutionID != "" {
		in.QueryExecutionId = pointers.String(r.QueryExecutionID)
	}
	if r.MaxResults > 0 {
		in.MaxResults = pointers.Long64(r.MaxResults)
	}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	return in
}

type GetQueryResultsResult struct {
	NextToken   string
	ResultSet   ResultSet
	UpdateCount int64
}

func NewGetQueryResultsResult(o *SDK.GetQueryResultsResponse) *GetQueryResultsResult {
	r := &GetQueryResultsResult{}
	if o == nil {
		return r
	}
	if o.GetQueryResultsOutput == nil {
		return r
	}
	oo := o.GetQueryResultsOutput

	if oo.NextToken != nil {
		r.NextToken = *oo.NextToken
	}
	if oo.UpdateCount != nil {
		r.UpdateCount = *oo.UpdateCount
	}
	r.ResultSet = NewResultSet(oo.ResultSet)
	return r
}
