package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateQueue executes `CreateQueue` operation.
func (svc *SQS) CreateQueue(ctx context.Context, r CreateQueueRequest) (*CreateQueueResult, error) {
	out, err := svc.RawCreateQueue(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateQueue",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateQueueResult(out), nil
}

// CreateQueueRequest has parameters for `CreateQueue` operation.
type CreateQueueRequest struct {
	QueueName string

	// optional
	Attributes map[string]string
	Tags       map[string]string
}

func (r CreateQueueRequest) ToInput() *SDK.CreateQueueInput {
	in := &SDK.CreateQueueInput{}

	if r.QueueName != "" {
		in.QueueName = pointers.String(r.QueueName)
	}

	in.Attributes = r.Attributes
	in.Tags = r.Tags
	return in
}

type CreateQueueResult struct {
	QueueURL string
}

func NewCreateQueueResult(o *SDK.CreateQueueResponse) *CreateQueueResult {
	result := &CreateQueueResult{}
	if o == nil {
		return result
	}

	if o.QueueUrl != nil {
		result.QueueURL = *o.QueueUrl
	}
	return result
}
