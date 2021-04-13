package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// TestRenderTemplate executes `TestRenderTemplate` operation.
func (svc *SES) TestRenderTemplate(ctx context.Context, r TestRenderTemplateRequest) (*TestRenderTemplateResult, error) {
	out, err := svc.RawTestRenderTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "TestRenderTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewTestRenderTemplateResult(out), nil
}

// TestRenderTemplateRequest has parameters for `TestRenderTemplate` operation.
type TestRenderTemplateRequest struct {
	TemplateData string
	TemplateName string
}

func (r TestRenderTemplateRequest) ToInput() *SDK.TestRenderTemplateInput {
	in := &SDK.TestRenderTemplateInput{}

	if r.TemplateData != "" {
		in.TemplateData = pointers.String(r.TemplateData)
	}
	if r.TemplateName != "" {
		in.TemplateName = pointers.String(r.TemplateName)
	}
	return in
}

type TestRenderTemplateResult struct {
	RenderedTemplate string
}

func NewTestRenderTemplateResult(o *SDK.TestRenderTemplateResponse) *TestRenderTemplateResult {
	result := &TestRenderTemplateResult{}
	if o == nil {
		return result
	}

	if o.RenderedTemplate != nil {
		result.RenderedTemplate = *o.RenderedTemplate
	}
	return result
}
