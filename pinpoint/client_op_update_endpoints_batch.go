package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// UpdateEndpointsBatch executes `UpdateEndpointsBatch` operation.
func (svc *Pinpoint) UpdateEndpointsBatch(ctx context.Context, r UpdateEndpointsBatchRequest) (*UpdateEndpointsBatchResult, error) {
	out, err := svc.RawUpdateEndpointsBatch(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "UpdateEndpointsBatch",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewUpdateEndpointsBatchResult(out), nil
}

// UpdateEndpointsBatchRequest has parameters for `UpdateEndpointsBatch` operation.
type UpdateEndpointsBatchRequest struct {
	ApplicationID        string
	EndpointBatchRequest EndpointBatchRequest
}

func (r UpdateEndpointsBatchRequest) ToInput() *SDK.UpdateEndpointsBatchInput {
	in := &SDK.UpdateEndpointsBatchInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}

	in.EndpointBatchRequest = r.EndpointBatchRequest.ToSDK()
	return in
}

type UpdateEndpointsBatchResult struct {
	Message   string
	RequestID string
}

func NewUpdateEndpointsBatchResult(o *SDK.UpdateEndpointsBatchResponse) *UpdateEndpointsBatchResult {
	result := &UpdateEndpointsBatchResult{}
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
