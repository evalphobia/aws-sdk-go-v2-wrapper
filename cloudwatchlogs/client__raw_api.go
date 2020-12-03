package cloudwatchlogs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// RawAssociateKmsKey executes `AssociateKmsKey` raw operation.
func (svc *CloudwatchLogs) RawAssociateKmsKey(ctx context.Context, in *SDK.AssociateKmsKeyInput) (*SDK.AssociateKmsKeyResponse, error) {
	return svc.client.AssociateKmsKeyRequest(in).Send(ctx)
}

// RawCancelExportTask executes `CancelExportTask` raw operation.
func (svc *CloudwatchLogs) RawCancelExportTask(ctx context.Context, in *SDK.CancelExportTaskInput) (*SDK.CancelExportTaskResponse, error) {
	return svc.client.CancelExportTaskRequest(in).Send(ctx)
}

// RawCreateExportTask executes `CreateExportTask` raw operation.
func (svc *CloudwatchLogs) RawCreateExportTask(ctx context.Context, in *SDK.CreateExportTaskInput) (*SDK.CreateExportTaskResponse, error) {
	return svc.client.CreateExportTaskRequest(in).Send(ctx)
}

// RawCreateLogGroup executes `CreateLogGroup` raw operation.
func (svc *CloudwatchLogs) RawCreateLogGroup(ctx context.Context, in *SDK.CreateLogGroupInput) (*SDK.CreateLogGroupResponse, error) {
	return svc.client.CreateLogGroupRequest(in).Send(ctx)
}

// RawCreateLogStream executes `CreateLogStream` raw operation.
func (svc *CloudwatchLogs) RawCreateLogStream(ctx context.Context, in *SDK.CreateLogStreamInput) (*SDK.CreateLogStreamResponse, error) {
	return svc.client.CreateLogStreamRequest(in).Send(ctx)
}

// RawDeleteDestination executes `DeleteDestination` raw operation.
func (svc *CloudwatchLogs) RawDeleteDestination(ctx context.Context, in *SDK.DeleteDestinationInput) (*SDK.DeleteDestinationResponse, error) {
	return svc.client.DeleteDestinationRequest(in).Send(ctx)
}

// RawDeleteLogGroup executes `DeleteLogGroup` raw operation.
func (svc *CloudwatchLogs) RawDeleteLogGroup(ctx context.Context, in *SDK.DeleteLogGroupInput) (*SDK.DeleteLogGroupResponse, error) {
	return svc.client.DeleteLogGroupRequest(in).Send(ctx)
}

// RawDeleteLogStream executes `DeleteLogStream` raw operation.
func (svc *CloudwatchLogs) RawDeleteLogStream(ctx context.Context, in *SDK.DeleteLogStreamInput) (*SDK.DeleteLogStreamResponse, error) {
	return svc.client.DeleteLogStreamRequest(in).Send(ctx)
}

// RawDeleteMetricFilter executes `DeleteMetricFilter` raw operation.
func (svc *CloudwatchLogs) RawDeleteMetricFilter(ctx context.Context, in *SDK.DeleteMetricFilterInput) (*SDK.DeleteMetricFilterResponse, error) {
	return svc.client.DeleteMetricFilterRequest(in).Send(ctx)
}

// RawDeleteQueryDefinition executes `DeleteQueryDefinition` raw operation.
func (svc *CloudwatchLogs) RawDeleteQueryDefinition(ctx context.Context, in *SDK.DeleteQueryDefinitionInput) (*SDK.DeleteQueryDefinitionResponse, error) {
	return svc.client.DeleteQueryDefinitionRequest(in).Send(ctx)
}

// RawDeleteResourcePolicy executes `DeleteResourcePolicy` raw operation.
func (svc *CloudwatchLogs) RawDeleteResourcePolicy(ctx context.Context, in *SDK.DeleteResourcePolicyInput) (*SDK.DeleteResourcePolicyResponse, error) {
	return svc.client.DeleteResourcePolicyRequest(in).Send(ctx)
}

// RawDeleteRetentionPolicy executes `DeleteRetentionPolicy` raw operation.
func (svc *CloudwatchLogs) RawDeleteRetentionPolicy(ctx context.Context, in *SDK.DeleteRetentionPolicyInput) (*SDK.DeleteRetentionPolicyResponse, error) {
	return svc.client.DeleteRetentionPolicyRequest(in).Send(ctx)
}

// RawDeleteSubscriptionFilter executes `DeleteSubscriptionFilter` raw operation.
func (svc *CloudwatchLogs) RawDeleteSubscriptionFilter(ctx context.Context, in *SDK.DeleteSubscriptionFilterInput) (*SDK.DeleteSubscriptionFilterResponse, error) {
	return svc.client.DeleteSubscriptionFilterRequest(in).Send(ctx)
}

