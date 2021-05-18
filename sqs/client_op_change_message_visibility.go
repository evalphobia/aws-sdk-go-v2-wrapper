package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ChangeMessageVisibility executes `ChangeMessageVisibility` operation.
func (svc *SQS) ChangeMessageVisibility(ctx context.Context, r ChangeMessageVisibilityRequest) (*ChangeMessageVisibilityResult, error) {
	out, err := svc.RawChangeMessageVisibility(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ChangeMessageVisibility",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewChangeMessageVisibilityResult(out), nil
}

// ChangeMessageVisibilityRequest has parameters for `ChangeMessageVisibility` operation.
type ChangeMessageVisibilityRequest struct {
	QueueURL          string
	ReceiptHandle     string
	VisibilityTimeout int64
}

func (r ChangeMessageVisibilityRequest) ToInput() *SDK.ChangeMessageVisibilityInput {
	in := &SDK.ChangeMessageVisibilityInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}
	if r.ReceiptHandle != "" {
		in.ReceiptHandle = pointers.String(r.ReceiptHandle)
	}
	in.VisibilityTimeout = pointers.Long64(r.VisibilityTimeout)
	return in
}

type ChangeMessageVisibilityResult struct{}

func NewChangeMessageVisibilityResult(o *SDK.ChangeMessageVisibilityResponse) *ChangeMessageVisibilityResult {
	result := &ChangeMessageVisibilityResult{}
	if o == nil {
		return result
	}

	return result
}
