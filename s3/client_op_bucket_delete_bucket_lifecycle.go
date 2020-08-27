package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteBucketLifecycle executes `DeleteBucketLifecycle` operation.
func (svc *S3) DeleteBucketLifecycle(ctx context.Context, r DeleteBucketLifecycleRequest) error {
	_, err := svc.RawDeleteBucketLifecycle(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteBucketLifecycle",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// DeleteBucketLifecycleRequest has parameters for `DeleteBucketLifecycle` operation.
type DeleteBucketLifecycleRequest struct {
	Bucket string
}

func (r DeleteBucketLifecycleRequest) ToInput() *SDK.DeleteBucketLifecycleInput {
	in := &SDK.DeleteBucketLifecycleInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}
