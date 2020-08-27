package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// HeadBucket executes `HeadBucket` operation.
func (svc *S3) HeadBucket(ctx context.Context, r HeadBucketRequest) (bool, error) {
	_, err := svc.RawHeadBucket(ctx, r.ToInput())
	if err == nil {
		return true, nil
	}
	if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == "NotFound" {
			return false, nil
		}
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "HeadBucket",
	})
	svc.Errorf(err.Error())
	return false, err
}

// HeadBucketRequest has parameters for `HeadBucket` operation.
type HeadBucketRequest struct {
	Bucket string
}

func (r HeadBucketRequest) ToInput() *SDK.HeadBucketInput {
	in := &SDK.HeadBucketInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}
