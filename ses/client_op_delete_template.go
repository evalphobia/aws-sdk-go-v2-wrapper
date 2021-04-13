package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteTemplate executes `DeleteTemplate` operation.
func (svc *SES) DeleteTemplate(ctx context.Context, r DeleteTemplateRequest) (*DeleteTemplateResult, error) {
	out, err := svc.RawDeleteTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteTemplateResult(out), nil
}

// DeleteTemplateRequest has parameters for `DeleteTemplate` operation.
type DeleteTemplateRequest struct {
	TemplateName string
}

func (r DeleteTemplateRequest) ToInput() *SDK.DeleteTemplateInput {
	in := &SDK.DeleteTemplateInput{}

	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	return in
}

type DeleteTemplateResult struct{}

func NewDeleteTemplateResult(o *SDK.DeleteTemplateResponse) *DeleteTemplateResult {
	result := &DeleteTemplateResult{}
	if o == nil {
		return result
	}

	return result
}
