package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type AttributeDimension struct {
	Values []string

	// optional
	AttributeType AttributeType
}

func newAttributeDimensionMap(in map[string]SDK.AttributeDimension) map[string]AttributeDimension {
	result := make(map[string]AttributeDimension, len(in))
	for key, val := range in {
		result[key] = newAttributeDimension(&val)
	}
	return result
}

func newAttributeDimension(o *SDK.AttributeDimension) AttributeDimension {
	result := AttributeDimension{}
	if o == nil {
		return result
	}

	result.Values = o.Values
	result.AttributeType = AttributeType(o.AttributeType)
	return result
}

func attributeDimensionMapToSDK(in map[string]AttributeDimension) map[string]SDK.AttributeDimension {
	result := make(map[string]SDK.AttributeDimension, len(in))
	for key, val := range in {
		result[key] = val.ToSDK()
	}
	return result
}

func (r AttributeDimension) ToSDK() SDK.AttributeDimension {
	o := SDK.AttributeDimension{}

	o.Values = r.Values
	o.AttributeType = SDK.AttributeType(r.AttributeType)
	return o
}

type CampaignCustomMessage struct {
	Data string
}

func newCampaignCustomMessage(o *SDK.CampaignCustomMessage) CampaignCustomMessage {
	result := CampaignCustomMessage{}
	if o == nil {
		return result
	}

	if o.Data != nil {
		result.Data = *o.Data
	}
	return result
}

func (r CampaignCustomMessage) ToSDK() *SDK.CampaignCustomMessage {
	o := SDK.CampaignCustomMessage{}

	if r.Data != "" {
		o.Data = pointers.String(r.Data)
	}
	return &o
}

type CampaignEmailMessage struct {
	Body        string
	FromAddress string
	HTMLBody    string
	Title       string
}

func newCampaignEmailMessage(o *SDK.CampaignEmailMessage) CampaignEmailMessage {
	result := CampaignEmailMessage{}
	if o == nil {
		return result
	}

	if o.Body != nil {
		result.Body = *o.Body
	}
	if o.FromAddress != nil {
		result.FromAddress = *o.FromAddress
	}
	if o.HtmlBody != nil {
		result.HTMLBody = *o.HtmlBody
	}
	if o.Title != nil {
		result.Title = *o.Title
	}
	return result
}

func (r CampaignEmailMessage) ToSDK() *SDK.CampaignEmailMessage {
	o := SDK.CampaignEmailMessage{}

	if r.Body != "" {
		o.Body = pointers.String(r.Body)
	}
	if r.FromAddress != "" {
		o.FromAddress = pointers.String(r.FromAddress)
	}
	if r.HTMLBody != "" {
		o.HtmlBody = pointers.String(r.HTMLBody)
	}
	if r.Title != "" {
		o.Title = pointers.String(r.Title)
	}
	return &o
}

type CampaignEventFilter struct {
	Dimensions EventDimensions
	FilterType FilterType
}

func newCampaignEventFilter(o *SDK.CampaignEventFilter) CampaignEventFilter {
	result := CampaignEventFilter{}
	if o == nil {
		return result
	}

	result.Dimensions = newEventDimensions(o.Dimensions)
	result.FilterType = FilterType(o.FilterType)
	return result
}

func (r CampaignEventFilter) ToSDK() *SDK.CampaignEventFilter {
	o := SDK.CampaignEventFilter{}

	o.Dimensions = r.Dimensions.ToSDK()
	o.FilterType = SDK.FilterType(r.FilterType)
	return &o
}

type CampaignHook struct {
	LambdaFunctionName string
	Mode               Mode
	WebURL             string
}

func newCampaignHook(o *SDK.CampaignHook) CampaignHook {
	result := CampaignHook{}
	if o == nil {
		return result
	}

	if o.LambdaFunctionName != nil {
		result.LambdaFunctionName = *o.LambdaFunctionName
	}

	result.Mode = Mode(o.Mode)

	if o.WebUrl != nil {
		result.WebURL = *o.WebUrl
	}
	return result
}

func (r CampaignHook) ToSDK() *SDK.CampaignHook {
	o := SDK.CampaignHook{}

	if r.LambdaFunctionName != "" {
		o.LambdaFunctionName = pointers.String(r.LambdaFunctionName)
	}

	o.Mode = SDK.Mode(r.Mode)

	if r.WebURL != "" {
		o.WebUrl = pointers.String(r.WebURL)
	}
	return &o
}

