package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetObjectTagging executes `GetObjectTagging` operation.
func (svc *S3) GetObjectTagging(ctx context.Context, r GetObjectTaggingRequest) (*GetObjectTaggingResult, error) {
	out, err := svc.RawGetObjectTagging(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetObjectTagging",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetObjectTaggingResult(out), nil
}

// GetObjectTaggingRequest has parameters for `GetObjectTagging` operation.
type GetObjectTaggingRequest struct {
	Bucket string
	Key    string

	// optional
	VersionID string
}

func (r GetObjectTaggingRequest) ToInput() *SDK.GetObjectTaggingInput {
	in := &SDK.GetObjectTaggingInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}
	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}
	return in
}

// GetObjectTaggingResult contains results from `GetObjectTagging` operation.
type GetObjectTaggingResult struct {
	TagSet    []Tag
	VersionID string
}

func NewGetObjectTaggingResult(output *SDK.GetObjectTaggingResponse) *GetObjectTaggingResult {
	r := &GetObjectTaggingResult{}
	if output == nil {
		return r
	}

	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}

	r.TagSet = newTags(output.TagSet)
	return r
}
