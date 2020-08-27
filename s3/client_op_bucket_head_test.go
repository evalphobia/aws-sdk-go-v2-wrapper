package s3

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestHeadBucket(t *testing.T) {
	const nonExistedBucket = "non-existed-bucket--xxx"
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)

	_ = svc.DeleteBucketFromName(ctx, testPutBucketName)
	_ = svc.DeleteBucketFromName(ctx, nonExistedBucket)

	ok, err := svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: nonExistedBucket,
	})
	is.NoErr(err)
	is.True(!ok) // non-existed bucket shold be false

	_, _ = svc.CreateBucketFromName(ctx, testPutBucketName)
	ok, err = svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: testPutBucketName,
	})
	is.NoErr(err)
	is.True(ok) // existed bucket shold be true

	_ = svc.DeleteBucketFromName(ctx, testPutBucketName)
	ok, err = svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: testPutBucketName,
	})
	is.NoErr(err)
	is.True(!ok) // deleted bucket shold be false
}