// RawDescribeDestinations executes `DescribeDestinations` raw operation.
func (svc *CloudwatchLogs) RawDescribeDestinations(ctx context.Context, in *SDK.DescribeDestinationsInput) (*SDK.DescribeDestinationsResponse, error) {
	return svc.client.DescribeDestinationsRequest(in).Send(ctx)
}

// RawDescribeExportTasks executes `DescribeExportTasks` raw operation.
func (svc *CloudwatchLogs) RawDescribeExportTasks(ctx context.Context, in *SDK.DescribeExportTasksInput) (*SDK.DescribeExportTasksResponse, error) {
	return svc.client.DescribeExportTasksRequest(in).Send(ctx)
}

// RawDescribeLogGroups executes `DescribeLogGroups` raw operation.
func (svc *CloudwatchLogs) RawDescribeLogGroups(ctx context.Context, in *SDK.DescribeLogGroupsInput) (*SDK.DescribeLogGroupsResponse, error) {
	return svc.client.DescribeLogGroupsRequest(in).Send(ctx)
}

// RawDescribeLogStreams executes `DescribeLogStreams` raw operation.
func (svc *CloudwatchLogs) RawDescribeLogStreams(ctx context.Context, in *SDK.DescribeLogStreamsInput) (*SDK.DescribeLogStreamsResponse, error) {
	return svc.client.DescribeLogStreamsRequest(in).Send(ctx)
}

// RawDescribeMetricFilters executes `DescribeMetricFilters` raw operation.
func (svc *CloudwatchLogs) RawDescribeMetricFilters(ctx context.Context, in *SDK.DescribeMetricFiltersInput) (*SDK.DescribeMetricFiltersResponse, error) {
	return svc.client.DescribeMetricFiltersRequest(in).Send(ctx)
}

// RawDescribeQueries executes `DescribeQueries` raw operation.
func (svc *CloudwatchLogs) RawDescribeQueries(ctx context.Context, in *SDK.DescribeQueriesInput) (*SDK.DescribeQueriesResponse, error) {
	return svc.client.DescribeQueriesRequest(in).Send(ctx)
}

// RawDescribeQueryDefinitions executes `DescribeQueryDefinitions` raw operation.
func (svc *CloudwatchLogs) RawDescribeQueryDefinitions(ctx context.Context, in *SDK.DescribeQueryDefinitionsInput) (*SDK.DescribeQueryDefinitionsResponse, error) {
	return svc.client.DescribeQueryDefinitionsRequest(in).Send(ctx)
}

// RawDescribeResourcePolicies executes `DescribeResourcePolicies` raw operation.
func (svc *CloudwatchLogs) RawDescribeResourcePolicies(ctx context.Context, in *SDK.DescribeResourcePoliciesInput) (*SDK.DescribeResourcePoliciesResponse, error) {
	return svc.client.DescribeResourcePoliciesRequest(in).Send(ctx)
}

// RawDescribeSubscriptionFilters executes `DescribeSubscriptionFilters` raw operation.
func (svc *CloudwatchLogs) RawDescribeSubscriptionFilters(ctx context.Context, in *SDK.DescribeSubscriptionFiltersInput) (*SDK.DescribeSubscriptionFiltersResponse, error) {
	return svc.client.DescribeSubscriptionFiltersRequest(in).Send(ctx)
}

// RawDisassociateKmsKey executes `DisassociateKmsKey` raw operation.
func (svc *CloudwatchLogs) RawDisassociateKmsKey(ctx context.Context, in *SDK.DisassociateKmsKeyInput) (*SDK.DisassociateKmsKeyResponse, error) {
	return svc.client.DisassociateKmsKeyRequest(in).Send(ctx)
}

// RawFilterLogEvents executes `FilterLogEvents` raw operation.
func (svc *CloudwatchLogs) RawFilterLogEvents(ctx context.Context, in *SDK.FilterLogEventsInput) (*SDK.FilterLogEventsResponse, error) {
	return svc.client.FilterLogEventsRequest(in).Send(ctx)
}

// RawGetLogEvents executes `GetLogEvents` raw operation.
func (svc *CloudwatchLogs) RawGetLogEvents(ctx context.Context, in *SDK.GetLogEventsInput) (*SDK.GetLogEventsResponse, error) {
	return svc.client.GetLogEventsRequest(in).Send(ctx)
}

// RawGetLogGroupFields executes `GetLogGroupFields` raw operation.
func (svc *CloudwatchLogs) RawGetLogGroupFields(ctx context.Context, in *SDK.GetLogGroupFieldsInput) (*SDK.GetLogGroupFieldsResponse, error) {
	return svc.client.GetLogGroupFieldsRequest(in).Send(ctx)
}

// RawGetLogRecord executes `GetLogRecord` raw operation.
func (svc *CloudwatchLogs) RawGetLogRecord(ctx context.Context, in *SDK.GetLogRecordInput) (*SDK.GetLogRecordResponse, error) {
	return svc.client.GetLogRecordRequest(in).Send(ctx)
}

