package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type AndroidPushNotificationTemplate struct {
	Action            Action
	Body              string
	ImageIconURL      string
	ImageURL          string
	RawContent        string
	SmallImageIconURL string
	Sound             string
	Title             string
	URL               string
}

func (r AndroidPushNotificationTemplate) ToSDK() *SDK.AndroidPushNotificationTemplate {
	o := SDK.AndroidPushNotificationTemplate{}

	o.Action = SDK.Action(r.Action)

	if r.Body != "" {
		o.Body = pointers.String(r.Body)
	}
	if r.ImageIconURL != "" {
		o.ImageIconUrl = pointers.String(r.ImageIconURL)
	}
	if r.ImageURL != "" {
		o.ImageUrl = pointers.String(r.ImageURL)
	}
	if r.RawContent != "" {
		o.RawContent = pointers.String(r.RawContent)
	}
	if r.SmallImageIconURL != "" {
		o.SmallImageIconUrl = pointers.String(r.SmallImageIconURL)
	}
	if r.Sound != "" {
		o.Sound = pointers.String(r.Sound)
	}
	if r.Title != "" {
		o.Title = pointers.String(r.Title)
	}
	if r.URL != "" {
		o.Url = pointers.String(r.URL)
	}
	return &o
}

type APNsPushNotificationTemplate struct {
	Action     Action
	Body       string
	MediaURL   string
	RawContent string
	Sound      string
	Title      string
	URL        string
}

func (r APNsPushNotificationTemplate) ToSDK() *SDK.APNSPushNotificationTemplate {
	o := SDK.APNSPushNotificationTemplate{}

	o.Action = SDK.Action(r.Action)

	if r.Body != "" {
		o.Body = pointers.String(r.Body)
	}
	if r.MediaURL != "" {
		o.MediaUrl = pointers.String(r.MediaURL)
	}
	if r.RawContent != "" {
		o.RawContent = pointers.String(r.RawContent)
	}
	if r.Sound != "" {
		o.Sound = pointers.String(r.Sound)
	}
	if r.Title != "" {
		o.Title = pointers.String(r.Title)
	}
	if r.URL != "" {
		o.Url = pointers.String(r.URL)
	}
	return &o
}

type CreateTemplateMessageBody struct {
	ARN       string
	Message   string
	RequestID string
}

func newCreateTemplateMessageBody(o *SDK.CreateTemplateMessageBody) CreateTemplateMessageBody {
	result := CreateTemplateMessageBody{}
	if o == nil {
		return result
	}

	if o.Arn != nil {
		result.ARN = *o.Arn
	}
	if o.Message != nil {
		result.Message = *o.Message
	}
	if o.RequestID != nil {
		result.RequestID = *o.RequestID
	}

	return result
}

type DefaultPushNotificationTemplate struct {
	Action Action
	Body   string
	Sound  string
	Title  string
	URL    string
}

func (r DefaultPushNotificationTemplate) ToSDK() *SDK.DefaultPushNotificationTemplate {
	o := SDK.DefaultPushNotificationTemplate{}

	o.Action = SDK.Action(r.Action)

	if r.Body != "" {
		o.Body = pointers.String(r.Body)
	}
	if r.Sound != "" {
		o.Sound = pointers.String(r.Sound)
	}
	if r.Title != "" {
		o.Title = pointers.String(r.Title)
	}
	if r.URL != "" {
		o.Url = pointers.String(r.URL)
	}
	return &o
}

type EmailTemplateRequest struct {
	DefaultSubstitutions string
	HTMLPart             string
	RecommenderID        string
	Subject              string
	Tags                 map[string]string
	TemplateDescription  string
	TextPart             string
}

