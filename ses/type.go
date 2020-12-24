package ses

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type MessageTag struct {
	Name  string
	Value string
}

func (r MessageTag) ToSDK() SDK.MessageTag {
	o := SDK.MessageTag{}

	if r.Name != "" {
		o.Name = pointers.String(r.Name)
	}
	if r.Value != "" {
		o.Value = pointers.String(r.Value)
	}
	return o
}
