package athena

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/athena"
)

// RawBatchGetNamedQuery executes `BatchGetNamedQuery` raw operation.
func (svc *Athena) RawBatchGetNamedQuery(ctx context.Context, in *SDK.BatchGetNamedQueryInput) (*SDK.BatchGetNamedQueryResponse, error) {
	return svc.client.BatchGetNamedQueryRequest(in).Send(ctx)
}

// RawBatchGetQueryExecution executes `BatchGetQueryExecution` raw operation.
func (svc *Athena) RawBatchGetQueryExecution(ctx context.Context, in *SDK.BatchGetQueryExecutionInput) (*SDK.BatchGetQueryExecutionResponse, error) {
	return svc.client.BatchGetQueryExecutionRequest(in).Send(ctx)
}

// RawCreateDataCatalog executes `CreateDataCatalog` raw operation.
func (svc *Athena) RawCreateDataCatalog(ctx context.Context, in *SDK.CreateDataCatalogInput) (*SDK.CreateDataCatalogResponse, error) {
	return svc.client.CreateDataCatalogRequest(in).Send(ctx)
}

// RawCreateNamedQuery executes `CreateNamedQuery` raw operation.
func (svc *Athena) RawCreateNamedQuery(ctx context.Context, in *SDK.CreateNamedQueryInput) (*SDK.CreateNamedQueryResponse, error) {
	return svc.client.CreateNamedQueryRequest(in).Send(ctx)
}

// RawCreateWorkGroup executes `CreateWorkGroup` raw operation.
func (svc *Athena) RawCreateWorkGroup(ctx context.Context, in *SDK.CreateWorkGroupInput) (*SDK.CreateWorkGroupResponse, error) {
	return svc.client.CreateWorkGroupRequest(in).Send(ctx)
}

// RawDeleteDataCatalog executes `DeleteDataCatalog` raw operation.
func (svc *Athena) RawDeleteDataCatalog(ctx context.Context, in *SDK.DeleteDataCatalogInput) (*SDK.DeleteDataCatalogResponse, error) {
	return svc.client.DeleteDataCatalogRequest(in).Send(ctx)
}

// RawDeleteNamedQuery executes `DeleteNamedQuery` raw operation.
func (svc *Athena) RawDeleteNamedQuery(ctx context.Context, in *SDK.DeleteNamedQueryInput) (*SDK.DeleteNamedQueryResponse, error) {
	return svc.client.DeleteNamedQueryRequest(in).Send(ctx)
}

// RawDeleteWorkGroup executes `DeleteWorkGroup` raw operation.
func (svc *Athena) RawDeleteWorkGroup(ctx context.Context, in *SDK.DeleteWorkGroupInput) (*SDK.DeleteWorkGroupResponse, error) {
	return svc.client.DeleteWorkGroupRequest(in).Send(ctx)
}

// RawGetDataCatalog executes `GetDataCatalog` raw operation.
func (svc *Athena) RawGetDataCatalog(ctx context.Context, in *SDK.GetDataCatalogInput) (*SDK.GetDataCatalogResponse, error) {
	return svc.client.GetDataCatalogRequest(in).Send(ctx)
}

// RawGetDatabase executes `GetDatabase` raw operation.
func (svc *Athena) RawGetDatabase(ctx context.Context, in *SDK.GetDatabaseInput) (*SDK.GetDatabaseResponse, error) {
	return svc.client.GetDatabaseRequest(in).Send(ctx)
}

// RawGetNamedQuery executes `GetNamedQuery` raw operation.
func (svc *Athena) RawGetNamedQuery(ctx context.Context, in *SDK.GetNamedQueryInput) (*SDK.GetNamedQueryResponse, error) {
	return svc.client.GetNamedQueryRequest(in).Send(ctx)
}

// RawGetQueryExecution executes `GetQueryExecution` raw operation.
func (svc *Athena) RawGetQueryExecution(ctx context.Context, in *SDK.GetQueryExecutionInput) (*SDK.GetQueryExecutionResponse, error) {
	return svc.client.GetQueryExecutionRequest(in).Send(ctx)
}

