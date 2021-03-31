package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteUserEndpoints executes `DeleteUserEndpoints` operation.
func (svc *Pinpoint) DeleteUserEndpoints(ctx context.Context, r DeleteUserEndpointsRequest) (*DeleteUserEndpointsResult, error) {
	out, err := svc.RawDeleteUserEndpoints(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteUserEndpoints",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteUserEndpointsResult(out), nil
}

// DeleteUserEndpointsRequest has parameters for `DeleteUserEndpoints` operation.
type DeleteUserEndpointsRequest struct {
	ApplicationID string
	UserID        string
}

func (r DeleteUserEndpointsRequest) ToInput() *SDK.DeleteUserEndpointsInput {
	in := &SDK.DeleteUserEndpointsInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.UserID != "" {
		in.UserId = pointers.String(r.UserID)
	}
	return in
}

type DeleteUserEndpointsResult struct {
	EndpointsResponse
}

func NewDeleteUserEndpointsResult(o *SDK.DeleteUserEndpointsResponse) *DeleteUserEndpointsResult {
	result := &DeleteUserEndpointsResult{}
	if o == nil {
		return result
	}

	result.EndpointsResponse = newEndpointsResponse(o.EndpointsResponse)
	return result
}
