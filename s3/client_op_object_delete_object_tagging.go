package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteObjectTagging executes `DeleteObjectTagging` operation.
func (svc *S3) DeleteObjectTagging(ctx context.Context, r DeleteObjectTaggingRequest) (*DeleteObjectTaggingResult, error) {
	out, err := svc.RawDeleteObjectTagging(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteObjectTagging",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteObjectTaggingResult(out), nil
}

// DeleteObjectTaggingRequest has parameters for `DeleteObjectTagging` operation.
type DeleteObjectTaggingRequest struct {
	Bucket string
	Key    string

	// optional
	VersionID string
}

func (r DeleteObjectTaggingRequest) ToInput() *SDK.DeleteObjectTaggingInput {
	in := &SDK.DeleteObjectTaggingInput{}
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

// DeleteObjectTaggingResult contains results from `DeleteObjectTagging` operation.
type DeleteObjectTaggingResult struct {
	VersionID string
}

func NewDeleteObjectTaggingResult(output *SDK.DeleteObjectTaggingResponse) *DeleteObjectTaggingResult {
	r := &DeleteObjectTaggingResult{}
	if output == nil {
		return r
	}
	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}
	return r
}
