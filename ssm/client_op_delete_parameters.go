package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// DeleteParameters executes `DeleteParameters` operation.
func (svc *SSM) DeleteParameters(ctx context.Context, r DeleteParametersRequest) (*DeleteParametersResult, error) {
	out, err := svc.RawDeleteParameters(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteParameters",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewDeleteParametersResult(out), err
}

// DeleteParametersRequest has parameters for `DeleteParameters` operation.
type DeleteParametersRequest struct {
	Names []string
}

func (r DeleteParametersRequest) ToInput() *SDK.DeleteParametersInput {
	in := &SDK.DeleteParametersInput{}
	in.Names = r.Names
	return in
}

type DeleteParametersResult struct {
	DeletedParameters []string
	InvalidParameters []string
}

func NewDeleteParametersResult(o *SDK.DeleteParametersResponse) *DeleteParametersResult {
	result := &DeleteParametersResult{}
	if o == nil {
		return result
	}

	result.DeletedParameters = o.DeletedParameters
	result.InvalidParameters = o.InvalidParameters
	return result
}
