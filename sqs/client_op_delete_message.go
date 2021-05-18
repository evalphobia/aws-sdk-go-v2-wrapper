package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteMessage executes `DeleteMessage` operation.
func (svc *SQS) DeleteMessage(ctx context.Context, r DeleteMessageRequest) (*DeleteMessageResult, error) {
	out, err := svc.RawDeleteMessage(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteMessage",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteMessageResult(out), nil
}

// DeleteMessageRequest has parameters for `DeleteMessage` operation.
type DeleteMessageRequest struct {
	QueueURL      string
	ReceiptHandle string
}

func (r DeleteMessageRequest) ToInput() *SDK.DeleteMessageInput {
	in := &SDK.DeleteMessageInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	if r.ReceiptHandle != "" {
		in.ReceiptHandle = pointers.String(r.ReceiptHandle)
	}
	return in
}

type DeleteMessageResult struct{}

func NewDeleteMessageResult(o *SDK.DeleteMessageResponse) *DeleteMessageResult {
	result := &DeleteMessageResult{}
	if o == nil {
		return result
	}

	return result
}
