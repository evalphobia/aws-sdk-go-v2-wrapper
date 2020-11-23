package pinpointemail

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpointemail"
)

// RawCreateConfigurationSet executes `CreateConfigurationSet` raw operation.
func (svc *PinpointEmail) RawCreateConfigurationSet(ctx context.Context, in *SDK.CreateConfigurationSetInput) (*SDK.CreateConfigurationSetResponse, error) {
	return svc.client.CreateConfigurationSetRequest(in).Send(ctx)
}

// RawCreateConfigurationSetEventDestination executes `CreateConfigurationSetEventDestination` raw operation.
func (svc *PinpointEmail) RawCreateConfigurationSetEventDestination(ctx context.Context, in *SDK.CreateConfigurationSetEventDestinationInput) (*SDK.CreateConfigurationSetEventDestinationResponse, error) {
	return svc.client.CreateConfigurationSetEventDestinationRequest(in).Send(ctx)
}

// RawCreateDedicatedIpPool executes `CreateDedicatedIpPool` raw operation.
func (svc *PinpointEmail) RawCreateDedicatedIpPool(ctx context.Context, in *SDK.CreateDedicatedIpPoolInput) (*SDK.CreateDedicatedIpPoolResponse, error) {
	return svc.client.CreateDedicatedIpPoolRequest(in).Send(ctx)
}

// RawCreateDeliverabilityTestReport executes `CreateDeliverabilityTestReport` raw operation.
func (svc *PinpointEmail) RawCreateDeliverabilityTestReport(ctx context.Context, in *SDK.CreateDeliverabilityTestReportInput) (*SDK.CreateDeliverabilityTestReportResponse, error) {
	return svc.client.CreateDeliverabilityTestReportRequest(in).Send(ctx)
}

// RawCreateEmailIdentity executes `CreateEmailIdentity` raw operation.
func (svc *PinpointEmail) RawCreateEmailIdentity(ctx context.Context, in *SDK.CreateEmailIdentityInput) (*SDK.CreateEmailIdentityResponse, error) {
	return svc.client.CreateEmailIdentityRequest(in).Send(ctx)
}

// RawDeleteConfigurationSet executes `DeleteConfigurationSet` raw operation.
func (svc *PinpointEmail) RawDeleteConfigurationSet(ctx context.Context, in *SDK.DeleteConfigurationSetInput) (*SDK.DeleteConfigurationSetResponse, error) {
	return svc.client.DeleteConfigurationSetRequest(in).Send(ctx)
}

// RawDeleteConfigurationSetEventDestination executes `DeleteConfigurationSetEventDestination` raw operation.
func (svc *PinpointEmail) RawDeleteConfigurationSetEventDestination(ctx context.Context, in *SDK.DeleteConfigurationSetEventDestinationInput) (*SDK.DeleteConfigurationSetEventDestinationResponse, error) {
	return svc.client.DeleteConfigurationSetEventDestinationRequest(in).Send(ctx)
}

// RawDeleteDedicatedIpPool executes `DeleteDedicatedIpPool` raw operation.
func (svc *PinpointEmail) RawDeleteDedicatedIpPool(ctx context.Context, in *SDK.DeleteDedicatedIpPoolInput) (*SDK.DeleteDedicatedIpPoolResponse, error) {
	return svc.client.DeleteDedicatedIpPoolRequest(in).Send(ctx)
}

// RawDeleteEmailIdentity executes `DeleteEmailIdentity` raw operation.
func (svc *PinpointEmail) RawDeleteEmailIdentity(ctx context.Context, in *SDK.DeleteEmailIdentityInput) (*SDK.DeleteEmailIdentityResponse, error) {
	return svc.client.DeleteEmailIdentityRequest(in).Send(ctx)
}

// RawGetAccount executes `GetAccount` raw operation.
func (svc *PinpointEmail) RawGetAccount(ctx context.Context, in *SDK.GetAccountInput) (*SDK.GetAccountResponse, error) {
	return svc.client.GetAccountRequest(in).Send(ctx)
}

// RawGetBlacklistReports executes `GetBlacklistReports` raw operation.
func (svc *PinpointEmail) RawGetBlacklistReports(ctx context.Context, in *SDK.GetBlacklistReportsInput) (*SDK.GetBlacklistReportsResponse, error) {
	return svc.client.GetBlacklistReportsRequest(in).Send(ctx)
}

