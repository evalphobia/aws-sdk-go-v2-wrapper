package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreatePushTemplate executes `CreatePushTemplate` operation.
func (svc *Pinpoint) CreatePushTemplate(ctx context.Context, r CreatePushTemplateRequest) (*CreatePushTemplateResult, error) {
	out, err := svc.RawCreatePushTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreatePushTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreatePushTemplateResult(out), nil
}

// CreatePushTemplateRequest has parameters for `CreatePushTemplate` operation.
type CreatePushTemplateRequest struct {
	PushNotificationTemplateRequest PushNotificationTemplateRequest
	TemplateName                    string
}

func (r CreatePushTemplateRequest) ToInput() *SDK.CreatePushTemplateInput {
	in := &SDK.CreatePushTemplateInput{}

	in.PushNotificationTemplateRequest = r.PushNotificationTemplateRequest.ToSDK()

	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	return in
}

type CreatePushTemplateResult struct {
	CreateTemplateMessageBody
}

func NewCreatePushTemplateResult(o *SDK.CreatePushTemplateResponse) *CreatePushTemplateResult {
	result := &CreatePushTemplateResult{}
	if o == nil {
		return result
	}

	result.CreateTemplateMessageBody = newCreateTemplateMessageBody(o.CreateTemplateMessageBody)
	return result
}
