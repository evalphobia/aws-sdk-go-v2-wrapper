package s3

import (
	"context"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetObjectRetention executes `GetObjectRetention` operation.
func (svc *S3) GetObjectRetention(ctx context.Context, r GetObjectRetentionRequest) (*GetObjectRetentionResult, error) {
	out, err := svc.RawGetObjectRetention(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetObjectRetention",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetObjectRetentionResult(out), nil
}

// GetObjectRetentionRequest has parameters for `GetObjectRetention` operation.
type GetObjectRetentionRequest struct {
	Bucket string
	Key    string

	// optional
	RequestPayer RequestPayer
	VersionID    string
}

func (r GetObjectRetentionRequest) ToInput() *SDK.GetObjectRetentionInput {
	in := &SDK.GetObjectRetentionInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}

	in.RequestPayer = SDK.RequestPayer(r.RequestPayer)

	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}
	return in
}

// GetObjectRetentionResult contains results from `GetObjectRetention` operation.
type GetObjectRetentionResult struct {
	Mode            ObjectLockRetentionMode
	RetainUntilDate time.Time
}

func NewGetObjectRetentionResult(output *SDK.GetObjectRetentionResponse) *GetObjectRetentionResult {
	r := &GetObjectRetentionResult{}
	if output == nil {
		return r
	}

	rr := output.Retention
	if rr == nil {
		return r
	}

	if rr.Mode != "" {
		r.Mode = ObjectLockRetentionMode(rr.Mode)
	}
	if rr.RetainUntilDate != nil {
		r.RetainUntilDate = *rr.RetainUntilDate
	}
	return r
}
