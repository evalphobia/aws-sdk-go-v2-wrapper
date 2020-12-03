package cloudwatchlogs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// StopQuery executes `StopQuery` operation.
func (svc *CloudwatchLogs) StopQuery(ctx context.Context, r StopQueryRequest) (*StopQueryResult, error) {
	out, err := svc.RawStopQuery(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "StopQuery",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewStopQueryResult(out), nil
}

// StopQueryRequest has parameters for `StopQuery` operation.
type StopQueryRequest struct {
	QueryID string
}

func (r StopQueryRequest) ToInput() *SDK.StopQueryInput {
	in := &SDK.StopQueryInput{}

	if r.QueryID != "" {
		in.QueryId = pointers.String(r.QueryID)
	}
	return in
}

type StopQueryResult struct {
	Success bool
}

func NewStopQueryResult(o *SDK.StopQueryResponse) *StopQueryResult {
	result := &StopQueryResult{}
	if o == nil {
		return result
	}

	if o.Success != nil {
		result.Success = *o.Success
	}
	return result
}
