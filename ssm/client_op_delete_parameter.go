package ssm

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteParameter executes `DeleteParameter` operation.
func (svc *SSM) DeleteParameter(ctx context.Context, r DeleteParameterRequest) error {
	_, err := svc.RawDeleteParameter(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "DeleteParameter",
		})
		svc.Errorf(err.Error())
	}
	return err
}

// DeleteParameterRequest has parameters for `DeleteParameter` operation.
type DeleteParameterRequest struct {
	Name string
}

func (r DeleteParameterRequest) ToInput() *SDK.DeleteParameterInput {
	in := &SDK.DeleteParameterInput{}
	if r.Name != "" {
		in.Name = pointers.String(r.Name)
	}
	return in
}
