package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateBucket executes `CreateBucket` operation.
func (svc *S3) CreateBucket(ctx context.Context, r CreateBucketRequest) (*CreateBucketResult, error) {
	out, err := svc.RawCreateBucket(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateBucket",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateBucketResult(out), nil
}

// CreateBucketRequest has parameters for `CreateBucket` operation.
type CreateBucketRequest struct {
	Bucket string

	// optional
	ACL                        BucketCannedACL
	LocationConstraint         BucketLocationConstraint
	GrantFullControl           string
	GrantRead                  string
	GrantReadACP               string
	GrantWrite                 string
	GrantWriteACP              string
	ObjectLockEnabledForBucket bool
}

func (r CreateBucketRequest) ToInput() *SDK.CreateBucketInput {
	in := &SDK.CreateBucketInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}

	in.ACL = SDK.BucketCannedACL(r.ACL)

	if r.LocationConstraint != "" {
		in.CreateBucketConfiguration = &SDK.CreateBucketConfiguration{
			LocationConstraint: SDK.BucketLocationConstraint(r.LocationConstraint),
		}
	}

	if r.GrantFullControl != "" {
		in.GrantFullControl = pointers.String(r.GrantFullControl)
	}
	if r.GrantRead != "" {
		in.GrantRead = pointers.String(r.GrantRead)
	}
	if r.GrantReadACP != "" {
		in.GrantReadACP = pointers.String(r.GrantReadACP)
	}
	if r.GrantWrite != "" {
		in.GrantWrite = pointers.String(r.GrantWrite)
	}
	if r.GrantWriteACP != "" {
		in.GrantWriteACP = pointers.String(r.GrantWriteACP)
	}
	if r.ObjectLockEnabledForBucket {
		in.ObjectLockEnabledForBucket = pointers.Bool(r.ObjectLockEnabledForBucket)
	}
	return in
}

// CreateBucketResult contains results from `CreateBucket` operation.
type CreateBucketResult struct {
	Location string
}

func NewCreateBucketResult(output *SDK.CreateBucketResponse) *CreateBucketResult {
	r := &CreateBucketResult{}
	if output == nil {
		return r
	}

	if output.Location != nil {
		r.Location = *output.Location
	}
	return r
}
