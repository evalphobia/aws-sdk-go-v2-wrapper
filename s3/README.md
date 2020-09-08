aws-sdk-go-v2-wrapper | S3
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/s3"
)

func main() {
	svc, err := s3.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	has, err := svc.XExistObject(ctx, "my-bucket", "users/101/data")
	if err != nil {
		panic(err)
	}
	if !has {
		panic("cannot find the data")
	}

	result, err := svc.XGetObjectFromPath(ctx, "my-bucket", "users/101/data")
	if err != nil {
		panic(err)
	}
	if !result.Exists {
		panic("cannot find the data")
	}

	byt, err := fresult.ToBytes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[my-bucket/users/101/data] has %d bytes", len(byt))

	presignedURL, err := svc.XGetPresignURL(ctx, "my-bucket", "users/101/data", 10 * time.Minute)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Download here: %s (this link expires after 10 minutes)", presignedURL)
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XCreateBucketFromName` | creates a bucket. |
| `XDeleteBucketFromName` | deletes a bucket. |
| `XDeleteObjectFromPath` | deletes an object. |
| `XExistBucket` | checks if the bucket already exists or not. |
| `XExistObject` | checks if the object exists or not. |
| `XGetObjectFromPath` | gets an object. |
| `XGetPresignURL` | gets an presigned url of the object for GET. |
| `XPutObjectToPath` | puts an object to the `path` and bytes data. |
