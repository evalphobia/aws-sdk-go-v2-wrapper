package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"
)

// RawCloneReceiptRuleSet executes `CloneReceiptRuleSet` raw operation.
func (svc *SES) RawCloneReceiptRuleSet(ctx context.Context, in *SDK.CloneReceiptRuleSetInput) (*SDK.CloneReceiptRuleSetResponse, error) {
	return svc.client.CloneReceiptRuleSetRequest(in).Send(ctx)
}

// RawCreateConfigurationSet executes `CreateConfigurationSet` raw operation.
func (svc *SES) RawCreateConfigurationSet(ctx context.Context, in *SDK.CreateConfigurationSetInput) (*SDK.CreateConfigurationSetResponse, error) {
	return svc.client.CreateConfigurationSetRequest(in).Send(ctx)
}

// RawCreateConfigurationSetEventDestination executes `CreateConfigurationSetEventDestination` raw operation.
func (svc *SES) RawCreateConfigurationSetEventDestination(ctx context.Context, in *SDK.CreateConfigurationSetEventDestinationInput) (*SDK.CreateConfigurationSetEventDestinationResponse, error) {
	return svc.client.CreateConfigurationSetEventDestinationRequest(in).Send(ctx)
}

// RawCreateConfigurationSetTrackingOptions executes `CreateConfigurationSetTrackingOptions` raw operation.
func (svc *SES) RawCreateConfigurationSetTrackingOptions(ctx context.Context, in *SDK.CreateConfigurationSetTrackingOptionsInput) (*SDK.CreateConfigurationSetTrackingOptionsResponse, error) {
	return svc.client.CreateConfigurationSetTrackingOptionsRequest(in).Send(ctx)
}

// RawCreateCustomVerificationEmailTemplate executes `CreateCustomVerificationEmailTemplate` raw operation.
func (svc *SES) RawCreateCustomVerificationEmailTemplate(ctx context.Context, in *SDK.CreateCustomVerificationEmailTemplateInput) (*SDK.CreateCustomVerificationEmailTemplateResponse, error) {
	return svc.client.CreateCustomVerificationEmailTemplateRequest(in).Send(ctx)
}

// RawCreateReceiptFilter executes `CreateReceiptFilter` raw operation.
func (svc *SES) RawCreateReceiptFilter(ctx context.Context, in *SDK.CreateReceiptFilterInput) (*SDK.CreateReceiptFilterResponse, error) {
	return svc.client.CreateReceiptFilterRequest(in).Send(ctx)
}

// RawCreateReceiptRule executes `CreateReceiptRule` raw operation.
func (svc *SES) RawCreateReceiptRule(ctx context.Context, in *SDK.CreateReceiptRuleInput) (*SDK.CreateReceiptRuleResponse, error) {
	return svc.client.CreateReceiptRuleRequest(in).Send(ctx)
}

// RawCreateReceiptRuleSet executes `CreateReceiptRuleSet` raw operation.
func (svc *SES) RawCreateReceiptRuleSet(ctx context.Context, in *SDK.CreateReceiptRuleSetInput) (*SDK.CreateReceiptRuleSetResponse, error) {
	return svc.client.CreateReceiptRuleSetRequest(in).Send(ctx)
}

// RawCreateTemplate executes `CreateTemplate` raw operation.
func (svc *SES) RawCreateTemplate(ctx context.Context, in *SDK.CreateTemplateInput) (*SDK.CreateTemplateResponse, error) {
	return svc.client.CreateTemplateRequest(in).Send(ctx)
}

// RawDeleteConfigurationSet executes `DeleteConfigurationSet` raw operation.
func (svc *SES) RawDeleteConfigurationSet(ctx context.Context, in *SDK.DeleteConfigurationSetInput) (*SDK.DeleteConfigurationSetResponse, error) {
	return svc.client.DeleteConfigurationSetRequest(in).Send(ctx)
}

