package s3

import (
	"context"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutObjectRetention executes `PutObjectRetention` operation.
func (svc *S3) PutObjectRetention(ctx context.Context, r PutObjectRetentionRequest) (*PutObjectRetentionResult, error) {
	out, err := svc.RawPutObjectRetention(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutObjectRetention",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewPutObjectRetentionResult(out), nil
}

// PutObjectRetentionRequest has parameters for `PutObjectRetention` operation.
type PutObjectRetentionRequest struct {
	Bucket string
	Key    string

	// optional
	BypassGovernanceRetention bool
	RequestPayer              RequestPayer
	VersionID                 string

	ObjectLockRetentionMode   ObjectLockRetentionMode
	ObjectLockRetainUntilDate time.Time
}

func (r PutObjectRetentionRequest) ToInput() *SDK.PutObjectRetentionInput {
	in := &SDK.PutObjectRetentionInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}
	if r.BypassGovernanceRetention {
		in.BypassGovernanceRetention = pointers.Bool(r.BypassGovernanceRetention)
	}
	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}

	if r.RequestPayer != "" {
		in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
	}

	switch {
	case r.ObjectLockRetentionMode != "",
		!r.ObjectLockRetainUntilDate.IsZero():
		v := &SDK.ObjectLockRetention{}
		if r.ObjectLockRetentionMode != "" {
			v.Mode = SDK.ObjectLockRetentionMode(r.ObjectLockRetentionMode)
		}
		if !r.ObjectLockRetainUntilDate.IsZero() {
			v.RetainUntilDate = &r.ObjectLockRetainUntilDate
		}
		in.Retention = v
	}

	return in
}

// PutObjectRetentionResult contains results from `PutObjectRetention` operation.
type PutObjectRetentionResult struct {
	RequestCharged RequestCharged
}

func NewPutObjectRetentionResult(output *SDK.PutObjectRetentionResponse) *PutObjectRetentionResult {
	r := &PutObjectRetentionResult{}
	if output == nil {
		return r
	}

	if output.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	return r
}
