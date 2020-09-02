package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DeleteTable executes `DeleteTable` operation.
func (svc *DynamoDB) DeleteTable(ctx context.Context, r DeleteTableRequest) (*DeleteTableResult, error) {
	out, err := svc.RawDeleteTable(ctx, r.ToInput())
	if err == nil {
		return NewDeleteTableResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "DeleteTable",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// DeleteTableRequest has parameters for `DeleteTable` operation.
type DeleteTableRequest struct {
	TableName string
}

func (r DeleteTableRequest) ToInput() *SDK.DeleteTableInput {
	in := &SDK.DeleteTableInput{}

	if r.TableName != "" {
		in.TableName = pointers.String(r.TableName)
	}
	return in
}

// DeleteTableResult contains results from `DeleteTable` operation.
type DeleteTableResult struct {
	TableDescription TableDescription
}

func NewDeleteTableResult(output *SDK.DeleteTableResponse) *DeleteTableResult {
	r := &DeleteTableResult{}
	if output == nil {
		return r
	}

	if output.TableDescription != nil {
		r.TableDescription = newTableDescription(*output.TableDescription)
	}
	return r
}
