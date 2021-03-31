package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateImportJob executes `CreateImportJob` operation.
func (svc *Pinpoint) CreateImportJob(ctx context.Context, r CreateImportJobRequest) (*CreateImportJobResult, error) {
	out, err := svc.RawCreateImportJob(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "CreateImportJob",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewCreateImportJobResult(out), nil
}

// CreateImportJobRequest has parameters for `CreateImportJob` operation.
type CreateImportJobRequest struct {
	ApplicationID    string
	ImportJobRequest ImportJobRequest
}

func (r CreateImportJobRequest) ToInput() *SDK.CreateImportJobInput {
	in := &SDK.CreateImportJobInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}

	in.ImportJobRequest = r.ImportJobRequest.ToSDK()
	return in
}

type CreateImportJobResult struct {
	ImportJobResponse
}

func NewCreateImportJobResult(o *SDK.CreateImportJobResponse) *CreateImportJobResult {
	result := &CreateImportJobResult{}
	if o == nil {
		return result
	}

	result.ImportJobResponse = newImportJobResponse(o.ImportJobResponse)
	return result
}
