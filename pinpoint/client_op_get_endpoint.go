package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetEndpoint executes `GetEndpoint` operation.
func (svc *Pinpoint) GetEndpoint(ctx context.Context, r GetEndpointRequest) (*GetEndpointResult, error) {
	out, err := svc.RawGetEndpoint(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetEndpoint",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetEndpointResult(out), nil
}

// GetEndpointRequest has parameters for `GetEndpoint` operation.
type GetEndpointRequest struct {
	ApplicationID string
	EndpointID    string
}

func (r GetEndpointRequest) ToInput() *SDK.GetEndpointInput {
	in := &SDK.GetEndpointInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.EndpointID != "" {
		in.EndpointId = pointers.String(r.EndpointID)
	}
	return in
}

type GetEndpointResult struct {
	EndpointResponse
}

func NewGetEndpointResult(o *SDK.GetEndpointResponse) *GetEndpointResult {
	result := &GetEndpointResult{}
	if o == nil {
		return result
	}

	result.EndpointResponse = newEndpointResponse(o.EndpointResponse)
	return result
}
