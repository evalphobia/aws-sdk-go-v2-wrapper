package cloudwatchlogs

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	defaultMaxRetry = 5
)

var defaultWaitFunc = func(i int) {
	const initialWaitSec = 1
	expo := math.Pow(2, float64(i)) * initialWaitSec
	time.Sleep(time.Duration(rand.Float64()*expo) * time.Second) // #nosec G404
}

// XQueryResults executes a query and waits for fetching complete results.
func (svc *CloudwatchLogs) XQueryResults(ctx context.Context, r XQueryResultsRequest) (*GetQueryResultsResult, error) {
	out, err := svc.StartQuery(ctx, StartQueryRequest{
		QueryString:   r.QueryString,
		StartTime:     r.StartTime,
		StartTimeInt:  r.StartTimeInt,
		EndTime:       r.EndTime,
		EndTimeInt:    r.EndTimeInt,
		Limit:         r.Limit,
		LogGroupName:  r.LogGroupName,
		LogGroupNames: r.LogGroupNames,
	})
	if err != nil {
		return nil, err
	}

	maxRetry := r.MaxRetry
	if maxRetry == 0 {
		maxRetry = defaultMaxRetry
	}
	waitFn := r.WaitFunc
	if waitFn == nil {
		waitFn = defaultWaitFunc
	}

	return svc.waitQueryResult(ctx, out.QueryID, maxRetry, waitFn)
}

type XQueryResultsRequest struct {
	QueryString  string
	StartTime    time.Time
	StartTimeInt int64
	EndTime      time.Time
	EndTimeInt   int64

	// optional
	Limit         int64
	LogGroupName  string
	LogGroupNames []string

	// extension
	MaxRetry int         // default=5
	WaitFunc func(i int) // waiting strategy, default=exponential backoff with full jitter from 1sec
}

func (svc *CloudwatchLogs) waitQueryResult(ctx context.Context, queryID string, maxRetry int, waitFn func(int)) (*GetQueryResultsResult, error) {
	for i := 0; i < maxRetry; i++ {
		waitFn(i) // wait

		res, err := svc.GetQueryResults(ctx, GetQueryResultsRequest{
			QueryID: queryID,
		})

		switch {
		case err != nil:
			return nil, err
		case res.Status.IsComplete():
			return res, nil
		case res.Status.IsRunning(),
			res.Status.IsScheduled():
			// do nothing; continue loop
		default:
			return nil, fmt.Errorf("Invalid query status: id:[%s], status:[%s]", queryID, string(res.Status))
		}
	}
	return nil, fmt.Errorf("Timeout error: id:[%s]", queryID)
}
