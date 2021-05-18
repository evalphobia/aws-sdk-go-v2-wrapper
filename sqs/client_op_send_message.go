package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendMessage executes `SendMessage` operation.
func (svc *SQS) SendMessage(ctx context.Context, r SendMessageRequest) (*SendMessageResult, error) {
	out, err := svc.RawSendMessage(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SendMessage",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSendMessageResult(out), nil
}

// SendMessageRequest has parameters for `SendMessage` operation.
type SendMessageRequest struct {
	MessageBody string
	QueueURL    string

	// for FIFO
	MessageDeduplicationID string
	MessageGroupID         string

	DelaySeconds            int64
	MessageAttributes       map[string]MessageAttributeValue
	MessageSystemAttributes map[string]MessageSystemAttributeValue
}

func (r SendMessageRequest) ToInput() *SDK.SendMessageInput {
	in := &SDK.SendMessageInput{}

	if r.MessageBody != "" {
		in.MessageBody = pointers.String(r.MessageBody)
	}
	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}

	if r.MessageDeduplicationID != "" {
		in.MessageDeduplicationId = pointers.String(r.MessageDeduplicationID)
	}
	if r.MessageGroupID != "" {
		in.MessageGroupId = pointers.String(r.MessageGroupID)
	}

	if r.DelaySeconds != 0 {
		in.DelaySeconds = pointers.Long64(r.DelaySeconds)
	}
	if len(r.MessageAttributes) != 0 {
		m := make(map[string]SDK.MessageAttributeValue)
		for k, v := range r.MessageAttributes {
			m[k] = v.ToSDK()
		}
		in.MessageAttributes = m
	}
	if len(r.MessageSystemAttributes) != 0 {
		m := make(map[string]SDK.MessageSystemAttributeValue)
		for k, v := range r.MessageSystemAttributes {
			m[k] = v.ToSDK()
		}
		in.MessageSystemAttributes = m
	}
	return in
}

type SendMessageResult struct {
	MD5OfMessageAttributes       string
	MD5OfMessageBody             string
	MD5OfMessageSystemAttributes string
	MessageID                    string

	// for FIFO
	SequenceNumber string
}

func NewSendMessageResult(o *SDK.SendMessageResponse) *SendMessageResult {
	result := &SendMessageResult{}
	if o == nil {
		return result
	}

	if o.MD5OfMessageAttributes != nil {
		result.MD5OfMessageAttributes = *o.MD5OfMessageAttributes
	}
	if o.MD5OfMessageBody != nil {
		result.MD5OfMessageBody = *o.MD5OfMessageBody
	}
	if o.MD5OfMessageSystemAttributes != nil {
		result.MD5OfMessageSystemAttributes = *o.MD5OfMessageSystemAttributes
	}
	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}
	if o.SequenceNumber != nil {
		result.SequenceNumber = *o.SequenceNumber
	}
	return result
}
