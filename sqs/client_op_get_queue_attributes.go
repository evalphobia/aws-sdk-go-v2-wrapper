package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetQueueAttributes executes `GetQueueAttributes` operation.
func (svc *SQS) GetQueueAttributes(ctx context.Context, r GetQueueAttributesRequest) (*GetQueueAttributesResult, error) {
	out, err := svc.RawGetQueueAttributes(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetQueueAttributes",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetQueueAttributesResult(out), nil
}

// GetQueueAttributesRequest has parameters for `GetQueueAttributes` operation.
type GetQueueAttributesRequest struct {
	QueueURL string

	// optional
	AttributeNames []QueueAttributeName
}

func (r GetQueueAttributesRequest) ToInput() *SDK.GetQueueAttributesInput {
	in := &SDK.GetQueueAttributesInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}

	if len(r.AttributeNames) != 0 {
		list := make([]SDK.QueueAttributeName, len(r.AttributeNames))
		for i, v := range r.AttributeNames {
			list[i] = SDK.QueueAttributeName(v)
		}
		in.AttributeNames = list
	}
	return in
}

type GetQueueAttributesResult struct {
	Attributes map[string]string
}

func NewGetQueueAttributesResult(o *SDK.GetQueueAttributesResponse) *GetQueueAttributesResult {
	result := &GetQueueAttributesResult{}
	if o == nil {
		return result
	}

	result.Attributes = o.Attributes
	return result
}
