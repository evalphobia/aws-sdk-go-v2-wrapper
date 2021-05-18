package ses

import (
	"context"
	"encoding/json"
	"fmt"
)

// XSendBulkTemplatedEmail sends bulk template emails.
func (svc *SES) XSendBulkTemplatedEmail(ctx context.Context, r XSendBulkTemplatedEmailRequest) (*SendBulkTemplatedEmailResult, error) {
	req, err := r.ToRequest()
	if err != nil {
		return nil, err
	}
	if req.ConfigurationSetName != "" {
		req.ConfigurationSetName = svc.defaultConfigurationSet
	}

	return svc.SendBulkTemplatedEmail(ctx, req)
}

// XSendBulkTemplatedEmailEachList sends bulk template emails.
func (svc *SES) XSendBulkTemplatedEmailEachList(ctx context.Context, template, from string, to []string, templateData []map[string]interface{}, tags []map[string]string) (*SendBulkTemplatedEmailResult, error) {
	if len(templateData) != 0 && len(to) != len(templateData) {
		return nil, fmt.Errorf("list size does not mutch: to:[%d], status:[%d]", len(to), len(templateData))
	}

	req := XSendBulkTemplatedEmailRequest{
		From:     from,
		Template: template,
	}

	req.Destinations = NewXBulkEmailDestinationList(to, templateData, tags)
	return svc.XSendBulkTemplatedEmail(ctx, req)
}

// XSendBulkTemplatedEmailRequest has parameters for `XSendBulkTemplatedEmail` function.
type XSendBulkTemplatedEmailRequest struct {
	Destinations         []XBulkEmailDestination
	From                 string
	Template             string
	ConfigurationSetName string

	// optional
	DefaultTemplateData map[string]interface{}
	DefaultTags         map[string]string
}

// NewXBulkEmailDestinationList creates []XBulkEmailDestination by the email destination and other data.
func NewXBulkEmailDestinationList(to []string, templateData []map[string]interface{}, tags []map[string]string) []XBulkEmailDestination {
	list := make([]XBulkEmailDestination, len(to))
	useTemplateData := len(templateData) != 0
	useTags := len(tags) != 0

	for i := range to {
		v := XBulkEmailDestination{
			To: []string{to[i]},
		}
		if useTemplateData {
			v.TemplateData = templateData[i]
		}
		if useTags {
			v.Tags = tags[i]
		}

		list[i] = v
	}
	return list
}

// ToRequest converts to SendBulkTemplatedEmailRequest.
func (r XSendBulkTemplatedEmailRequest) ToRequest() (SendBulkTemplatedEmailRequest, error) {
	req := SendBulkTemplatedEmailRequest{
		Template:             r.Template,
		Source:               r.From,
		ConfigurationSetName: r.ConfigurationSetName,
	}

	if len(r.Destinations) != 0 {
		list := make([]BulkEmailDestination, len(r.Destinations))
		for i, v := range r.Destinations {
			vv, err := v.ToRequest()
			if err != nil {
				return req, err
			}
			list[i] = vv
		}
		req.Destinations = list
	}

	req.DefaultTemplateData = "{}"
	if len(r.DefaultTemplateData) != 0 {
		b, err := json.Marshal(r.DefaultTemplateData)
		if err != nil {
			return req, err
		}
		req.DefaultTemplateData = string(b)
	}

	if len(r.DefaultTags) != 0 {
		tags := make([]MessageTag, 0, len(r.DefaultTags))
		for k, v := range r.DefaultTags {
			tags = append(tags, MessageTag{
				Name:  k,
				Value: v,
			})
		}
		req.DefaultTags = tags
	}
	return req, nil
}

type XBulkEmailDestination struct {
	To []string

	// optional
	Cc           []string
	Bcc          []string
	TemplateData map[string]interface{}
	Tags         map[string]string
}

func (r XBulkEmailDestination) ToRequest() (BulkEmailDestination, error) {
	req := BulkEmailDestination{
		Destination: Destination{
			ToAddresses:  r.To,
			CcAddresses:  r.Cc,
			BccAddresses: r.Bcc,
		},
	}

	req.ReplacementTemplateData = "{}"
	if len(r.TemplateData) != 0 {
		b, err := json.Marshal(r.TemplateData)
		if err != nil {
			return req, err
		}
		req.ReplacementTemplateData = string(b)
	}

	if len(r.Tags) != 0 {
		tags := make([]MessageTag, 0, len(r.Tags))
		for k, v := range r.Tags {
			tags = append(tags, MessageTag{
				Name:  k,
				Value: v,
			})
		}
		req.ReplacementTags = tags
	}
	return req, nil
}
