package pinpointemail

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpointemail"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendEmail executes `SendEmail` operation.
func (svc *PinpointEmail) SendEmail(ctx context.Context, r SendEmailRequest) (*SendEmailResult, error) {
	out, err := svc.RawSendEmail(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SendEmail",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSendEmailResult(out), nil
}

// SendEmailRequest has parameters for `SendEmail` operation.
type SendEmailRequest struct {
	To   []string
	Cc   []string
	Bcc  []string
	From string

	// for text or html
	Body           string
	BodyCharset    string
	Subject        string
	SubjectCharset string
	UseHTML        bool

	// for template
	TemplateARN  string
	TemplateData string

	// for raw
	RawMessageData []byte

	ConfigurationSetName           string
	EmailTags                      []Tag
	FeedbackForwardingEmailAddress string
	ReplyToAddresses               []string
}

func (r SendEmailRequest) ToInput() *SDK.SendEmailInput {
	in := &SDK.SendEmailInput{}
	switch {
	case len(r.To) != 0,
		len(r.Cc) != 0,
		len(r.Bcc) != 0:
		in.Destination = &SDK.Destination{
			ToAddresses:  r.To,
			CcAddresses:  r.Cc,
			BccAddresses: r.Bcc,
		}
	}
	if r.From != "" {
		in.FromEmailAddress = pointers.String(r.From)
	}

	if len(r.RawMessageData) != 0 {
		in.Content = &SDK.EmailContent{
			Raw: &SDK.RawMessage{
				Data: r.RawMessageData,
			},
		}
	}

	switch {
	case r.Subject != "",
		r.Body != "":
		if in.Content == nil {
			in.Content = &SDK.EmailContent{}
		}
		in.Content.Simple = &SDK.Message{}
	}

	if r.Subject != "" {
		content := SDK.Content{
			Data: pointers.String(r.Subject),
		}
		if r.SubjectCharset != "" {
			content.Charset = pointers.String(r.SubjectCharset)
		}
		in.Content.Simple.Subject = &content
	}

	if r.Body != "" {
		content := SDK.Content{
			Data: pointers.String(r.Body),
		}
		if r.BodyCharset != "" {
			content.Charset = pointers.String(r.BodyCharset)
		}

		body := SDK.Body{}
		switch {
		case r.UseHTML:
			body.Html = &content
		default:
			body.Text = &content
		}
		in.Content.Simple.Body = &body
	}

	if r.TemplateARN != "" {
		template := SDK.Template{
			TemplateArn: pointers.String(r.TemplateARN),
		}
		if r.TemplateData != "" {
			template.TemplateData = pointers.String(r.TemplateData)
		}
		in.Content.Template = &template
	}

	if r.ConfigurationSetName != "" {
		in.ConfigurationSetName = pointers.String(r.ConfigurationSetName)
	}
	if len(r.EmailTags) != 0 {
		list := make([]SDK.MessageTag, len(r.EmailTags))
		for i, v := range r.EmailTags {
			list[i] = v.ToEmailTag()
		}
	}
	if r.FeedbackForwardingEmailAddress != "" {
		in.FeedbackForwardingEmailAddress = pointers.String(r.FeedbackForwardingEmailAddress)
	}
	in.ReplyToAddresses = r.ReplyToAddresses
	return in
}

type SendEmailResult struct {
	MessageID string
}

func NewSendEmailResult(o *SDK.SendEmailResponse) *SendEmailResult {
	result := &SendEmailResult{}
	if o == nil {
		return result
	}

	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}
	return result
}
