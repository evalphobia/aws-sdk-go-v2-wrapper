package ses

import (
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
