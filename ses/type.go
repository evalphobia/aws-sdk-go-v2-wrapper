package ses

import (
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type BulkEmailDestination struct {
	Destination             Destination
	ReplacementTags         []MessageTag
	ReplacementTemplateData string
}

func (r BulkEmailDestination) ToSDK() SDK.BulkEmailDestination {
	o := SDK.BulkEmailDestination{}

	o.Destination = r.Destination.ToSDK()

	if len(r.ReplacementTags) != 0 {
		list := make([]SDK.MessageTag, len(r.ReplacementTags))
		for i, v := range r.ReplacementTags {
			list[i] = v.ToSDK()
		}
		o.ReplacementTags = list
	}

	if r.ReplacementTemplateData != "" {
		o.ReplacementTemplateData = pointers.String(r.ReplacementTemplateData)
	}
	return o
}

type BulkEmailDestinationStatus struct {
	Error     string
	MessageID string
	Status    BulkEmailStatus
}

func newBulkEmailDestinationStatuses(list []SDK.BulkEmailDestinationStatus) []BulkEmailDestinationStatus {
	if len(list) == 0 {
		return nil
	}

	results := make([]BulkEmailDestinationStatus, len(list))
	for i, v := range list {
		results[i] = newBulkEmailDestinationStatus(v)
	}
	return results
}

func newBulkEmailDestinationStatus(o SDK.BulkEmailDestinationStatus) BulkEmailDestinationStatus {
	result := BulkEmailDestinationStatus{}

	if o.Error != nil {
		result.Error = *o.Error
	}
	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}

	result.Status = BulkEmailStatus(o.Status)
	return result
}

type Destination struct {
	BccAddresses []string
	CcAddresses  []string
	ToAddresses  []string
}

func (r Destination) ToSDK() *SDK.Destination {
	o := SDK.Destination{}

	o.BccAddresses = r.BccAddresses
	o.CcAddresses = r.CcAddresses
	o.ToAddresses = r.ToAddresses
	return &o
}

type MessageTag struct {
	Name  string
	Value string
}

func (r MessageTag) ToSDK() SDK.MessageTag {
	o := SDK.MessageTag{}

	if r.Name != "" {
		o.Name = pointers.String(r.Name)
	}
	if r.Value != "" {
		o.Value = pointers.String(r.Value)
	}
	return o
}

type Template struct {
	TemplateName string

	// optional
	HTMLPart    string
	SubjectPart string
	TextPart    string
}

func newTemplate(o *SDK.Template) Template {
	result := Template{}
	if o == nil {
		return result
	}

	if o.TemplateName != nil {
		result.TemplateName = *o.TemplateName
	}
	if o.HtmlPart != nil {
		result.HTMLPart = *o.HtmlPart
	}
	if o.SubjectPart != nil {
		result.SubjectPart = *o.SubjectPart
	}
	if o.TextPart != nil {
		result.TextPart = *o.TextPart
	}
	return result
}

func (r Template) ToSDK() *SDK.Template {
	o := SDK.Template{}

	if r.TemplateName != "" {
		o.TemplateName = pointers.String(r.TemplateName)
	}
	if r.HTMLPart != "" {
		o.HtmlPart = pointers.String(r.HTMLPart)
	}
	if r.SubjectPart != "" {
		o.SubjectPart = pointers.String(r.SubjectPart)
	}
	if r.TextPart != "" {
		o.TextPart = pointers.String(r.TextPart)
	}
	return &o
}

type TemplateMetadata struct {
	CreatedTimestamp time.Time
	Name             string
}

func newTemplateMetadata(o *SDK.TemplateMetadata) TemplateMetadata {
	result := TemplateMetadata{}
	if o == nil {
		return result
	}

	if o.CreatedTimestamp != nil {
		result.CreatedTimestamp = *o.CreatedTimestamp
	}
	if o.Name != nil {
		result.Name = *o.Name
	}
	return result
}
