aws-sdk-go-v2-wrapper | SQS
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/sqs"
)

func main() {
	svc, err := sqs.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()


	// create queue
	queueName := "my-queue"
	queueURL, err := svc.XCreateQueue(ctx, "my-queue")
	if err != nil {
		panic(err)
	}

	err := svc.XSendMessage(ctx, queueURL, "foo")
	if err != nil {
		panic(err)
	}

	messages, err := svc.XReceiveMessage(ctx, queueURL)
	if err != nil {
		panic(err)
	}


	err = svc.XPurgeQueue(ctx, queueURL)
	if err != nil {
		panic(err)
	}

	err = svc.XPurgeQueueByName(ctx, queueName)
	if err != nil {
		panic(err)
	}

	err = svc.XDeleteQueueByName(ctx, queueName)
	if err != nil {
		panic(err)
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XCreateQueue` | creates a SQS queue. |
| `XDeleteQueue` | deletes a SQS queue. |
| `XDeleteQueueByName` | deletes a SQS queue by queue name. |
| `XExistsQueue` | checks if the SQS queue exists or not. |
| `XGetQueueURL` | gets a SQS queue url. |
| `XListAllQueueURLs` | gets all of the SQS queue url. |
| `XPurgeQueue` | purges message in the queue. |
| `XPurgeQueueByName` | purges message in the queue by queue name. |
| `XDeleteMessage` | deletes a message. |
| `XReceiveMessage` |receives messages. |
| `XReceiveMessageLongPolling` | receives messages by long polling. |
| `XSendMessage` | sends a message. |
| `XIsEmptyQueue` | checks if the queue have any message or not. |
