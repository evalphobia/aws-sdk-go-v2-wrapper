aws-sdk-go-v2-wrapper | KMS
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/kms"
)

func main() {
	svc, err := kms.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	base64Text, err := svc.XEncryptString(ctx, "alias/my-key", "I love you!")
	if err != nil {
		panic(err)
	}

	plainText, err := svc.XDecryptString(ctx, base64Text)
	if err != nil {
		panic(err)
	}

	if plainText != "I love you!" {
		panic("'plainText' should be 'I love you!'")
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XDeleteKey` | deletes the key using 'ScheduleKeyDeletion'. |
| `XEncryptBytes` | encrypts the 'plainData' using the 'key' and returns 'encryptedData'. |
| `XEncryptString` | encrypts the 'plainText' using the 'key' and returns encrypted 'base64Text'. |
| `XDecryptBytes` | decrypts the 'encryptedData'. |
| `XDecryptString` | decrypts the 'base64Text'. |
| `XReEncryptBytes` | re-encrypts the 'encryptedData' using 'destinationKey'. |
| `XReEncryptString` | re-encrypts the 'base64Text' using 'destinationKey'. |
