package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendTemplatedEmail executes `SendTemplatedEmail` operation.
func (svc *SES) SendTemplatedEmail(ctx context.Context, r SendTemplatedEmailRequest) (*SendTemplatedEmailResult, error) {
	out, err := svc.RawSendTemplatedEmail(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SendTemplatedEmail",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSendTemplatedEmailResult(out), nil
}

// SendTemplatedEmailRequest has parameters for `SendTemplatedEmail` operation.
type SendTemplatedEmailRequest struct {
	Destination  Destination
	Source       string
	Template     string
	TemplateData string

	ConfigurationSetName string
	ReplyToAddresses     []string
	ReturnPath           string
	ReturnPathARN        string
	SourceARN            string
	Tags                 []MessageTag
	TemplateARN          string
}

func (r SendTemplatedEmailRequest) ToInput() *SDK.SendTemplatedEmailInput {
	in := &SDK.SendTemplatedEmailInput{}

	in.Destination = r.Destination.ToSDK()

	if r.Source != "" {
		in.Source = pointers.String(r.Source)
	}
	if r.Template != "" {
		in.Template = pointers.String(r.Template)
	}
	if r.TemplateData != "" {
		in.TemplateData = pointers.String(r.TemplateData)
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

	if r.TemplateARN != "" {
		in.TemplateArn = pointers.String(r.TemplateARN)
	}
	return in
}

type SendTemplatedEmailResult struct {
	MessageID string
}

func NewSendTemplatedEmailResult(o *SDK.SendTemplatedEmailResponse) *SendTemplatedEmailResult {
	result := &SendTemplatedEmailResult{}
	if o == nil {
		return result
	}

	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}
	return result
}
