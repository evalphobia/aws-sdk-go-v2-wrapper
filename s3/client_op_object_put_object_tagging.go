package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutObjectTagging executes `PutObjectTagging` operation.
func (svc *S3) PutObjectTagging(ctx context.Context, r PutObjectTaggingRequest) (*PutObjectTaggingResult, error) {
	out, err := svc.RawPutObjectTagging(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutObjectTagging",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewPutObjectTaggingResult(out), nil
}

// PutObjectTaggingRequest has parameters for `PutObjectTagging` operation.
type PutObjectTaggingRequest struct {
	Bucket        string
	Key           string
	TaggingTagSet []Tag

	// optional
	VersionID string
}

func (r PutObjectTaggingRequest) ToInput() *SDK.PutObjectTaggingInput {
	in := &SDK.PutObjectTaggingInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}
	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
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

// PutObjectTaggingResult contains results from `PutObjectTagging` operation.
type PutObjectTaggingResult struct {
	VersionID string
}

func NewPutObjectTaggingResult(output *SDK.PutObjectTaggingResponse) *PutObjectTaggingResult {
	r := &PutObjectTaggingResult{}
	if output == nil {
		return r
	}

	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}
	return r
}
