package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetBucketPolicy executes `GetBucketPolicy` operation.
func (svc *S3) GetBucketPolicy(ctx context.Context, r GetBucketPolicyRequest) (*GetBucketPolicyResult, error) {
	out, err := svc.RawGetBucketPolicy(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetBucketPolicy",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetBucketPolicyResult(out), nil
}

// GetBucketPolicyRequest has parameters for `GetBucketPolicy` operation.
type GetBucketPolicyRequest struct {
	Bucket string
}

func (r GetBucketPolicyRequest) ToInput() *SDK.GetBucketPolicyInput {
	in := &SDK.GetBucketPolicyInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}

// GetBucketPolicyResult contains results from `GetBucketPolicy` operation.
type GetBucketPolicyResult struct {
	Policy string
}

func NewGetBucketPolicyResult(output *SDK.GetBucketPolicyResponse) *GetBucketPolicyResult {
	r := &GetBucketPolicyResult{}
	if output == nil {
		return r
	}

	if output.Policy != nil {
		r.Policy = *output.Policy
	}
	return r
}
