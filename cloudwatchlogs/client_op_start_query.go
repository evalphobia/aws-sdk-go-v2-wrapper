package cloudwatchlogs

import (
	"context"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// StartQuery executes `StartQuery` operation.
func (svc *CloudwatchLogs) StartQuery(ctx context.Context, r StartQueryRequest) (*StartQueryResult, error) {
	out, err := svc.RawStartQuery(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "StartQuery",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewStartQueryResult(out), nil
}

// StartQueryRequest has parameters for `StartQuery` operation.
type StartQueryRequest struct {
	QueryString  string
	StartTime    time.Time
	StartTimeInt int64
	EndTime      time.Time
	EndTimeInt   int64

	// optional
	Limit         int64
	LogGroupName  string
	LogGroupNames []string
}

func (r StartQueryRequest) ToInput() *SDK.StartQueryInput {
	in := &SDK.StartQueryInput{}

	if r.QueryString != "" {
		in.QueryString = pointers.String(r.QueryString)
	}
	switch {
	case !r.StartTime.IsZero():
		in.StartTime = pointers.Long64(r.StartTime.Unix())
	case r.StartTimeInt != 0:
		in.StartTime = pointers.Long64(r.StartTimeInt)
	}
	switch {
	case !r.EndTime.IsZero():
		in.EndTime = pointers.Long64(r.EndTime.Unix())
	case r.EndTimeInt != 0:
		in.EndTime = pointers.Long64(r.EndTimeInt)
	}

	if r.Limit != 0 {
		in.Limit = pointers.Long64(r.Limit)
	}
	if r.LogGroupName != "" {
		in.LogGroupName = pointers.String(r.LogGroupName)
	}
	in.LogGroupNames = r.LogGroupNames
	return in
}

type StartQueryResult struct {
	QueryID string
}

func NewStartQueryResult(o *SDK.StartQueryResponse) *StartQueryResult {
	result := &StartQueryResult{}
	if o == nil {
		return result
	}

	if o.QueryId != nil {
		result.QueryID = *o.QueryId
	}
	return result
}
