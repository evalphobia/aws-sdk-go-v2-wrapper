package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// RawBatchGetItem executes `BatchGetItem` raw operation.
func (svc *DynamoDB) RawBatchGetItem(ctx context.Context, in *SDK.BatchGetItemInput) (*SDK.BatchGetItemResponse, error) {
	return svc.client.BatchGetItemRequest(in).Send(ctx)
}

// RawBatchWriteItem executes `BatchWriteItem` raw operation.
func (svc *DynamoDB) RawBatchWriteItem(ctx context.Context, in *SDK.BatchWriteItemInput) (*SDK.BatchWriteItemResponse, error) {
	return svc.client.BatchWriteItemRequest(in).Send(ctx)
}

// RawCreateBackup executes `CreateBackup` raw operation.
func (svc *DynamoDB) RawCreateBackup(ctx context.Context, in *SDK.CreateBackupInput) (*SDK.CreateBackupResponse, error) {
	return svc.client.CreateBackupRequest(in).Send(ctx)
}

// RawCreateGlobalTable executes `CreateGlobalTable` raw operation.
func (svc *DynamoDB) RawCreateGlobalTable(ctx context.Context, in *SDK.CreateGlobalTableInput) (*SDK.CreateGlobalTableResponse, error) {
	return svc.client.CreateGlobalTableRequest(in).Send(ctx)
}

// RawCreateTable executes `CreateTable` raw operation.
func (svc *DynamoDB) RawCreateTable(ctx context.Context, in *SDK.CreateTableInput) (*SDK.CreateTableResponse, error) {
	return svc.client.CreateTableRequest(in).Send(ctx)
}

// RawDeleteBackup executes `DeleteBackup` raw operation.
func (svc *DynamoDB) RawDeleteBackup(ctx context.Context, in *SDK.DeleteBackupInput) (*SDK.DeleteBackupResponse, error) {
	return svc.client.DeleteBackupRequest(in).Send(ctx)
}

// RawDeleteItem executes `DeleteItem` raw operation.
func (svc *DynamoDB) RawDeleteItem(ctx context.Context, in *SDK.DeleteItemInput) (*SDK.DeleteItemResponse, error) {
	return svc.client.DeleteItemRequest(in).Send(ctx)
}

// RawDeleteTable executes `DeleteTable` raw operation.
func (svc *DynamoDB) RawDeleteTable(ctx context.Context, in *SDK.DeleteTableInput) (*SDK.DeleteTableResponse, error) {
	return svc.client.DeleteTableRequest(in).Send(ctx)
}

// RawDescribeBackup executes `DescribeBackup` raw operation.
func (svc *DynamoDB) RawDescribeBackup(ctx context.Context, in *SDK.DescribeBackupInput) (*SDK.DescribeBackupResponse, error) {
	return svc.client.DescribeBackupRequest(in).Send(ctx)
}

// RawDescribeContinuousBackups executes `DescribeContinuousBackups` raw operation.
func (svc *DynamoDB) RawDescribeContinuousBackups(ctx context.Context, in *SDK.DescribeContinuousBackupsInput) (*SDK.DescribeContinuousBackupsResponse, error) {
	return svc.client.DescribeContinuousBackupsRequest(in).Send(ctx)
}

// RawDescribeContributorInsights executes `DescribeContributorInsights` raw operation.
func (svc *DynamoDB) RawDescribeContributorInsights(ctx context.Context, in *SDK.DescribeContributorInsightsInput) (*SDK.DescribeContributorInsightsResponse, error) {
	return svc.client.DescribeContributorInsightsRequest(in).Send(ctx)
}

// RawDescribeEndpoints executes `DescribeEndpoints` raw operation.
func (svc *DynamoDB) RawDescribeEndpoints(ctx context.Context, in *SDK.DescribeEndpointsInput) (*SDK.DescribeEndpointsResponse, error) {
	return svc.client.DescribeEndpointsRequest(in).Send(ctx)
}

// RawDescribeGlobalTable executes `DescribeGlobalTable` raw operation.
func (svc *DynamoDB) RawDescribeGlobalTable(ctx context.Context, in *SDK.DescribeGlobalTableInput) (*SDK.DescribeGlobalTableResponse, error) {
	return svc.client.DescribeGlobalTableRequest(in).Send(ctx)
}

// RawDescribeGlobalTableSettings executes `DescribeGlobalTableSettings` raw operation.
func (svc *DynamoDB) RawDescribeGlobalTableSettings(ctx context.Context, in *SDK.DescribeGlobalTableSettingsInput) (*SDK.DescribeGlobalTableSettingsResponse, error) {
	return svc.client.DescribeGlobalTableSettingsRequest(in).Send(ctx)
}

