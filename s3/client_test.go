package s3

import (
	"context"
	"testing"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/matryer/is"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	defaultEndpoint = "http://localhost:4567"

	testEmptyBucketName = "test-empty-bucket"
	testPutBucketName   = "test-put-bucket"
	testCopyBucketName  = "test-copy-bucket"

	testS3Path  = "test_path"
	testBaseURL = "http://localhost:4567/" + testPutBucketName
)

func TestNew(t *testing.T) {
	is := is.NewRelaxed(t)

	// custom endpoint
	svc, err := New(getTestConfig())
	is.NoErr(err)
	is.True(svc.client != nil)
	is.Equal(SDK.ServiceName, svc.client.Metadata.ServiceName)
	is.Equal(SDK.ServiceID, svc.client.Metadata.ServiceID)
	is.Equal(SDK.EndpointsID, svc.client.Metadata.EndpointsID)

	ep, err := svc.client.EndpointResolver.ResolveEndpoint("", "")
	is.NoErr(err)
	is.Equal(defaultEndpoint, ep.URL)

	// plain region
	svc, err = New(config.Config{})
	is.NoErr(err)
	is.True(svc.client != nil)

	region := "us-west-1"
	expectedEndpoint := "https://s3." + region + ".amazonaws.com"
	ep, err = svc.client.EndpointResolver.ResolveEndpoint(SDK.EndpointsID, region)
	is.NoErr(err)
	is.Equal(expectedEndpoint, ep.URL)

	_, err = svc.client.EndpointResolver.ResolveEndpoint("", "")
	is.True(err != nil)
}

func TestSetLogger(t *testing.T) {
	is := is.NewRelaxed(t)
	svc := getTestClient(t)
	is.Equal(log.DefaultLogger, svc.logger)

	stdLogger := &log.StdLogger{}
	svc.SetLogger(stdLogger)
	is.Equal(stdLogger, svc.logger)
}

func getTestConfig() config.Config {
	return config.Config{
		AccessKey:        "access",
		SecretKey:        "secret",
		CommonEndpoint:   defaultEndpoint,
		S3ForcePathStyle: true,
	}
}

func getTestClient(t *testing.T) *S3 {
	svc, err := New(getTestConfig())
	if err != nil {
		t.Errorf("error on create client; error=%s;", err.Error())
		t.FailNow()
	}
	return svc
}

func createBucket(name string) error {
	svc, err := New(getTestConfig())
	if err != nil {
		return err
	}

	ctx := context.Background()
	ok, err := svc.IsExistBucket(ctx, name)
	switch {
	case err != nil:
		return err
	case ok:
		return nil
	}

	_, err = svc.CreateBucketFromName(ctx, name)
	return err
}

// func testPutObject(t *testing.T) {
// 	is := is.NewRelaxed(t)
// 	createBucket(testPutBucketName)
// 	f := openFile(t)
// 	defer f.Close() // nolint:gosec

// 	svc := getTestClient(t)
// 	b, err := svc.GetBucket(testPutBucketName)
// 	is.NoErr(err)

// 	obj := NewPutObject(f)
// 	b.PutOne(obj, testS3Path, ACLPublicRead)

// 	err = b.PutAll()
// 	is.NoErr(err)
// }
