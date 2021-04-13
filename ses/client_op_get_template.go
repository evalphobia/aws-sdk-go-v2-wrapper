package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetTemplate executes `GetTemplate` operation.
func (svc *SES) GetTemplate(ctx context.Context, r GetTemplateRequest) (*GetTemplateResult, error) {
	out, err := svc.RawGetTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetTemplateResult(out), nil
}

// GetTemplateRequest has parameters for `GetTemplate` operation.
type GetTemplateRequest struct {
	TemplateName string
}

func (r GetTemplateRequest) ToInput() *SDK.GetTemplateInput {
	in := &SDK.GetTemplateInput{}

	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	return in
}

type GetTemplateResult struct {
	Template Template
}

func NewGetTemplateResult(o *SDK.GetTemplateResponse) *GetTemplateResult {
	result := &GetTemplateResult{}
	if o == nil {
		return result
	}

	result.Template = newTemplate(o.Template)
	return result
}
