package s3

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestClientXAPIObject(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)
	_ = createTestBucket(testPutBucketName)

	const path = "TestClientXAPIObject/test"

	// check non-existed object
	_, err := svc.XDeleteObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)

	ok, err := svc.XExistObject(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.Equal(ok, false)

	resGet, err := svc.XGetObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.Equal(resGet.Exists, false)

	// check existed object
	resPut, err := svc.XPutObjectToPath(ctx, testPutBucketName, path, []byte("001"))
	is.NoErr(err)
	is.True(resPut.ETag != "") // ETag must be set

	ok, err = svc.XExistObject(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.True(ok)

	resGet, err = svc.XGetObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.Equal(resGet.ETag, resPut.ETag) // ETag must be same
	byt, err := resGet.ToBytes()
	is.NoErr(err)
	is.Equal(byt, []byte("001")) // content must be updated

	// update existed object
	resPut, err = svc.XPutObjectToPath(ctx, testPutBucketName, path, []byte("002"))
	is.NoErr(err)
	is.True(resPut.ETag != "") // ETag must be set

	resGet, err = svc.XGetObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.Equal(resGet.ETag, resPut.ETag) // ETag must be same
	byt, err = resGet.ToBytes()
	is.NoErr(err)
	is.Equal(byt, []byte("002")) // content must be updated

	// check deleted object
	_, err = svc.XDeleteObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)

	ok, err = svc.XExistObject(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.Equal(ok, false)

	resGet, err = svc.XGetObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)
	is.Equal(resGet.Exists, false)

	_, err = svc.XDeleteObjectFromPath(ctx, testPutBucketName, path)
	is.NoErr(err)
}
