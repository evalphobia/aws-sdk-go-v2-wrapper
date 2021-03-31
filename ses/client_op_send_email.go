package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendEmail executes `SendEmail` operation.
func (svc *SES) SendEmail(ctx context.Context, r SendEmailRequest) (*SendEmailResult, error) {
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
	Destination Destination
	Source      string

	Subject        string
	SubjectCharset string // optional

	// Either HTMLBody or TextBody is requred
	HTMLBody    string
	HTMLCharset string
	TextBody    string
	TextCharset string

	// optional
	ConfigurationSetName string
	ReplyToAddresses     []string
	ReturnPath           string
	ReturnPathARN        string
	SourceARN            string
	Tags                 []MessageTag
}

func (r SendEmailRequest) ToInput() *SDK.SendEmailInput {
	in := &SDK.SendEmailInput{
		Message: &SDK.Message{},
	}

	in.Destination = r.Destination.ToSDK()

	if r.Source != "" {
		in.Source = pointers.String(r.Source)
	}

	if r.Subject != "" {
		content := SDK.Content{
			Data: pointers.String(r.Subject),
		}
		if r.SubjectCharset != "" {
			content.Charset = pointers.String(r.SubjectCharset)
		}
		in.Message.Subject = &content
	}

	if r.HTMLBody != "" || r.TextBody != "" {
		body := SDK.Body{}
		if r.HTMLBody != "" {
			content := SDK.Content{}
			content.Data = pointers.String(r.HTMLBody)
			if r.HTMLCharset != "" {
				content.Charset = pointers.String(r.HTMLCharset)
			}
			body.Html = &content
		}
		if r.TextBody != "" {
			content := SDK.Content{}
			content.Data = pointers.String(r.TextBody)
			if r.TextCharset != "" {
				content.Charset = pointers.String(r.TextCharset)
			}
			body.Text = &content
		}
		in.Message.Body = &body
	}

	if r.ConfigurationSetName != "" {
		in.ConfigurationSetName = pointers.String(r.ConfigurationSetName)
	}

	in.ReplyToAddresses = r.ReplyToAddresses

	if r.ReturnPath != "" {
		in.ReturnPath = pointers.String(r.ReturnPath)
	}
	if r.ReturnPathARN != "" {
		in.ReturnPathArn = pointers.String(r.ReturnPathARN)
	}
	if r.SourceARN != "" {
		in.SourceArn = pointers.String(r.SourceARN)
	}

	if len(r.Tags) != 0 {
		list := make([]SDK.MessageTag, len(r.Tags))
		for i, v := range r.Tags {
			list[i] = v.ToSDK()
		}
		in.Tags = list
	}
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
