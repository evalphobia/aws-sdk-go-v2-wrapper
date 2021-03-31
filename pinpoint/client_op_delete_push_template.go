package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeletePushTemplate executes `DeletePushTemplate` operation.
func (svc *Pinpoint) DeletePushTemplate(ctx context.Context, r DeletePushTemplateRequest) (*DeletePushTemplateResult, error) {
	out, err := svc.RawDeletePushTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeletePushTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeletePushTemplateResult(out), nil
}

// DeletePushTemplateRequest has parameters for `DeletePushTemplate` operation.
type DeletePushTemplateRequest struct {
	TemplateName string

	// optional
	Version string
}

func (r DeletePushTemplateRequest) ToInput() *SDK.DeletePushTemplateInput {
	in := &SDK.DeletePushTemplateInput{}
	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	if r.Version != "" {
		in.Version = pointers.String(r.Version)
	}
	return in
}

type DeletePushTemplateResult struct {
	MessageBody
}

func NewDeletePushTemplateResult(o *SDK.DeletePushTemplateResponse) *DeletePushTemplateResult {
	result := &DeletePushTemplateResult{}
	if o == nil {
		return result
	}

	result.MessageBody = newMessageBody(o.MessageBody)
	return result
}
