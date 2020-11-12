package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"
)

// RawAddTagsToResource executes `AddTagsToResource` raw operation.
func (svc *SSM) RawAddTagsToResource(ctx context.Context, in *SDK.AddTagsToResourceInput) (*SDK.AddTagsToResourceResponse, error) {
	return svc.client.AddTagsToResourceRequest(in).Send(ctx)
}

// RawCancelCommand executes `CancelCommand` raw operation.
func (svc *SSM) RawCancelCommand(ctx context.Context, in *SDK.CancelCommandInput) (*SDK.CancelCommandResponse, error) {
	return svc.client.CancelCommandRequest(in).Send(ctx)
}

// RawCancelMaintenanceWindowExecution executes `CancelMaintenanceWindowExecution` raw operation.
func (svc *SSM) RawCancelMaintenanceWindowExecution(ctx context.Context, in *SDK.CancelMaintenanceWindowExecutionInput) (*SDK.CancelMaintenanceWindowExecutionResponse, error) {
	return svc.client.CancelMaintenanceWindowExecutionRequest(in).Send(ctx)
}

// RawCreateActivation executes `CreateActivation` raw operation.
func (svc *SSM) RawCreateActivation(ctx context.Context, in *SDK.CreateActivationInput) (*SDK.CreateActivationResponse, error) {
	return svc.client.CreateActivationRequest(in).Send(ctx)
}

// RawCreateAssociation executes `CreateAssociation` raw operation.
func (svc *SSM) RawCreateAssociation(ctx context.Context, in *SDK.CreateAssociationInput) (*SDK.CreateAssociationResponse, error) {
	return svc.client.CreateAssociationRequest(in).Send(ctx)
}

// RawCreateAssociationBatch executes `CreateAssociationBatch` raw operation.
func (svc *SSM) RawCreateAssociationBatch(ctx context.Context, in *SDK.CreateAssociationBatchInput) (*SDK.CreateAssociationBatchResponse, error) {
	return svc.client.CreateAssociationBatchRequest(in).Send(ctx)
}

// RawCreateDocument executes `CreateDocument` raw operation.
func (svc *SSM) RawCreateDocument(ctx context.Context, in *SDK.CreateDocumentInput) (*SDK.CreateDocumentResponse, error) {
	return svc.client.CreateDocumentRequest(in).Send(ctx)
}

// RawCreateMaintenanceWindow executes `CreateMaintenanceWindow` raw operation.
func (svc *SSM) RawCreateMaintenanceWindow(ctx context.Context, in *SDK.CreateMaintenanceWindowInput) (*SDK.CreateMaintenanceWindowResponse, error) {
	return svc.client.CreateMaintenanceWindowRequest(in).Send(ctx)
}

// RawCreateOpsItem executes `CreateOpsItem` raw operation.
func (svc *SSM) RawCreateOpsItem(ctx context.Context, in *SDK.CreateOpsItemInput) (*SDK.CreateOpsItemResponse, error) {
	return svc.client.CreateOpsItemRequest(in).Send(ctx)
}

// RawCreatePatchBaseline executes `CreatePatchBaseline` raw operation.
func (svc *SSM) RawCreatePatchBaseline(ctx context.Context, in *SDK.CreatePatchBaselineInput) (*SDK.CreatePatchBaselineResponse, error) {
	return svc.client.CreatePatchBaselineRequest(in).Send(ctx)
}

// RawCreateResourceDataSync executes `CreateResourceDataSync` raw operation.
func (svc *SSM) RawCreateResourceDataSync(ctx context.Context, in *SDK.CreateResourceDataSyncInput) (*SDK.CreateResourceDataSyncResponse, error) {
	return svc.client.CreateResourceDataSyncRequest(in).Send(ctx)
}

// RawDeleteActivation executes `DeleteActivation` raw operation.
func (svc *SSM) RawDeleteActivation(ctx context.Context, in *SDK.DeleteActivationInput) (*SDK.DeleteActivationResponse, error) {
	return svc.client.DeleteActivationRequest(in).Send(ctx)
}

// RawDeleteAssociation executes `DeleteAssociation` raw operation.
func (svc *SSM) RawDeleteAssociation(ctx context.Context, in *SDK.DeleteAssociationInput) (*SDK.DeleteAssociationResponse, error) {
	return svc.client.DeleteAssociationRequest(in).Send(ctx)
}

// RawDeleteDocument executes `DeleteDocument` raw operation.
func (svc *SSM) RawDeleteDocument(ctx context.Context, in *SDK.DeleteDocumentInput) (*SDK.DeleteDocumentResponse, error) {
	return svc.client.DeleteDocumentRequest(in).Send(ctx)
}

// RawDeleteInventory executes `DeleteInventory` raw operation.
func (svc *SSM) RawDeleteInventory(ctx context.Context, in *SDK.DeleteInventoryInput) (*SDK.DeleteInventoryResponse, error) {
	return svc.client.DeleteInventoryRequest(in).Send(ctx)
}

// RawDeleteMaintenanceWindow executes `DeleteMaintenanceWindow` raw operation.
func (svc *SSM) RawDeleteMaintenanceWindow(ctx context.Context, in *SDK.DeleteMaintenanceWindowInput) (*SDK.DeleteMaintenanceWindowResponse, error) {
	return svc.client.DeleteMaintenanceWindowRequest(in).Send(ctx)
}