type CampaignLimits struct {
	Daily             int64
	MaximumDuration   int64
	MessagesPerSecond int64
	Total             int64
}

func newCampaignLimits(o *SDK.CampaignLimits) CampaignLimits {
	result := CampaignLimits{}
	if o == nil {
		return result
	}

	if o.Daily != nil {
		result.Daily = *o.Daily
	}
	if o.MaximumDuration != nil {
		result.MaximumDuration = *o.MaximumDuration
	}
	if o.MessagesPerSecond != nil {
		result.MessagesPerSecond = *o.MessagesPerSecond
	}
	if o.Total != nil {
		result.Total = *o.Total
	}
	return result
}

func (r CampaignLimits) ToSDK() *SDK.CampaignLimits {
	o := SDK.CampaignLimits{}

	if r.Daily != 0 {
		o.Daily = pointers.Long64(r.Daily)
	}
	if r.MaximumDuration != 0 {
		o.MaximumDuration = pointers.Long64(r.MaximumDuration)
	}
	if r.MessagesPerSecond != 0 {
		o.MessagesPerSecond = pointers.Long64(r.MessagesPerSecond)
	}
	if r.Total != 0 {
		o.Total = pointers.Long64(r.Total)
	}
	return &o
}

type CampaignResponse struct {
	ApplicationID    string
	ARN              string
	CreationDate     string
	ID               string
	LastModifiedDate string
	SegmentID        string
	SegmentVersion   int64

	AdditionalTreatments        []TreatmentResource
	CustomDeliveryConfiguration CustomDeliveryConfiguration
	DefaultState                CampaignStatus
	Description                 string
	HoldoutPercent              int64
	Hook                        CampaignHook
	IsPaused                    bool
	Limits                      CampaignLimits
	MessageConfiguration        MessageConfiguration
	Name                        string
	Schedule                    Schedule
	State                       CampaignStatus
	Tags                        map[string]string
	TemplateConfiguration       TemplateConfiguration
	TreatmentDescription        string
	TreatmentName               string
	Version                     int64
}

func newCampaignResponse(o *SDK.CampaignResponse) CampaignResponse {
	result := CampaignResponse{}
	if o == nil {
		return result
	}

	if o.ApplicationId != nil {
		result.ApplicationID = *o.ApplicationId
	}
	if o.Arn != nil {
		result.ARN = *o.Arn
	}
	if o.CreationDate != nil {
		result.CreationDate = *o.CreationDate
	}
	if o.Id != nil {
		result.ID = *o.Id
	}
	if o.LastModifiedDate != nil {
		result.LastModifiedDate = *o.LastModifiedDate
	}
	if o.SegmentId != nil {
		result.SegmentID = *o.SegmentId
	}
	if o.SegmentVersion != nil {
		result.SegmentVersion = *o.SegmentVersion
	}

	if len(o.AdditionalTreatments) != 0 {
		list := make([]TreatmentResource, len(o.AdditionalTreatments))
		for i, v := range o.AdditionalTreatments {
			list[i] = newTreatmentResource(&v)
		}
		result.AdditionalTreatments = list
	}

	result.CustomDeliveryConfiguration = newCustomDeliveryConfiguration(o.CustomDeliveryConfiguration)

	if o.DefaultState != nil {
		result.DefaultState = CampaignStatus(o.DefaultState.CampaignStatus)
	}

	if o.Description != nil {
		result.Description = *o.Description
	}
	if o.HoldoutPercent != nil {
		result.HoldoutPercent = *o.HoldoutPercent
	}

	result.Hook = newCampaignHook(o.Hook)

	if o.IsPaused != nil {
		result.IsPaused = *o.IsPaused
	}

	result.Limits = newCampaignLimits(o.Limits)
	result.MessageConfiguration = newMessageConfiguration(o.MessageConfiguration)

	if o.Name != nil {
		result.Name = *o.Name
	}

	result.Schedule = newSchedule(o.Schedule)

	if o.State != nil {
		result.State = CampaignStatus(o.State.CampaignStatus)
	}

	result.Tags = o.Tags
	result.TemplateConfiguration = newTemplateConfiguration(o.TemplateConfiguration)

	if o.TreatmentDescription != nil {
		result.TreatmentDescription = *o.TreatmentDescription
	}
	if o.TreatmentName != nil {
		result.TreatmentName = *o.TreatmentName
	}
	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}