// RawDeleteConfigurationSetEventDestination executes `DeleteConfigurationSetEventDestination` raw operation.
func (svc *SES) RawDeleteConfigurationSetEventDestination(ctx context.Context, in *SDK.DeleteConfigurationSetEventDestinationInput) (*SDK.DeleteConfigurationSetEventDestinationResponse, error) {
	return svc.client.DeleteConfigurationSetEventDestinationRequest(in).Send(ctx)
}

// RawDeleteConfigurationSetTrackingOptions executes `DeleteConfigurationSetTrackingOptions` raw operation.
func (svc *SES) RawDeleteConfigurationSetTrackingOptions(ctx context.Context, in *SDK.DeleteConfigurationSetTrackingOptionsInput) (*SDK.DeleteConfigurationSetTrackingOptionsResponse, error) {
	return svc.client.DeleteConfigurationSetTrackingOptionsRequest(in).Send(ctx)
}

// RawDeleteCustomVerificationEmailTemplate executes `DeleteCustomVerificationEmailTemplate` raw operation.
func (svc *SES) RawDeleteCustomVerificationEmailTemplate(ctx context.Context, in *SDK.DeleteCustomVerificationEmailTemplateInput) (*SDK.DeleteCustomVerificationEmailTemplateResponse, error) {
	return svc.client.DeleteCustomVerificationEmailTemplateRequest(in).Send(ctx)
}

// RawDeleteIdentity executes `DeleteIdentity` raw operation.
func (svc *SES) RawDeleteIdentity(ctx context.Context, in *SDK.DeleteIdentityInput) (*SDK.DeleteIdentityResponse, error) {
	return svc.client.DeleteIdentityRequest(in).Send(ctx)
}

// RawDeleteIdentityPolicy executes `DeleteIdentityPolicy` raw operation.
func (svc *SES) RawDeleteIdentityPolicy(ctx context.Context, in *SDK.DeleteIdentityPolicyInput) (*SDK.DeleteIdentityPolicyResponse, error) {
	return svc.client.DeleteIdentityPolicyRequest(in).Send(ctx)
}

// RawDeleteReceiptFilter executes `DeleteReceiptFilter` raw operation.
func (svc *SES) RawDeleteReceiptFilter(ctx context.Context, in *SDK.DeleteReceiptFilterInput) (*SDK.DeleteReceiptFilterResponse, error) {
	return svc.client.DeleteReceiptFilterRequest(in).Send(ctx)
}

// RawDeleteReceiptRule executes `DeleteReceiptRule` raw operation.
func (svc *SES) RawDeleteReceiptRule(ctx context.Context, in *SDK.DeleteReceiptRuleInput) (*SDK.DeleteReceiptRuleResponse, error) {
	return svc.client.DeleteReceiptRuleRequest(in).Send(ctx)
}

// RawDeleteReceiptRuleSet executes `DeleteReceiptRuleSet` raw operation.
func (svc *SES) RawDeleteReceiptRuleSet(ctx context.Context, in *SDK.DeleteReceiptRuleSetInput) (*SDK.DeleteReceiptRuleSetResponse, error) {
	return svc.client.DeleteReceiptRuleSetRequest(in).Send(ctx)
}

// RawDeleteTemplate executes `DeleteTemplate` raw operation.
func (svc *SES) RawDeleteTemplate(ctx context.Context, in *SDK.DeleteTemplateInput) (*SDK.DeleteTemplateResponse, error) {
	return svc.client.DeleteTemplateRequest(in).Send(ctx)
}

// RawDeleteVerifiedEmailAddress executes `DeleteVerifiedEmailAddress` raw operation.
func (svc *SES) RawDeleteVerifiedEmailAddress(ctx context.Context, in *SDK.DeleteVerifiedEmailAddressInput) (*SDK.DeleteVerifiedEmailAddressResponse, error) {
	return svc.client.DeleteVerifiedEmailAddressRequest(in).Send(ctx)
}