// RawDescribeLimits executes `DescribeLimits` raw operation.
func (svc *DynamoDB) RawDescribeLimits(ctx context.Context, in *SDK.DescribeLimitsInput) (*SDK.DescribeLimitsResponse, error) {
	return svc.client.DescribeLimitsRequest(in).Send(ctx)
}

// RawDescribeTable executes `DescribeTable` raw operation.
func (svc *DynamoDB) RawDescribeTable(ctx context.Context, in *SDK.DescribeTableInput) (*SDK.DescribeTableResponse, error) {
	return svc.client.DescribeTableRequest(in).Send(ctx)
}

// RawDescribeTableReplicaAutoScaling executes `DescribeTableReplicaAutoScaling` raw operation.
func (svc *DynamoDB) RawDescribeTableReplicaAutoScaling(ctx context.Context, in *SDK.DescribeTableReplicaAutoScalingInput) (*SDK.DescribeTableReplicaAutoScalingResponse, error) {
	return svc.client.DescribeTableReplicaAutoScalingRequest(in).Send(ctx)
}

// RawDescribeTimeToLive executes `DescribeTimeToLive` raw operation.
func (svc *DynamoDB) RawDescribeTimeToLive(ctx context.Context, in *SDK.DescribeTimeToLiveInput) (*SDK.DescribeTimeToLiveResponse, error) {
	return svc.client.DescribeTimeToLiveRequest(in).Send(ctx)
}

// RawGetItem executes `GetItem` raw operation.
func (svc *DynamoDB) RawGetItem(ctx context.Context, in *SDK.GetItemInput) (*SDK.GetItemResponse, error) {
	return svc.client.GetItemRequest(in).Send(ctx)
}

// RawListBackups executes `ListBackups` raw operation.
func (svc *DynamoDB) RawListBackups(ctx context.Context, in *SDK.ListBackupsInput) (*SDK.ListBackupsResponse, error) {
	return svc.client.ListBackupsRequest(in).Send(ctx)
}

// RawListContributorInsights executes `ListContributorInsights` raw operation.
func (svc *DynamoDB) RawListContributorInsights(ctx context.Context, in *SDK.ListContributorInsightsInput) (*SDK.ListContributorInsightsResponse, error) {
	return svc.client.ListContributorInsightsRequest(in).Send(ctx)
}

// RawListGlobalTables executes `ListGlobalTables` raw operation.
func (svc *DynamoDB) RawListGlobalTables(ctx context.Context, in *SDK.ListGlobalTablesInput) (*SDK.ListGlobalTablesResponse, error) {
	return svc.client.ListGlobalTablesRequest(in).Send(ctx)
}

// RawListTables executes `ListTables` raw operation.
func (svc *DynamoDB) RawListTables(ctx context.Context, in *SDK.ListTablesInput) (*SDK.ListTablesResponse, error) {
	return svc.client.ListTablesRequest(in).Send(ctx)
}

// RawListTagsOfResource executes `ListTagsOfResource` raw operation.
func (svc *DynamoDB) RawListTagsOfResource(ctx context.Context, in *SDK.ListTagsOfResourceInput) (*SDK.ListTagsOfResourceResponse, error) {
	return svc.client.ListTagsOfResourceRequest(in).Send(ctx)
}

// RawPutItem executes `PutItem` raw operation.
func (svc *DynamoDB) RawPutItem(ctx context.Context, in *SDK.PutItemInput) (*SDK.PutItemResponse, error) {
	return svc.client.PutItemRequest(in).Send(ctx)
}

// RawQuery executes `Query` raw operation.
func (svc *DynamoDB) RawQuery(ctx context.Context, in *SDK.QueryInput) (*SDK.QueryResponse, error) {
	return svc.client.QueryRequest(in).Send(ctx)
}

// RawRestoreTableFromBackup executes `RestoreTableFromBackup` raw operation.
func (svc *DynamoDB) RawRestoreTableFromBackup(ctx context.Context, in *SDK.RestoreTableFromBackupInput) (*SDK.RestoreTableFromBackupResponse, error) {
	return svc.client.RestoreTableFromBackupRequest(in).Send(ctx)
}

// RawRestoreTableToPointInTime executes `RestoreTableToPointInTime` raw operation.
func (svc *DynamoDB) RawRestoreTableToPointInTime(ctx context.Context, in *SDK.RestoreTableToPointInTimeInput) (*SDK.RestoreTableToPointInTimeResponse, error) {
	return svc.client.RestoreTableToPointInTimeRequest(in).Send(ctx)
}