type CampaignSMSMessage struct {
	Body        string
	MessageType MessageType
	SenderID    string
}

func newCampaignSMSMessage(o *SDK.CampaignSmsMessage) CampaignSMSMessage {
	result := CampaignSMSMessage{}
	if o == nil {
		return result
	}

	if o.Body != nil {
		result.Body = *o.Body
	}

	result.MessageType = MessageType(o.MessageType)

	if o.SenderId != nil {
		result.SenderID = *o.SenderId
	}
	return result
}

func (r CampaignSMSMessage) ToSDK() *SDK.CampaignSmsMessage {
	o := SDK.CampaignSmsMessage{}

	if r.Body != "" {
		o.Body = pointers.String(r.Body)
	}

	o.MessageType = SDK.MessageType(r.MessageType)

	if r.SenderID != "" {
		o.SenderId = pointers.String(r.SenderID)
	}
	return &o
}

type CustomDeliveryConfiguration struct {
	DeliveryURI string

	// optional
	EndpointTypes []EndpointTypesElement
}

func newCustomDeliveryConfiguration(o *SDK.CustomDeliveryConfiguration) CustomDeliveryConfiguration {
	result := CustomDeliveryConfiguration{}
	if o == nil {
		return result
	}

	if o.DeliveryUri != nil {
		result.DeliveryURI = *o.DeliveryUri
	}

	if len(o.EndpointTypes) != 0 {
		list := make([]EndpointTypesElement, len(o.EndpointTypes))
		for i, v := range o.EndpointTypes {
			list[i] = EndpointTypesElement(v)
		}
		result.EndpointTypes = list
	}
	return result
}

func (r CustomDeliveryConfiguration) ToSDK() *SDK.CustomDeliveryConfiguration {
	o := SDK.CustomDeliveryConfiguration{}

	if r.DeliveryURI != "" {
		o.DeliveryUri = pointers.String(r.DeliveryURI)
	}

	if len(r.EndpointTypes) != 0 {
		list := make([]SDK.EndpointTypesElement, len(r.EndpointTypes))
		for i, v := range r.EndpointTypes {
			list[i] = SDK.EndpointTypesElement(v)
		}
		o.EndpointTypes = list
	}
	return &o
}

type EventDimensions struct {
	Attributes map[string]AttributeDimension
	EventType  SetDimension
	Metrics    map[string]MetricDimension
}

func newEventDimensions(o *SDK.EventDimensions) EventDimensions {
	result := EventDimensions{}
	if o == nil {
		return result
	}

	result.Attributes = newAttributeDimensionMap(o.Attributes)
	result.EventType = newSetDimension(o.EventType)
	result.Metrics = newMetricDimensionMap(o.Metrics)
	return result
}

func (r EventDimensions) ToSDK() *SDK.EventDimensions {
	o := SDK.EventDimensions{}

	o.Attributes = attributeDimensionMapToSDK(r.Attributes)
	o.EventType = r.EventType.ToSDK()
	o.Metrics = metricDimensionMapToSDK(r.Metrics)
	return &o
}

type Message struct {
	Action            Action
	Body              string
	ImageIconURL      string
	ImageSmallIconURL string
	ImageURL          string
	JSONBody          string
	MediaURL          string
	RawContent        string
	SilentPush        bool
	TimeToLive        int64
	Title             string
	URL               string
}

func newMessage(o *SDK.Message) Message {
	result := Message{}
	if o == nil {
		return result
	}

	result.Action = Action(o.Action)

	if o.Body != nil {
		result.Body = *o.Body
	}
	if o.ImageIconUrl != nil {
		result.ImageIconURL = *o.ImageIconUrl
	}
	if o.ImageSmallIconUrl != nil {
		result.ImageSmallIconURL = *o.ImageSmallIconUrl
	}
	if o.ImageUrl != nil {
		result.ImageURL = *o.ImageUrl
	}
	if o.JsonBody != nil {
		result.JSONBody = *o.JsonBody
	}
	if o.MediaUrl != nil {
		result.MediaURL = *o.MediaUrl
	}
	if o.RawContent != nil {
		result.RawContent = *o.RawContent
	}
	if o.SilentPush != nil {
		result.SilentPush = *o.SilentPush
	}
	if o.TimeToLive != nil {
		result.TimeToLive = *o.TimeToLive
	}
	if o.Title != nil {
		result.Title = *o.Title
	}
	if o.Url != nil {
		result.URL = *o.Url
	}
	return result
}

