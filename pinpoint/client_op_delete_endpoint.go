package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteEndpoint executes `DeleteEndpoint` operation.
func (svc *Pinpoint) DeleteEndpoint(ctx context.Context, r DeleteEndpointRequest) (*DeleteEndpointResult, error) {
	out, err := svc.RawDeleteEndpoint(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteEndpoint",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteEndpointResult(out), nil
}

// DeleteEndpointRequest has parameters for `DeleteEndpoint` operation.
type DeleteEndpointRequest struct {
	ApplicationID string
	EndpointID    string
}

func (r DeleteEndpointRequest) ToInput() *SDK.DeleteEndpointInput {
	in := &SDK.DeleteEndpointInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.EndpointID != "" {
		in.EndpointId = pointers.String(r.EndpointID)
	}
	return in
}

type DeleteEndpointResult struct {
	EndpointResponse
}

func NewDeleteEndpointResult(o *SDK.DeleteEndpointResponse) *DeleteEndpointResult {
	result := &DeleteEndpointResult{}
	if o == nil {
		return result
	}

	result.EndpointResponse = newEndpointResponse(o.EndpointResponse)
	return result
}
