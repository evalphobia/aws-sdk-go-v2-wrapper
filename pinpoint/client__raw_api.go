package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

// RawCreateApp executes `CreateApp` raw operation.
func (svc *Pinpoint) RawCreateApp(ctx context.Context, in *SDK.CreateAppInput) (*SDK.CreateAppResponse, error) {
	return svc.client.CreateAppRequest(in).Send(ctx)
}

// RawCreateCampaign executes `CreateCampaign` raw operation.
func (svc *Pinpoint) RawCreateCampaign(ctx context.Context, in *SDK.CreateCampaignInput) (*SDK.CreateCampaignResponse, error) {
	return svc.client.CreateCampaignRequest(in).Send(ctx)
}

// RawCreateEmailTemplate executes `CreateEmailTemplate` raw operation.
func (svc *Pinpoint) RawCreateEmailTemplate(ctx context.Context, in *SDK.CreateEmailTemplateInput) (*SDK.CreateEmailTemplateResponse, error) {
	return svc.client.CreateEmailTemplateRequest(in).Send(ctx)
}

// RawCreateExportJob executes `CreateExportJob` raw operation.
func (svc *Pinpoint) RawCreateExportJob(ctx context.Context, in *SDK.CreateExportJobInput) (*SDK.CreateExportJobResponse, error) {
	return svc.client.CreateExportJobRequest(in).Send(ctx)
}

// RawCreateImportJob executes `CreateImportJob` raw operation.
func (svc *Pinpoint) RawCreateImportJob(ctx context.Context, in *SDK.CreateImportJobInput) (*SDK.CreateImportJobResponse, error) {
	return svc.client.CreateImportJobRequest(in).Send(ctx)
}

// RawCreateJourney executes `CreateJourney` raw operation.
func (svc *Pinpoint) RawCreateJourney(ctx context.Context, in *SDK.CreateJourneyInput) (*SDK.CreateJourneyResponse, error) {
	return svc.client.CreateJourneyRequest(in).Send(ctx)
}

// RawCreatePushTemplate executes `CreatePushTemplate` raw operation.
func (svc *Pinpoint) RawCreatePushTemplate(ctx context.Context, in *SDK.CreatePushTemplateInput) (*SDK.CreatePushTemplateResponse, error) {
	return svc.client.CreatePushTemplateRequest(in).Send(ctx)
}

// RawCreateRecommenderConfiguration executes `CreateRecommenderConfiguration` raw operation.
func (svc *Pinpoint) RawCreateRecommenderConfiguration(ctx context.Context, in *SDK.CreateRecommenderConfigurationInput) (*SDK.CreateRecommenderConfigurationResponse, error) {
	return svc.client.CreateRecommenderConfigurationRequest(in).Send(ctx)
}

// RawCreateSegment executes `CreateSegment` raw operation.
func (svc *Pinpoint) RawCreateSegment(ctx context.Context, in *SDK.CreateSegmentInput) (*SDK.CreateSegmentResponse, error) {
	return svc.client.CreateSegmentRequest(in).Send(ctx)
}

// RawCreateSmsTemplate executes `CreateSmsTemplate` raw operation.
func (svc *Pinpoint) RawCreateSmsTemplate(ctx context.Context, in *SDK.CreateSmsTemplateInput) (*SDK.CreateSmsTemplateResponse, error) {
	return svc.client.CreateSmsTemplateRequest(in).Send(ctx)
}

// RawCreateVoiceTemplate executes `CreateVoiceTemplate` raw operation.
func (svc *Pinpoint) RawCreateVoiceTemplate(ctx context.Context, in *SDK.CreateVoiceTemplateInput) (*SDK.CreateVoiceTemplateResponse, error) {
	return svc.client.CreateVoiceTemplateRequest(in).Send(ctx)
}

// RawDeleteAdmChannel executes `DeleteAdmChannel` raw operation.
func (svc *Pinpoint) RawDeleteAdmChannel(ctx context.Context, in *SDK.DeleteAdmChannelInput) (*SDK.DeleteAdmChannelResponse, error) {
	return svc.client.DeleteAdmChannelRequest(in).Send(ctx)
}

// RawDeleteApnsChannel executes `DeleteApnsChannel` raw operation.
func (svc *Pinpoint) RawDeleteApnsChannel(ctx context.Context, in *SDK.DeleteApnsChannelInput) (*SDK.DeleteApnsChannelResponse, error) {
	return svc.client.DeleteApnsChannelRequest(in).Send(ctx)
}

// RawDeleteApnsSandboxChannel executes `DeleteApnsSandboxChannel` raw operation.
func (svc *Pinpoint) RawDeleteApnsSandboxChannel(ctx context.Context, in *SDK.DeleteApnsSandboxChannelInput) (*SDK.DeleteApnsSandboxChannelResponse, error) {
	return svc.client.DeleteApnsSandboxChannelRequest(in).Send(ctx)
}

