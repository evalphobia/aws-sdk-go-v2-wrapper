package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutBucketLifecycleConfiguration executes `PutBucketLifecycleConfiguration` operation.
func (svc *S3) PutBucketLifecycleConfiguration(ctx context.Context, r PutBucketLifecycleConfigurationRequest) error {
	_, err := svc.RawPutBucketLifecycleConfiguration(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutBucketLifecycleConfiguration",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// PutBucketLifecycleConfigurationRequest has parameters for `PutBucketLifecycleConfiguration` operation.
type PutBucketLifecycleConfigurationRequest struct {
	Bucket string
	Rules  []LifecycleRule
}

func (r PutBucketLifecycleConfigurationRequest) ToInput() *SDK.PutBucketLifecycleConfigurationInput {
	in := &SDK.PutBucketLifecycleConfigurationInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if len(r.Rules) != 0 {
		list := make([]SDK.LifecycleRule, len(r.Rules))
		for i, v := range r.Rules {
			list[i] = v.ToSDK()
		}
		in.LifecycleConfiguration = &SDK.BucketLifecycleConfiguration{
			Rules: list,
		}
	}
	return in
}