// RawGetConfigurationSet executes `GetConfigurationSet` raw operation.
func (svc *PinpointEmail) RawGetConfigurationSet(ctx context.Context, in *SDK.GetConfigurationSetInput) (*SDK.GetConfigurationSetResponse, error) {
	return svc.client.GetConfigurationSetRequest(in).Send(ctx)
}

// RawGetConfigurationSetEventDestinations executes `GetConfigurationSetEventDestinations` raw operation.
func (svc *PinpointEmail) RawGetConfigurationSetEventDestinations(ctx context.Context, in *SDK.GetConfigurationSetEventDestinationsInput) (*SDK.GetConfigurationSetEventDestinationsResponse, error) {
	return svc.client.GetConfigurationSetEventDestinationsRequest(in).Send(ctx)
}

// RawGetDedicatedIp executes `GetDedicatedIp` raw operation.
func (svc *PinpointEmail) RawGetDedicatedIp(ctx context.Context, in *SDK.GetDedicatedIpInput) (*SDK.GetDedicatedIpResponse, error) {
	return svc.client.GetDedicatedIpRequest(in).Send(ctx)
}

// RawGetDedicatedIps executes `GetDedicatedIps` raw operation.
func (svc *PinpointEmail) RawGetDedicatedIps(ctx context.Context, in *SDK.GetDedicatedIpsInput) (*SDK.GetDedicatedIpsResponse, error) {
	return svc.client.GetDedicatedIpsRequest(in).Send(ctx)
}

// RawGetDeliverabilityDashboardOptions executes `GetDeliverabilityDashboardOptions` raw operation.
func (svc *PinpointEmail) RawGetDeliverabilityDashboardOptions(ctx context.Context, in *SDK.GetDeliverabilityDashboardOptionsInput) (*SDK.GetDeliverabilityDashboardOptionsResponse, error) {
	return svc.client.GetDeliverabilityDashboardOptionsRequest(in).Send(ctx)
}

// RawGetDeliverabilityTestReport executes `GetDeliverabilityTestReport` raw operation.
func (svc *PinpointEmail) RawGetDeliverabilityTestReport(ctx context.Context, in *SDK.GetDeliverabilityTestReportInput) (*SDK.GetDeliverabilityTestReportResponse, error) {
	return svc.client.GetDeliverabilityTestReportRequest(in).Send(ctx)
}

// RawGetDomainDeliverabilityCampaign executes `GetDomainDeliverabilityCampaign` raw operation.
func (svc *PinpointEmail) RawGetDomainDeliverabilityCampaign(ctx context.Context, in *SDK.GetDomainDeliverabilityCampaignInput) (*SDK.GetDomainDeliverabilityCampaignResponse, error) {
	return svc.client.GetDomainDeliverabilityCampaignRequest(in).Send(ctx)
}

// RawGetDomainStatisticsReport executes `GetDomainStatisticsReport` raw operation.
func (svc *PinpointEmail) RawGetDomainStatisticsReport(ctx context.Context, in *SDK.GetDomainStatisticsReportInput) (*SDK.GetDomainStatisticsReportResponse, error) {
	return svc.client.GetDomainStatisticsReportRequest(in).Send(ctx)
}

// RawGetEmailIdentity executes `GetEmailIdentity` raw operation.
func (svc *PinpointEmail) RawGetEmailIdentity(ctx context.Context, in *SDK.GetEmailIdentityInput) (*SDK.GetEmailIdentityResponse, error) {
	return svc.client.GetEmailIdentityRequest(in).Send(ctx)
}

// RawListConfigurationSets executes `ListConfigurationSets` raw operation.
func (svc *PinpointEmail) RawListConfigurationSets(ctx context.Context, in *SDK.ListConfigurationSetsInput) (*SDK.ListConfigurationSetsResponse, error) {
	return svc.client.ListConfigurationSetsRequest(in).Send(ctx)
}

// RawListDedicatedIpPools executes `ListDedicatedIpPools` raw operation.
func (svc *PinpointEmail) RawListDedicatedIpPools(ctx context.Context, in *SDK.ListDedicatedIpPoolsInput) (*SDK.ListDedicatedIpPoolsResponse, error) {
	return svc.client.ListDedicatedIpPoolsRequest(in).Send(ctx)
}

