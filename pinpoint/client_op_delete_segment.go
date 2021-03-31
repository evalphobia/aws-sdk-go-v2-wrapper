package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteSegment executes `DeleteSegment` operation.
func (svc *Pinpoint) DeleteSegment(ctx context.Context, r DeleteSegmentRequest) (*DeleteSegmentResult, error) {
	out, err := svc.RawDeleteSegment(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteSegment",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteSegmentResult(out), nil
}

// DeleteSegmentRequest has parameters for `DeleteSegment` operation.
type DeleteSegmentRequest struct {
	ApplicationID string
	SegmentID     string
}

func (r DeleteSegmentRequest) ToInput() *SDK.DeleteSegmentInput {
	in := &SDK.DeleteSegmentInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.SegmentID != "" {
		in.SegmentId = pointers.String(r.SegmentID)
	}
	return in
}

type DeleteSegmentResult struct {
	SegmentResponse
}

func NewDeleteSegmentResult(o *SDK.DeleteSegmentResponse) *DeleteSegmentResult {
	result := &DeleteSegmentResult{}
	if o == nil {
		return result
	}

	result.SegmentResponse = newSegmentResponse(o.SegmentResponse)
	return result
}
