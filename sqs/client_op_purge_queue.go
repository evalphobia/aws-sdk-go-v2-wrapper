package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PurgeQueue executes `PurgeQueue` operation.
func (svc *SQS) PurgeQueue(ctx context.Context, r PurgeQueueRequest) (*PurgeQueueResult, error) {
	out, err := svc.RawPurgeQueue(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PurgeQueue",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewPurgeQueueResult(out), nil
}

// PurgeQueueRequest has parameters for `PurgeQueue` operation.
type PurgeQueueRequest struct {
	QueueURL string
}

func (r PurgeQueueRequest) ToInput() *SDK.PurgeQueueInput {
	in := &SDK.PurgeQueueInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	return in
}

type PurgeQueueResult struct{}

func NewPurgeQueueResult(o *SDK.PurgeQueueResponse) *PurgeQueueResult {
	result := &PurgeQueueResult{}
	if o == nil {
		return result
	}

	return result
}
