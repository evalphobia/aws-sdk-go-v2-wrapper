package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateEmailTemplate executes `CreateEmailTemplate` operation.
func (svc *Pinpoint) CreateEmailTemplate(ctx context.Context, r CreateEmailTemplateRequest) (*CreateEmailTemplateResult, error) {
	out, err := svc.RawCreateEmailTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateEmailTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateEmailTemplateResult(out), nil
}

// CreateEmailTemplateRequest has parameters for `CreateEmailTemplate` operation.
type CreateEmailTemplateRequest struct {
	EmailTemplateRequest EmailTemplateRequest
	TemplateName         string
}

func (r CreateEmailTemplateRequest) ToInput() *SDK.CreateEmailTemplateInput {
	in := &SDK.CreateEmailTemplateInput{}

	in.EmailTemplateRequest = r.EmailTemplateRequest.ToSDK()

	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	return in
}

type CreateEmailTemplateResult struct {
	CreateTemplateMessageBody
}

func NewCreateEmailTemplateResult(o *SDK.CreateEmailTemplateResponse) *CreateEmailTemplateResult {
	result := &CreateEmailTemplateResult{}
	if o == nil {
		return result
	}

	result.CreateTemplateMessageBody = newCreateTemplateMessageBody(o.CreateTemplateMessageBody)
	return result
}
