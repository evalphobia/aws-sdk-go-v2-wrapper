package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteQueue executes `DeleteQueue` operation.
func (svc *SQS) DeleteQueue(ctx context.Context, r DeleteQueueRequest) (*DeleteQueueResult, error) {
	out, err := svc.RawDeleteQueue(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteQueue",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteQueueResult(out), nil
}

// DeleteQueueRequest has parameters for `DeleteQueue` operation.
type DeleteQueueRequest struct {
	QueueURL string
}

func (r DeleteQueueRequest) ToInput() *SDK.DeleteQueueInput {
	in := &SDK.DeleteQueueInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	return in
}

type DeleteQueueResult struct{}

func NewDeleteQueueResult(o *SDK.DeleteQueueResponse) *DeleteQueueResult {
	result := &DeleteQueueResult{}
	if o == nil {
		return result
	}

	return result
}