// RawGetQueryResults executes `GetQueryResults` raw operation.
func (svc *Athena) RawGetQueryResults(ctx context.Context, in *SDK.GetQueryResultsInput) (*SDK.GetQueryResultsResponse, error) {
	return svc.client.GetQueryResultsRequest(in).Send(ctx)
}

// RawGetTableMetadata executes `GetTableMetadata` raw operation.
func (svc *Athena) RawGetTableMetadata(ctx context.Context, in *SDK.GetTableMetadataInput) (*SDK.GetTableMetadataResponse, error) {
	return svc.client.GetTableMetadataRequest(in).Send(ctx)
}

// RawGetWorkGroup executes `GetWorkGroup` raw operation.
func (svc *Athena) RawGetWorkGroup(ctx context.Context, in *SDK.GetWorkGroupInput) (*SDK.GetWorkGroupResponse, error) {
	return svc.client.GetWorkGroupRequest(in).Send(ctx)
}

// RawListDataCatalogs executes `ListDataCatalogs` raw operation.
func (svc *Athena) RawListDataCatalogs(ctx context.Context, in *SDK.ListDataCatalogsInput) (*SDK.ListDataCatalogsResponse, error) {
	return svc.client.ListDataCatalogsRequest(in).Send(ctx)
}

// RawListDatabases executes `ListDatabases` raw operation.
func (svc *Athena) RawListDatabases(ctx context.Context, in *SDK.ListDatabasesInput) (*SDK.ListDatabasesResponse, error) {
	return svc.client.ListDatabasesRequest(in).Send(ctx)
}

// RawListNamedQueries executes `ListNamedQueries` raw operation.
func (svc *Athena) RawListNamedQueries(ctx context.Context, in *SDK.ListNamedQueriesInput) (*SDK.ListNamedQueriesResponse, error) {
	return svc.client.ListNamedQueriesRequest(in).Send(ctx)
}

// RawListQueryExecutions executes `ListQueryExecutions` raw operation.
func (svc *Athena) RawListQueryExecutions(ctx context.Context, in *SDK.ListQueryExecutionsInput) (*SDK.ListQueryExecutionsResponse, error) {
	return svc.client.ListQueryExecutionsRequest(in).Send(ctx)
}

// RawListTableMetadata executes `ListTableMetadata` raw operation.
func (svc *Athena) RawListTableMetadata(ctx context.Context, in *SDK.ListTableMetadataInput) (*SDK.ListTableMetadataResponse, error) {
	return svc.client.ListTableMetadataRequest(in).Send(ctx)
}

// RawListTagsForResource executes `ListTagsForResource` raw operation.
func (svc *Athena) RawListTagsForResource(ctx context.Context, in *SDK.ListTagsForResourceInput) (*SDK.ListTagsForResourceResponse, error) {
	return svc.client.ListTagsForResourceRequest(in).Send(ctx)
}

// RawListWorkGroups executes `ListWorkGroups` raw operation.
func (svc *Athena) RawListWorkGroups(ctx context.Context, in *SDK.ListWorkGroupsInput) (*SDK.ListWorkGroupsResponse, error) {
	return svc.client.ListWorkGroupsRequest(in).Send(ctx)
}

// RawStartQueryExecution executes `StartQueryExecution` raw operation.
func (svc *Athena) RawStartQueryExecution(ctx context.Context, in *SDK.StartQueryExecutionInput) (*SDK.StartQueryExecutionResponse, error) {
	return svc.client.StartQueryExecutionRequest(in).Send(ctx)
}

// RawStopQueryExecution executes `StopQueryExecution` raw operation.
func (svc *Athena) RawStopQueryExecution(ctx context.Context, in *SDK.StopQueryExecutionInput) (*SDK.StopQueryExecutionResponse, error) {
	return svc.client.StopQueryExecutionRequest(in).Send(ctx)
}
