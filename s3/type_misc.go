package s3

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/s3"
)

func newCommonPrefixes(list []SDK.CommonPrefix) []string {
	if len(list) == 0 {
		return nil
	}

	result := make([]string, 0, len(list))
	for _, v := range list {
		if v.Prefix != nil {
			result = append(result, *v.Prefix)
		}
	}
	return result
}
