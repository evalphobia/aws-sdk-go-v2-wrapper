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

	_ = svc.XDeleteBucketFromName(ctx, testPutBucketName)
	_ = svc.XDeleteBucketFromName(ctx, nonExistedBucket)

	ok, err := svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: nonExistedBucket,
	})
	is.NoErr(err)
	is.True(!ok) // non-existed bucket shold be false

	_, _ = svc.XCreateBucketFromName(ctx, testPutBucketName)
	ok, err = svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: testPutBucketName,
	})
	is.NoErr(err)
	is.True(ok) // existed bucket shold be true

	_ = svc.XDeleteBucketFromName(ctx, testPutBucketName)
	ok, err = svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: testPutBucketName,
	})
	is.NoErr(err)
	is.True(!ok) // deleted bucket shold be false
}
