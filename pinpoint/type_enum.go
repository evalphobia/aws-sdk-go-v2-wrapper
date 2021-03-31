package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

type Action string

const (
	ActionOpenApp  Action = Action(SDK.ActionOpenApp)
	ActionDeepLink Action = Action(SDK.ActionDeepLink)
	ActionURL      Action = Action(SDK.ActionUrl)
)

type AttributeType string

const (
	AttributeTypeInclusive AttributeType = AttributeType(SDK.AttributeTypeInclusive)
	AttributeTypeExclusive AttributeType = AttributeType(SDK.AttributeTypeExclusive)
)

type CampaignStatus string

const (
	CampaignStatusScheduled      CampaignStatus = CampaignStatus(SDK.CampaignStatusScheduled)
	CampaignStatusExecuting      CampaignStatus = CampaignStatus(SDK.CampaignStatusExecuting)
	CampaignStatusPendingNextRun CampaignStatus = CampaignStatus(SDK.CampaignStatusPendingNextRun)
	CampaignStatusCompleted      CampaignStatus = CampaignStatus(SDK.CampaignStatusCompleted)
	CampaignStatusPaused         CampaignStatus = CampaignStatus(SDK.CampaignStatusPaused)
	CampaignStatusDeleted        CampaignStatus = CampaignStatus(SDK.CampaignStatusDeleted)
)

type ChannelType string

const (
	ChannelTypePush            ChannelType = ChannelType(SDK.ChannelTypePush)
	ChannelTypeGCM             ChannelType = ChannelType(SDK.ChannelTypeGcm)
	ChannelTypeAPNs            ChannelType = ChannelType(SDK.ChannelTypeApns)
	ChannelTypeAPNsSandbox     ChannelType = ChannelType(SDK.ChannelTypeApnsSandbox)
	ChannelTypeAPNsVoIP        ChannelType = ChannelType(SDK.ChannelTypeApnsVoip)
	ChannelTypeAPNsVoIPSandbox ChannelType = ChannelType(SDK.ChannelTypeApnsVoipSandbox)
	ChannelTypeADM             ChannelType = ChannelType(SDK.ChannelTypeAdm)
	ChannelTypeSMS             ChannelType = ChannelType(SDK.ChannelTypeSms)
	ChannelTypeVoice           ChannelType = ChannelType(SDK.ChannelTypeVoice)
	ChannelTypeEmail           ChannelType = ChannelType(SDK.ChannelTypeEmail)
	ChannelTypeBaidu           ChannelType = ChannelType(SDK.ChannelTypeBaidu)
	ChannelTypeCustom          ChannelType = ChannelType(SDK.ChannelTypeCustom)
)

type DimensionType string

const (
	DimensionTypeInclusive DimensionType = DimensionType(SDK.DimensionTypeInclusive)
	DimensionTypeExclusive DimensionType = DimensionType(SDK.DimensionTypeExclusive)
)

type Duration string

const (
	DurationHr24  Duration = Duration(SDK.DurationHr24)
	DurationDay7  Duration = Duration(SDK.DurationDay7)
	DurationDay14 Duration = Duration(SDK.DurationDay14)
	DurationDay30 Duration = Duration(SDK.DurationDay30)
)

type EndpointTypesElement string

const (
	EndpointTypesElementPush            ChannelType = ChannelType(SDK.EndpointTypesElementPush)
	EndpointTypesElementGCM             ChannelType = ChannelType(SDK.EndpointTypesElementGcm)
	EndpointTypesElementAPNs            ChannelType = ChannelType(SDK.EndpointTypesElementApns)
	EndpointTypesElementAPNsSandbox     ChannelType = ChannelType(SDK.EndpointTypesElementApnsSandbox)
	EndpointTypesElementAPNsVoIP        ChannelType = ChannelType(SDK.EndpointTypesElementApnsVoip)
	EndpointTypesElementAPNsVoIPSandbox ChannelType = ChannelType(SDK.EndpointTypesElementApnsVoipSandbox)
	EndpointTypesElementADM             ChannelType = ChannelType(SDK.EndpointTypesElementAdm)
	EndpointTypesElementSMS             ChannelType = ChannelType(SDK.EndpointTypesElementSms)
	EndpointTypesElementVoice           ChannelType = ChannelType(SDK.EndpointTypesElementVoice)
	EndpointTypesElementEmail           ChannelType = ChannelType(SDK.EndpointTypesElementEmail)
	EndpointTypesElementBaidu           ChannelType = ChannelType(SDK.EndpointTypesElementBaidu)
	EndpointTypesElementCustom          ChannelType = ChannelType(SDK.EndpointTypesElementCustom)
)

