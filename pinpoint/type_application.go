package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

type ApplicationsResponse struct {
	Item      []ApplicationResponse
	NextToken string
}

func newApplicationsResponse(o *SDK.ApplicationsResponse) ApplicationsResponse {
	result := ApplicationsResponse{}
	if o == nil {
		return result
	}

	if len(o.Item) != 0 {
		list := make([]ApplicationResponse, len(o.Item))
		for i, v := range o.Item {
			v := v
			list[i] = newApplicationResponse(&v)
		}
		result.Item = list
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}
	return result
}

type ApplicationResponse struct {
	ARN  string
	ID   string
	Name string
	Tags map[string]string
}

func newApplicationResponse(o *SDK.ApplicationResponse) ApplicationResponse {
	result := ApplicationResponse{}
	if o == nil {
		return result
	}

	if o.Arn != nil {
		result.ARN = *o.Arn
	}
	if o.Id != nil {
		result.ID = *o.Id
	}
	if o.Name != nil {
		result.Name = *o.Name
	}

	result.Tags = o.Tags
	return result
}
