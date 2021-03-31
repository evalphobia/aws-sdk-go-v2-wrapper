package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteEmailTemplate executes `DeleteEmailTemplate` operation.
func (svc *Pinpoint) DeleteEmailTemplate(ctx context.Context, r DeleteEmailTemplateRequest) (*DeleteEmailTemplateResult, error) {
	out, err := svc.RawDeleteEmailTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteEmailTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteEmailTemplateResult(out), nil
}

// DeleteEmailTemplateRequest has parameters for `DeleteEmailTemplate` operation.
type DeleteEmailTemplateRequest struct {
	TemplateName string

	// optional
	Version string
}

func (r DeleteEmailTemplateRequest) ToInput() *SDK.DeleteEmailTemplateInput {
	in := &SDK.DeleteEmailTemplateInput{}
	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	if r.Version != "" {
		in.Version = pointers.String(r.Version)
	}
	return in
}

type DeleteEmailTemplateResult struct {
	MessageBody
}

func NewDeleteEmailTemplateResult(o *SDK.DeleteEmailTemplateResponse) *DeleteEmailTemplateResult {
	result := &DeleteEmailTemplateResult{}
	if o == nil {
		return result
	}

	result.MessageBody = newMessageBody(o.MessageBody)
	return result
}
