package ses

import (
	"context"
	"encoding/json"
)

// XSendTemplatedEmail sends a template emails.
func (svc *SES) XSendTemplatedEmail(ctx context.Context, r XSendTemplatedEmailRequest) (*SendTemplatedEmailResult, error) {
	req, err := r.ToRequest()
	if err != nil {
		return nil, err
	}
	if req.ConfigurationSetName != "" {
		req.ConfigurationSetName = svc.defaultConfigurationSet
	}

	return svc.SendTemplatedEmail(ctx, req)
}

// XSendTemplatedEmailRequest has parameters for `XSendTemplatedEmail` function.
type XSendTemplatedEmailRequest struct {
	Destination          Destination
	From                 string
	Template             string
	ConfigurationSetName string

	// optional
	TemplateData map[string]interface{}
	Tags         map[string]string
}

// ToRequest converts to SendTemplatedEmailRequest.
func (r XSendTemplatedEmailRequest) ToRequest() (SendTemplatedEmailRequest, error) {
	req := SendTemplatedEmailRequest{
		Destination:          r.Destination,
		Template:             r.Template,
		Source:               r.From,
		ConfigurationSetName: r.ConfigurationSetName,
	}

	req.TemplateData = "{}"
	if len(r.TemplateData) != 0 {
		b, err := json.Marshal(r.TemplateData)
		if err != nil {
			return req, err
		}
		req.TemplateData = string(b)
	}

	if len(r.Tags) != 0 {
		tags := make([]MessageTag, 0, len(r.Tags))
		for k, v := range r.Tags {
			tags = append(tags, MessageTag{
				Name:  k,
				Value: v,
			})
		}
		req.Tags = tags
	}
	return req, nil
}