// RawDeleteParameter executes `DeleteParameter` raw operation.
func (svc *SSM) RawDeleteParameter(ctx context.Context, in *SDK.DeleteParameterInput) (*SDK.DeleteParameterResponse, error) {
	return svc.client.DeleteParameterRequest(in).Send(ctx)
}

// RawDeleteParameters executes `DeleteParameters` raw operation.
func (svc *SSM) RawDeleteParameters(ctx context.Context, in *SDK.DeleteParametersInput) (*SDK.DeleteParametersResponse, error) {
	return svc.client.DeleteParametersRequest(in).Send(ctx)
}

// RawDeletePatchBaseline executes `DeletePatchBaseline` raw operation.
func (svc *SSM) RawDeletePatchBaseline(ctx context.Context, in *SDK.DeletePatchBaselineInput) (*SDK.DeletePatchBaselineResponse, error) {
	return svc.client.DeletePatchBaselineRequest(in).Send(ctx)
}

// RawDeleteResourceDataSync executes `DeleteResourceDataSync` raw operation.
func (svc *SSM) RawDeleteResourceDataSync(ctx context.Context, in *SDK.DeleteResourceDataSyncInput) (*SDK.DeleteResourceDataSyncResponse, error) {
	return svc.client.DeleteResourceDataSyncRequest(in).Send(ctx)
}

// RawDeregisterManagedInstance executes `DeregisterManagedInstance` raw operation.
func (svc *SSM) RawDeregisterManagedInstance(ctx context.Context, in *SDK.DeregisterManagedInstanceInput) (*SDK.DeregisterManagedInstanceResponse, error) {
	return svc.client.DeregisterManagedInstanceRequest(in).Send(ctx)
}

// RawDeregisterPatchBaselineForPatchGroup executes `DeregisterPatchBaselineForPatchGroup` raw operation.
func (svc *SSM) RawDeregisterPatchBaselineForPatchGroup(ctx context.Context, in *SDK.DeregisterPatchBaselineForPatchGroupInput) (*SDK.DeregisterPatchBaselineForPatchGroupResponse, error) {
	return svc.client.DeregisterPatchBaselineForPatchGroupRequest(in).Send(ctx)
}

// RawDeregisterTargetFromMaintenanceWindow executes `DeregisterTargetFromMaintenanceWindow` raw operation.
func (svc *SSM) RawDeregisterTargetFromMaintenanceWindow(ctx context.Context, in *SDK.DeregisterTargetFromMaintenanceWindowInput) (*SDK.DeregisterTargetFromMaintenanceWindowResponse, error) {
	return svc.client.DeregisterTargetFromMaintenanceWindowRequest(in).Send(ctx)
}

// RawDeregisterTaskFromMaintenanceWindow executes `DeregisterTaskFromMaintenanceWindow` raw operation.
func (svc *SSM) RawDeregisterTaskFromMaintenanceWindow(ctx context.Context, in *SDK.DeregisterTaskFromMaintenanceWindowInput) (*SDK.DeregisterTaskFromMaintenanceWindowResponse, error) {
	return svc.client.DeregisterTaskFromMaintenanceWindowRequest(in).Send(ctx)
}

// RawDescribeActivations executes `DescribeActivations` raw operation.
func (svc *SSM) RawDescribeActivations(ctx context.Context, in *SDK.DescribeActivationsInput) (*SDK.DescribeActivationsResponse, error) {
	return svc.client.DescribeActivationsRequest(in).Send(ctx)
}

// RawDescribeAssociation executes `DescribeAssociation` raw operation.
func (svc *SSM) RawDescribeAssociation(ctx context.Context, in *SDK.DescribeAssociationInput) (*SDK.DescribeAssociationResponse, error) {
	return svc.client.DescribeAssociationRequest(in).Send(ctx)
}

// RawDescribeAssociationExecutionTargets executes `DescribeAssociationExecutionTargets` raw operation.
func (svc *SSM) RawDescribeAssociationExecutionTargets(ctx context.Context, in *SDK.DescribeAssociationExecutionTargetsInput) (*SDK.DescribeAssociationExecutionTargetsResponse, error) {
	return svc.client.DescribeAssociationExecutionTargetsRequest(in).Send(ctx)
}

// RawDescribeAssociationExecutions executes `DescribeAssociationExecutions` raw operation.
func (svc *SSM) RawDescribeAssociationExecutions(ctx context.Context, in *SDK.DescribeAssociationExecutionsInput) (*SDK.DescribeAssociationExecutionsResponse, error) {
	return svc.client.DescribeAssociationExecutionsRequest(in).Send(ctx)
}

// RawDescribeAutomationExecutions executes `DescribeAutomationExecutions` raw operation.
func (svc *SSM) RawDescribeAutomationExecutions(ctx context.Context, in *SDK.DescribeAutomationExecutionsInput) (*SDK.DescribeAutomationExecutionsResponse, error) {
	return svc.client.DescribeAutomationExecutionsRequest(in).Send(ctx)
}

// RawDescribeAutomationStepExecutions executes `DescribeAutomationStepExecutions` raw operation.
func (svc *SSM) RawDescribeAutomationStepExecutions(ctx context.Context, in *SDK.DescribeAutomationStepExecutionsInput) (*SDK.DescribeAutomationStepExecutionsResponse, error) {
	return svc.client.DescribeAutomationStepExecutionsRequest(in).Send(ctx)
}

