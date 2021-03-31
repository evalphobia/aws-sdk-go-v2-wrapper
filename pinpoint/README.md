aws-sdk-go-v2-wrapper | Pinpoint
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/pinpoint"
)

func main() {
	svc, err := pinpoint.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()


	appID := "abcd-efgh-0213-ijkl"
	// check the app exists or not
	ok, err := svc.XExistsApp(ctx, appID)
	if err != nil {
		panic(err)
	}

	if !ok {
		// gets app id from name
		appID, err = svc.XExistsApp(ctx, "my-cool-project")
		if err != nil {
			panic(err)
		}

		ok, err = svc.XExistsApp(ctx, appID)
		if err != nil {
			panic(err)
		}
		if !ok {
			panic("cannot find the app")
		}
	}

	attrs := map[string]string{
		"OS": "iOS",
	}
	userAttrs := map[string]string{
		"Country": "Canada",
	}
	// update the endpoint
	err = svc.XUpdateEndpointEmail(ctx, appID, "endpoint-id", "user-id", "example@example.com", attrs, userAttrs)
	if err != nil {
		panic(err)
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XExistsApp` | checks if the app (project) exists or not. |
| `XGetAppIDByName` | gets application id by name. |
| `XUpdateEndpointEmail` | creates or updates email of the endpoint. |