// RawDeleteApnsVoipChannel executes `DeleteApnsVoipChannel` raw operation.
func (svc *Pinpoint) RawDeleteApnsVoipChannel(ctx context.Context, in *SDK.DeleteApnsVoipChannelInput) (*SDK.DeleteApnsVoipChannelResponse, error) {
	return svc.client.DeleteApnsVoipChannelRequest(in).Send(ctx)
}

// RawDeleteApnsVoipSandboxChannel executes `DeleteApnsVoipSandboxChannel` raw operation.
func (svc *Pinpoint) RawDeleteApnsVoipSandboxChannel(ctx context.Context, in *SDK.DeleteApnsVoipSandboxChannelInput) (*SDK.DeleteApnsVoipSandboxChannelResponse, error) {
	return svc.client.DeleteApnsVoipSandboxChannelRequest(in).Send(ctx)
}

// RawDeleteApp executes `DeleteApp` raw operation.
func (svc *Pinpoint) RawDeleteApp(ctx context.Context, in *SDK.DeleteAppInput) (*SDK.DeleteAppResponse, error) {
	return svc.client.DeleteAppRequest(in).Send(ctx)
}

// RawDeleteBaiduChannel executes `DeleteBaiduChannel` raw operation.
func (svc *Pinpoint) RawDeleteBaiduChannel(ctx context.Context, in *SDK.DeleteBaiduChannelInput) (*SDK.DeleteBaiduChannelResponse, error) {
	return svc.client.DeleteBaiduChannelRequest(in).Send(ctx)
}

// RawDeleteCampaign executes `DeleteCampaign` raw operation.
func (svc *Pinpoint) RawDeleteCampaign(ctx context.Context, in *SDK.DeleteCampaignInput) (*SDK.DeleteCampaignResponse, error) {
	return svc.client.DeleteCampaignRequest(in).Send(ctx)
}

// RawDeleteEmailChannel executes `DeleteEmailChannel` raw operation.
func (svc *Pinpoint) RawDeleteEmailChannel(ctx context.Context, in *SDK.DeleteEmailChannelInput) (*SDK.DeleteEmailChannelResponse, error) {
	return svc.client.DeleteEmailChannelRequest(in).Send(ctx)
}

// RawDeleteEmailTemplate executes `DeleteEmailTemplate` raw operation.
func (svc *Pinpoint) RawDeleteEmailTemplate(ctx context.Context, in *SDK.DeleteEmailTemplateInput) (*SDK.DeleteEmailTemplateResponse, error) {
	return svc.client.DeleteEmailTemplateRequest(in).Send(ctx)
}

// RawDeleteEndpoint executes `DeleteEndpoint` raw operation.
func (svc *Pinpoint) RawDeleteEndpoint(ctx context.Context, in *SDK.DeleteEndpointInput) (*SDK.DeleteEndpointResponse, error) {
	return svc.client.DeleteEndpointRequest(in).Send(ctx)
}

// RawDeleteEventStream executes `DeleteEventStream` raw operation.
func (svc *Pinpoint) RawDeleteEventStream(ctx context.Context, in *SDK.DeleteEventStreamInput) (*SDK.DeleteEventStreamResponse, error) {
	return svc.client.DeleteEventStreamRequest(in).Send(ctx)
}

// RawDeleteGcmChannel executes `DeleteGcmChannel` raw operation.
func (svc *Pinpoint) RawDeleteGcmChannel(ctx context.Context, in *SDK.DeleteGcmChannelInput) (*SDK.DeleteGcmChannelResponse, error) {
	return svc.client.DeleteGcmChannelRequest(in).Send(ctx)
}

// RawDeleteJourney executes `DeleteJourney` raw operation.
func (svc *Pinpoint) RawDeleteJourney(ctx context.Context, in *SDK.DeleteJourneyInput) (*SDK.DeleteJourneyResponse, error) {
	return svc.client.DeleteJourneyRequest(in).Send(ctx)
}

// RawDeletePushTemplate executes `DeletePushTemplate` raw operation.
func (svc *Pinpoint) RawDeletePushTemplate(ctx context.Context, in *SDK.DeletePushTemplateInput) (*SDK.DeletePushTemplateResponse, error) {
	return svc.client.DeletePushTemplateRequest(in).Send(ctx)
}

// RawDeleteRecommenderConfiguration executes `DeleteRecommenderConfiguration` raw operation.
func (svc *Pinpoint) RawDeleteRecommenderConfiguration(ctx context.Context, in *SDK.DeleteRecommenderConfigurationInput) (*SDK.DeleteRecommenderConfigurationResponse, error) {
	return svc.client.DeleteRecommenderConfigurationRequest(in).Send(ctx)
}

// RawDeleteSegment executes `DeleteSegment` raw operation.
func (svc *Pinpoint) RawDeleteSegment(ctx context.Context, in *SDK.DeleteSegmentInput) (*SDK.DeleteSegmentResponse, error) {
	return svc.client.DeleteSegmentRequest(in).Send(ctx)
}

