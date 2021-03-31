package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetApp executes `GetApp` operation.
func (svc *Pinpoint) GetApp(ctx context.Context, r GetAppRequest) (*GetAppResult, error) {
	out, err := svc.RawGetApp(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetApp",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetAppResult(out), nil
}

// GetAppRequest has parameters for `GetApp` operation.
type GetAppRequest struct {
	ApplicationID string
}

func (r GetAppRequest) ToInput() *SDK.GetAppInput {
	in := &SDK.GetAppInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	return in
}

type GetAppResult struct {
	ApplicationResponse
}

func NewGetAppResult(o *SDK.GetAppResponse) *GetAppResult {
	result := &GetAppResult{}
	if o == nil {
		return result
	}

	result.ApplicationResponse = newApplicationResponse(o.ApplicationResponse)
	return result
}
