package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteBucketTagging executes `DeleteBucketTagging` operation.
func (svc *S3) DeleteBucketTagging(ctx context.Context, r DeleteBucketTaggingRequest) error {
	_, err := svc.RawDeleteBucketTagging(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteBucketTagging",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// DeleteBucketTaggingRequest has parameters for `DeleteBucketTagging` operation.
type DeleteBucketTaggingRequest struct {
	Bucket string
}

func (r DeleteBucketTaggingRequest) ToInput() *SDK.DeleteBucketTaggingInput {
	in := &SDK.DeleteBucketTaggingInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}
