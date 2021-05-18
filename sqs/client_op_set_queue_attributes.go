package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SetQueueAttributes executes `SetQueueAttributes` operation.
func (svc *SQS) SetQueueAttributes(ctx context.Context, r SetQueueAttributesRequest) (*SetQueueAttributesResult, error) {
	out, err := svc.RawSetQueueAttributes(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SetQueueAttributes",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSetQueueAttributesResult(out), nil
}

// SetQueueAttributesRequest has parameters for `SetQueueAttributes` operation.
type SetQueueAttributesRequest struct {
	Attributes map[string]string
	QueueURL   string
}

func (r SetQueueAttributesRequest) ToInput() *SDK.SetQueueAttributesInput {
	in := &SDK.SetQueueAttributesInput{}

	in.Attributes = r.Attributes
	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	return in
}

type SetQueueAttributesResult struct{}

func NewSetQueueAttributesResult(o *SDK.SetQueueAttributesResponse) *SetQueueAttributesResult {
	result := &SetQueueAttributesResult{}
	if o == nil {
		return result
	}

	return result
}
