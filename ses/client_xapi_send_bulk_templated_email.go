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

	return svc.SendBulkTemplatedEmail(ctx, req)
}

// XSendBulkTemplatedEmailEachList sends bulk template emails.
func (svc *SES) XSendBulkTemplatedEmailEachList(ctx context.Context, template, from string, to []string, templateData []map[string]interface{}) (*SendBulkTemplatedEmailResult, error) {
	if len(templateData) != 0 && len(to) != len(templateData) {
		return nil, fmt.Errorf("list size does not mutch: to:[%d], status:[%d]", len(to), len(templateData))
	}

	req := XSendBulkTemplatedEmailRequest{
		From:     from,
		Template: template,
	}

	if len(to) != 0 {
		list := make([]XBulkEmailDestination, len(to))
		useTemplateData := len(templateData) != 0
		for i := range to {
			v := XBulkEmailDestination{
				To: []string{to[i]},
			}
			if useTemplateData {
				v.TemplateData = templateData[i]
			}
			list[i] = v
		}
		req.Destinations = list
	}

	return svc.XSendBulkTemplatedEmail(ctx, req)
}

// XSendBulkTemplatedEmailRequest has parameters for `XSendBulkTemplatedEmail` function.
type XSendBulkTemplatedEmailRequest struct {
	Destinations []XBulkEmailDestination
	From         string
	Template     string

	// optional
	DefaultTemplateData map[string]interface{}
}

// ToRequest converts to SendBulkTemplatedEmailRequest.
func (r XSendBulkTemplatedEmailRequest) ToRequest() (SendBulkTemplatedEmailRequest, error) {
	req := SendBulkTemplatedEmailRequest{
		Template: r.Template,
		Source:   r.From,
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
	return req, nil
}

type XBulkEmailDestination struct {
	To []string

	// optional
	Cc           []string
	Bcc          []string
	TemplateData map[string]interface{}
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
	return req, nil
}