// RawDeleteSmsChannel executes `DeleteSmsChannel` raw operation.
func (svc *Pinpoint) RawDeleteSmsChannel(ctx context.Context, in *SDK.DeleteSmsChannelInput) (*SDK.DeleteSmsChannelResponse, error) {
	return svc.client.DeleteSmsChannelRequest(in).Send(ctx)
}

// RawDeleteSmsTemplate executes `DeleteSmsTemplate` raw operation.
func (svc *Pinpoint) RawDeleteSmsTemplate(ctx context.Context, in *SDK.DeleteSmsTemplateInput) (*SDK.DeleteSmsTemplateResponse, error) {
	return svc.client.DeleteSmsTemplateRequest(in).Send(ctx)
}

// RawDeleteUserEndpoints executes `DeleteUserEndpoints` raw operation.
func (svc *Pinpoint) RawDeleteUserEndpoints(ctx context.Context, in *SDK.DeleteUserEndpointsInput) (*SDK.DeleteUserEndpointsResponse, error) {
	return svc.client.DeleteUserEndpointsRequest(in).Send(ctx)
}

// RawDeleteVoiceChannel executes `DeleteVoiceChannel` raw operation.
func (svc *Pinpoint) RawDeleteVoiceChannel(ctx context.Context, in *SDK.DeleteVoiceChannelInput) (*SDK.DeleteVoiceChannelResponse, error) {
	return svc.client.DeleteVoiceChannelRequest(in).Send(ctx)
}

// RawDeleteVoiceTemplate executes `DeleteVoiceTemplate` raw operation.
func (svc *Pinpoint) RawDeleteVoiceTemplate(ctx context.Context, in *SDK.DeleteVoiceTemplateInput) (*SDK.DeleteVoiceTemplateResponse, error) {
	return svc.client.DeleteVoiceTemplateRequest(in).Send(ctx)
}

// RawGetAdmChannel executes `GetAdmChannel` raw operation.
func (svc *Pinpoint) RawGetAdmChannel(ctx context.Context, in *SDK.GetAdmChannelInput) (*SDK.GetAdmChannelResponse, error) {
	return svc.client.GetAdmChannelRequest(in).Send(ctx)
}

// RawGetApnsChannel executes `GetApnsChannel` raw operation.
func (svc *Pinpoint) RawGetApnsChannel(ctx context.Context, in *SDK.GetApnsChannelInput) (*SDK.GetApnsChannelResponse, error) {
	return svc.client.GetApnsChannelRequest(in).Send(ctx)
}

// RawGetApnsSandboxChannel executes `GetApnsSandboxChannel` raw operation.
func (svc *Pinpoint) RawGetApnsSandboxChannel(ctx context.Context, in *SDK.GetApnsSandboxChannelInput) (*SDK.GetApnsSandboxChannelResponse, error) {
	return svc.client.GetApnsSandboxChannelRequest(in).Send(ctx)
}

// RawGetApnsVoipChannel executes `GetApnsVoipChannel` raw operation.
func (svc *Pinpoint) RawGetApnsVoipChannel(ctx context.Context, in *SDK.GetApnsVoipChannelInput) (*SDK.GetApnsVoipChannelResponse, error) {
	return svc.client.GetApnsVoipChannelRequest(in).Send(ctx)
}

// RawGetApnsVoipSandboxChannel executes `GetApnsVoipSandboxChannel` raw operation.
func (svc *Pinpoint) RawGetApnsVoipSandboxChannel(ctx context.Context, in *SDK.GetApnsVoipSandboxChannelInput) (*SDK.GetApnsVoipSandboxChannelResponse, error) {
	return svc.client.GetApnsVoipSandboxChannelRequest(in).Send(ctx)
}

// RawGetApp executes `GetApp` raw operation.
func (svc *Pinpoint) RawGetApp(ctx context.Context, in *SDK.GetAppInput) (*SDK.GetAppResponse, error) {
	return svc.client.GetAppRequest(in).Send(ctx)
}

// RawGetApplicationDateRangeKpi executes `GetApplicationDateRangeKpi` raw operation.
func (svc *Pinpoint) RawGetApplicationDateRangeKpi(ctx context.Context, in *SDK.GetApplicationDateRangeKpiInput) (*SDK.GetApplicationDateRangeKpiResponse, error) {
	return svc.client.GetApplicationDateRangeKpiRequest(in).Send(ctx)
}

// RawGetApplicationSettings executes `GetApplicationSettings` raw operation.
func (svc *Pinpoint) RawGetApplicationSettings(ctx context.Context, in *SDK.GetApplicationSettingsInput) (*SDK.GetApplicationSettingsResponse, error) {
	return svc.client.GetApplicationSettingsRequest(in).Send(ctx)
}

