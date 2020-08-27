package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// ListBuckets executes `ListBuckets` operation.
func (svc *S3) ListBuckets(ctx context.Context) (*ListBucketsResult, error) {
	out, err := svc.RawListBuckets(ctx, &SDK.ListBucketsInput{})
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListBuckets",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListBucketsResult(out), nil
}

// ListBucketsResult contains results from `ListBuckets` operation.
type ListBucketsResult struct {
	Buckets []Bucket
	Owner   Owner
}

func NewListBucketsResult(output *SDK.ListBucketsResponse) *ListBucketsResult {
	r := &ListBucketsResult{}
	if output == nil {
		return r
	}

	r.Owner = newOwner(output.Owner)
	r.Buckets = newBuckets(output.Buckets)
	return r
}
