package s3

import (
	"context"
)

// CreateBucketFromName creates a bucket of `name`.
func (svc *S3) CreateBucketFromName(ctx context.Context, name string) (*CreateBucketResult, error) {
	return svc.CreateBucket(ctx, CreateBucketRequest{
		Bucket: name,
	})
}

// DeleteBucketFromName deletes a bucket of `name`.
func (svc *S3) DeleteBucketFromName(ctx context.Context, name string) error {
	return svc.DeleteBucket(ctx, DeleteBucketRequest{
		Bucket: name,
	})
}

// IsExistBucket checks if the bucket already exists or not.
func (svc *S3) IsExistBucket(ctx context.Context, name string) (bool, error) {
	return svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: name,
	})
}