// RawDescribeActiveReceiptRuleSet executes `DescribeActiveReceiptRuleSet` raw operation.
func (svc *SES) RawDescribeActiveReceiptRuleSet(ctx context.Context, in *SDK.DescribeActiveReceiptRuleSetInput) (*SDK.DescribeActiveReceiptRuleSetResponse, error) {
	return svc.client.DescribeActiveReceiptRuleSetRequest(in).Send(ctx)
}

// RawDescribeConfigurationSet executes `DescribeConfigurationSet` raw operation.
func (svc *SES) RawDescribeConfigurationSet(ctx context.Context, in *SDK.DescribeConfigurationSetInput) (*SDK.DescribeConfigurationSetResponse, error) {
	return svc.client.DescribeConfigurationSetRequest(in).Send(ctx)
}

// RawDescribeReceiptRule executes `DescribeReceiptRule` raw operation.
func (svc *SES) RawDescribeReceiptRule(ctx context.Context, in *SDK.DescribeReceiptRuleInput) (*SDK.DescribeReceiptRuleResponse, error) {
	return svc.client.DescribeReceiptRuleRequest(in).Send(ctx)
}

// RawDescribeReceiptRuleSet executes `DescribeReceiptRuleSet` raw operation.
func (svc *SES) RawDescribeReceiptRuleSet(ctx context.Context, in *SDK.DescribeReceiptRuleSetInput) (*SDK.DescribeReceiptRuleSetResponse, error) {
	return svc.client.DescribeReceiptRuleSetRequest(in).Send(ctx)
}

// RawGetAccountSendingEnabled executes `GetAccountSendingEnabled` raw operation.
func (svc *SES) RawGetAccountSendingEnabled(ctx context.Context, in *SDK.GetAccountSendingEnabledInput) (*SDK.GetAccountSendingEnabledResponse, error) {
	return svc.client.GetAccountSendingEnabledRequest(in).Send(ctx)
}

// RawGetCustomVerificationEmailTemplate executes `GetCustomVerificationEmailTemplate` raw operation.
func (svc *SES) RawGetCustomVerificationEmailTemplate(ctx context.Context, in *SDK.GetCustomVerificationEmailTemplateInput) (*SDK.GetCustomVerificationEmailTemplateResponse, error) {
	return svc.client.GetCustomVerificationEmailTemplateRequest(in).Send(ctx)
}

// RawGetIdentityDkimAttributes executes `GetIdentityDkimAttributes` raw operation.
func (svc *SES) RawGetIdentityDkimAttributes(ctx context.Context, in *SDK.GetIdentityDkimAttributesInput) (*SDK.GetIdentityDkimAttributesResponse, error) {
	return svc.client.GetIdentityDkimAttributesRequest(in).Send(ctx)
}

// RawGetIdentityMailFromDomainAttributes executes `GetIdentityMailFromDomainAttributes` raw operation.
func (svc *SES) RawGetIdentityMailFromDomainAttributes(ctx context.Context, in *SDK.GetIdentityMailFromDomainAttributesInput) (*SDK.GetIdentityMailFromDomainAttributesResponse, error) {
	return svc.client.GetIdentityMailFromDomainAttributesRequest(in).Send(ctx)
}

// RawGetIdentityNotificationAttributes executes `GetIdentityNotificationAttributes` raw operation.
func (svc *SES) RawGetIdentityNotificationAttributes(ctx context.Context, in *SDK.GetIdentityNotificationAttributesInput) (*SDK.GetIdentityNotificationAttributesResponse, error) {
	return svc.client.GetIdentityNotificationAttributesRequest(in).Send(ctx)
}

// RawGetIdentityPolicies executes `GetIdentityPolicies` raw operation.
func (svc *SES) RawGetIdentityPolicies(ctx context.Context, in *SDK.GetIdentityPoliciesInput) (*SDK.GetIdentityPoliciesResponse, error) {
	return svc.client.GetIdentityPoliciesRequest(in).Send(ctx)
}

