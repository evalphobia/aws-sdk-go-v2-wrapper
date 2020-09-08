aws-sdk-go-v2-wrapper | DynamoDB
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/dynamodb"
)

func main() {
	svc, err := dynamodb.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	item, err := svc.GetSingleItem(ctx, dynamodb.XGetSingleItem{
		TableName:     "users",
		HashKeyName:   "user_id",
		HashKeyValue:  101,
		RangeKeyName:  "session_id",
		RangeKeyValue: "sess-abcdefg123456",
	})
	if err != nil {
		panic(err)
	}

	mm := dynamodb.ToMapValue(item)
	if mm["user_id"] != 101 {
		panic("`user_id` should be `101`")
	}
	if mm["session_id"] != "sess-abcdefg123456" {
		panic("`session_id` should be `sess-abcdefg123456`")
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XBatchDeleteItems` | deletes multiple items using 'BatchWriteItems'. |
| `XDeleteTableFromName` | deletes a table. |
| `XExistTable` | checks if the table already exists or not. |
| `XForceDeleteAll` | deletes all data in the table.  |
| `XGetSingleItem` | gets single item. |

# Other Examples

## PutItem

```go
data := map[string]interface{}{
    "user_id": 101,
    "session_id": "sess-abcdefg123456",
}
m, err := dynamodb.MarshalToMap(data)
if err != nil {
    return err
}

svc.PutItem(ctx, dynamodb.PutItemRequest{
	TableName: "users",
	Item:      m,
})
```


## Query

```go
result, err := svc.Query(ctx, dynamodb.QueryRequest{
	TableName: "users",
	IndexName: "gsi-foobar",
	// Select: dynamodb.SelectCount,
	Limit:     1,
	XConditions: dynamodb.XConditions{
		KeyConditions: []dynamodb.XCondition{
			{
				Name:     "user_id",
				Value:    101,
				Operator: dynamodb.ComparisonOperatorEq,
			},
			{
				Name:     "session_id",
				Value:    "sess-abcdefg123456",
				Operator: dynamodb.ComparisonOperatorEq,
			},
		},
	},
})
if err != nil {
    return err
}

if result.Count != 1 {
	panic("result should be one item")
}

mm := result.ToSliceMap()
if len(mm) != 1 {
	panic("result should be one item")
}

if mm["user_id"] != 101 {
    panic("`user_id` should be `101`")
}
if mm["session_id"] != "sess-abcdefg123456" {
    panic("`session_id` should be `sess-abcdefg123456`")
}
```


## UpdateItem

```go
_, err = svc.UpdateItem(ctx, dynamodb.UpdateItemRequest{
	TableName: "users",
	// [WHERE user_id = 101 AND session_id = 'sess-abcdefg123456'] in SQL
	Key: map[string]dynamodb.AttributeValue{
		"user_id": dynamodb.AttributeValue{
			Number: 101,
		},
		"session_id": dynamodb.AttributeValue{
			String: "sess-abcdefg123456",
		},
	},
	ReturnValues: dynamodb.ReturnValueNone,
	XConditions: dynamodb.XConditions{
		Updates: []dynamodb.XUpdateCondition{
			// change email address
			// [UPDATE SET email = 'example@example.com'] in SQL
			{
				Name:      "email",
				Value:     "example@example.com",
				Operation: dynamodb.OperationModeSET, // you can omit `set` parameter
			},
			// increment age
			// [UPDATE SET age = age + 1] in SQL
			{
				Name:      "age",
				Value:     1,
				Operation: dynamodb.OperationModeADD,
			},
		},
		// DO NOT create new item when p-key does not exist.
		Conditions: []dynamodb.XCondition{
			{
				Name:     "user_id",
				Operator: dynamodb.ComparisonOperatorAttrExists,
			},
		},
	},
})
```
