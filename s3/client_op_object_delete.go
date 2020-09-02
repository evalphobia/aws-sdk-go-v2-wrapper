package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteObject executes `DeleteObject` operation.
func (svc *S3) DeleteObject(ctx context.Context, r DeleteObjectRequest) (*DeleteObjectResult, error) {
	out, err := svc.RawDeleteObject(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteObject",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteObjectResult(out), nil
}

// DeleteObjectRequest has parameters for `DeleteObject` operation.
type DeleteObjectRequest struct {
	Bucket string
	Key    string

	// optional
	BypassGovernanceRetention bool
	MFA                       string
	RequestPayer              RequestPayer
	VersionID                 string
}

func (r DeleteObjectRequest) ToInput() *SDK.DeleteObjectInput {
	in := &SDK.DeleteObjectInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}

	if r.BypassGovernanceRetention {
		in.BypassGovernanceRetention = pointers.Bool(r.BypassGovernanceRetention)
	}
	if r.MFA != "" {
		in.MFA = pointers.String(r.MFA)
	}

	in.RequestPayer = SDK.RequestPayer(r.RequestPayer)

	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}
	return in
}

// DeleteObjectResult contains results from `DeleteObject` operation.
type DeleteObjectResult struct {
	DeleteMarker   bool
	RequestCharged RequestCharged
	VersionID      string
}

func NewDeleteObjectResult(output *SDK.DeleteObjectResponse) *DeleteObjectResult {
	r := &DeleteObjectResult{}
	if output == nil {
		return r
	}
	if output.DeleteMarker != nil {
		r.DeleteMarker = *output.DeleteMarker
	}
	if output.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}
	return r
}
