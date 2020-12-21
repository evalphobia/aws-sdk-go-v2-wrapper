package athena

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
	const initialWaitSec = 10
	expo := math.Pow(2, float64(i)) * initialWaitSec
	time.Sleep(time.Duration(rand.Float64()*expo) * time.Second) // #nosec G404
}

// XQueryResults executes a query and waits for fetching complete results.
func (svc *Athena) XQueryResults(ctx context.Context, r XQueryResultsRequest) (*GetQueryResultsResult, error) {
	out, err := svc.StartQueryExecution(ctx, StartQueryExecutionRequest{
		QueryString:           r.QueryString,
		ClientRequestToken:    r.ClientRequestToken,
		QueryExecutionContext: r.QueryExecutionContext,
		ResultConfiguration:   r.ResultConfiguration,
		WorkGroup:             r.WorkGroup,
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

	return svc.waitQueryResult(ctx, out.QueryExecutionID, maxRetry, waitFn)
}

type XQueryResultsRequest struct {
	QueryString string

	// optional
	ClientRequestToken    string
	QueryExecutionContext QueryExecutionContext
	ResultConfiguration   ResultConfiguration
	WorkGroup             string

	// extension
	MaxRetry int         // default=5
	WaitFunc func(i int) // waiting strategy, default=exponential backoff with full jitter from 10sec
}

func (svc *Athena) waitQueryResult(ctx context.Context, queryID string, maxRetry int, waitFn func(int)) (*GetQueryResultsResult, error) {
	for i := 0; i < maxRetry; i++ {
		waitFn(i) // wait

		res, err := svc.GetQueryExecution(ctx, GetQueryExecutionRequest{
			QueryExecutionID: queryID,
		})
		if err != nil {
			return nil, err
		}

		s := res.Status.State
		switch {
		case s.IsSucceeded():
			return svc.XGetQueryResultsAll(ctx, queryID)
		case s.IsRunning(),
			s.IsQueued():
			// do nothing; continue loop
		default:
			return nil, fmt.Errorf("Invalid query status: id:[%s], status:[%s]", queryID, string(s))
		}
	}
	return nil, fmt.Errorf("Timeout error: id:[%s]", queryID)
}

func (svc *Athena) XGetQueryResultsAll(ctx context.Context, queryID string) (*GetQueryResultsResult, error) {
	out, err := svc.GetQueryResults(ctx, GetQueryResultsRequest{
		QueryExecutionID: queryID,
	})
	if err != nil {
		return nil, err
	}

	rows := out.ResultSet.Rows
	for out.NextToken != "" {
		out, err = svc.GetQueryResults(ctx, GetQueryResultsRequest{
			QueryExecutionID: queryID,
			NextToken:        out.NextToken,
		})
		if err != nil {
			return nil, err
		}
		rows = append(rows, out.ResultSet.Rows...)
	}
	out.ResultSet.Rows = rows
	return out, nil
}
