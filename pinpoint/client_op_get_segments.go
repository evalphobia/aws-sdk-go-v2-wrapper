package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetSegments executes `GetSegments` operation.
func (svc *Pinpoint) GetSegments(ctx context.Context, r GetSegmentsRequest) (*GetSegmentsResult, error) {
	out, err := svc.RawGetSegments(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetSegments",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetSegmentsResult(out), nil
}

// GetSegmentsRequest has parameters for `GetSegments` operation.
type GetSegmentsRequest struct {
	ApplicationID string
	PageSize      string
	Token         string
}

func (r GetSegmentsRequest) ToInput() *SDK.GetSegmentsInput {
	in := &SDK.GetSegmentsInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.PageSize != "" {
		in.PageSize = pointers.String(r.PageSize)
	}
	if r.Token != "" {
		in.Token = pointers.String(r.Token)
	}
	return in
}

type GetSegmentsResult struct {
	SegmentsResponse
}

func NewGetSegmentsResult(o *SDK.GetSegmentsResponse) *GetSegmentsResult {
	result := &GetSegmentsResult{}
	if o == nil {
		return result
	}

	result.SegmentsResponse = newSegmentsResponse(o.SegmentsResponse)
	return result
}
