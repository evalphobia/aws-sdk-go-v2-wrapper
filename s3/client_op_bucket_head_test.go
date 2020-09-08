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

	_ = svc.XForceDeleteBucketFromName(ctx, testPutBucketName)
	_ = svc.XForceDeleteBucketFromName(ctx, nonExistedBucket)

	ok, err := svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: nonExistedBucket,
	})
	is.NoErr(err)
	is.True(!ok) // non-existed bucket should be false

	_, _ = svc.XCreateBucketFromName(ctx, testPutBucketName)
	ok, err = svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: testPutBucketName,
	})
	is.NoErr(err)
	is.True(ok) // existed bucket should be true

	err = svc.XDeleteBucketFromName(ctx, testPutBucketName)
	is.NoErr(err)

	ok, err = svc.HeadBucket(ctx, HeadBucketRequest{
		Bucket: testPutBucketName,
	})
	is.NoErr(err)
	is.True(!ok) // deleted bucket should be false
}
