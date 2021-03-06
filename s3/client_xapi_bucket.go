package s3

import (
	"context"
)

// XCreateBucketFromName creates a bucket of `name`.
func (svc *S3) XCreateBucketFromName(ctx context.Context, name string) (*CreateBucketResult, error) {
	return svc.CreateBucket(ctx, CreateBucketRequest{
		Bucket: name,
	})
}

// XDeleteBucketFromName deletes a bucket of `name`.
func (svc *S3) XDeleteBucketFromName(ctx context.Context, name string) error {
	return svc.DeleteBucket(ctx, DeleteBucketRequest{
		Bucket: name,
	})
}

// XExistBucket checks if the bucket already exists or not.
func (svc *S3) XExistBucket(ctx context.Context, name string) (bool, error) {
	return svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: name,
	})
}

// XForceDeleteBucketFromName deletes a bucket with deleting all of objects.
func (svc *S3) XForceDeleteBucketFromName(ctx context.Context, name string) error {
	_, err := svc.XDeleteAllObjects(ctx, name)
	if err != nil {
		return err
	}

	return svc.XDeleteBucketFromName(ctx, name)
}
