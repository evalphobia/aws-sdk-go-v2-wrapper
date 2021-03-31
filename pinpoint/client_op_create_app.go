package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateApp executes `CreateApp` operation.
func (svc *Pinpoint) CreateApp(ctx context.Context, r CreateAppRequest) (*CreateAppResult, error) {
	out, err := svc.RawCreateApp(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateApp",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateAppResult(out), nil
}

// CreateAppRequest has parameters for `CreateApp` operation.
type CreateAppRequest struct {
	Name string

	// optional
	Tags map[string]string
}

func (r CreateAppRequest) ToInput() *SDK.CreateAppInput {
	in := &SDK.CreateAppInput{
		CreateApplicationRequest: &SDK.CreateApplicationRequest{},
	}
	if r.Name != "" {
		in.CreateApplicationRequest.Name = pointers.String(r.Name)
	}
	in.CreateApplicationRequest.Tags = r.Tags
	return in
}

type CreateAppResult struct {
	ApplicationResponse
}

func NewCreateAppResult(o *SDK.CreateAppResponse) *CreateAppResult {
	result := &CreateAppResult{}
	if o == nil {
		return result
	}

	result.ApplicationResponse = newApplicationResponse(o.ApplicationResponse)
	return result
}
