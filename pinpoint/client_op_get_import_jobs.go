package pinpoint

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetImportJobs executes `GetImportJobs` operation.
func (svc *Pinpoint) GetImportJobs(ctx context.Context, r GetImportJobsRequest) (*GetImportJobsResult, error) {
	out, err := svc.RawGetImportJobs(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetImportJobs",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetImportJobsResult(out), nil
}

// GetImportJobsRequest has parameters for `GetImportJobs` operation.
type GetImportJobsRequest struct {
	ApplicationID string
	PageSize      string
	Token         string
}

func (r GetImportJobsRequest) ToInput() *SDK.GetImportJobsInput {
	in := &SDK.GetImportJobsInput{}
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

type GetImportJobsResult struct {
	ImportJobsResponse
}

func NewGetImportJobsResult(o *SDK.GetImportJobsResponse) *GetImportJobsResult {
	result := &GetImportJobsResult{}
	if o == nil {
		return result
	}

	result.ImportJobsResponse = newImportJobsResponse(o.ImportJobsResponse)
	return result
}