// RawGetApps executes `GetApps` raw operation.
func (svc *Pinpoint) RawGetApps(ctx context.Context, in *SDK.GetAppsInput) (*SDK.GetAppsResponse, error) {
	return svc.client.GetAppsRequest(in).Send(ctx)
}

// RawGetBaiduChannel executes `GetBaiduChannel` raw operation.
func (svc *Pinpoint) RawGetBaiduChannel(ctx context.Context, in *SDK.GetBaiduChannelInput) (*SDK.GetBaiduChannelResponse, error) {
	return svc.client.GetBaiduChannelRequest(in).Send(ctx)
}

// RawGetCampaign executes `GetCampaign` raw operation.
func (svc *Pinpoint) RawGetCampaign(ctx context.Context, in *SDK.GetCampaignInput) (*SDK.GetCampaignResponse, error) {
	return svc.client.GetCampaignRequest(in).Send(ctx)
}

// RawGetCampaignActivities executes `GetCampaignActivities` raw operation.
func (svc *Pinpoint) RawGetCampaignActivities(ctx context.Context, in *SDK.GetCampaignActivitiesInput) (*SDK.GetCampaignActivitiesResponse, error) {
	return svc.client.GetCampaignActivitiesRequest(in).Send(ctx)
}

// RawGetCampaignDateRangeKpi executes `GetCampaignDateRangeKpi` raw operation.
func (svc *Pinpoint) RawGetCampaignDateRangeKpi(ctx context.Context, in *SDK.GetCampaignDateRangeKpiInput) (*SDK.GetCampaignDateRangeKpiResponse, error) {
	return svc.client.GetCampaignDateRangeKpiRequest(in).Send(ctx)
}

// RawGetCampaignVersion executes `GetCampaignVersion` raw operation.
func (svc *Pinpoint) RawGetCampaignVersion(ctx context.Context, in *SDK.GetCampaignVersionInput) (*SDK.GetCampaignVersionResponse, error) {
	return svc.client.GetCampaignVersionRequest(in).Send(ctx)
}

// RawGetCampaignVersions executes `GetCampaignVersions` raw operation.
func (svc *Pinpoint) RawGetCampaignVersions(ctx context.Context, in *SDK.GetCampaignVersionsInput) (*SDK.GetCampaignVersionsResponse, error) {
	return svc.client.GetCampaignVersionsRequest(in).Send(ctx)
}

// RawGetCampaigns executes `GetCampaigns` raw operation.
func (svc *Pinpoint) RawGetCampaigns(ctx context.Context, in *SDK.GetCampaignsInput) (*SDK.GetCampaignsResponse, error) {
	return svc.client.GetCampaignsRequest(in).Send(ctx)
}

// RawGetChannels executes `GetChannels` raw operation.
func (svc *Pinpoint) RawGetChannels(ctx context.Context, in *SDK.GetChannelsInput) (*SDK.GetChannelsResponse, error) {
	return svc.client.GetChannelsRequest(in).Send(ctx)
}

// RawGetEmailChannel executes `GetEmailChannel` raw operation.
func (svc *Pinpoint) RawGetEmailChannel(ctx context.Context, in *SDK.GetEmailChannelInput) (*SDK.GetEmailChannelResponse, error) {
	return svc.client.GetEmailChannelRequest(in).Send(ctx)
}

// RawGetEmailTemplate executes `GetEmailTemplate` raw operation.
func (svc *Pinpoint) RawGetEmailTemplate(ctx context.Context, in *SDK.GetEmailTemplateInput) (*SDK.GetEmailTemplateResponse, error) {
	return svc.client.GetEmailTemplateRequest(in).Send(ctx)
}

// RawGetEndpoint executes `GetEndpoint` raw operation.
func (svc *Pinpoint) RawGetEndpoint(ctx context.Context, in *SDK.GetEndpointInput) (*SDK.GetEndpointResponse, error) {
	return svc.client.GetEndpointRequest(in).Send(ctx)
}

// RawGetEventStream executes `GetEventStream` raw operation.
func (svc *Pinpoint) RawGetEventStream(ctx context.Context, in *SDK.GetEventStreamInput) (*SDK.GetEventStreamResponse, error) {
	return svc.client.GetEventStreamRequest(in).Send(ctx)
}

// RawGetExportJob executes `GetExportJob` raw operation.
func (svc *Pinpoint) RawGetExportJob(ctx context.Context, in *SDK.GetExportJobInput) (*SDK.GetExportJobResponse, error) {
	return svc.client.GetExportJobRequest(in).Send(ctx)
}

// RawGetExportJobs executes `GetExportJobs` raw operation.
func (svc *Pinpoint) RawGetExportJobs(ctx context.Context, in *SDK.GetExportJobsInput) (*SDK.GetExportJobsResponse, error) {
	return svc.client.GetExportJobsRequest(in).Send(ctx)
}

