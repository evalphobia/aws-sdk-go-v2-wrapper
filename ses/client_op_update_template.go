package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// UpdateTemplate executes `UpdateTemplate` operation.
func (svc *SES) UpdateTemplate(ctx context.Context, r UpdateTemplateRequest) (*UpdateTemplateResult, error) {
	out, err := svc.RawUpdateTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "UpdateTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewUpdateTemplateResult(out), nil
}

// UpdateTemplateRequest has parameters for `UpdateTemplate` operation.
type UpdateTemplateRequest struct {
	Template Template
}

func (r UpdateTemplateRequest) ToInput() *SDK.UpdateTemplateInput {
	in := &SDK.UpdateTemplateInput{}

	in.Template = r.Template.ToSDK()
	return in
}

type UpdateTemplateResult struct {
}

func NewUpdateTemplateResult(o *SDK.UpdateTemplateResponse) *UpdateTemplateResult {
	result := &UpdateTemplateResult{}
	if o == nil {
		return result
	}

	return result
}
