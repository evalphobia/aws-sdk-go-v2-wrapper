package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutBucketTagging executes `PutBucketTagging` operation.
func (svc *S3) PutBucketTagging(ctx context.Context, r PutBucketTaggingRequest) error {
	_, err := svc.RawPutBucketTagging(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutBucketTagging",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// PutBucketTaggingRequest has parameters for `PutBucketTagging` operation.
type PutBucketTaggingRequest struct {
	Bucket        string
	TaggingTagSet []Tag
}

func (r PutBucketTaggingRequest) ToInput() *SDK.PutBucketTaggingInput {
	in := &SDK.PutBucketTaggingInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if len(r.TaggingTagSet) != 0 {
		list := make([]SDK.Tag, len(r.TaggingTagSet))
		for i, v := range r.TaggingTagSet {
			list[i] = v.ToSDK()
		}
		in.Tagging = &SDK.Tagging{
			TagSet: list,
		}
	}
	return in
}
