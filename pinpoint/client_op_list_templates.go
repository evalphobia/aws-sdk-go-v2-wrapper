package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListTemplates executes `ListTemplates` operation.
func (svc *Pinpoint) ListTemplates(ctx context.Context, r ListTemplatesRequest) (*ListTemplatesResult, error) {
	out, err := svc.RawListTemplates(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListTemplates",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListTemplatesResult(out), nil
}

// ListTemplatesRequest has parameters for `ListTemplates` operation.
type ListTemplatesRequest struct {
	NextToken    string
	PageSize     string
	Prefix       string
	TemplateType string
}

func (r ListTemplatesRequest) ToInput() *SDK.ListTemplatesInput {
	in := &SDK.ListTemplatesInput{}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	if r.PageSize != "" {
		in.PageSize = pointers.String(r.PageSize)
	}
	if r.Prefix != "" {
		in.Prefix = pointers.String(r.Prefix)
	}
	if r.TemplateType != "" {
		in.TemplateType = pointers.String(r.TemplateType)
	}
	return in
}

type ListTemplatesResult struct {
	TemplatesResponse
}

func NewListTemplatesResult(o *SDK.ListTemplatesResponse) *ListTemplatesResult {
	result := &ListTemplatesResult{}
	if o == nil {
		return result
	}

	result.TemplatesResponse = newTemplatesResponse(o.TemplatesResponse)
	return result
}