type FilterType string

const (
	FilterTypeSystem   FilterType = FilterType(SDK.FilterTypeSystem)
	FilterTypeEndpoint FilterType = FilterType(SDK.FilterTypeEndpoint)
)

type Format string

const (
	FormatCSV  Format = Format(SDK.FormatCsv)
	FormatJSON Format = Format(SDK.FormatJson)
)

type Frequency string

const (
	FrequencyOnce    Frequency = Frequency(SDK.FrequencyOnce)
	FrequencyHourly  Frequency = Frequency(SDK.FrequencyHourly)
	FrequencyDaily   Frequency = Frequency(SDK.FrequencyDaily)
	FrequencyWeekly  Frequency = Frequency(SDK.FrequencyWeekly)
	FrequencyMonthly Frequency = Frequency(SDK.FrequencyMonthly)
	FrequencyEvent   Frequency = Frequency(SDK.FrequencyEvent)
)

type Include string

const (
	IncludeAll  Include = Include(SDK.IncludeAll)
	IncludeAny  Include = Include(SDK.IncludeAny)
	IncludeNone Include = Include(SDK.IncludeNone)
)

type JobStatus string

const (
	JobStatusCreated                    JobStatus = JobStatus(SDK.JobStatusCreated)
	JobStatusPreparingForInitialization JobStatus = JobStatus(SDK.JobStatusPreparingForInitialization)
	JobStatusInitializing               JobStatus = JobStatus(SDK.JobStatusInitializing)
	JobStatusProcessing                 JobStatus = JobStatus(SDK.JobStatusProcessing)
	JobStatusPendingJob                 JobStatus = JobStatus(SDK.JobStatusPendingJob)
	JobStatusCompleting                 JobStatus = JobStatus(SDK.JobStatusCompleting)
	JobStatusCompleted                  JobStatus = JobStatus(SDK.JobStatusCompleted)
	JobStatusFailing                    JobStatus = JobStatus(SDK.JobStatusFailing)
	JobStatusFailed                     JobStatus = JobStatus(SDK.JobStatusFailed)
)

type MessageType string

const (
	MessageTypeTransactional MessageType = MessageType(SDK.MessageTypeTransactional)
	MessageTypePromotional   MessageType = MessageType(SDK.MessageTypePromotional)
)

type Mode string

const (
	ModeDelivery Mode = Mode(SDK.ModeDelivery)
	ModeFilter   Mode = Mode(SDK.ModeFilter)
)

type RecencyType string

const (
	RecencyTypeActive   RecencyType = RecencyType(SDK.RecencyTypeActive)
	RecencyTypeInactive RecencyType = RecencyType(SDK.RecencyTypeInactive)
)

type SegmentType string

const (
	SegmentTypeDimensional SegmentType = SegmentType(SDK.SegmentTypeDimensional)
	SegmentTypeImport      SegmentType = SegmentType(SDK.SegmentTypeImport)
)

type SourceType string

const (
	SourceTypeAll  SourceType = SourceType(SDK.SourceTypeAll)
	SourceTypeAny  SourceType = SourceType(SDK.SourceTypeAny)
	SourceTypeNone SourceType = SourceType(SDK.SourceTypeNone)
)

type TemplateType string

const (
	TemplateTypeEmail TemplateType = TemplateType(SDK.TemplateTypeEmail)
	TemplateTypeSMS   TemplateType = TemplateType(SDK.TemplateTypeSms)
	TemplateTypeVoice TemplateType = TemplateType(SDK.TemplateTypeVoice)
	TemplateTypePush  TemplateType = TemplateType(SDK.TemplateTypePush)
)

type Type string

const (
	TypeAll  Type = Type(SDK.TypeAll)
	TypeAny  Type = Type(SDK.TypeAny)
	TypeNone Type = Type(SDK.TypeNone)
)