// RawGetIdentityVerificationAttributes executes `GetIdentityVerificationAttributes` raw operation.
func (svc *SES) RawGetIdentityVerificationAttributes(ctx context.Context, in *SDK.GetIdentityVerificationAttributesInput) (*SDK.GetIdentityVerificationAttributesResponse, error) {
	return svc.client.GetIdentityVerificationAttributesRequest(in).Send(ctx)
}

// RawGetSendQuota executes `GetSendQuota` raw operation.
func (svc *SES) RawGetSendQuota(ctx context.Context, in *SDK.GetSendQuotaInput) (*SDK.GetSendQuotaResponse, error) {
	return svc.client.GetSendQuotaRequest(in).Send(ctx)
}

// RawGetSendStatistics executes `GetSendStatistics` raw operation.
func (svc *SES) RawGetSendStatistics(ctx context.Context, in *SDK.GetSendStatisticsInput) (*SDK.GetSendStatisticsResponse, error) {
	return svc.client.GetSendStatisticsRequest(in).Send(ctx)
}

// RawGetTemplate executes `GetTemplate` raw operation.
func (svc *SES) RawGetTemplate(ctx context.Context, in *SDK.GetTemplateInput) (*SDK.GetTemplateResponse, error) {
	return svc.client.GetTemplateRequest(in).Send(ctx)
}

// RawListConfigurationSets executes `ListConfigurationSets` raw operation.
func (svc *SES) RawListConfigurationSets(ctx context.Context, in *SDK.ListConfigurationSetsInput) (*SDK.ListConfigurationSetsResponse, error) {
	return svc.client.ListConfigurationSetsRequest(in).Send(ctx)
}

// RawListCustomVerificationEmailTemplates executes `ListCustomVerificationEmailTemplates` raw operation.
func (svc *SES) RawListCustomVerificationEmailTemplates(ctx context.Context, in *SDK.ListCustomVerificationEmailTemplatesInput) (*SDK.ListCustomVerificationEmailTemplatesResponse, error) {
	return svc.client.ListCustomVerificationEmailTemplatesRequest(in).Send(ctx)
}

// RawListIdentities executes `ListIdentities` raw operation.
func (svc *SES) RawListIdentities(ctx context.Context, in *SDK.ListIdentitiesInput) (*SDK.ListIdentitiesResponse, error) {
	return svc.client.ListIdentitiesRequest(in).Send(ctx)
}

// RawListIdentityPolicies executes `ListIdentityPolicies` raw operation.
func (svc *SES) RawListIdentityPolicies(ctx context.Context, in *SDK.ListIdentityPoliciesInput) (*SDK.ListIdentityPoliciesResponse, error) {
	return svc.client.ListIdentityPoliciesRequest(in).Send(ctx)
}

// RawListReceiptFilters executes `ListReceiptFilters` raw operation.
func (svc *SES) RawListReceiptFilters(ctx context.Context, in *SDK.ListReceiptFiltersInput) (*SDK.ListReceiptFiltersResponse, error) {
	return svc.client.ListReceiptFiltersRequest(in).Send(ctx)
}

// RawListReceiptRuleSets executes `ListReceiptRuleSets` raw operation.
func (svc *SES) RawListReceiptRuleSets(ctx context.Context, in *SDK.ListReceiptRuleSetsInput) (*SDK.ListReceiptRuleSetsResponse, error) {
	return svc.client.ListReceiptRuleSetsRequest(in).Send(ctx)
}

// RawListTemplates executes `ListTemplates` raw operation.
func (svc *SES) RawListTemplates(ctx context.Context, in *SDK.ListTemplatesInput) (*SDK.ListTemplatesResponse, error) {
	return svc.client.ListTemplatesRequest(in).Send(ctx)
}