// RawGetQueryResults executes `GetQueryResults` raw operation.
func (svc *CloudwatchLogs) RawGetQueryResults(ctx context.Context, in *SDK.GetQueryResultsInput) (*SDK.GetQueryResultsResponse, error) {
	return svc.client.GetQueryResultsRequest(in).Send(ctx)
}

// RawListTagsLogGroup executes `ListTagsLogGroup` raw operation.
func (svc *CloudwatchLogs) RawListTagsLogGroup(ctx context.Context, in *SDK.ListTagsLogGroupInput) (*SDK.ListTagsLogGroupResponse, error) {
	return svc.client.ListTagsLogGroupRequest(in).Send(ctx)
}

// RawPutDestination executes `PutDestination` raw operation.
func (svc *CloudwatchLogs) RawPutDestination(ctx context.Context, in *SDK.PutDestinationInput) (*SDK.PutDestinationResponse, error) {
	return svc.client.PutDestinationRequest(in).Send(ctx)
}

// RawPutDestinationPolicy executes `PutDestinationPolicy` raw operation.
func (svc *CloudwatchLogs) RawPutDestinationPolicy(ctx context.Context, in *SDK.PutDestinationPolicyInput) (*SDK.PutDestinationPolicyResponse, error) {
	return svc.client.PutDestinationPolicyRequest(in).Send(ctx)
}

// RawPutLogEvents executes `PutLogEvents` raw operation.
func (svc *CloudwatchLogs) RawPutLogEvents(ctx context.Context, in *SDK.PutLogEventsInput) (*SDK.PutLogEventsResponse, error) {
	return svc.client.PutLogEventsRequest(in).Send(ctx)
}

// RawPutMetricFilter executes `PutMetricFilter` raw operation.
func (svc *CloudwatchLogs) RawPutMetricFilter(ctx context.Context, in *SDK.PutMetricFilterInput) (*SDK.PutMetricFilterResponse, error) {
	return svc.client.PutMetricFilterRequest(in).Send(ctx)
}

// RawPutQueryDefinition executes `PutQueryDefinition` raw operation.
func (svc *CloudwatchLogs) RawPutQueryDefinition(ctx context.Context, in *SDK.PutQueryDefinitionInput) (*SDK.PutQueryDefinitionResponse, error) {
	return svc.client.PutQueryDefinitionRequest(in).Send(ctx)
}

// RawPutResourcePolicy executes `PutResourcePolicy` raw operation.
func (svc *CloudwatchLogs) RawPutResourcePolicy(ctx context.Context, in *SDK.PutResourcePolicyInput) (*SDK.PutResourcePolicyResponse, error) {
	return svc.client.PutResourcePolicyRequest(in).Send(ctx)
}

// RawPutRetentionPolicy executes `PutRetentionPolicy` raw operation.
func (svc *CloudwatchLogs) RawPutRetentionPolicy(ctx context.Context, in *SDK.PutRetentionPolicyInput) (*SDK.PutRetentionPolicyResponse, error) {
	return svc.client.PutRetentionPolicyRequest(in).Send(ctx)
}

// RawPutSubscriptionFilter executes `PutSubscriptionFilter` raw operation.
func (svc *CloudwatchLogs) RawPutSubscriptionFilter(ctx context.Context, in *SDK.PutSubscriptionFilterInput) (*SDK.PutSubscriptionFilterResponse, error) {
	return svc.client.PutSubscriptionFilterRequest(in).Send(ctx)
}

// RawStartQuery executes `StartQuery` raw operation.
func (svc *CloudwatchLogs) RawStartQuery(ctx context.Context, in *SDK.StartQueryInput) (*SDK.StartQueryResponse, error) {
	return svc.client.StartQueryRequest(in).Send(ctx)
}

// RawStopQuery executes `StopQuery` raw operation.
func (svc *CloudwatchLogs) RawStopQuery(ctx context.Context, in *SDK.StopQueryInput) (*SDK.StopQueryResponse, error) {
	return svc.client.StopQueryRequest(in).Send(ctx)
}

// RawTagLogGroup executes `TagLogGroup` raw operation.
func (svc *CloudwatchLogs) RawTagLogGroup(ctx context.Context, in *SDK.TagLogGroupInput) (*SDK.TagLogGroupResponse, error) {
	return svc.client.TagLogGroupRequest(in).Send(ctx)
}

// RawTestMetricFilter executes `TestMetricFilter` raw operation.
func (svc *CloudwatchLogs) RawTestMetricFilter(ctx context.Context, in *SDK.TestMetricFilterInput) (*SDK.TestMetricFilterResponse, error) {
	return svc.client.TestMetricFilterRequest(in).Send(ctx)
}

// RawUntagLogGroup executes `UntagLogGroup` raw operation.
func (svc *CloudwatchLogs) RawUntagLogGroup(ctx context.Context, in *SDK.UntagLogGroupInput) (*SDK.UntagLogGroupResponse, error) {
	return svc.client.UntagLogGroupRequest(in).Send(ctx)
}