// RawDescribeAvailablePatches executes `DescribeAvailablePatches` raw operation.
func (svc *SSM) RawDescribeAvailablePatches(ctx context.Context, in *SDK.DescribeAvailablePatchesInput) (*SDK.DescribeAvailablePatchesResponse, error) {
	return svc.client.DescribeAvailablePatchesRequest(in).Send(ctx)
}

// RawDescribeDocument executes `DescribeDocument` raw operation.
func (svc *SSM) RawDescribeDocument(ctx context.Context, in *SDK.DescribeDocumentInput) (*SDK.DescribeDocumentResponse, error) {
	return svc.client.DescribeDocumentRequest(in).Send(ctx)
}

// RawDescribeDocumentPermission executes `DescribeDocumentPermission` raw operation.
func (svc *SSM) RawDescribeDocumentPermission(ctx context.Context, in *SDK.DescribeDocumentPermissionInput) (*SDK.DescribeDocumentPermissionResponse, error) {
	return svc.client.DescribeDocumentPermissionRequest(in).Send(ctx)
}

// RawDescribeEffectiveInstanceAssociations executes `DescribeEffectiveInstanceAssociations` raw operation.
func (svc *SSM) RawDescribeEffectiveInstanceAssociations(ctx context.Context, in *SDK.DescribeEffectiveInstanceAssociationsInput) (*SDK.DescribeEffectiveInstanceAssociationsResponse, error) {
	return svc.client.DescribeEffectiveInstanceAssociationsRequest(in).Send(ctx)
}

// RawDescribeEffectivePatchesForPatchBaseline executes `DescribeEffectivePatchesForPatchBaseline` raw operation.
func (svc *SSM) RawDescribeEffectivePatchesForPatchBaseline(ctx context.Context, in *SDK.DescribeEffectivePatchesForPatchBaselineInput) (*SDK.DescribeEffectivePatchesForPatchBaselineResponse, error) {
	return svc.client.DescribeEffectivePatchesForPatchBaselineRequest(in).Send(ctx)
}

// RawDescribeInstanceAssociationsStatus executes `DescribeInstanceAssociationsStatus` raw operation.
func (svc *SSM) RawDescribeInstanceAssociationsStatus(ctx context.Context, in *SDK.DescribeInstanceAssociationsStatusInput) (*SDK.DescribeInstanceAssociationsStatusResponse, error) {
	return svc.client.DescribeInstanceAssociationsStatusRequest(in).Send(ctx)
}

// RawDescribeInstanceInformation executes `DescribeInstanceInformation` raw operation.
func (svc *SSM) RawDescribeInstanceInformation(ctx context.Context, in *SDK.DescribeInstanceInformationInput) (*SDK.DescribeInstanceInformationResponse, error) {
	return svc.client.DescribeInstanceInformationRequest(in).Send(ctx)
}

// RawDescribeInstancePatchStates executes `DescribeInstancePatchStates` raw operation.
func (svc *SSM) RawDescribeInstancePatchStates(ctx context.Context, in *SDK.DescribeInstancePatchStatesInput) (*SDK.DescribeInstancePatchStatesResponse, error) {
	return svc.client.DescribeInstancePatchStatesRequest(in).Send(ctx)
}

// RawDescribeInstancePatchStatesForPatchGroup executes `DescribeInstancePatchStatesForPatchGroup` raw operation.
func (svc *SSM) RawDescribeInstancePatchStatesForPatchGroup(ctx context.Context, in *SDK.DescribeInstancePatchStatesForPatchGroupInput) (*SDK.DescribeInstancePatchStatesForPatchGroupResponse, error) {
	return svc.client.DescribeInstancePatchStatesForPatchGroupRequest(in).Send(ctx)
}

// RawDescribeInstancePatches executes `DescribeInstancePatches` raw operation.
func (svc *SSM) RawDescribeInstancePatches(ctx context.Context, in *SDK.DescribeInstancePatchesInput) (*SDK.DescribeInstancePatchesResponse, error) {
	return svc.client.DescribeInstancePatchesRequest(in).Send(ctx)
}

