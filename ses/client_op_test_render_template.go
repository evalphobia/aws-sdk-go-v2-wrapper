package ses

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// GetSendQuota executes `GetSendQuota` operation.
func (svc *SES) GetSendQuota(ctx context.Context, r GetSendQuotaRequest) (*GetSendQuotaResult, error) {
	out, err := svc.RawGetSendQuota(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetSendQuota",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetSendQuotaResult(out), nil
}

// GetSendQuotaRequest has parameters for `GetSendQuota` operation.
type GetSendQuotaRequest struct {
}

func (r GetSendQuotaRequest) ToInput() *SDK.GetSendQuotaInput {
	in := &SDK.GetSendQuotaInput{}
	return in
}

type GetSendQuotaResult struct {
	Max24HourSend   float64
	MaxSendRate     float64
	SentLast24Hours float64
}

func NewGetSendQuotaResult(o *SDK.GetSendQuotaResponse) *GetSendQuotaResult {
	result := &GetSendQuotaResult{}
	if o == nil {
		return result
	}

	if o.Max24HourSend != nil {
		result.Max24HourSend = *o.Max24HourSend
	}
	if o.MaxSendRate != nil {
		result.MaxSendRate = *o.MaxSendRate
	}
	if o.SentLast24Hours != nil {
		result.SentLast24Hours = *o.SentLast24Hours
	}
	return result
}
