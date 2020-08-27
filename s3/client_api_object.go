package s3

import (
	"context"
)

// IsExistObject checks if the object exists or not.
func (svc *S3) IsExistObject(ctx context.Context, bucket, path string) (bool, error) {
	res, err := svc.HeadObject(ctx, HeadObjectRequest{
		Bucket: bucket,
		Key:    path,
	})
	if err != nil {
		return false, err
	}
	return res.IsExist, nil
}

// GetObjectFromPath gets an object from `path`.
func (svc *S3) GetObjectFromPath(ctx context.Context, bucket, path string) (*GetObjectResult, error) {
	return svc.GetObject(ctx, GetObjectRequest{
		Bucket: bucket,
		Key:    path,
	})
}

// DeleteObjectFromPath deletes an object from `path`.
func (svc *S3) DeleteObjectFromPath(ctx context.Context, bucket, path string) (*DeleteObjectResult, error) {
	return svc.DeleteObject(ctx, DeleteObjectRequest{
		Bucket: bucket,
		Key:    path,
	})
}

// PutObjectToPath puts an object to the `path` and bytes data.
func (svc *S3) PutObjectToPath(ctx context.Context, bucket, path string, data []byte) (*PutObjectResult, error) {
	return svc.PutObject(ctx, PutObjectRequest{
		Bucket:    bucket,
		Key:       path,
		BodyBytes: data,
	})
}