func (r Message) ToSDK() *SDK.Message {
	o := SDK.Message{}

	o.Action = SDK.Action(r.Action)

	if r.Body != "" {
		o.Body = pointers.String(r.Body)
	}
	if r.ImageIconURL != "" {
		o.ImageIconUrl = pointers.String(r.ImageIconURL)
	}
	if r.ImageSmallIconURL != "" {
		o.ImageSmallIconUrl = pointers.String(r.ImageSmallIconURL)
	}
	if r.ImageURL != "" {
		o.ImageUrl = pointers.String(r.ImageURL)
	}
	if r.JSONBody != "" {
		o.JsonBody = pointers.String(r.JSONBody)
	}
	if r.MediaURL != "" {
		o.MediaUrl = pointers.String(r.MediaURL)
	}
	if r.RawContent != "" {
		o.RawContent = pointers.String(r.RawContent)
	}
	if r.SilentPush {
		o.SilentPush = pointers.Bool(r.SilentPush)
	}
	if r.TimeToLive != 0 {
		o.TimeToLive = pointers.Long64(r.TimeToLive)
	}
	if r.Title != "" {
		o.Title = pointers.String(r.Title)
	}
	if r.URL != "" {
		o.Url = pointers.String(r.URL)
	}
	return &o
}

type MessageConfiguration struct {
	ADMMessage     Message
	APNsMessage    Message
	BaiduMessage   Message
	CustomMessage  CampaignCustomMessage
	DefaultMessage Message
	EmailMessage   CampaignEmailMessage
	GCMMessage     Message
	SMSMessage     CampaignSMSMessage
}

func newMessageConfiguration(o *SDK.MessageConfiguration) MessageConfiguration {
	result := MessageConfiguration{}
	if o == nil {
		return result
	}

	result.ADMMessage = newMessage(o.ADMMessage)
	result.APNsMessage = newMessage(o.APNSMessage)
	result.BaiduMessage = newMessage(o.BaiduMessage)
	result.CustomMessage = newCampaignCustomMessage(o.CustomMessage)
	result.DefaultMessage = newMessage(o.DefaultMessage)
	result.EmailMessage = newCampaignEmailMessage(o.EmailMessage)
	result.GCMMessage = newMessage(o.GCMMessage)
	result.SMSMessage = newCampaignSMSMessage(o.SMSMessage)
	return result
}

func (r MessageConfiguration) ToSDK() *SDK.MessageConfiguration {
	o := SDK.MessageConfiguration{}

	o.ADMMessage = r.ADMMessage.ToSDK()
	o.APNSMessage = r.APNsMessage.ToSDK()
	o.BaiduMessage = r.BaiduMessage.ToSDK()
	o.CustomMessage = r.CustomMessage.ToSDK()
	o.DefaultMessage = r.DefaultMessage.ToSDK()
	o.EmailMessage = r.EmailMessage.ToSDK()
	o.GCMMessage = r.GCMMessage.ToSDK()
	o.SMSMessage = r.SMSMessage.ToSDK()
	return &o
}

type MetricDimension struct {
	ComparisonOperator string
	Value              float64
}

func newMetricDimensionMap(in map[string]SDK.MetricDimension) map[string]MetricDimension {
	result := make(map[string]MetricDimension, len(in))
	for key, val := range in {
		result[key] = newMetricDimension(&val)
	}
	return result
}

func newMetricDimension(o *SDK.MetricDimension) MetricDimension {
	result := MetricDimension{}

	if o.ComparisonOperator != nil {
		result.ComparisonOperator = *o.ComparisonOperator
	}
	if o.Value != nil {
		result.Value = *o.Value
	}
	return result
}

func metricDimensionMapToSDK(in map[string]MetricDimension) map[string]SDK.MetricDimension {
	result := make(map[string]SDK.MetricDimension, len(in))
	for key, val := range in {
		result[key] = val.ToSDK()
	}
	return result
}
func (r MetricDimension) ToSDK() SDK.MetricDimension {
	o := SDK.MetricDimension{}

	if r.ComparisonOperator != "" {
		o.ComparisonOperator = pointers.String(r.ComparisonOperator)
	}

	o.Value = pointers.Float64(r.Value)
	return o
}

type QuietTime struct {
	End   string
	Start string
}

