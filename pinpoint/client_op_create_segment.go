package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateSegment executes `CreateSegment` operation.
func (svc *Pinpoint) CreateSegment(ctx context.Context, r CreateSegmentRequest) (*CreateSegmentResult, error) {
	out, err := svc.RawCreateSegment(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateSegment",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateSegmentResult(out), nil
}

// CreateSegmentRequest has parameters for `CreateSegment` operation.
type CreateSegmentRequest struct {
	ApplicationID       string
	WriteSegmentRequest WriteSegmentRequest
}

func (r CreateSegmentRequest) ToInput() *SDK.CreateSegmentInput {
	in := &SDK.CreateSegmentInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}

	in.WriteSegmentRequest = r.WriteSegmentRequest.ToSDK()
	return in
}

type CreateSegmentResult struct {
	SegmentResponse
}

func NewCreateSegmentResult(o *SDK.CreateSegmentResponse) *CreateSegmentResult {
	result := &CreateSegmentResult{}
	if o == nil {
		return result
	}

	result.SegmentResponse = newSegmentResponse(o.SegmentResponse)
	return result
}