// RawListVerifiedEmailAddresses executes `ListVerifiedEmailAddresses` raw operation.
func (svc *SES) RawListVerifiedEmailAddresses(ctx context.Context, in *SDK.ListVerifiedEmailAddressesInput) (*SDK.ListVerifiedEmailAddressesResponse, error) {
	return svc.client.ListVerifiedEmailAddressesRequest(in).Send(ctx)
}

// RawPutConfigurationSetDeliveryOptions executes `PutConfigurationSetDeliveryOptions` raw operation.
func (svc *SES) RawPutConfigurationSetDeliveryOptions(ctx context.Context, in *SDK.PutConfigurationSetDeliveryOptionsInput) (*SDK.PutConfigurationSetDeliveryOptionsResponse, error) {
	return svc.client.PutConfigurationSetDeliveryOptionsRequest(in).Send(ctx)
}

// RawPutIdentityPolicy executes `PutIdentityPolicy` raw operation.
func (svc *SES) RawPutIdentityPolicy(ctx context.Context, in *SDK.PutIdentityPolicyInput) (*SDK.PutIdentityPolicyResponse, error) {
	return svc.client.PutIdentityPolicyRequest(in).Send(ctx)
}

// RawReorderReceiptRuleSet executes `ReorderReceiptRuleSet` raw operation.
func (svc *SES) RawReorderReceiptRuleSet(ctx context.Context, in *SDK.ReorderReceiptRuleSetInput) (*SDK.ReorderReceiptRuleSetResponse, error) {
	return svc.client.ReorderReceiptRuleSetRequest(in).Send(ctx)
}

// RawSendBounce executes `SendBounce` raw operation.
func (svc *SES) RawSendBounce(ctx context.Context, in *SDK.SendBounceInput) (*SDK.SendBounceResponse, error) {
	return svc.client.SendBounceRequest(in).Send(ctx)
}

// RawSendBulkTemplatedEmail executes `SendBulkTemplatedEmail` raw operation.
func (svc *SES) RawSendBulkTemplatedEmail(ctx context.Context, in *SDK.SendBulkTemplatedEmailInput) (*SDK.SendBulkTemplatedEmailResponse, error) {
	return svc.client.SendBulkTemplatedEmailRequest(in).Send(ctx)
}

// RawSendCustomVerificationEmail executes `SendCustomVerificationEmail` raw operation.
func (svc *SES) RawSendCustomVerificationEmail(ctx context.Context, in *SDK.SendCustomVerificationEmailInput) (*SDK.SendCustomVerificationEmailResponse, error) {
	return svc.client.SendCustomVerificationEmailRequest(in).Send(ctx)
}

// RawSendEmail executes `SendEmail` raw operation.
func (svc *SES) RawSendEmail(ctx context.Context, in *SDK.SendEmailInput) (*SDK.SendEmailResponse, error) {
	return svc.client.SendEmailRequest(in).Send(ctx)
}

// RawSendRawEmail executes `SendRawEmail` raw operation.
func (svc *SES) RawSendRawEmail(ctx context.Context, in *SDK.SendRawEmailInput) (*SDK.SendRawEmailResponse, error) {
	return svc.client.SendRawEmailRequest(in).Send(ctx)
}

// RawSendTemplatedEmail executes `SendTemplatedEmail` raw operation.
func (svc *SES) RawSendTemplatedEmail(ctx context.Context, in *SDK.SendTemplatedEmailInput) (*SDK.SendTemplatedEmailResponse, error) {
	return svc.client.SendTemplatedEmailRequest(in).Send(ctx)
}

// RawSetActiveReceiptRuleSet executes `SetActiveReceiptRuleSet` raw operation.
func (svc *SES) RawSetActiveReceiptRuleSet(ctx context.Context, in *SDK.SetActiveReceiptRuleSetInput) (*SDK.SetActiveReceiptRuleSetResponse, error) {
	return svc.client.SetActiveReceiptRuleSetRequest(in).Send(ctx)
}

