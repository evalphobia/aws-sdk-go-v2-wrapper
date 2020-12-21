aws-sdk-go-v2-wrapper | Athena
----


# Quick Usage

```go
import (
	"context"
	"encoding/json"
	"strings"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/athena"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
)

func main() {
	svc, err := athena.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	sql := "SELECT col1, col2 FROM db.table WHERE col3 = 'foobar'"
	// fetch error level logs
	result, err := svc.XQueryResults(ctx, cloudwatchlogs.XQueryResultsRequest{
		QueryString: sql,
		WorkGroup:   "primary",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("mapResult: [%+v]", result.ResultSet.ToMapString())
	fmt.Printf("sliceResult: [%+v]", result.ResultSet.ToListString())
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XQueryResults` | executes a query and waits for fetching complete results. |
| `XGetQueryResultsAll` | waits for fetching complete results. |
