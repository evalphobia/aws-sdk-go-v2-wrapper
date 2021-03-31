package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendBulkTemplatedEmail executes `SendBulkTemplatedEmail` operation.
func (svc *SES) SendBulkTemplatedEmail(ctx context.Context, r SendBulkTemplatedEmailRequest) (*SendBulkTemplatedEmailResult, error) {
	out, err := svc.RawSendBulkTemplatedEmail(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SendBulkTemplatedEmail",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSendBulkTemplatedEmailResult(out), nil
}

// SendBulkTemplatedEmailRequest has parameters for `SendBulkTemplatedEmail` operation.
type SendBulkTemplatedEmailRequest struct {
	Destinations []BulkEmailDestination
	Source       string
	Template     string

	// optional
	ConfigurationSetName string
	DefaultTags          []MessageTag
	DefaultTemplateData  string
	ReplyToAddresses     []string
	ReturnPath           string
	ReturnPathARN        string
	SourceARN            string
	TemplateARN          string
}

func (r SendBulkTemplatedEmailRequest) ToInput() *SDK.SendBulkTemplatedEmailInput {
	in := &SDK.SendBulkTemplatedEmailInput{}

	if len(r.Destinations) != 0 {
		list := make([]SDK.BulkEmailDestination, len(r.Destinations))
		for i, v := range r.Destinations {
			list[i] = v.ToSDK()
		}
		in.Destinations = list
	}

	if r.Source != "" {
		in.Source = pointers.String(r.Source)
	}
	if r.Template != "" {
		in.Template = pointers.String(r.Template)
	}

	if r.ConfigurationSetName != "" {
		in.ConfigurationSetName = pointers.String(r.ConfigurationSetName)
	}

	if len(r.DefaultTags) != 0 {
		list := make([]SDK.MessageTag, len(r.DefaultTags))
		for i, v := range r.DefaultTags {
			list[i] = v.ToSDK()
		}
		in.DefaultTags = list
	}

	if r.DefaultTemplateData != "" {
		in.DefaultTemplateData = pointers.String(r.DefaultTemplateData)
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
	if r.TemplateARN != "" {
		in.TemplateArn = pointers.String(r.TemplateARN)
	}
	return in
}

type SendBulkTemplatedEmailResult struct {
	Status []BulkEmailDestinationStatus
}

func NewSendBulkTemplatedEmailResult(o *SDK.SendBulkTemplatedEmailResponse) *SendBulkTemplatedEmailResult {
	result := &SendBulkTemplatedEmailResult{}
	if o == nil {
		return result
	}

	result.Status = newBulkEmailDestinationStatuses(o.Status)
	return result
}
