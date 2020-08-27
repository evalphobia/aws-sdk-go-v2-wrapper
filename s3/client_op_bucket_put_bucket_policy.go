package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutBucketPolicy executes `PutBucketPolicy` operation.
func (svc *S3) PutBucketPolicy(ctx context.Context, r PutBucketPolicyRequest) error {
	_, err := svc.RawPutBucketPolicy(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutBucketPolicy",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// PutBucketPolicyRequest has parameters for `PutBucketPolicy` operation.
type PutBucketPolicyRequest struct {
	Bucket                        string
	ConfirmRemoveSelfBucketAccess bool
	Policy                        string
}

func (r PutBucketPolicyRequest) ToInput() *SDK.PutBucketPolicyInput {
	in := &SDK.PutBucketPolicyInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.ConfirmRemoveSelfBucketAccess {
		in.ConfirmRemoveSelfBucketAccess = pointers.Bool(r.ConfirmRemoveSelfBucketAccess)
	}
	if r.Policy != "" {
		in.Policy = pointers.String(r.Policy)
	}
	return in
}