// RawListDeliverabilityTestReports executes `ListDeliverabilityTestReports` raw operation.
func (svc *PinpointEmail) RawListDeliverabilityTestReports(ctx context.Context, in *SDK.ListDeliverabilityTestReportsInput) (*SDK.ListDeliverabilityTestReportsResponse, error) {
	return svc.client.ListDeliverabilityTestReportsRequest(in).Send(ctx)
}

// RawListDomainDeliverabilityCampaigns executes `ListDomainDeliverabilityCampaigns` raw operation.
func (svc *PinpointEmail) RawListDomainDeliverabilityCampaigns(ctx context.Context, in *SDK.ListDomainDeliverabilityCampaignsInput) (*SDK.ListDomainDeliverabilityCampaignsResponse, error) {
	return svc.client.ListDomainDeliverabilityCampaignsRequest(in).Send(ctx)
}

// RawListEmailIdentities executes `ListEmailIdentities` raw operation.
func (svc *PinpointEmail) RawListEmailIdentities(ctx context.Context, in *SDK.ListEmailIdentitiesInput) (*SDK.ListEmailIdentitiesResponse, error) {
	return svc.client.ListEmailIdentitiesRequest(in).Send(ctx)
}

// RawListTagsForResource executes `ListTagsForResource` raw operation.
func (svc *PinpointEmail) RawListTagsForResource(ctx context.Context, in *SDK.ListTagsForResourceInput) (*SDK.ListTagsForResourceResponse, error) {
	return svc.client.ListTagsForResourceRequest(in).Send(ctx)
}

// RawPutAccountDedicatedIpWarmupAttributes executes `PutAccountDedicatedIpWarmupAttributes` raw operation.
func (svc *PinpointEmail) RawPutAccountDedicatedIpWarmupAttributes(ctx context.Context, in *SDK.PutAccountDedicatedIpWarmupAttributesInput) (*SDK.PutAccountDedicatedIpWarmupAttributesResponse, error) {
	return svc.client.PutAccountDedicatedIpWarmupAttributesRequest(in).Send(ctx)
}

// RawPutAccountSendingAttributes executes `PutAccountSendingAttributes` raw operation.
func (svc *PinpointEmail) RawPutAccountSendingAttributes(ctx context.Context, in *SDK.PutAccountSendingAttributesInput) (*SDK.PutAccountSendingAttributesResponse, error) {
	return svc.client.PutAccountSendingAttributesRequest(in).Send(ctx)
}

// RawPutConfigurationSetDeliveryOptions executes `PutConfigurationSetDeliveryOptions` raw operation.
func (svc *PinpointEmail) RawPutConfigurationSetDeliveryOptions(ctx context.Context, in *SDK.PutConfigurationSetDeliveryOptionsInput) (*SDK.PutConfigurationSetDeliveryOptionsResponse, error) {
	return svc.client.PutConfigurationSetDeliveryOptionsRequest(in).Send(ctx)
}

// RawPutConfigurationSetReputationOptions executes `PutConfigurationSetReputationOptions` raw operation.
func (svc *PinpointEmail) RawPutConfigurationSetReputationOptions(ctx context.Context, in *SDK.PutConfigurationSetReputationOptionsInput) (*SDK.PutConfigurationSetReputationOptionsResponse, error) {
	return svc.client.PutConfigurationSetReputationOptionsRequest(in).Send(ctx)
}

// RawPutConfigurationSetSendingOptions executes `PutConfigurationSetSendingOptions` raw operation.
func (svc *PinpointEmail) RawPutConfigurationSetSendingOptions(ctx context.Context, in *SDK.PutConfigurationSetSendingOptionsInput) (*SDK.PutConfigurationSetSendingOptionsResponse, error) {
	return svc.client.PutConfigurationSetSendingOptionsRequest(in).Send(ctx)
}

// RawPutConfigurationSetTrackingOptions executes `PutConfigurationSetTrackingOptions` raw operation.
func (svc *PinpointEmail) RawPutConfigurationSetTrackingOptions(ctx context.Context, in *SDK.PutConfigurationSetTrackingOptionsInput) (*SDK.PutConfigurationSetTrackingOptionsResponse, error) {
	return svc.client.PutConfigurationSetTrackingOptionsRequest(in).Send(ctx)
}

