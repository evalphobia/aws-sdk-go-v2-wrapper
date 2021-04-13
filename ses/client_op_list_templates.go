package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListTemplates executes `ListTemplates` operation.
func (svc *SES) ListTemplates(ctx context.Context, r ListTemplatesRequest) (*ListTemplatesResult, error) {
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
	MaxItems  int64
	NextToken string
}

func (r ListTemplatesRequest) ToInput() *SDK.ListTemplatesInput {
	in := &SDK.ListTemplatesInput{}

	if r.MaxItems != 0 {
		in.MaxItems = pointers.Long64(r.MaxItems)
	}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	return in
}

type ListTemplatesResult struct {
	NextToken         string
	TemplatesMetadata []TemplateMetadata
}

func NewListTemplatesResult(o *SDK.ListTemplatesResponse) *ListTemplatesResult {
	result := &ListTemplatesResult{}
	if o == nil {
		return result
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}

	if len(o.TemplatesMetadata) != 0 {
		list := make([]TemplateMetadata, len(o.TemplatesMetadata))
		for i, v := range o.TemplatesMetadata {
			v := v
			list[i] = newTemplateMetadata(&v)
		}
		result.TemplatesMetadata = list
	}
	return result
}