// RawGetGcmChannel executes `GetGcmChannel` raw operation.
func (svc *Pinpoint) RawGetGcmChannel(ctx context.Context, in *SDK.GetGcmChannelInput) (*SDK.GetGcmChannelResponse, error) {
	return svc.client.GetGcmChannelRequest(in).Send(ctx)
}

// RawGetImportJob executes `GetImportJob` raw operation.
func (svc *Pinpoint) RawGetImportJob(ctx context.Context, in *SDK.GetImportJobInput) (*SDK.GetImportJobResponse, error) {
	return svc.client.GetImportJobRequest(in).Send(ctx)
}

// RawGetImportJobs executes `GetImportJobs` raw operation.
func (svc *Pinpoint) RawGetImportJobs(ctx context.Context, in *SDK.GetImportJobsInput) (*SDK.GetImportJobsResponse, error) {
	return svc.client.GetImportJobsRequest(in).Send(ctx)
}

// RawGetJourney executes `GetJourney` raw operation.
func (svc *Pinpoint) RawGetJourney(ctx context.Context, in *SDK.GetJourneyInput) (*SDK.GetJourneyResponse, error) {
	return svc.client.GetJourneyRequest(in).Send(ctx)
}

// RawGetJourneyDateRangeKpi executes `GetJourneyDateRangeKpi` raw operation.
func (svc *Pinpoint) RawGetJourneyDateRangeKpi(ctx context.Context, in *SDK.GetJourneyDateRangeKpiInput) (*SDK.GetJourneyDateRangeKpiResponse, error) {
	return svc.client.GetJourneyDateRangeKpiRequest(in).Send(ctx)
}

// RawGetJourneyExecutionActivityMetrics executes `GetJourneyExecutionActivityMetrics` raw operation.
func (svc *Pinpoint) RawGetJourneyExecutionActivityMetrics(ctx context.Context, in *SDK.GetJourneyExecutionActivityMetricsInput) (*SDK.GetJourneyExecutionActivityMetricsResponse, error) {
	return svc.client.GetJourneyExecutionActivityMetricsRequest(in).Send(ctx)
}

// RawGetJourneyExecutionMetrics executes `GetJourneyExecutionMetrics` raw operation.
func (svc *Pinpoint) RawGetJourneyExecutionMetrics(ctx context.Context, in *SDK.GetJourneyExecutionMetricsInput) (*SDK.GetJourneyExecutionMetricsResponse, error) {
	return svc.client.GetJourneyExecutionMetricsRequest(in).Send(ctx)
}

// RawGetPushTemplate executes `GetPushTemplate` raw operation.
func (svc *Pinpoint) RawGetPushTemplate(ctx context.Context, in *SDK.GetPushTemplateInput) (*SDK.GetPushTemplateResponse, error) {
	return svc.client.GetPushTemplateRequest(in).Send(ctx)
}

// RawGetRecommenderConfiguration executes `GetRecommenderConfiguration` raw operation.
func (svc *Pinpoint) RawGetRecommenderConfiguration(ctx context.Context, in *SDK.GetRecommenderConfigurationInput) (*SDK.GetRecommenderConfigurationResponse, error) {
	return svc.client.GetRecommenderConfigurationRequest(in).Send(ctx)
}

// RawGetRecommenderConfigurations executes `GetRecommenderConfigurations` raw operation.
func (svc *Pinpoint) RawGetRecommenderConfigurations(ctx context.Context, in *SDK.GetRecommenderConfigurationsInput) (*SDK.GetRecommenderConfigurationsResponse, error) {
	return svc.client.GetRecommenderConfigurationsRequest(in).Send(ctx)
}

// RawGetSegment executes `GetSegment` raw operation.
func (svc *Pinpoint) RawGetSegment(ctx context.Context, in *SDK.GetSegmentInput) (*SDK.GetSegmentResponse, error) {
	return svc.client.GetSegmentRequest(in).Send(ctx)
}

// RawGetSegmentExportJobs executes `GetSegmentExportJobs` raw operation.
func (svc *Pinpoint) RawGetSegmentExportJobs(ctx context.Context, in *SDK.GetSegmentExportJobsInput) (*SDK.GetSegmentExportJobsResponse, error) {
	return svc.client.GetSegmentExportJobsRequest(in).Send(ctx)
}

// RawGetSegmentImportJobs executes `GetSegmentImportJobs` raw operation.
func (svc *Pinpoint) RawGetSegmentImportJobs(ctx context.Context, in *SDK.GetSegmentImportJobsInput) (*SDK.GetSegmentImportJobsResponse, error) {
	return svc.client.GetSegmentImportJobsRequest(in).Send(ctx)
}

// RawGetSegmentVersion executes `GetSegmentVersion` raw operation.
func (svc *Pinpoint) RawGetSegmentVersion(ctx context.Context, in *SDK.GetSegmentVersionInput) (*SDK.GetSegmentVersionResponse, error) {
	return svc.client.GetSegmentVersionRequest(in).Send(ctx)
}

