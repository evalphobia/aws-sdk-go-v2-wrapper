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
		_ = svc.XDeleteBucketFromName(ctx, testPutBucketName)
		res, err := svc.XCreateBucketFromName(ctx, testPutBucketName)
		is.NoErr(err)              // success creation
		is.Equal(res.Location, "") // empty value on FakeS3

		_, err = svc.XCreateBucketFromName(ctx, testPutBucketName)
		// is.True(err != nil) // already created bucket
		is.NoErr(err) // already created bucket does not error on FakeS3
	})

	t.Run("XDeleteBucketFromName", func(t *testing.T) {
		_ = svc.XDeleteBucketFromName(ctx, testPutBucketName)
		_, _ = svc.XCreateBucketFromName(ctx, testPutBucketName)

		err := svc.XDeleteBucketFromName(ctx, testPutBucketName)
		is.NoErr(err) // success deletion

		err = svc.XDeleteBucketFromName(ctx, testPutBucketName)
		is.True(err != nil) // already deleted bucket
	})
	t.Run("XExistBucket", func(t *testing.T) {
		_ = svc.XDeleteBucketFromName(ctx, testPutBucketName)

		ok, err := svc.XExistBucket(ctx, testPutBucketName)
		is.True(!ok)
		is.NoErr(err)

		_, _ = svc.XCreateBucketFromName(ctx, testPutBucketName)
		ok, err = svc.XExistBucket(ctx, testPutBucketName)
		is.True(ok)
		is.NoErr(err)
	})
}
