package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetBucketLifecycleConfiguration executes `GetBucketLifecycleConfiguration` operation.
func (svc *S3) GetBucketLifecycleConfiguration(ctx context.Context, r GetBucketLifecycleConfigurationRequest) (*GetBucketLifecycleConfigurationResult, error) {
	out, err := svc.RawGetBucketLifecycleConfiguration(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetBucketLifecycleConfiguration",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetBucketLifecycleConfigurationResult(out), nil
}

// GetBucketLifecycleConfigurationRequest has parameters for `GetBucketLifecycleConfiguration` operation.
type GetBucketLifecycleConfigurationRequest struct {
	Bucket string
}

func (r GetBucketLifecycleConfigurationRequest) ToInput() *SDK.GetBucketLifecycleConfigurationInput {
	in := &SDK.GetBucketLifecycleConfigurationInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	return in
}

// GetBucketLifecycleConfigurationResult contains results from `GetBucketLifecycleConfiguration` operation.
type GetBucketLifecycleConfigurationResult struct {
	Rules []LifecycleRule
}

func NewGetBucketLifecycleConfigurationResult(output *SDK.GetBucketLifecycleConfigurationResponse) *GetBucketLifecycleConfigurationResult {
	r := &GetBucketLifecycleConfigurationResult{}
	if output == nil {
		return r
	}

	r.Rules = newLifecycleRules(output.Rules)
	return r
}
