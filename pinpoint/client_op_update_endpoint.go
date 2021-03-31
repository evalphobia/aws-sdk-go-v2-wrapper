package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// UpdateEndpoint executes `UpdateEndpoint` operation.
func (svc *Pinpoint) UpdateEndpoint(ctx context.Context, r UpdateEndpointRequest) (*UpdateEndpointResult, error) {
	out, err := svc.RawUpdateEndpoint(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "UpdateEndpoint",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewUpdateEndpointResult(out), nil
}

// UpdateEndpointRequest has parameters for `UpdateEndpoint` operation.
type UpdateEndpointRequest struct {
	ApplicationID   string
	EndpointID      string
	EndpointRequest EndpointRequest
}

func (r UpdateEndpointRequest) ToInput() *SDK.UpdateEndpointInput {
	in := &SDK.UpdateEndpointInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.EndpointID != "" {
		in.EndpointId = pointers.String(r.EndpointID)
	}

	in.EndpointRequest = r.EndpointRequest.ToSDK()
	return in
}

type UpdateEndpointResult struct {
	Message   string
	RequestID string
}

func NewUpdateEndpointResult(o *SDK.UpdateEndpointResponse) *UpdateEndpointResult {
	result := &UpdateEndpointResult{}
	if o == nil {
		return result
	}

	if o.MessageBody != nil {
		v := o.MessageBody
		result.Message = *v.Message
		result.RequestID = *v.RequestID
	}
	return result
}
