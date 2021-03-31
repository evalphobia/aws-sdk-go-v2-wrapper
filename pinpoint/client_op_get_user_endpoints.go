package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetUserEndpoints executes `GetUserEndpoints` operation.
func (svc *Pinpoint) GetUserEndpoints(ctx context.Context, r GetUserEndpointsRequest) (*GetUserEndpointsResult, error) {
	out, err := svc.RawGetUserEndpoints(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetUserEndpoints",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetUserEndpointsResult(out), nil
}

// GetUserEndpointsRequest has parameters for `GetUserEndpoints` operation.
type GetUserEndpointsRequest struct {
	ApplicationID string
	UserID        string
}

func (r GetUserEndpointsRequest) ToInput() *SDK.GetUserEndpointsInput {
	in := &SDK.GetUserEndpointsInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.UserID != "" {
		in.UserId = pointers.String(r.UserID)
	}
	return in
}

type GetUserEndpointsResult struct {
	EndpointsResponse
}

func NewGetUserEndpointsResult(o *SDK.GetUserEndpointsResponse) *GetUserEndpointsResult {
	result := &GetUserEndpointsResult{}
	if o == nil {
		return result
	}

	result.EndpointsResponse = newEndpointsResponse(o.EndpointsResponse)
	return result
}
