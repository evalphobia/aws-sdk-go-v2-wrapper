package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteBucketPolicy executes `DeleteBucketPolicy` operation.
func (svc *S3) DeleteBucketPolicy(ctx context.Context, r DeleteBucketPolicyRequest) error {
	_, err := svc.RawDeleteBucketPolicy(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteBucketPolicy",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// DeleteBucketPolicyRequest has parameters for `DeleteBucketPolicy` operation.
type DeleteBucketPolicyRequest struct {
	Bucket string
}

func (r DeleteBucketPolicyRequest) ToInput() *SDK.DeleteBucketPolicyInput {
	in := &SDK.DeleteBucketPolicyInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}
