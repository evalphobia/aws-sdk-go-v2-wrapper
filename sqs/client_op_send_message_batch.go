package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// SendMessageBatch executes `SendMessageBatch` operation.
func (svc *SQS) SendMessageBatch(ctx context.Context, r SendMessageBatchRequest) (*SendMessageBatchResult, error) {
	out, err := svc.RawSendMessageBatch(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "SendMessageBatch",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewSendMessageBatchResult(out), nil
}

// SendMessageBatchRequest has parameters for `SendMessageBatch` operation.
type SendMessageBatchRequest struct {
	Entries  []SendMessageBatchRequestEntry
	QueueURL string
}

func (r SendMessageBatchRequest) ToInput() *SDK.SendMessageBatchInput {
	in := &SDK.SendMessageBatchInput{}

	if len(r.Entries) != 0 {
		list := make([]SDK.SendMessageBatchRequestEntry, len(r.Entries))
		for i, v := range r.Entries {
			list[i] = v.ToSDK()
		}
		in.Entries = list
	}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	return in
}

type SendMessageBatchResult struct {
	Failed     []BatchResultErrorEntry
	Successful []SendMessageBatchResultEntry
}

func NewSendMessageBatchResult(o *SDK.SendMessageBatchResponse) *SendMessageBatchResult {
	result := &SendMessageBatchResult{}
	if o == nil {
		return result
	}

	result.Failed = newBatchResultErrorEntryList(o.Failed)
	result.Successful = newSendMessageBatchResultEntryList(o.Successful)
	return result
}
