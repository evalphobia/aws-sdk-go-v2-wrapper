package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetSegment executes `GetSegment` operation.
func (svc *Pinpoint) GetSegment(ctx context.Context, r GetSegmentRequest) (*GetSegmentResult, error) {
	out, err := svc.RawGetSegment(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetSegment",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetSegmentResult(out), nil
}

// GetSegmentRequest has parameters for `GetSegment` operation.
type GetSegmentRequest struct {
	ApplicationID string
	SegmentID     string
}

func (r GetSegmentRequest) ToInput() *SDK.GetSegmentInput {
	in := &SDK.GetSegmentInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.SegmentID != "" {
		in.SegmentId = pointers.String(r.SegmentID)
	}
	return in
}

type GetSegmentResult struct {
	SegmentResponse
}

func NewGetSegmentResult(o *SDK.GetSegmentResponse) *GetSegmentResult {
	result := &GetSegmentResult{}
	if o == nil {
		return result
	}

	result.SegmentResponse = newSegmentResponse(o.SegmentResponse)
	return result
}
