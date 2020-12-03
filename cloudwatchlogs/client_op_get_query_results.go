package cloudwatchlogs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetQueryResults executes `GetQueryResults` operation.
func (svc *CloudwatchLogs) GetQueryResults(ctx context.Context, r GetQueryResultsRequest) (*GetQueryResultsResult, error) {
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
	QueryID string
}

func (r GetQueryResultsRequest) ToInput() *SDK.GetQueryResultsInput {
	in := &SDK.GetQueryResultsInput{}

	if r.QueryID != "" {
		in.QueryId = pointers.String(r.QueryID)
	}
	return in
}

type GetQueryResultsResult struct {
	Results                  QueryResults
	StatisticsBytesScanned   float64
	StatisticsRecordsMatched float64
	StatisticsRecordsScanned float64
	Status                   QueryStatus
}

func NewGetQueryResultsResult(o *SDK.GetQueryResultsResponse) *GetQueryResultsResult {
	r := &GetQueryResultsResult{}
	if o == nil {
		return r
	}
	r.Results = NewQueryResults(o.Results)
	r.Status = QueryStatus(o.Status)
	if o.Statistics != nil {
		v := o.Statistics
		if v.BytesScanned != nil {
			r.StatisticsBytesScanned = *v.BytesScanned
		}
		if v.RecordsMatched != nil {
			r.StatisticsRecordsMatched = *v.RecordsMatched
		}
		if v.RecordsScanned != nil {
			r.StatisticsRecordsScanned = *v.RecordsScanned
		}
	}
	return r
}
