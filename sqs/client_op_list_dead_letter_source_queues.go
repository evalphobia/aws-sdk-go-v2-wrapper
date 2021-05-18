package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListDeadLetterSourceQueues executes `ListDeadLetterSourceQueues` operation.
func (svc *SQS) ListDeadLetterSourceQueues(ctx context.Context, r ListDeadLetterSourceQueuesRequest) (*ListDeadLetterSourceQueuesResult, error) {
	out, err := svc.RawListDeadLetterSourceQueues(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListDeadLetterSourceQueues",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListDeadLetterSourceQueuesResult(out), nil
}

// ListDeadLetterSourceQueuesRequest has parameters for `ListDeadLetterSourceQueues` operation.
type ListDeadLetterSourceQueuesRequest struct {
	QueueURL string

	// optional
	MaxResults int64
	NextToken  string
}

func (r ListDeadLetterSourceQueuesRequest) ToInput() *SDK.ListDeadLetterSourceQueuesInput {
	in := &SDK.ListDeadLetterSourceQueuesInput{}

	if r.QueueURL != "" {
		in.QueueUrl = pointers.String(r.QueueURL)
	}

	if r.MaxResults != 0 {
		in.MaxResults = pointers.Long64(r.MaxResults)
	}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	return in
}

type ListDeadLetterSourceQueuesResult struct {
	NextToken string
	QueueURLs []string
}

func NewListDeadLetterSourceQueuesResult(o *SDK.ListDeadLetterSourceQueuesResponse) *ListDeadLetterSourceQueuesResult {
	result := &ListDeadLetterSourceQueuesResult{}
	if o == nil {
		return result
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}
	result.QueueURLs = o.QueueUrls
	return result
}
