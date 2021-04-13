package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// CreateTemplate executes `CreateTemplate` operation.
func (svc *SES) CreateTemplate(ctx context.Context, r CreateTemplateRequest) (*CreateTemplateResult, error) {
	out, err := svc.RawCreateTemplate(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateTemplate",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateTemplateResult(out), nil
}

// CreateTemplateRequest has parameters for `CreateTemplate` operation.
type CreateTemplateRequest struct {
	Template Template
}

func (r CreateTemplateRequest) ToInput() *SDK.CreateTemplateInput {
	in := &SDK.CreateTemplateInput{}

	in.Template = r.Template.ToSDK()
	return in
}

type CreateTemplateResult struct{}

func NewCreateTemplateResult(o *SDK.CreateTemplateResponse) *CreateTemplateResult {
	result := &CreateTemplateResult{}
	if o == nil {
		return result
	}

	return result
}