// RawGetSegmentVersions executes `GetSegmentVersions` raw operation.
func (svc *Pinpoint) RawGetSegmentVersions(ctx context.Context, in *SDK.GetSegmentVersionsInput) (*SDK.GetSegmentVersionsResponse, error) {
	return svc.client.GetSegmentVersionsRequest(in).Send(ctx)
}

// RawGetSegments executes `GetSegments` raw operation.
func (svc *Pinpoint) RawGetSegments(ctx context.Context, in *SDK.GetSegmentsInput) (*SDK.GetSegmentsResponse, error) {
	return svc.client.GetSegmentsRequest(in).Send(ctx)
}

// RawGetSmsChannel executes `GetSmsChannel` raw operation.
func (svc *Pinpoint) RawGetSmsChannel(ctx context.Context, in *SDK.GetSmsChannelInput) (*SDK.GetSmsChannelResponse, error) {
	return svc.client.GetSmsChannelRequest(in).Send(ctx)
}

// RawGetSmsTemplate executes `GetSmsTemplate` raw operation.
func (svc *Pinpoint) RawGetSmsTemplate(ctx context.Context, in *SDK.GetSmsTemplateInput) (*SDK.GetSmsTemplateResponse, error) {
	return svc.client.GetSmsTemplateRequest(in).Send(ctx)
}

// RawGetUserEndpoints executes `GetUserEndpoints` raw operation.
func (svc *Pinpoint) RawGetUserEndpoints(ctx context.Context, in *SDK.GetUserEndpointsInput) (*SDK.GetUserEndpointsResponse, error) {
	return svc.client.GetUserEndpointsRequest(in).Send(ctx)
}

// RawGetVoiceChannel executes `GetVoiceChannel` raw operation.
func (svc *Pinpoint) RawGetVoiceChannel(ctx context.Context, in *SDK.GetVoiceChannelInput) (*SDK.GetVoiceChannelResponse, error) {
	return svc.client.GetVoiceChannelRequest(in).Send(ctx)
}

// RawGetVoiceTemplate executes `GetVoiceTemplate` raw operation.
func (svc *Pinpoint) RawGetVoiceTemplate(ctx context.Context, in *SDK.GetVoiceTemplateInput) (*SDK.GetVoiceTemplateResponse, error) {
	return svc.client.GetVoiceTemplateRequest(in).Send(ctx)
}

// RawListJourneys executes `ListJourneys` raw operation.
func (svc *Pinpoint) RawListJourneys(ctx context.Context, in *SDK.ListJourneysInput) (*SDK.ListJourneysResponse, error) {
	return svc.client.ListJourneysRequest(in).Send(ctx)
}

// RawListTagsForResource executes `ListTagsForResource` raw operation.
func (svc *Pinpoint) RawListTagsForResource(ctx context.Context, in *SDK.ListTagsForResourceInput) (*SDK.ListTagsForResourceResponse, error) {
	return svc.client.ListTagsForResourceRequest(in).Send(ctx)
}

// RawListTemplateVersions executes `ListTemplateVersions` raw operation.
func (svc *Pinpoint) RawListTemplateVersions(ctx context.Context, in *SDK.ListTemplateVersionsInput) (*SDK.ListTemplateVersionsResponse, error) {
	return svc.client.ListTemplateVersionsRequest(in).Send(ctx)
}

// RawListTemplates executes `ListTemplates` raw operation.
func (svc *Pinpoint) RawListTemplates(ctx context.Context, in *SDK.ListTemplatesInput) (*SDK.ListTemplatesResponse, error) {
	return svc.client.ListTemplatesRequest(in).Send(ctx)
}

// RawPhoneNumberValidate executes `PhoneNumberValidate` raw operation.
func (svc *Pinpoint) RawPhoneNumberValidate(ctx context.Context, in *SDK.PhoneNumberValidateInput) (*SDK.PhoneNumberValidateResponse, error) {
	return svc.client.PhoneNumberValidateRequest(in).Send(ctx)
}

// RawPutEventStream executes `PutEventStream` raw operation.
func (svc *Pinpoint) RawPutEventStream(ctx context.Context, in *SDK.PutEventStreamInput) (*SDK.PutEventStreamResponse, error) {
	return svc.client.PutEventStreamRequest(in).Send(ctx)
}

// RawPutEvents executes `PutEvents` raw operation.
func (svc *Pinpoint) RawPutEvents(ctx context.Context, in *SDK.PutEventsInput) (*SDK.PutEventsResponse, error) {
	return svc.client.PutEventsRequest(in).Send(ctx)
}

// RawRemoveAttributes executes `RemoveAttributes` raw operation.
func (svc *Pinpoint) RawRemoveAttributes(ctx context.Context, in *SDK.RemoveAttributesInput) (*SDK.RemoveAttributesResponse, error) {
	return svc.client.RemoveAttributesRequest(in).Send(ctx)
}

