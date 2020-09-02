package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// DescribeTable executes `DescribeTable` operation.
func (svc *DynamoDB) DescribeTable(ctx context.Context, r DescribeTableRequest) (*DescribeTableResult, error) {
	out, err := svc.RawDescribeTable(ctx, r.ToInput())
	if err == nil {
		return NewDescribeTableResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "DescribeTable",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// DescribeTableRequest has parameters for `DescribeTable` operation.
type DescribeTableRequest struct {
	TableName string
}

func (r DescribeTableRequest) ToInput() *SDK.DescribeTableInput {
	in := &SDK.DescribeTableInput{}

	if r.TableName != "" {
		in.TableName = pointers.String(r.TableName)
	}
	return in
}

// DescribeTableResult contains results from `DescribeTable` operation.
type DescribeTableResult struct {
	Table TableDescription
}

func NewDescribeTableResult(output *SDK.DescribeTableResponse) *DescribeTableResult {
	r := &DescribeTableResult{}
	if output == nil {
		return r
	}

	if output.Table != nil {
		r.Table = newTableDescription(*output.Table)
	}
	return r
}
