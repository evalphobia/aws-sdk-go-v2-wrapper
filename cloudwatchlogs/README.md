aws-sdk-go-v2-wrapper | Cloudwatch Logs
----


# Quick Usage

```go
import (
	"context"
	"encoding/json"
	"strings"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/cloudwatchlogs"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
)

func main() {
	svc, err := cloudwatchlogs.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	filterValues := []string{
		`fields @timestamp, @message, level`,
		`| filter level = "error"`,
		`| sort @timestamp asc`,
		`| limit 20`,
	}
	// fetch error level logs
	result, err := svc.XQueryResults(ctx, cloudwatchlogs.XQueryResultsRequest{
		LogGroupName: "/aws/lambda/my-lambda-func",
		QueryString:  strings.Join(filterValues, " "),
		StartTime:    time.Now().Add(-1 * 5 * time.Minute), // past 5 mins
		EndTime:      time.Now(),
	})
	switch {
	case err != nil:
		panic(err)
	case len(result.Results) == 0:
		fmt.Printf("No errors are detected.")
	default:
		byt, _ := json.Marshal(res.Results)
		fmt.Printf("Error Hits:[%d], Error Data:[%s]", len(res.Results), string(byt))
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XQueryResults` | executes a query and waits for fetching complete results. |
