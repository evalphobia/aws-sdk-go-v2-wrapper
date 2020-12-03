package cloudwatchlogs

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

type QueryStatus string

const (
	QueryStatusScheduled QueryStatus = QueryStatus(SDK.QueryStatusScheduled)
	QueryStatusRunning   QueryStatus = QueryStatus(SDK.QueryStatusRunning)
	QueryStatusComplete  QueryStatus = QueryStatus(SDK.QueryStatusComplete)
	QueryStatusFailed    QueryStatus = QueryStatus(SDK.QueryStatusFailed)
	QueryStatusCancelled QueryStatus = QueryStatus(SDK.QueryStatusCancelled)
)

func (v QueryStatus) IsScheduled() bool {
	return v == QueryStatusScheduled
}

func (v QueryStatus) IsRunning() bool {
	return v == QueryStatusRunning
}

func (v QueryStatus) IsComplete() bool {
	return v == QueryStatusComplete
}

func (v QueryStatus) IsFailed() bool {
	return v == QueryStatusFailed
}

func (v QueryStatus) IsCancelled() bool {
	return v == QueryStatusCancelled
}
