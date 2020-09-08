package s3

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestClientXAPIBucket(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)

	t.Run("XCreateBucketFromName", func(t *testing.T) {
		_ = svc.XForceDeleteBucketFromName(ctx, testPutBucketName)
		res, err := svc.XCreateBucketFromName(ctx, testPutBucketName)
		is.NoErr(err) // success creation
		is.Equal(res.Location, "/test-put-bucket")

		_, err = svc.XCreateBucketFromName(ctx, testPutBucketName)
		is.True(err != nil) // already created bucket
	})

	t.Run("XDeleteBucketFromName", func(t *testing.T) {
		_ = svc.XForceDeleteBucketFromName(ctx, testPutBucketName)
		_, _ = svc.XCreateBucketFromName(ctx, testPutBucketName)

		err := svc.XDeleteBucketFromName(ctx, testPutBucketName)
		is.NoErr(err) // success deletion

		err = svc.XDeleteBucketFromName(ctx, testPutBucketName)
		is.True(err != nil) // already deleted bucket
	})
	t.Run("XExistBucket", func(t *testing.T) {
		_ = svc.XForceDeleteBucketFromName(ctx, testPutBucketName)

		ok, err := svc.XExistBucket(ctx, testPutBucketName)
		is.True(!ok)
		is.NoErr(err)

		_, _ = svc.XCreateBucketFromName(ctx, testPutBucketName)
		ok, err = svc.XExistBucket(ctx, testPutBucketName)
		is.True(ok)
		is.NoErr(err)
	})
}
