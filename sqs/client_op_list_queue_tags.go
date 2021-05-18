package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListQueueTags executes `ListQueueTags` operation.
func (svc *SQS) ListQueueTags(ctx context.Context, r ListQueueTagsRequest) (*ListQueueTagsResult, error) {
	out, err := svc.RawListQueueTags(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListQueueTags",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListQueueTagsResult(out), nil
}

// ListQueueTagsRequest has parameters for `ListQueueTags` operation.
type ListQueueTagsRequest struct {
	QueueURL string
}

func (r ListQueueTagsRequest) ToInput() *SDK.ListQueueTagsInput {
	in := &SDK.ListQueueTagsInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	return in
}

type ListQueueTagsResult struct {
	Tags map[string]string
}

func NewListQueueTagsResult(o *SDK.ListQueueTagsResponse) *ListQueueTagsResult {
	result := &ListQueueTagsResult{}
	if o == nil {
		return result
	}

	result.Tags = o.Tags
	return result
}