// RawSendMessages executes `SendMessages` raw operation.
func (svc *Pinpoint) RawSendMessages(ctx context.Context, in *SDK.SendMessagesInput) (*SDK.SendMessagesResponse, error) {
	return svc.client.SendMessagesRequest(in).Send(ctx)
}

// RawSendUsersMessages executes `SendUsersMessages` raw operation.
func (svc *Pinpoint) RawSendUsersMessages(ctx context.Context, in *SDK.SendUsersMessagesInput) (*SDK.SendUsersMessagesResponse, error) {
	return svc.client.SendUsersMessagesRequest(in).Send(ctx)
}

// RawTagResource executes `TagResource` raw operation.
func (svc *Pinpoint) RawTagResource(ctx context.Context, in *SDK.TagResourceInput) (*SDK.TagResourceResponse, error) {
	return svc.client.TagResourceRequest(in).Send(ctx)
}

// RawUntagResource executes `UntagResource` raw operation.
func (svc *Pinpoint) RawUntagResource(ctx context.Context, in *SDK.UntagResourceInput) (*SDK.UntagResourceResponse, error) {
	return svc.client.UntagResourceRequest(in).Send(ctx)
}

// RawUpdateAdmChannel executes `UpdateAdmChannel` raw operation.
func (svc *Pinpoint) RawUpdateAdmChannel(ctx context.Context, in *SDK.UpdateAdmChannelInput) (*SDK.UpdateAdmChannelResponse, error) {
	return svc.client.UpdateAdmChannelRequest(in).Send(ctx)
}

// RawUpdateApnsChannel executes `UpdateApnsChannel` raw operation.
func (svc *Pinpoint) RawUpdateApnsChannel(ctx context.Context, in *SDK.UpdateApnsChannelInput) (*SDK.UpdateApnsChannelResponse, error) {
	return svc.client.UpdateApnsChannelRequest(in).Send(ctx)
}

// RawUpdateApnsSandboxChannel executes `UpdateApnsSandboxChannel` raw operation.
func (svc *Pinpoint) RawUpdateApnsSandboxChannel(ctx context.Context, in *SDK.UpdateApnsSandboxChannelInput) (*SDK.UpdateApnsSandboxChannelResponse, error) {
	return svc.client.UpdateApnsSandboxChannelRequest(in).Send(ctx)
}

// RawUpdateApnsVoipChannel executes `UpdateApnsVoipChannel` raw operation.
func (svc *Pinpoint) RawUpdateApnsVoipChannel(ctx context.Context, in *SDK.UpdateApnsVoipChannelInput) (*SDK.UpdateApnsVoipChannelResponse, error) {
	return svc.client.UpdateApnsVoipChannelRequest(in).Send(ctx)
}

// RawUpdateApnsVoipSandboxChannel executes `UpdateApnsVoipSandboxChannel` raw operation.
func (svc *Pinpoint) RawUpdateApnsVoipSandboxChannel(ctx context.Context, in *SDK.UpdateApnsVoipSandboxChannelInput) (*SDK.UpdateApnsVoipSandboxChannelResponse, error) {
	return svc.client.UpdateApnsVoipSandboxChannelRequest(in).Send(ctx)
}

// RawUpdateApplicationSettings executes `UpdateApplicationSettings` raw operation.
func (svc *Pinpoint) RawUpdateApplicationSettings(ctx context.Context, in *SDK.UpdateApplicationSettingsInput) (*SDK.UpdateApplicationSettingsResponse, error) {
	return svc.client.UpdateApplicationSettingsRequest(in).Send(ctx)
}

// RawUpdateBaiduChannel executes `UpdateBaiduChannel` raw operation.
func (svc *Pinpoint) RawUpdateBaiduChannel(ctx context.Context, in *SDK.UpdateBaiduChannelInput) (*SDK.UpdateBaiduChannelResponse, error) {
	return svc.client.UpdateBaiduChannelRequest(in).Send(ctx)
}

// RawUpdateCampaign executes `UpdateCampaign` raw operation.
func (svc *Pinpoint) RawUpdateCampaign(ctx context.Context, in *SDK.UpdateCampaignInput) (*SDK.UpdateCampaignResponse, error) {
	return svc.client.UpdateCampaignRequest(in).Send(ctx)
}

// RawUpdateEmailChannel executes `UpdateEmailChannel` raw operation.
func (svc *Pinpoint) RawUpdateEmailChannel(ctx context.Context, in *SDK.UpdateEmailChannelInput) (*SDK.UpdateEmailChannelResponse, error) {
	return svc.client.UpdateEmailChannelRequest(in).Send(ctx)
}

// RawUpdateEmailTemplate executes `UpdateEmailTemplate` raw operation.
func (svc *Pinpoint) RawUpdateEmailTemplate(ctx context.Context, in *SDK.UpdateEmailTemplateInput) (*SDK.UpdateEmailTemplateResponse, error) {
	return svc.client.UpdateEmailTemplateRequest(in).Send(ctx)
}

