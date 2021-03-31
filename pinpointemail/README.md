aws-sdk-go-v2-wrapper | Pinpoint Email
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/pinpointemail"
)

func main() {
	svc, err := pinpointemail.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()


	// text email
	err := svc.XSendEmailText(ctx, "My email report", "I love you!", "example@example.com", "example@example.com")
	if err != nil {
		panic(err)
	}

	// HTML email
	htmlBody := `<!DOCTYPE html><html lang="en"><body><p style="text-align: center;">I love you!</p></body></html>`
	err = svc.XSendEmailHTML(ctx, "My email report", htmlBody, "example@example.com", "example@example.com")
	if err != nil {
		panic(err)
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XSendEmailHTML` | sends HTML type email. |
| `XSendEmailText` | sends text type email. |
