package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteMessageBatch executes `DeleteMessageBatch` operation.
func (svc *SQS) DeleteMessageBatch(ctx context.Context, r DeleteMessageBatchRequest) (*DeleteMessageBatchResult, error) {
	out, err := svc.RawDeleteMessageBatch(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteMessageBatch",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteMessageBatchResult(out), nil
}

// DeleteMessageBatchRequest has parameters for `DeleteMessageBatch` operation.
type DeleteMessageBatchRequest struct {
	Entries  []DeleteMessageBatchRequestEntry
	QueueURL string
}

func (r DeleteMessageBatchRequest) ToInput() *SDK.DeleteMessageBatchInput {
	in := &SDK.DeleteMessageBatchInput{}

	if len(r.Entries) != 0 {
		list := make([]SDK.DeleteMessageBatchRequestEntry, len(r.Entries))
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

type DeleteMessageBatchResult struct {
	Failed     []BatchResultErrorEntry
	Successful []DeleteMessageBatchResultEntry
}

func NewDeleteMessageBatchResult(o *SDK.DeleteMessageBatchResponse) *DeleteMessageBatchResult {
	result := &DeleteMessageBatchResult{}
	if o == nil {
		return result
	}

	result.Failed = newBatchResultErrorEntryList(o.Failed)
	result.Successful = newDeleteMessageBatchResultEntryList(o.Successful)
	return result
}
