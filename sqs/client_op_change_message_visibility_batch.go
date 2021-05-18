package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ChangeMessageVisibilityBatch executes `ChangeMessageVisibilityBatch` operation.
func (svc *SQS) ChangeMessageVisibilityBatch(ctx context.Context, r ChangeMessageVisibilityBatchRequest) (*ChangeMessageVisibilityBatchResult, error) {
	out, err := svc.RawChangeMessageVisibilityBatch(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ChangeMessageVisibilityBatch",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewChangeMessageVisibilityBatchResult(out), nil
}

// ChangeMessageVisibilityBatchRequest has parameters for `ChangeMessageVisibilityBatch` operation.
type ChangeMessageVisibilityBatchRequest struct {
	Entries  []ChangeMessageVisibilityBatchRequestEntry
	QueueURL string
}

func (r ChangeMessageVisibilityBatchRequest) ToInput() *SDK.ChangeMessageVisibilityBatchInput {
	in := &SDK.ChangeMessageVisibilityBatchInput{}

	if len(r.Entries) != 0 {
		list := make([]SDK.ChangeMessageVisibilityBatchRequestEntry, len(r.Entries))
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

type ChangeMessageVisibilityBatchResult struct {
	Failed     []BatchResultErrorEntry                   `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`
	Successful []ChangeMessageVisibilityBatchResultEntry `locationNameList:"ChangeMessageVisibilityBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

func NewChangeMessageVisibilityBatchResult(o *SDK.ChangeMessageVisibilityBatchResponse) *ChangeMessageVisibilityBatchResult {
	result := &ChangeMessageVisibilityBatchResult{}
	if o == nil {
		return result
	}

	result.Failed = newBatchResultErrorEntryList(o.Failed)
	result.Successful = newChangeMessageVisibilityBatchResultEntryList(o.Successful)
	return result
}