// RawDescribeInventoryDeletions executes `DescribeInventoryDeletions` raw operation.
func (svc *SSM) RawDescribeInventoryDeletions(ctx context.Context, in *SDK.DescribeInventoryDeletionsInput) (*SDK.DescribeInventoryDeletionsResponse, error) {
	return svc.client.DescribeInventoryDeletionsRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowExecutionTaskInvocations executes `DescribeMaintenanceWindowExecutionTaskInvocations` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowExecutionTaskInvocations(ctx context.Context, in *SDK.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*SDK.DescribeMaintenanceWindowExecutionTaskInvocationsResponse, error) {
	return svc.client.DescribeMaintenanceWindowExecutionTaskInvocationsRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowExecutionTasks executes `DescribeMaintenanceWindowExecutionTasks` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowExecutionTasks(ctx context.Context, in *SDK.DescribeMaintenanceWindowExecutionTasksInput) (*SDK.DescribeMaintenanceWindowExecutionTasksResponse, error) {
	return svc.client.DescribeMaintenanceWindowExecutionTasksRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowExecutions executes `DescribeMaintenanceWindowExecutions` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowExecutions(ctx context.Context, in *SDK.DescribeMaintenanceWindowExecutionsInput) (*SDK.DescribeMaintenanceWindowExecutionsResponse, error) {
	return svc.client.DescribeMaintenanceWindowExecutionsRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowSchedule executes `DescribeMaintenanceWindowSchedule` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowSchedule(ctx context.Context, in *SDK.DescribeMaintenanceWindowScheduleInput) (*SDK.DescribeMaintenanceWindowScheduleResponse, error) {
	return svc.client.DescribeMaintenanceWindowScheduleRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowTargets executes `DescribeMaintenanceWindowTargets` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowTargets(ctx context.Context, in *SDK.DescribeMaintenanceWindowTargetsInput) (*SDK.DescribeMaintenanceWindowTargetsResponse, error) {
	return svc.client.DescribeMaintenanceWindowTargetsRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowTasks executes `DescribeMaintenanceWindowTasks` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowTasks(ctx context.Context, in *SDK.DescribeMaintenanceWindowTasksInput) (*SDK.DescribeMaintenanceWindowTasksResponse, error) {
	return svc.client.DescribeMaintenanceWindowTasksRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindows executes `DescribeMaintenanceWindows` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindows(ctx context.Context, in *SDK.DescribeMaintenanceWindowsInput) (*SDK.DescribeMaintenanceWindowsResponse, error) {
	return svc.client.DescribeMaintenanceWindowsRequest(in).Send(ctx)
}

// RawDescribeMaintenanceWindowsForTarget executes `DescribeMaintenanceWindowsForTarget` raw operation.
func (svc *SSM) RawDescribeMaintenanceWindowsForTarget(ctx context.Context, in *SDK.DescribeMaintenanceWindowsForTargetInput) (*SDK.DescribeMaintenanceWindowsForTargetResponse, error) {
	return svc.client.DescribeMaintenanceWindowsForTargetRequest(in).Send(ctx)
}

// RawDescribeOpsItems executes `DescribeOpsItems` raw operation.
func (svc *SSM) RawDescribeOpsItems(ctx context.Context, in *SDK.DescribeOpsItemsInput) (*SDK.DescribeOpsItemsResponse, error) {
	return svc.client.DescribeOpsItemsRequest(in).Send(ctx)
}

// RawDescribeParameters executes `DescribeParameters` raw operation.
func (svc *SSM) RawDescribeParameters(ctx context.Context, in *SDK.DescribeParametersInput) (*SDK.DescribeParametersResponse, error) {
	return svc.client.DescribeParametersRequest(in).Send(ctx)
}

// RawDescribePatchBaselines executes `DescribePatchBaselines` raw operation.
func (svc *SSM) RawDescribePatchBaselines(ctx context.Context, in *SDK.DescribePatchBaselinesInput) (*SDK.DescribePatchBaselinesResponse, error) {
	return svc.client.DescribePatchBaselinesRequest(in).Send(ctx)
}

// RawDescribePatchGroupState executes `DescribePatchGroupState` raw operation.
func (svc *SSM) RawDescribePatchGroupState(ctx context.Context, in *SDK.DescribePatchGroupStateInput) (*SDK.DescribePatchGroupStateResponse, error) {
	return svc.client.DescribePatchGroupStateRequest(in).Send(ctx)
}

// RawDescribePatchGroups executes `DescribePatchGroups` raw operation.
func (svc *SSM) RawDescribePatchGroups(ctx context.Context, in *SDK.DescribePatchGroupsInput) (*SDK.DescribePatchGroupsResponse, error) {
	return svc.client.DescribePatchGroupsRequest(in).Send(ctx)
}

// RawDescribePatchProperties executes `DescribePatchProperties` raw operation.
func (svc *SSM) RawDescribePatchProperties(ctx context.Context, in *SDK.DescribePatchPropertiesInput) (*SDK.DescribePatchPropertiesResponse, error) {
	return svc.client.DescribePatchPropertiesRequest(in).Send(ctx)
}

// RawDescribeSessions executes `DescribeSessions` raw operation.
func (svc *SSM) RawDescribeSessions(ctx context.Context, in *SDK.DescribeSessionsInput) (*SDK.DescribeSessionsResponse, error) {
	return svc.client.DescribeSessionsRequest(in).Send(ctx)
}

// RawGetAutomationExecution executes `GetAutomationExecution` raw operation.
func (svc *SSM) RawGetAutomationExecution(ctx context.Context, in *SDK.GetAutomationExecutionInput) (*SDK.GetAutomationExecutionResponse, error) {
	return svc.client.GetAutomationExecutionRequest(in).Send(ctx)
}

// RawGetCalendarState executes `GetCalendarState` raw operation.
func (svc *SSM) RawGetCalendarState(ctx context.Context, in *SDK.GetCalendarStateInput) (*SDK.GetCalendarStateResponse, error) {
	return svc.client.GetCalendarStateRequest(in).Send(ctx)
}

// RawGetCommandInvocation executes `GetCommandInvocation` raw operation.
func (svc *SSM) RawGetCommandInvocation(ctx context.Context, in *SDK.GetCommandInvocationInput) (*SDK.GetCommandInvocationResponse, error) {
	return svc.client.GetCommandInvocationRequest(in).Send(ctx)
}

// RawGetConnectionStatus executes `GetConnectionStatus` raw operation.
func (svc *SSM) RawGetConnectionStatus(ctx context.Context, in *SDK.GetConnectionStatusInput) (*SDK.GetConnectionStatusResponse, error) {
	return svc.client.GetConnectionStatusRequest(in).Send(ctx)
}

// RawGetDefaultPatchBaseline executes `GetDefaultPatchBaseline` raw operation.
func (svc *SSM) RawGetDefaultPatchBaseline(ctx context.Context, in *SDK.GetDefaultPatchBaselineInput) (*SDK.GetDefaultPatchBaselineResponse, error) {
	return svc.client.GetDefaultPatchBaselineRequest(in).Send(ctx)
}

// RawGetDeployablePatchSnapshotForInstance executes `GetDeployablePatchSnapshotForInstance` raw operation.
func (svc *SSM) RawGetDeployablePatchSnapshotForInstance(ctx context.Context, in *SDK.GetDeployablePatchSnapshotForInstanceInput) (*SDK.GetDeployablePatchSnapshotForInstanceResponse, error) {
	return svc.client.GetDeployablePatchSnapshotForInstanceRequest(in).Send(ctx)
}

// RawGetDocument executes `GetDocument` raw operation.
func (svc *SSM) RawGetDocument(ctx context.Context, in *SDK.GetDocumentInput) (*SDK.GetDocumentResponse, error) {
	return svc.client.GetDocumentRequest(in).Send(ctx)
}

// RawGetInventory executes `GetInventory` raw operation.
func (svc *SSM) RawGetInventory(ctx context.Context, in *SDK.GetInventoryInput) (*SDK.GetInventoryResponse, error) {
	return svc.client.GetInventoryRequest(in).Send(ctx)
}

// RawGetInventorySchema executes `GetInventorySchema` raw operation.
func (svc *SSM) RawGetInventorySchema(ctx context.Context, in *SDK.GetInventorySchemaInput) (*SDK.GetInventorySchemaResponse, error) {
	return svc.client.GetInventorySchemaRequest(in).Send(ctx)
}

// RawGetMaintenanceWindow executes `GetMaintenanceWindow` raw operation.
func (svc *SSM) RawGetMaintenanceWindow(ctx context.Context, in *SDK.GetMaintenanceWindowInput) (*SDK.GetMaintenanceWindowResponse, error) {
	return svc.client.GetMaintenanceWindowRequest(in).Send(ctx)
}

// RawGetMaintenanceWindowExecution executes `GetMaintenanceWindowExecution` raw operation.
func (svc *SSM) RawGetMaintenanceWindowExecution(ctx context.Context, in *SDK.GetMaintenanceWindowExecutionInput) (*SDK.GetMaintenanceWindowExecutionResponse, error) {
	return svc.client.GetMaintenanceWindowExecutionRequest(in).Send(ctx)
}

// RawGetMaintenanceWindowExecutionTask executes `GetMaintenanceWindowExecutionTask` raw operation.
func (svc *SSM) RawGetMaintenanceWindowExecutionTask(ctx context.Context, in *SDK.GetMaintenanceWindowExecutionTaskInput) (*SDK.GetMaintenanceWindowExecutionTaskResponse, error) {
	return svc.client.GetMaintenanceWindowExecutionTaskRequest(in).Send(ctx)
}

// RawGetMaintenanceWindowExecutionTaskInvocation executes `GetMaintenanceWindowExecutionTaskInvocation` raw operation.
func (svc *SSM) RawGetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, in *SDK.GetMaintenanceWindowExecutionTaskInvocationInput) (*SDK.GetMaintenanceWindowExecutionTaskInvocationResponse, error) {
	return svc.client.GetMaintenanceWindowExecutionTaskInvocationRequest(in).Send(ctx)
}

// RawGetMaintenanceWindowTask executes `GetMaintenanceWindowTask` raw operation.
func (svc *SSM) RawGetMaintenanceWindowTask(ctx context.Context, in *SDK.GetMaintenanceWindowTaskInput) (*SDK.GetMaintenanceWindowTaskResponse, error) {
	return svc.client.GetMaintenanceWindowTaskRequest(in).Send(ctx)
}

// RawGetOpsItem executes `GetOpsItem` raw operation.
func (svc *SSM) RawGetOpsItem(ctx context.Context, in *SDK.GetOpsItemInput) (*SDK.GetOpsItemResponse, error) {
	return svc.client.GetOpsItemRequest(in).Send(ctx)
}

// RawGetOpsSummary executes `GetOpsSummary` raw operation.
func (svc *SSM) RawGetOpsSummary(ctx context.Context, in *SDK.GetOpsSummaryInput) (*SDK.GetOpsSummaryResponse, error) {
	return svc.client.GetOpsSummaryRequest(in).Send(ctx)
}

// RawGetParameter executes `GetParameter` raw operation.
func (svc *SSM) RawGetParameter(ctx context.Context, in *SDK.GetParameterInput) (*SDK.GetParameterResponse, error) {
	return svc.client.GetParameterRequest(in).Send(ctx)
}

// RawGetParameterHistory executes `GetParameterHistory` raw operation.
func (svc *SSM) RawGetParameterHistory(ctx context.Context, in *SDK.GetParameterHistoryInput) (*SDK.GetParameterHistoryResponse, error) {
	return svc.client.GetParameterHistoryRequest(in).Send(ctx)
}

// RawGetParameters executes `GetParameters` raw operation.
func (svc *SSM) RawGetParameters(ctx context.Context, in *SDK.GetParametersInput) (*SDK.GetParametersResponse, error) {
	return svc.client.GetParametersRequest(in).Send(ctx)
}

// RawGetParametersByPath executes `GetParametersByPath` raw operation.
func (svc *SSM) RawGetParametersByPath(ctx context.Context, in *SDK.GetParametersByPathInput) (*SDK.GetParametersByPathResponse, error) {
	return svc.client.GetParametersByPathRequest(in).Send(ctx)
}

// RawGetPatchBaseline executes `GetPatchBaseline` raw operation.
func (svc *SSM) RawGetPatchBaseline(ctx context.Context, in *SDK.GetPatchBaselineInput) (*SDK.GetPatchBaselineResponse, error) {
	return svc.client.GetPatchBaselineRequest(in).Send(ctx)
}

// RawGetPatchBaselineForPatchGroup executes `GetPatchBaselineForPatchGroup` raw operation.
func (svc *SSM) RawGetPatchBaselineForPatchGroup(ctx context.Context, in *SDK.GetPatchBaselineForPatchGroupInput) (*SDK.GetPatchBaselineForPatchGroupResponse, error) {
	return svc.client.GetPatchBaselineForPatchGroupRequest(in).Send(ctx)
}

// RawGetServiceSetting executes `GetServiceSetting` raw operation.
func (svc *SSM) RawGetServiceSetting(ctx context.Context, in *SDK.GetServiceSettingInput) (*SDK.GetServiceSettingResponse, error) {
	return svc.client.GetServiceSettingRequest(in).Send(ctx)
}

// RawLabelParameterVersion executes `LabelParameterVersion` raw operation.
func (svc *SSM) RawLabelParameterVersion(ctx context.Context, in *SDK.LabelParameterVersionInput) (*SDK.LabelParameterVersionResponse, error) {
	return svc.client.LabelParameterVersionRequest(in).Send(ctx)
}

// RawListAssociationVersions executes `ListAssociationVersions` raw operation.
func (svc *SSM) RawListAssociationVersions(ctx context.Context, in *SDK.ListAssociationVersionsInput) (*SDK.ListAssociationVersionsResponse, error) {
	return svc.client.ListAssociationVersionsRequest(in).Send(ctx)
}

// RawListAssociations executes `ListAssociations` raw operation.
func (svc *SSM) RawListAssociations(ctx context.Context, in *SDK.ListAssociationsInput) (*SDK.ListAssociationsResponse, error) {
	return svc.client.ListAssociationsRequest(in).Send(ctx)
}

// RawListCommandInvocations executes `ListCommandInvocations` raw operation.
func (svc *SSM) RawListCommandInvocations(ctx context.Context, in *SDK.ListCommandInvocationsInput) (*SDK.ListCommandInvocationsResponse, error) {
	return svc.client.ListCommandInvocationsRequest(in).Send(ctx)
}

// RawListCommands executes `ListCommands` raw operation.
func (svc *SSM) RawListCommands(ctx context.Context, in *SDK.ListCommandsInput) (*SDK.ListCommandsResponse, error) {
	return svc.client.ListCommandsRequest(in).Send(ctx)
}

// RawListComplianceItems executes `ListComplianceItems` raw operation.
func (svc *SSM) RawListComplianceItems(ctx context.Context, in *SDK.ListComplianceItemsInput) (*SDK.ListComplianceItemsResponse, error) {
	return svc.client.ListComplianceItemsRequest(in).Send(ctx)
}

// RawListComplianceSummaries executes `ListComplianceSummaries` raw operation.
func (svc *SSM) RawListComplianceSummaries(ctx context.Context, in *SDK.ListComplianceSummariesInput) (*SDK.ListComplianceSummariesResponse, error) {
	return svc.client.ListComplianceSummariesRequest(in).Send(ctx)
}

// RawListDocumentVersions executes `ListDocumentVersions` raw operation.
func (svc *SSM) RawListDocumentVersions(ctx context.Context, in *SDK.ListDocumentVersionsInput) (*SDK.ListDocumentVersionsResponse, error) {
	return svc.client.ListDocumentVersionsRequest(in).Send(ctx)
}

// RawListDocuments executes `ListDocuments` raw operation.
func (svc *SSM) RawListDocuments(ctx context.Context, in *SDK.ListDocumentsInput) (*SDK.ListDocumentsResponse, error) {
	return svc.client.ListDocumentsRequest(in).Send(ctx)
}

// RawListInventoryEntries executes `ListInventoryEntries` raw operation.
func (svc *SSM) RawListInventoryEntries(ctx context.Context, in *SDK.ListInventoryEntriesInput) (*SDK.ListInventoryEntriesResponse, error) {
	return svc.client.ListInventoryEntriesRequest(in).Send(ctx)
}

// RawListResourceComplianceSummaries executes `ListResourceComplianceSummaries` raw operation.
func (svc *SSM) RawListResourceComplianceSummaries(ctx context.Context, in *SDK.ListResourceComplianceSummariesInput) (*SDK.ListResourceComplianceSummariesResponse, error) {
	return svc.client.ListResourceComplianceSummariesRequest(in).Send(ctx)
}

// RawListResourceDataSync executes `ListResourceDataSync` raw operation.
func (svc *SSM) RawListResourceDataSync(ctx context.Context, in *SDK.ListResourceDataSyncInput) (*SDK.ListResourceDataSyncResponse, error) {
	return svc.client.ListResourceDataSyncRequest(in).Send(ctx)
}

// RawListTagsForResource executes `ListTagsForResource` raw operation.
func (svc *SSM) RawListTagsForResource(ctx context.Context, in *SDK.ListTagsForResourceInput) (*SDK.ListTagsForResourceResponse, error) {
	return svc.client.ListTagsForResourceRequest(in).Send(ctx)
}

// RawModifyDocumentPermission executes `ModifyDocumentPermission` raw operation.
func (svc *SSM) RawModifyDocumentPermission(ctx context.Context, in *SDK.ModifyDocumentPermissionInput) (*SDK.ModifyDocumentPermissionResponse, error) {
	return svc.client.ModifyDocumentPermissionRequest(in).Send(ctx)
}

// RawPutComplianceItems executes `PutComplianceItems` raw operation.
func (svc *SSM) RawPutComplianceItems(ctx context.Context, in *SDK.PutComplianceItemsInput) (*SDK.PutComplianceItemsResponse, error) {
	return svc.client.PutComplianceItemsRequest(in).Send(ctx)
}

// RawPutInventory executes `PutInventory` raw operation.
func (svc *SSM) RawPutInventory(ctx context.Context, in *SDK.PutInventoryInput) (*SDK.PutInventoryResponse, error) {
	return svc.client.PutInventoryRequest(in).Send(ctx)
}

// RawPutParameter executes `PutParameter` raw operation.
func (svc *SSM) RawPutParameter(ctx context.Context, in *SDK.PutParameterInput) (*SDK.PutParameterResponse, error) {
	return svc.client.PutParameterRequest(in).Send(ctx)
}

// RawRegisterDefaultPatchBaseline executes `RegisterDefaultPatchBaseline` raw operation.
func (svc *SSM) RawRegisterDefaultPatchBaseline(ctx context.Context, in *SDK.RegisterDefaultPatchBaselineInput) (*SDK.RegisterDefaultPatchBaselineResponse, error) {
	return svc.client.RegisterDefaultPatchBaselineRequest(in).Send(ctx)
}

// RawRegisterPatchBaselineForPatchGroup executes `RegisterPatchBaselineForPatchGroup` raw operation.
func (svc *SSM) RawRegisterPatchBaselineForPatchGroup(ctx context.Context, in *SDK.RegisterPatchBaselineForPatchGroupInput) (*SDK.RegisterPatchBaselineForPatchGroupResponse, error) {
	return svc.client.RegisterPatchBaselineForPatchGroupRequest(in).Send(ctx)
}

// RawRegisterTargetWithMaintenanceWindow executes `RegisterTargetWithMaintenanceWindow` raw operation.
func (svc *SSM) RawRegisterTargetWithMaintenanceWindow(ctx context.Context, in *SDK.RegisterTargetWithMaintenanceWindowInput) (*SDK.RegisterTargetWithMaintenanceWindowResponse, error) {
	return svc.client.RegisterTargetWithMaintenanceWindowRequest(in).Send(ctx)
}

// RawRegisterTaskWithMaintenanceWindow executes `RegisterTaskWithMaintenanceWindow` raw operation.
func (svc *SSM) RawRegisterTaskWithMaintenanceWindow(ctx context.Context, in *SDK.RegisterTaskWithMaintenanceWindowInput) (*SDK.RegisterTaskWithMaintenanceWindowResponse, error) {
	return svc.client.RegisterTaskWithMaintenanceWindowRequest(in).Send(ctx)
}

// RawRemoveTagsFromResource executes `RemoveTagsFromResource` raw operation.
func (svc *SSM) RawRemoveTagsFromResource(ctx context.Context, in *SDK.RemoveTagsFromResourceInput) (*SDK.RemoveTagsFromResourceResponse, error) {
	return svc.client.RemoveTagsFromResourceRequest(in).Send(ctx)
}

// RawResetServiceSetting executes `ResetServiceSetting` raw operation.
func (svc *SSM) RawResetServiceSetting(ctx context.Context, in *SDK.ResetServiceSettingInput) (*SDK.ResetServiceSettingResponse, error) {
	return svc.client.ResetServiceSettingRequest(in).Send(ctx)
}

// RawResumeSession executes `ResumeSession` raw operation.
func (svc *SSM) RawResumeSession(ctx context.Context, in *SDK.ResumeSessionInput) (*SDK.ResumeSessionResponse, error) {
	return svc.client.ResumeSessionRequest(in).Send(ctx)
}

// RawSendAutomationSignal executes `SendAutomationSignal` raw operation.
func (svc *SSM) RawSendAutomationSignal(ctx context.Context, in *SDK.SendAutomationSignalInput) (*SDK.SendAutomationSignalResponse, error) {
	return svc.client.SendAutomationSignalRequest(in).Send(ctx)
}

// RawSendCommand executes `SendCommand` raw operation.
func (svc *SSM) RawSendCommand(ctx context.Context, in *SDK.SendCommandInput) (*SDK.SendCommandResponse, error) {
	return svc.client.SendCommandRequest(in).Send(ctx)
}

// RawStartAssociationsOnce executes `StartAssociationsOnce` raw operation.
func (svc *SSM) RawStartAssociationsOnce(ctx context.Context, in *SDK.StartAssociationsOnceInput) (*SDK.StartAssociationsOnceResponse, error) {
	return svc.client.StartAssociationsOnceRequest(in).Send(ctx)
}

// RawStartAutomationExecution executes `StartAutomationExecution` raw operation.
func (svc *SSM) RawStartAutomationExecution(ctx context.Context, in *SDK.StartAutomationExecutionInput) (*SDK.StartAutomationExecutionResponse, error) {
	return svc.client.StartAutomationExecutionRequest(in).Send(ctx)
}

// RawStartSession executes `StartSession` raw operation.
func (svc *SSM) RawStartSession(ctx context.Context, in *SDK.StartSessionInput) (*SDK.StartSessionResponse, error) {
	return svc.client.StartSessionRequest(in).Send(ctx)
}

// RawStopAutomationExecution executes `StopAutomationExecution` raw operation.
func (svc *SSM) RawStopAutomationExecution(ctx context.Context, in *SDK.StopAutomationExecutionInput) (*SDK.StopAutomationExecutionResponse, error) {
	return svc.client.StopAutomationExecutionRequest(in).Send(ctx)
}

// RawTerminateSession executes `TerminateSession` raw operation.
func (svc *SSM) RawTerminateSession(ctx context.Context, in *SDK.TerminateSessionInput) (*SDK.TerminateSessionResponse, error) {
	return svc.client.TerminateSessionRequest(in).Send(ctx)
}

// RawUpdateAssociation executes `UpdateAssociation` raw operation.
func (svc *SSM) RawUpdateAssociation(ctx context.Context, in *SDK.UpdateAssociationInput) (*SDK.UpdateAssociationResponse, error) {
	return svc.client.UpdateAssociationRequest(in).Send(ctx)
}

// RawUpdateAssociationStatus executes `UpdateAssociationStatus` raw operation.
func (svc *SSM) RawUpdateAssociationStatus(ctx context.Context, in *SDK.UpdateAssociationStatusInput) (*SDK.UpdateAssociationStatusResponse, error) {
	return svc.client.UpdateAssociationStatusRequest(in).Send(ctx)
}

// RawUpdateDocument executes `UpdateDocument` raw operation.
func (svc *SSM) RawUpdateDocument(ctx context.Context, in *SDK.UpdateDocumentInput) (*SDK.UpdateDocumentResponse, error) {
	return svc.client.UpdateDocumentRequest(in).Send(ctx)
}

// RawUpdateDocumentDefaultVersion executes `UpdateDocumentDefaultVersion` raw operation.
func (svc *SSM) RawUpdateDocumentDefaultVersion(ctx context.Context, in *SDK.UpdateDocumentDefaultVersionInput) (*SDK.UpdateDocumentDefaultVersionResponse, error) {
	return svc.client.UpdateDocumentDefaultVersionRequest(in).Send(ctx)
}

// RawUpdateMaintenanceWindow executes `UpdateMaintenanceWindow` raw operation.
func (svc *SSM) RawUpdateMaintenanceWindow(ctx context.Context, in *SDK.UpdateMaintenanceWindowInput) (*SDK.UpdateMaintenanceWindowResponse, error) {
	return svc.client.UpdateMaintenanceWindowRequest(in).Send(ctx)
}

// RawUpdateMaintenanceWindowTarget executes `UpdateMaintenanceWindowTarget` raw operation.
func (svc *SSM) RawUpdateMaintenanceWindowTarget(ctx context.Context, in *SDK.UpdateMaintenanceWindowTargetInput) (*SDK.UpdateMaintenanceWindowTargetResponse, error) {
	return svc.client.UpdateMaintenanceWindowTargetRequest(in).Send(ctx)
}

// RawUpdateMaintenanceWindowTask executes `UpdateMaintenanceWindowTask` raw operation.
func (svc *SSM) RawUpdateMaintenanceWindowTask(ctx context.Context, in *SDK.UpdateMaintenanceWindowTaskInput) (*SDK.UpdateMaintenanceWindowTaskResponse, error) {
	return svc.client.UpdateMaintenanceWindowTaskRequest(in).Send(ctx)
}

// RawUpdateManagedInstanceRole executes `UpdateManagedInstanceRole` raw operation.
func (svc *SSM) RawUpdateManagedInstanceRole(ctx context.Context, in *SDK.UpdateManagedInstanceRoleInput) (*SDK.UpdateManagedInstanceRoleResponse, error) {
	return svc.client.UpdateManagedInstanceRoleRequest(in).Send(ctx)
}

// RawUpdateOpsItem executes `UpdateOpsItem` raw operation.
func (svc *SSM) RawUpdateOpsItem(ctx context.Context, in *SDK.UpdateOpsItemInput) (*SDK.UpdateOpsItemResponse, error) {
	return svc.client.UpdateOpsItemRequest(in).Send(ctx)
}

// RawUpdatePatchBaseline executes `UpdatePatchBaseline` raw operation.
func (svc *SSM) RawUpdatePatchBaseline(ctx context.Context, in *SDK.UpdatePatchBaselineInput) (*SDK.UpdatePatchBaselineResponse, error) {
	return svc.client.UpdatePatchBaselineRequest(in).Send(ctx)
}

// RawUpdateResourceDataSync executes `UpdateResourceDataSync` raw operation.
func (svc *SSM) RawUpdateResourceDataSync(ctx context.Context, in *SDK.UpdateResourceDataSyncInput) (*SDK.UpdateResourceDataSyncResponse, error) {
	return svc.client.UpdateResourceDataSyncRequest(in).Send(ctx)
}

// RawUpdateServiceSetting executes `UpdateServiceSetting` raw operation.
func (svc *SSM) RawUpdateServiceSetting(ctx context.Context, in *SDK.UpdateServiceSettingInput) (*SDK.UpdateServiceSettingResponse, error) {
	return svc.client.UpdateServiceSettingRequest(in).Send(ctx)
}