func newQuietTime(o *SDK.QuietTime) QuietTime {
	result := QuietTime{}
	if o == nil {
		return result
	}

	if o.End != nil {
		result.End = *o.End
	}
	if o.Start != nil {
		result.Start = *o.Start
	}
	return result
}

func (r QuietTime) ToSDK() *SDK.QuietTime {
	o := SDK.QuietTime{}

	if r.End != "" {
		o.End = pointers.String(r.End)
	}
	if r.Start != "" {
		o.Start = pointers.String(r.Start)
	}
	return &o
}

type Schedule struct {
	StartTime string

	// optional
	EndTime     string
	EventFilter CampaignEventFilter
	Frequency   Frequency
	IsLocalTime bool
	QuietTime   QuietTime
	Timezone    string
}

func newSchedule(o *SDK.Schedule) Schedule {
	result := Schedule{}
	if o == nil {
		return result
	}

	if o.StartTime != nil {
		result.StartTime = *o.StartTime
	}

	result.EventFilter = newCampaignEventFilter(o.EventFilter)
	result.Frequency = Frequency(o.Frequency)

	if o.IsLocalTime != nil {
		result.IsLocalTime = *o.IsLocalTime
	}

	result.QuietTime = newQuietTime(o.QuietTime)

	if o.Timezone != nil {
		result.Timezone = *o.Timezone
	}
	return result
}

func (r Schedule) ToSDK() *SDK.Schedule {
	o := SDK.Schedule{}

	if r.StartTime != "" {
		o.StartTime = pointers.String(r.StartTime)
	}
	if r.EndTime != "" {
		o.EndTime = pointers.String(r.EndTime)
	}

	o.EventFilter = r.EventFilter.ToSDK()
	o.Frequency = SDK.Frequency(r.Frequency)

	if r.IsLocalTime {
		o.IsLocalTime = pointers.Bool(r.IsLocalTime)
	}

	o.QuietTime = r.QuietTime.ToSDK()

	if r.Timezone != "" {
		o.Timezone = pointers.String(r.Timezone)
	}
	return &o
}

type SetDimension struct {
	Values []string

	// optional
	DimensionType DimensionType
}

func newSetDimension(o *SDK.SetDimension) SetDimension {
	result := SetDimension{}
	if o == nil {
		return result
	}

	result.Values = o.Values
	result.DimensionType = DimensionType(o.DimensionType)
	return result
}

func (r SetDimension) ToSDK() *SDK.SetDimension {
	o := SDK.SetDimension{}

	o.Values = r.Values
	o.DimensionType = SDK.DimensionType(r.DimensionType)
	return &o
}

type Template struct {
	Name    string
	Version string
}

func newTemplate(o *SDK.Template) Template {
	result := Template{}
	if o == nil {
		return result
	}

	if o.Name != nil {
		result.Name = *o.Name
	}
	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}

func (r Template) ToSDK() *SDK.Template {
	o := SDK.Template{}

	if r.Name != "" {
		o.Name = pointers.String(r.Name)
	}
	if r.Version != "" {
		o.Version = pointers.String(r.Version)
	}
	return &o
}

type TemplateConfiguration struct {
	EmailTemplate Template
	PushTemplate  Template
	SMSTemplate   Template
	VoiceTemplate Template
}

func newTemplateConfiguration(o *SDK.TemplateConfiguration) TemplateConfiguration {
	result := TemplateConfiguration{}
	if o == nil {
		return result
	}

	result.EmailTemplate = newTemplate(o.EmailTemplate)
	result.PushTemplate = newTemplate(o.PushTemplate)
	result.SMSTemplate = newTemplate(o.SMSTemplate)
	result.VoiceTemplate = newTemplate(o.VoiceTemplate)
	return result
}

func (r TemplateConfiguration) ToSDK() *SDK.TemplateConfiguration {
	o := SDK.TemplateConfiguration{}

	o.EmailTemplate = r.EmailTemplate.ToSDK()
	o.PushTemplate = r.PushTemplate.ToSDK()
	o.SMSTemplate = r.SMSTemplate.ToSDK()
	o.VoiceTemplate = r.VoiceTemplate.ToSDK()
	return &o
}

type TreatmentResource struct {
	ID          string
	SizePercent int64

	// optional
	CustomDeliveryConfiguration CustomDeliveryConfiguration
	MessageConfiguration        MessageConfiguration
	Schedule                    Schedule
	State                       CampaignStatus
	TemplateConfiguration       TemplateConfiguration
	TreatmentDescription        string
	TreatmentName               string
}

