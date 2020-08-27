package s3

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestPutObject(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)
	createBucket(testPutBucketName)

	// 1st put
	resPut, err := svc.PutObject(ctx, PutObjectRequest{
		Bucket:    testPutBucketName,
		Key:       "TestPutObject/test-file-01",
		BodyBytes: []byte("001"),
	})
	is.NoErr(err)
	is.True(resPut.ETag != "") // ETag must be set

	resGet, err := svc.GetObjectFromPath(ctx, testPutBucketName, "TestPutObject/test-file-01")
	is.NoErr(err)
	is.Equal(resGet.ETag, resPut.ETag) // ETag must be same
	byt, err := resGet.ToBytes()
	is.NoErr(err)
	is.Equal(byt, []byte("001"))

	// 2nd put (update)
	resPut, err = svc.PutObject(ctx, PutObjectRequest{
		Bucket:    testPutBucketName,
		Key:       "TestPutObject/test-file-01",
		BodyBytes: []byte("002"),
	})
	is.NoErr(err)
	is.True(resPut.ETag != "") // ETag must be set

	resGet, err = svc.GetObjectFromPath(ctx, testPutBucketName, "TestPutObject/test-file-01")
	is.NoErr(err)
	is.Equal(resGet.ETag, resPut.ETag) // ETag must be same
	byt, err = resGet.ToBytes()
	is.NoErr(err)
	is.Equal(byt, []byte("002")) // content must be updated
}
