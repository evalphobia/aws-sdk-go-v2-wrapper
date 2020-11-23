package pinpointemail

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpointemail"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type Tag struct {
	Key   string
	Value string
}

func (r Tag) ToSDK() SDK.Tag {
	o := SDK.Tag{}

	if r.Key != "" {
		o.Key = pointers.String(r.Key)
	}
	if r.Value != "" {
		o.Value = pointers.String(r.Value)
	}
	return o
}

func (r Tag) ToEmailTag() SDK.MessageTag {
	o := SDK.MessageTag{}

	if r.Key != "" {
		o.Name = pointers.String(r.Key)
	}
	if r.Value != "" {
		o.Value = pointers.String(r.Value)
	}
	return o
}
