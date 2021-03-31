package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateCampaign executes `CreateCampaign` operation.
func (svc *Pinpoint) CreateCampaign(ctx context.Context, r CreateCampaignRequest) (*CreateCampaignResult, error) {
	out, err := svc.RawCreateCampaign(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateCampaign",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateCampaignResult(out), nil
}

// CreateCampaignRequest has parameters for `CreateCampaign` operation.
type CreateCampaignRequest struct {
	ApplicationID        string
	WriteCampaignRequest WriteCampaignRequest
}

func (r CreateCampaignRequest) ToInput() *SDK.CreateCampaignInput {
	in := &SDK.CreateCampaignInput{}

	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	in.WriteCampaignRequest = r.WriteCampaignRequest.ToSDK()
	return in
}

type CreateCampaignResult struct {
	CampaignResponse
}

func NewCreateCampaignResult(o *SDK.CreateCampaignResponse) *CreateCampaignResult {
	result := &CreateCampaignResult{}
	if o == nil {
		return result
	}

	result.CampaignResponse = newCampaignResponse(o.CampaignResponse)
	return result
}
