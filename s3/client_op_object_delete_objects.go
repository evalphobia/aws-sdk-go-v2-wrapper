package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteObjects executes `DeleteObjects` operation.
func (svc *S3) DeleteObjects(ctx context.Context, r DeleteObjectsRequest) (*DeleteObjectsResult, error) {
	out, err := svc.RawDeleteObjects(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteObjects",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteObjectsResult(out), nil
}

// DeleteObjectsRequest has parameters for `DeleteObjects` operation.
type DeleteObjectsRequest struct {
	Bucket string
	Delete Delete

	// optional
	BypassGovernanceRetention bool
	MFA                       string
	RequestPayer              RequestPayer
}

func (r DeleteObjectsRequest) ToInput() *SDK.DeleteObjectsInput {
	in := &SDK.DeleteObjectsInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	v := r.Delete.ToSDK()
	in.Delete = &v

	if r.BypassGovernanceRetention {
		in.BypassGovernanceRetention = pointers.Bool(r.BypassGovernanceRetention)
	}
	if r.MFA != "" {
		in.MFA = pointers.String(r.MFA)
	}

	in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
	return in
}

// DeleteObjectsResult contains results from `DeleteObjects` operation.
type DeleteObjectsResult struct {
	Deleted        []DeletedObject
	Errors         []Error
	RequestCharged RequestCharged
}

func NewDeleteObjectsResult(output *SDK.DeleteObjectsResponse) *DeleteObjectsResult {
	r := &DeleteObjectsResult{}
	if output == nil {
		return r
	}

	r.Deleted = newDeletedObjects(output.Deleted)
	r.Errors = newErrors(output.Errors)

	if output.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	return r
}