// RawScan executes `Scan` raw operation.
func (svc *DynamoDB) RawScan(ctx context.Context, in *SDK.ScanInput) (*SDK.ScanResponse, error) {
	return svc.client.ScanRequest(in).Send(ctx)
}

// RawTagResource executes `TagResource` raw operation.
func (svc *DynamoDB) RawTagResource(ctx context.Context, in *SDK.TagResourceInput) (*SDK.TagResourceResponse, error) {
	return svc.client.TagResourceRequest(in).Send(ctx)
}

// RawTransactGetItems executes `TransactGetItems` raw operation.
func (svc *DynamoDB) RawTransactGetItems(ctx context.Context, in *SDK.TransactGetItemsInput) (*SDK.TransactGetItemsResponse, error) {
	return svc.client.TransactGetItemsRequest(in).Send(ctx)
}

// RawTransactWriteItems executes `TransactWriteItems` raw operation.
func (svc *DynamoDB) RawTransactWriteItems(ctx context.Context, in *SDK.TransactWriteItemsInput) (*SDK.TransactWriteItemsResponse, error) {
	return svc.client.TransactWriteItemsRequest(in).Send(ctx)
}

// RawUntagResource executes `UntagResource` raw operation.
func (svc *DynamoDB) RawUntagResource(ctx context.Context, in *SDK.UntagResourceInput) (*SDK.UntagResourceResponse, error) {
	return svc.client.UntagResourceRequest(in).Send(ctx)
}

// RawUpdateContinuousBackups executes `UpdateContinuousBackups` raw operation.
func (svc *DynamoDB) RawUpdateContinuousBackups(ctx context.Context, in *SDK.UpdateContinuousBackupsInput) (*SDK.UpdateContinuousBackupsResponse, error) {
	return svc.client.UpdateContinuousBackupsRequest(in).Send(ctx)
}

// RawUpdateContributorInsights executes `UpdateContributorInsights` raw operation.
func (svc *DynamoDB) RawUpdateContributorInsights(ctx context.Context, in *SDK.UpdateContributorInsightsInput) (*SDK.UpdateContributorInsightsResponse, error) {
	return svc.client.UpdateContributorInsightsRequest(in).Send(ctx)
}

// RawUpdateGlobalTable executes `UpdateGlobalTable` raw operation.
func (svc *DynamoDB) RawUpdateGlobalTable(ctx context.Context, in *SDK.UpdateGlobalTableInput) (*SDK.UpdateGlobalTableResponse, error) {
	return svc.client.UpdateGlobalTableRequest(in).Send(ctx)
}

// RawUpdateGlobalTableSettings executes `UpdateGlobalTableSettings` raw operation.
func (svc *DynamoDB) RawUpdateGlobalTableSettings(ctx context.Context, in *SDK.UpdateGlobalTableSettingsInput) (*SDK.UpdateGlobalTableSettingsResponse, error) {
	return svc.client.UpdateGlobalTableSettingsRequest(in).Send(ctx)
}

// RawUpdateItem executes `UpdateItem` raw operation.
func (svc *DynamoDB) RawUpdateItem(ctx context.Context, in *SDK.UpdateItemInput) (*SDK.UpdateItemResponse, error) {
	return svc.client.UpdateItemRequest(in).Send(ctx)
}

// RawUpdateTable executes `UpdateTable` raw operation.
func (svc *DynamoDB) RawUpdateTable(ctx context.Context, in *SDK.UpdateTableInput) (*SDK.UpdateTableResponse, error) {
	return svc.client.UpdateTableRequest(in).Send(ctx)
}

// RawUpdateTableReplicaAutoScaling executes `UpdateTableReplicaAutoScaling` raw operation.
func (svc *DynamoDB) RawUpdateTableReplicaAutoScaling(ctx context.Context, in *SDK.UpdateTableReplicaAutoScalingInput) (*SDK.UpdateTableReplicaAutoScalingResponse, error) {
	return svc.client.UpdateTableReplicaAutoScalingRequest(in).Send(ctx)
}

// RawUpdateTimeToLive executes `UpdateTimeToLive` raw operation.
func (svc *DynamoDB) RawUpdateTimeToLive(ctx context.Context, in *SDK.UpdateTimeToLiveInput) (*SDK.UpdateTimeToLiveResponse, error) {
	return svc.client.UpdateTimeToLiveRequest(in).Send(ctx)
}
