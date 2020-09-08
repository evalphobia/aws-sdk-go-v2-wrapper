package s3

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestDeleteObject(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)
	_ = createTestBucket(testPutBucketName)

	_, err := svc.PutObject(ctx, PutObjectRequest{
		Bucket:    testPutBucketName,
		Key:       "TestDeleteObject/test-file-01",
		BodyBytes: []byte("-"),
	})
	is.NoErr(err)

	ok, err := svc.XExistObject(ctx, testPutBucketName, "TestDeleteObject/test-file-01")
	is.NoErr(err)
	is.True(ok)

	// after deletion
	_, err = svc.DeleteObject(ctx, DeleteObjectRequest{
		Bucket: testPutBucketName,
		Key:    "TestDeleteObject/test-file-01",
	})
	is.NoErr(err)

	ok, err = svc.XExistObject(ctx, testPutBucketName, "TestDeleteObject/test-file-01")
	is.NoErr(err)
	is.True(!ok)

	// delete non-existed file
	_, err = svc.DeleteObject(ctx, DeleteObjectRequest{
		Bucket: testPutBucketName,
		Key:    "TestDeleteObject/test-file-01",
	})
	is.NoErr(err)
}
