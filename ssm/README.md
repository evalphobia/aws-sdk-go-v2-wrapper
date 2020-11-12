aws-sdk-go-v2-wrapper | SSM
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/ssm"
)

func main() {
	svc, err := ssm.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()


	// plain values
	alreadyExists, err := svc.XPutParameter(ctx, "plain-text", "I love you!")
	if err != nil {
		panic(err)
	}
	_ = alreadyExists
	// if alreadyExists {
	// 	panic("'plain-text' should not set")
	// }

	value, notFound, err := svc.XGetParameterValue(ctx, "plain-text")
	if err != nil {
		panic(err)
	}
	if notFound {
		panic("'value' should have value")
	}
	if value != "I love you!" {
		panic("'value' should be 'I love you!'")
	}

	// encrypted values
	alreadyExists, err = svc.XPutParameterSecrure(ctx, "cipher-text", "I love you!")
	if err != nil {
		panic(err)
	}
	_ = alreadyExists
	// if alreadyExists {
	// 	panic("'cipher-text' should not set")
	// }

	value, notFound, err = svc.XGetParameterValueWithDecryption(ctx, "cipher-text")
	if err != nil {
		panic(err)
	}
	if notFound {
		panic("'value' should have value")
	}
	if value != "I love you!" {
		panic("'value' should be 'I love you!'")
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XGetParameterValue` | gets a value from SSM Parameter Store. |
| `XGetParameterValueWithDecryption` | gets a decrypted value from SSM Parameter Store. |
| `XGetParameterValueList` | gets value list from SSM Parameter Store. |
| `XGetParameterValueListWithDecryption` | gets decrypted value list from SSM Parameter Store. |
| `XPutParameter` | puts a value to SSM Parameter Store. |
| `XPutParameterList` | puts values (StringList) to SSM Parameter Store. |
| `XPutParameterSecrure` | puts a values (SecureString) to SSM Parameter Store with an encryption of default AWS KMS. |
| `XPutParameterSecureCMK` | puts a values (SecureString) to SSM Parameter Store with an encryption of your own CMK. |