func (r EmailTemplateRequest) ToSDK() *SDK.EmailTemplateRequest {
	o := SDK.EmailTemplateRequest{}

	if r.DefaultSubstitutions != "" {
		o.DefaultSubstitutions = pointers.String(r.DefaultSubstitutions)
	}
	if r.HTMLPart != "" {
		o.HtmlPart = pointers.String(r.HTMLPart)
	}
	if r.RecommenderID != "" {
		o.RecommenderId = pointers.String(r.RecommenderID)
	}
	if r.Subject != "" {
		o.Subject = pointers.String(r.Subject)
	}

	o.Tags = r.Tags

	if r.TemplateDescription != "" {
		o.TemplateDescription = pointers.String(r.TemplateDescription)
	}
	if r.TextPart != "" {
		o.TextPart = pointers.String(r.TextPart)
	}
	return &o
}

type PushNotificationTemplateRequest struct {
	ADM                  AndroidPushNotificationTemplate
	APNS                 APNsPushNotificationTemplate
	Baidu                AndroidPushNotificationTemplate
	Default              DefaultPushNotificationTemplate
	DefaultSubstitutions string
	GCM                  AndroidPushNotificationTemplate
	RecommenderID        string
	Tags                 map[string]string
	TemplateDescription  string
}

func (r PushNotificationTemplateRequest) ToSDK() *SDK.PushNotificationTemplateRequest {
	o := SDK.PushNotificationTemplateRequest{}

	if r.ADM.Body != "" {
		o.ADM = r.ADM.ToSDK()
	}
	if r.APNS.Body != "" {
		o.APNS = r.APNS.ToSDK()
	}
	if r.Baidu.Body != "" {
		o.Baidu = r.Baidu.ToSDK()
	}
	if r.Default.Body != "" {
		o.Default = r.Default.ToSDK()
	}
	if r.DefaultSubstitutions != "" {
		o.DefaultSubstitutions = pointers.String(r.DefaultSubstitutions)
	}
	if r.GCM.Body != "" {
		o.GCM = r.GCM.ToSDK()
	}
	if r.RecommenderID != "" {
		o.RecommenderId = pointers.String(r.RecommenderID)
	}

	o.Tags = r.Tags

	if r.TemplateDescription != "" {
		o.TemplateDescription = pointers.String(r.TemplateDescription)
	}
	return &o
}

type MessageBody struct {
	Message   string
	RequestID string
}

func newMessageBody(o *SDK.MessageBody) MessageBody {
	result := MessageBody{}
	if o == nil {
		return result
	}

	if o.Message != nil {
		result.Message = *o.Message
	}
	if o.RequestID != nil {
		result.RequestID = *o.RequestID
	}
	return result
}

type TemplateResponse struct {
	CreationDate     string
	LastModifiedDate string
	TemplateName     string
	TemplateType     TemplateType

	ARN                  string
	DefaultSubstitutions string
	Tags                 map[string]string
	TemplateDescription  string
	Version              string
}

func newTemplateResponse(o *SDK.TemplateResponse) TemplateResponse {
	result := TemplateResponse{}
	if o == nil {
		return result
	}

	if o.CreationDate != nil {
		result.CreationDate = *o.CreationDate
	}
	if o.LastModifiedDate != nil {
		result.LastModifiedDate = *o.LastModifiedDate
	}
	if o.TemplateName != nil {
		result.TemplateName = *o.TemplateName
	}

	result.TemplateType = TemplateType(o.TemplateType)

	if o.Arn != nil {
		result.ARN = *o.Arn
	}
	if o.DefaultSubstitutions != nil {
		result.DefaultSubstitutions = *o.DefaultSubstitutions
	}

	result.Tags = o.Tags

	if o.TemplateDescription != nil {
		result.TemplateDescription = *o.TemplateDescription
	}
	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}

type TemplatesResponse struct {
	Item []TemplateResponse

	// optional
	NextToken string
}

func newTemplatesResponse(o *SDK.TemplatesResponse) TemplatesResponse {
	result := TemplatesResponse{}
	if o == nil {
		return result
	}

	if len(o.Item) != 0 {
		list := make([]TemplateResponse, len(o.Item))
		for i, v := range o.Item {
			list[i] = newTemplateResponse(&v)
		}
		result.Item = list
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}
	return result
}
