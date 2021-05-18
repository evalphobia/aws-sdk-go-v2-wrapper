package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ReceiveMessage executes `ReceiveMessage` operation.
func (svc *SQS) ReceiveMessage(ctx context.Context, r ReceiveMessageRequest) (*ReceiveMessageResult, error) {
	out, err := svc.RawReceiveMessage(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ReceiveMessage",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewReceiveMessageResult(out), nil
}

// ReceiveMessageRequest has parameters for `ReceiveMessage` operation.
type ReceiveMessageRequest struct {
	QueueURL string

	AttributeNames          []QueueAttributeName
	MaxNumberOfMessages     int64
	MessageAttributeNames   []string
	ReceiveRequestAttemptID string
	VisibilityTimeout       int64
	WaitTimeSeconds         int64

	// set if you want to force to use short polling
	UseShortPolling bool
}

func (r ReceiveMessageRequest) ToInput() *SDK.ReceiveMessageInput {
	in := &SDK.ReceiveMessageInput{}

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
	if r.MaxNumberOfMessages != 0 {
		in.MaxNumberOfMessages = pointers.Long64(r.MaxNumberOfMessages)
	}
	r.MessageAttributeNames = in.MessageAttributeNames
	if r.ReceiveRequestAttemptID != "" {
		in.ReceiveRequestAttemptId = pointers.String(r.ReceiveRequestAttemptID)
	}
	if r.VisibilityTimeout != 0 {
		in.VisibilityTimeout = pointers.Long64(r.VisibilityTimeout)
	}
	if r.WaitTimeSeconds != 0 {
		in.WaitTimeSeconds = pointers.Long64(r.WaitTimeSeconds)
	}
	if r.UseShortPolling {
		in.WaitTimeSeconds = pointers.Long64(0)
	}
	return in
}

type ReceiveMessageResult struct {
	Messages []Message
}

func NewReceiveMessageResult(o *SDK.ReceiveMessageResponse) *ReceiveMessageResult {
	result := &ReceiveMessageResult{}
	if o == nil {
		return result
	}

	if len(o.Messages) != 0 {
		list := make([]Message, len(o.Messages))
		for i, v := range o.Messages {
			list[i] = newMessage(v)
		}
		result.Messages = list
	}
	return result
}