// RawSetIdentityDkimEnabled executes `SetIdentityDkimEnabled` raw operation.
func (svc *SES) RawSetIdentityDkimEnabled(ctx context.Context, in *SDK.SetIdentityDkimEnabledInput) (*SDK.SetIdentityDkimEnabledResponse, error) {
	return svc.client.SetIdentityDkimEnabledRequest(in).Send(ctx)
}

// RawSetIdentityFeedbackForwardingEnabled executes `SetIdentityFeedbackForwardingEnabled` raw operation.
func (svc *SES) RawSetIdentityFeedbackForwardingEnabled(ctx context.Context, in *SDK.SetIdentityFeedbackForwardingEnabledInput) (*SDK.SetIdentityFeedbackForwardingEnabledResponse, error) {
	return svc.client.SetIdentityFeedbackForwardingEnabledRequest(in).Send(ctx)
}

// RawSetIdentityHeadersInNotificationsEnabled executes `SetIdentityHeadersInNotificationsEnabled` raw operation.
func (svc *SES) RawSetIdentityHeadersInNotificationsEnabled(ctx context.Context, in *SDK.SetIdentityHeadersInNotificationsEnabledInput) (*SDK.SetIdentityHeadersInNotificationsEnabledResponse, error) {
	return svc.client.SetIdentityHeadersInNotificationsEnabledRequest(in).Send(ctx)
}

// RawSetIdentityMailFromDomain executes `SetIdentityMailFromDomain` raw operation.
func (svc *SES) RawSetIdentityMailFromDomain(ctx context.Context, in *SDK.SetIdentityMailFromDomainInput) (*SDK.SetIdentityMailFromDomainResponse, error) {
	return svc.client.SetIdentityMailFromDomainRequest(in).Send(ctx)
}

// RawSetIdentityNotificationTopic executes `SetIdentityNotificationTopic` raw operation.
func (svc *SES) RawSetIdentityNotificationTopic(ctx context.Context, in *SDK.SetIdentityNotificationTopicInput) (*SDK.SetIdentityNotificationTopicResponse, error) {
	return svc.client.SetIdentityNotificationTopicRequest(in).Send(ctx)
}

// RawSetReceiptRulePosition executes `SetReceiptRulePosition` raw operation.
func (svc *SES) RawSetReceiptRulePosition(ctx context.Context, in *SDK.SetReceiptRulePositionInput) (*SDK.SetReceiptRulePositionResponse, error) {
	return svc.client.SetReceiptRulePositionRequest(in).Send(ctx)
}

// RawTestRenderTemplate executes `TestRenderTemplate` raw operation.
func (svc *SES) RawTestRenderTemplate(ctx context.Context, in *SDK.TestRenderTemplateInput) (*SDK.TestRenderTemplateResponse, error) {
	return svc.client.TestRenderTemplateRequest(in).Send(ctx)
}

// RawUpdateAccountSendingEnabled executes `UpdateAccountSendingEnabled` raw operation.
func (svc *SES) RawUpdateAccountSendingEnabled(ctx context.Context, in *SDK.UpdateAccountSendingEnabledInput) (*SDK.UpdateAccountSendingEnabledResponse, error) {
	return svc.client.UpdateAccountSendingEnabledRequest(in).Send(ctx)
}

// RawUpdateConfigurationSetEventDestination executes `UpdateConfigurationSetEventDestination` raw operation.
func (svc *SES) RawUpdateConfigurationSetEventDestination(ctx context.Context, in *SDK.UpdateConfigurationSetEventDestinationInput) (*SDK.UpdateConfigurationSetEventDestinationResponse, error) {
	return svc.client.UpdateConfigurationSetEventDestinationRequest(in).Send(ctx)
}