func newTreatmentResource(o *SDK.TreatmentResource) TreatmentResource {
	result := TreatmentResource{}
	if o == nil {
		return result
	}

	if o.Id != nil {
		result.ID = *o.Id
	}
	if o.SizePercent != nil {
		result.SizePercent = *o.SizePercent
	}

	result.CustomDeliveryConfiguration = newCustomDeliveryConfiguration(o.CustomDeliveryConfiguration)
	result.MessageConfiguration = newMessageConfiguration(o.MessageConfiguration)
	result.Schedule = newSchedule(o.Schedule)
	if o.State != nil {
		result.State = CampaignStatus(o.State.CampaignStatus)
	}
	result.TemplateConfiguration = newTemplateConfiguration(o.TemplateConfiguration)

	if o.TreatmentDescription != nil {
		result.TreatmentDescription = *o.TreatmentDescription
	}
	if o.TreatmentName != nil {
		result.TreatmentName = *o.TreatmentName
	}
	return result
}

type WriteCampaignRequest struct {
	AdditionalTreatments        []WriteTreatmentResource
	CustomDeliveryConfiguration CustomDeliveryConfiguration
	Description                 string
	HoldoutPercent              int64
	Hook                        CampaignHook
	IsPaused                    bool
	Limits                      CampaignLimits
	MessageConfiguration        MessageConfiguration
	Name                        string
	Schedule                    Schedule
	SegmentID                   string
	SegmentVersion              int64
	Tags                        map[string]string
	TemplateConfiguration       TemplateConfiguration
	TreatmentDescription        string
	TreatmentName               string
}

func (r WriteCampaignRequest) ToSDK() *SDK.WriteCampaignRequest {
	o := SDK.WriteCampaignRequest{}

	if len(r.AdditionalTreatments) != 0 {
		list := make([]SDK.WriteTreatmentResource, len(r.AdditionalTreatments))
		for i, v := range r.AdditionalTreatments {
			list[i] = v.ToSDK()
		}
		o.AdditionalTreatments = list
	}
	o.CustomDeliveryConfiguration = r.CustomDeliveryConfiguration.ToSDK()

	if r.Description != "" {
		o.Description = pointers.String(r.Description)
	}
	if r.HoldoutPercent != 0 {
		o.HoldoutPercent = pointers.Long64(r.HoldoutPercent)
	}

	o.Hook = r.Hook.ToSDK()

	if r.IsPaused {
		o.IsPaused = pointers.Bool(r.IsPaused)
	}
	if r.Name != "" {
		o.Name = pointers.String(r.Name)
	}

	if r.SegmentID != "" {
		o.SegmentId = pointers.String(r.SegmentID)
	}
	if r.SegmentVersion != 0 {
		o.SegmentVersion = pointers.Long64(r.SegmentVersion)
	}

	o.Tags = r.Tags

	if r.TreatmentDescription != "" {
		o.TreatmentDescription = pointers.String(r.TreatmentDescription)
	}
	if r.TreatmentName != "" {
		o.TreatmentName = pointers.String(r.TreatmentName)
	}
	return &o
}

type WriteTreatmentResource struct {
	SizePercent int64

	// optional
	CustomDeliveryConfiguration CustomDeliveryConfiguration
	MessageConfiguration        MessageConfiguration
	Schedule                    Schedule
	TemplateConfiguration       TemplateConfiguration
	TreatmentDescription        string
	TreatmentName               string
}

func (r WriteTreatmentResource) ToSDK() SDK.WriteTreatmentResource {
	o := SDK.WriteTreatmentResource{}

	if r.SizePercent != 0 {
		o.SizePercent = pointers.Long64(r.SizePercent)
	}

	o.CustomDeliveryConfiguration = r.CustomDeliveryConfiguration.ToSDK()
	o.MessageConfiguration = r.MessageConfiguration.ToSDK()
	o.Schedule = r.Schedule.ToSDK()
	o.TemplateConfiguration = r.TemplateConfiguration.ToSDK()

	if r.TreatmentDescription != "" {
		o.TreatmentDescription = pointers.String(r.TreatmentDescription)
	}
	if r.TreatmentName != "" {
		o.TreatmentName = pointers.String(r.TreatmentName)
	}
	return o
}
