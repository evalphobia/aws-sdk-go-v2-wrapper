package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetApps executes `GetApps` operation.
func (svc *Pinpoint) GetApps(ctx context.Context, r GetAppsRequest) (*GetAppsResult, error) {
	out, err := svc.RawGetApps(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetApps",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetAppsResult(out), nil
}

// GetAppsRequest has parameters for `GetApps` operation.
type GetAppsRequest struct {
	PageSize string
	Token    string
}

func (r GetAppsRequest) ToInput() *SDK.GetAppsInput {
	in := &SDK.GetAppsInput{}
	if r.PageSize != "" {
		in.PageSize = pointers.String(r.PageSize)
	}
	if r.Token != "" {
		in.Token = pointers.String(r.Token)
	}
	return in
}

type GetAppsResult struct {
	ApplicationsResponse
}

func NewGetAppsResult(o *SDK.GetAppsResponse) *GetAppsResult {
	result := &GetAppsResult{}
	if o == nil {
		return result
	}

	result.ApplicationsResponse = newApplicationsResponse(o.ApplicationsResponse)
	return result
}
