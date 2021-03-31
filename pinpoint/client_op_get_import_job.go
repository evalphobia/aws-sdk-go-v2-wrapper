package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetImportJob executes `GetImportJob` operation.
func (svc *Pinpoint) GetImportJob(ctx context.Context, r GetImportJobRequest) (*GetImportJobResult, error) {
	out, err := svc.RawGetImportJob(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetImportJob",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetImportJobResult(out), nil
}

// GetImportJobRequest has parameters for `GetImportJob` operation.
type GetImportJobRequest struct {
	ApplicationID string
	JobID         string
}

func (r GetImportJobRequest) ToInput() *SDK.GetImportJobInput {
	in := &SDK.GetImportJobInput{}
	if r.ApplicationID != "" {
		in.ApplicationId = pointers.String(r.ApplicationID)
	}
	if r.JobID != "" {
		in.JobId = pointers.String(r.JobID)
	}
	return in
}

type GetImportJobResult struct {
	ImportJobResponse
}

func NewGetImportJobResult(o *SDK.GetImportJobResponse) *GetImportJobResult {
	result := &GetImportJobResult{}
	if o == nil {
		return result
	}

	result.ImportJobResponse = newImportJobResponse(o.ImportJobResponse)
	return result
}
