package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListQueues executes `ListQueues` operation.
func (svc *SQS) ListQueues(ctx context.Context, r ListQueuesRequest) (*ListQueuesResult, error) {
	out, err := svc.RawListQueues(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListQueues",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListQueuesResult(out), nil
}

// ListQueuesRequest has parameters for `ListQueues` operation.
type ListQueuesRequest struct {
	// optional
	MaxResults      int64
	NextToken       string
	QueueNamePrefix string
}

func (r ListQueuesRequest) ToInput() *SDK.ListQueuesInput {
	in := &SDK.ListQueuesInput{}

	if r.MaxResults != 0 {
		in.MaxResults = pointers.Long64(r.MaxResults)
	}
	if r.NextToken != "" {
		in.NextToken = pointers.String(r.NextToken)
	}
	if r.QueueNamePrefix != "" {
		in.QueueNamePrefix = pointers.String(r.QueueNamePrefix)
	}
	return in
}

type ListQueuesResult struct {
	NextToken string
	QueueURLs []string
}

func NewListQueuesResult(o *SDK.ListQueuesResponse) *ListQueuesResult {
	result := &ListQueuesResult{}
	if o == nil {
		return result
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}
	result.QueueURLs = o.QueueUrls
	return result
}
