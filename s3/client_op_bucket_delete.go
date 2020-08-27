package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteBucket executes `DeleteBucket` operation.
func (svc *S3) DeleteBucket(ctx context.Context, r DeleteBucketRequest) error {
	_, err := svc.RawDeleteBucket(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteBucket",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// DeleteBucketRequest has parameters for `DeleteBucket` operation.
type DeleteBucketRequest struct {
	Bucket string
}

func (r DeleteBucketRequest) ToInput() *SDK.DeleteBucketInput {
	in := &SDK.DeleteBucketInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}