// RawUpdateEndpoint executes `UpdateEndpoint` raw operation.
func (svc *Pinpoint) RawUpdateEndpoint(ctx context.Context, in *SDK.UpdateEndpointInput) (*SDK.UpdateEndpointResponse, error) {
	return svc.client.UpdateEndpointRequest(in).Send(ctx)
}

// RawUpdateEndpointsBatch executes `UpdateEndpointsBatch` raw operation.
func (svc *Pinpoint) RawUpdateEndpointsBatch(ctx context.Context, in *SDK.UpdateEndpointsBatchInput) (*SDK.UpdateEndpointsBatchResponse, error) {
	return svc.client.UpdateEndpointsBatchRequest(in).Send(ctx)
}

// RawUpdateGcmChannel executes `UpdateGcmChannel` raw operation.
func (svc *Pinpoint) RawUpdateGcmChannel(ctx context.Context, in *SDK.UpdateGcmChannelInput) (*SDK.UpdateGcmChannelResponse, error) {
	return svc.client.UpdateGcmChannelRequest(in).Send(ctx)
}

// RawUpdateJourney executes `UpdateJourney` raw operation.
func (svc *Pinpoint) RawUpdateJourney(ctx context.Context, in *SDK.UpdateJourneyInput) (*SDK.UpdateJourneyResponse, error) {
	return svc.client.UpdateJourneyRequest(in).Send(ctx)
}

// RawUpdateJourneyState executes `UpdateJourneyState` raw operation.
func (svc *Pinpoint) RawUpdateJourneyState(ctx context.Context, in *SDK.UpdateJourneyStateInput) (*SDK.UpdateJourneyStateResponse, error) {
	return svc.client.UpdateJourneyStateRequest(in).Send(ctx)
}

// RawUpdatePushTemplate executes `UpdatePushTemplate` raw operation.
func (svc *Pinpoint) RawUpdatePushTemplate(ctx context.Context, in *SDK.UpdatePushTemplateInput) (*SDK.UpdatePushTemplateResponse, error) {
	return svc.client.UpdatePushTemplateRequest(in).Send(ctx)
}

// RawUpdateRecommenderConfiguration executes `UpdateRecommenderConfiguration` raw operation.
func (svc *Pinpoint) RawUpdateRecommenderConfiguration(ctx context.Context, in *SDK.UpdateRecommenderConfigurationInput) (*SDK.UpdateRecommenderConfigurationResponse, error) {
	return svc.client.UpdateRecommenderConfigurationRequest(in).Send(ctx)
}

// RawUpdateSegment executes `UpdateSegment` raw operation.
func (svc *Pinpoint) RawUpdateSegment(ctx context.Context, in *SDK.UpdateSegmentInput) (*SDK.UpdateSegmentResponse, error) {
	return svc.client.UpdateSegmentRequest(in).Send(ctx)
}

// RawUpdateSmsChannel executes `UpdateSmsChannel` raw operation.
func (svc *Pinpoint) RawUpdateSmsChannel(ctx context.Context, in *SDK.UpdateSmsChannelInput) (*SDK.UpdateSmsChannelResponse, error) {
	return svc.client.UpdateSmsChannelRequest(in).Send(ctx)
}

// RawUpdateSmsTemplate executes `UpdateSmsTemplate` raw operation.
func (svc *Pinpoint) RawUpdateSmsTemplate(ctx context.Context, in *SDK.UpdateSmsTemplateInput) (*SDK.UpdateSmsTemplateResponse, error) {
	return svc.client.UpdateSmsTemplateRequest(in).Send(ctx)
}

// RawUpdateTemplateActiveVersion executes `UpdateTemplateActiveVersion` raw operation.
func (svc *Pinpoint) RawUpdateTemplateActiveVersion(ctx context.Context, in *SDK.UpdateTemplateActiveVersionInput) (*SDK.UpdateTemplateActiveVersionResponse, error) {
	return svc.client.UpdateTemplateActiveVersionRequest(in).Send(ctx)
}

// RawUpdateVoiceChannel executes `UpdateVoiceChannel` raw operation.
func (svc *Pinpoint) RawUpdateVoiceChannel(ctx context.Context, in *SDK.UpdateVoiceChannelInput) (*SDK.UpdateVoiceChannelResponse, error) {
	return svc.client.UpdateVoiceChannelRequest(in).Send(ctx)
}

// RawUpdateVoiceTemplate executes `UpdateVoiceTemplate` raw operation.
func (svc *Pinpoint) RawUpdateVoiceTemplate(ctx context.Context, in *SDK.UpdateVoiceTemplateInput) (*SDK.UpdateVoiceTemplateResponse, error) {
	return svc.client.UpdateVoiceTemplateRequest(in).Send(ctx)
}
