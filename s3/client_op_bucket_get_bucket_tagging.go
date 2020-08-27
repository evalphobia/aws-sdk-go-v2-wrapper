package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetBucketTagging executes `GetBucketTagging` operation.
func (svc *S3) GetBucketTagging(ctx context.Context, r GetBucketTaggingRequest) (*GetBucketTaggingResult, error) {
	out, err := svc.RawGetBucketTagging(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetBucketTagging",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetBucketTaggingResult(out), nil
}

// GetBucketTaggingRequest has parameters for `GetBucketTagging` operation.
type GetBucketTaggingRequest struct {
	Bucket string
}

func (r GetBucketTaggingRequest) ToInput() *SDK.GetBucketTaggingInput {
	in := &SDK.GetBucketTaggingInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}

// GetBucketTaggingResult contains results from `GetBucketTagging` operation.
type GetBucketTaggingResult struct {
	TagSet []Tag
}

func NewGetBucketTaggingResult(output *SDK.GetBucketTaggingResponse) *GetBucketTaggingResult {
	r := &GetBucketTaggingResult{}
	if output == nil {
		return r
	}

	r.TagSet = newTags(output.TagSet)
	return r
}