// RawUpdateConfigurationSetReputationMetricsEnabled executes `UpdateConfigurationSetReputationMetricsEnabled` raw operation.
func (svc *SES) RawUpdateConfigurationSetReputationMetricsEnabled(ctx context.Context, in *SDK.UpdateConfigurationSetReputationMetricsEnabledInput) (*SDK.UpdateConfigurationSetReputationMetricsEnabledResponse, error) {
	return svc.client.UpdateConfigurationSetReputationMetricsEnabledRequest(in).Send(ctx)
}

// RawUpdateConfigurationSetSendingEnabled executes `UpdateConfigurationSetSendingEnabled` raw operation.
func (svc *SES) RawUpdateConfigurationSetSendingEnabled(ctx context.Context, in *SDK.UpdateConfigurationSetSendingEnabledInput) (*SDK.UpdateConfigurationSetSendingEnabledResponse, error) {
	return svc.client.UpdateConfigurationSetSendingEnabledRequest(in).Send(ctx)
}

// RawUpdateConfigurationSetTrackingOptions executes `UpdateConfigurationSetTrackingOptions` raw operation.
func (svc *SES) RawUpdateConfigurationSetTrackingOptions(ctx context.Context, in *SDK.UpdateConfigurationSetTrackingOptionsInput) (*SDK.UpdateConfigurationSetTrackingOptionsResponse, error) {
	return svc.client.UpdateConfigurationSetTrackingOptionsRequest(in).Send(ctx)
}

// RawUpdateCustomVerificationEmailTemplate executes `UpdateCustomVerificationEmailTemplate` raw operation.
func (svc *SES) RawUpdateCustomVerificationEmailTemplate(ctx context.Context, in *SDK.UpdateCustomVerificationEmailTemplateInput) (*SDK.UpdateCustomVerificationEmailTemplateResponse, error) {
	return svc.client.UpdateCustomVerificationEmailTemplateRequest(in).Send(ctx)
}

// RawUpdateReceiptRule executes `UpdateReceiptRule` raw operation.
func (svc *SES) RawUpdateReceiptRule(ctx context.Context, in *SDK.UpdateReceiptRuleInput) (*SDK.UpdateReceiptRuleResponse, error) {
	return svc.client.UpdateReceiptRuleRequest(in).Send(ctx)
}

// RawUpdateTemplate executes `UpdateTemplate` raw operation.
func (svc *SES) RawUpdateTemplate(ctx context.Context, in *SDK.UpdateTemplateInput) (*SDK.UpdateTemplateResponse, error) {
	return svc.client.UpdateTemplateRequest(in).Send(ctx)
}

// RawVerifyDomainDkim executes `VerifyDomainDkim` raw operation.
func (svc *SES) RawVerifyDomainDkim(ctx context.Context, in *SDK.VerifyDomainDkimInput) (*SDK.VerifyDomainDkimResponse, error) {
	return svc.client.VerifyDomainDkimRequest(in).Send(ctx)
}

// RawVerifyDomainIdentity executes `VerifyDomainIdentity` raw operation.
func (svc *SES) RawVerifyDomainIdentity(ctx context.Context, in *SDK.VerifyDomainIdentityInput) (*SDK.VerifyDomainIdentityResponse, error) {
	return svc.client.VerifyDomainIdentityRequest(in).Send(ctx)
}

// RawVerifyEmailAddress executes `VerifyEmailAddress` raw operation.
func (svc *SES) RawVerifyEmailAddress(ctx context.Context, in *SDK.VerifyEmailAddressInput) (*SDK.VerifyEmailAddressResponse, error) {
	return svc.client.VerifyEmailAddressRequest(in).Send(ctx)
}

// RawVerifyEmailIdentity executes `VerifyEmailIdentity` raw operation.
func (svc *SES) RawVerifyEmailIdentity(ctx context.Context, in *SDK.VerifyEmailIdentityInput) (*SDK.VerifyEmailIdentityResponse, error) {
	return svc.client.VerifyEmailIdentityRequest(in).Send(ctx)
}
