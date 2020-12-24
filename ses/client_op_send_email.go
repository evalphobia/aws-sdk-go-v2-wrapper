package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendRawEmail executes `SendRawEmail` operation.
func (svc *SES) SendRawEmail(ctx context.Context, r SendRawEmailRequest) (*SendRawEmailResult, error) {
	out, err := svc.RawSendRawEmail(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SendRawEmail",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSendRawEmailResult(out), nil
}

// SendRawEmailRequest has parameters for `SendRawEmail` operation.
type SendRawEmailRequest struct {
	RawMessageData []byte

	// optional
	ConfigurationSetName string
	Destinations         []string
	FromArn              string
	ReturnPathArn        string
	Source               string
	SourceArn            string
	Tags                 []MessageTag
}

func (r SendRawEmailRequest) ToInput() *SDK.SendRawEmailInput {
	in := &SDK.SendRawEmailInput{}

	if len(r.RawMessageData) != 0 {
		in.RawMessage = &SDK.RawMessage{
			Data: r.RawMessageData,
		}
	}

	if r.ConfigurationSetName != "" {
		in.ConfigurationSetName = pointers.String(r.ConfigurationSetName)
	}
	if r.FromArn != "" {
		in.FromArn = pointers.String(r.FromArn)
	}
	if r.ReturnPathArn != "" {
		in.ReturnPathArn = pointers.String(r.ReturnPathArn)
	}
	if r.Source != "" {
		in.Source = pointers.String(r.Source)
	}
	if r.SourceArn != "" {
		in.SourceArn = pointers.String(r.SourceArn)
	}

	if len(r.Tags) != 0 {
		list := make([]SDK.MessageTag, len(r.Tags))
		for i, v := range r.Tags {
			list[i] = v.ToSDK()
		}
		in.Tags = list
	}
	in.Destinations = r.Destinations
	return in
}

type SendRawEmailResult struct {
	MessageID string
}

func NewSendRawEmailResult(o *SDK.SendRawEmailResponse) *SendRawEmailResult {
	result := &SendRawEmailResult{}
	if o == nil {
		return result
	}

	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}
	return result
}
