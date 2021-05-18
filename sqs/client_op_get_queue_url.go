package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetQueueURL executes `GetQueueURL` operation.
func (svc *SQS) GetQueueURL(ctx context.Context, r GetQueueURLRequest) (*GetQueueURLResult, error) {
	out, err := svc.RawGetQueueUrl(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetQueueURL",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetQueueURLResult(out), nil
}

// GetQueueURLRequest has parameters for `GetQueueURL` operation.
type GetQueueURLRequest struct {
	QueueName string

	// optional
	QueueOwnerAWSAccountID string
}

func (r GetQueueURLRequest) ToInput() *SDK.GetQueueUrlInput {
	in := &SDK.GetQueueUrlInput{}

	if r.QueueName != "" {
		in.QueueName = pointers.String(r.QueueName)
	}

	if r.QueueOwnerAWSAccountID != "" {
		in.QueueOwnerAWSAccountId = pointers.String(r.QueueOwnerAWSAccountID)
	}
	return in
}

type GetQueueURLResult struct {
	QueueURL string
}

func NewGetQueueURLResult(o *SDK.GetQueueUrlResponse) *GetQueueURLResult {
	result := &GetQueueURLResult{}
	if o == nil {
		return result
	}

	if o.QueueUrl != nil {
		result.QueueURL = *o.QueueUrl
	}
	return result
}