// RawPutDedicatedIpInPool executes `PutDedicatedIpInPool` raw operation.
func (svc *PinpointEmail) RawPutDedicatedIpInPool(ctx context.Context, in *SDK.PutDedicatedIpInPoolInput) (*SDK.PutDedicatedIpInPoolResponse, error) {
	return svc.client.PutDedicatedIpInPoolRequest(in).Send(ctx)
}

// RawPutDedicatedIpWarmupAttributes executes `PutDedicatedIpWarmupAttributes` raw operation.
func (svc *PinpointEmail) RawPutDedicatedIpWarmupAttributes(ctx context.Context, in *SDK.PutDedicatedIpWarmupAttributesInput) (*SDK.PutDedicatedIpWarmupAttributesResponse, error) {
	return svc.client.PutDedicatedIpWarmupAttributesRequest(in).Send(ctx)
}

// RawPutDeliverabilityDashboardOption executes `PutDeliverabilityDashboardOption` raw operation.
func (svc *PinpointEmail) RawPutDeliverabilityDashboardOption(ctx context.Context, in *SDK.PutDeliverabilityDashboardOptionInput) (*SDK.PutDeliverabilityDashboardOptionResponse, error) {
	return svc.client.PutDeliverabilityDashboardOptionRequest(in).Send(ctx)
}

// RawPutEmailIdentityDkimAttributes executes `PutEmailIdentityDkimAttributes` raw operation.
func (svc *PinpointEmail) RawPutEmailIdentityDkimAttributes(ctx context.Context, in *SDK.PutEmailIdentityDkimAttributesInput) (*SDK.PutEmailIdentityDkimAttributesResponse, error) {
	return svc.client.PutEmailIdentityDkimAttributesRequest(in).Send(ctx)
}

// RawPutEmailIdentityFeedbackAttributes executes `PutEmailIdentityFeedbackAttributes` raw operation.
func (svc *PinpointEmail) RawPutEmailIdentityFeedbackAttributes(ctx context.Context, in *SDK.PutEmailIdentityFeedbackAttributesInput) (*SDK.PutEmailIdentityFeedbackAttributesResponse, error) {
	return svc.client.PutEmailIdentityFeedbackAttributesRequest(in).Send(ctx)
}

// RawPutEmailIdentityMailFromAttributes executes `PutEmailIdentityMailFromAttributes` raw operation.
func (svc *PinpointEmail) RawPutEmailIdentityMailFromAttributes(ctx context.Context, in *SDK.PutEmailIdentityMailFromAttributesInput) (*SDK.PutEmailIdentityMailFromAttributesResponse, error) {
	return svc.client.PutEmailIdentityMailFromAttributesRequest(in).Send(ctx)
}

// RawSendEmail executes `SendEmail` raw operation.
func (svc *PinpointEmail) RawSendEmail(ctx context.Context, in *SDK.SendEmailInput) (*SDK.SendEmailResponse, error) {
	return svc.client.SendEmailRequest(in).Send(ctx)
}

// RawTagResource executes `TagResource` raw operation.
func (svc *PinpointEmail) RawTagResource(ctx context.Context, in *SDK.TagResourceInput) (*SDK.TagResourceResponse, error) {
	return svc.client.TagResourceRequest(in).Send(ctx)
}

// RawUntagResource executes `UntagResource` raw operation.
func (svc *PinpointEmail) RawUntagResource(ctx context.Context, in *SDK.UntagResourceInput) (*SDK.UntagResourceResponse, error) {
	return svc.client.UntagResourceRequest(in).Send(ctx)
}

// RawUpdateConfigurationSetEventDestination executes `UpdateConfigurationSetEventDestination` raw operation.
func (svc *PinpointEmail) RawUpdateConfigurationSetEventDestination(ctx context.Context, in *SDK.UpdateConfigurationSetEventDestinationInput) (*SDK.UpdateConfigurationSetEventDestinationResponse, error) {
	return svc.client.UpdateConfigurationSetEventDestinationRequest(in).Send(ctx)
}
