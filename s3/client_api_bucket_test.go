package s3

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestClientAPIBucket(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)

	t.Run("CreateBucketFromName", func(t *testing.T) {
		_ = svc.DeleteBucketFromName(ctx, testPutBucketName)
		res, err := svc.CreateBucketFromName(ctx, testPutBucketName)
		is.NoErr(err)              // success creation
		is.Equal(res.Location, "") // empty value on FakeS3

		_, err = svc.CreateBucketFromName(ctx, testPutBucketName)
		// is.True(err != nil) // already created bucket
		is.NoErr(err) // already created bucket does not error on FakeS3
	})

	t.Run("DeleteBucketFromName", func(t *testing.T) {
		_ = svc.DeleteBucketFromName(ctx, testPutBucketName)
		_, _ = svc.CreateBucketFromName(ctx, testPutBucketName)

		err := svc.DeleteBucketFromName(ctx, testPutBucketName)
		is.NoErr(err) // success deletion

		err = svc.DeleteBucketFromName(ctx, testPutBucketName)
		is.True(err != nil) // already deleted bucket
	})
	t.Run("ExistBucket", func(t *testing.T) {
		_ = svc.DeleteBucketFromName(ctx, testPutBucketName)

		ok, err := svc.ExistBucket(ctx, testPutBucketName)
		is.True(!ok)
		is.NoErr(err)

		_, _ = svc.CreateBucketFromName(ctx, testPutBucketName)
		ok, err = svc.ExistBucket(ctx, testPutBucketName)
		is.True(ok)
		is.NoErr(err)
	})
}
