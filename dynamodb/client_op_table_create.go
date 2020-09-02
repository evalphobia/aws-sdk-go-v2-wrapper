package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// CreateTable executes `CreateTable` operation.
func (svc *DynamoDB) CreateTable(ctx context.Context, r CreateTableRequest) (*CreateTableResult, error) {
	out, err := svc.RawCreateTable(ctx, r.ToInput())
	if err == nil {
		return NewCreateTableResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "CreateTable",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// CreateTableRequest has parameters for `CreateTable` operation.
type CreateTableRequest struct {
	TableName            string
	AttributeDefinitions []AttributeDefinition
	KeySchema            []KeySchemaElement

	// optional
	BillingMode            BillingMode
	GlobalSecondaryIndexes []GlobalSecondaryIndex
	LocalSecondaryIndexes  []LocalSecondaryIndex
	ProvisionedThroughput  ProvisionedThroughput
	SSESpecification       SSESpecification
	StreamSpecification    StreamSpecification
	Tags                   []Tag
}

func (r CreateTableRequest) ToInput() *SDK.CreateTableInput {
	in := &SDK.CreateTableInput{}

	if r.TableName != "" {
		in.TableName = pointers.String(r.TableName)
	}
	if len(r.AttributeDefinitions) != 0 {
		list := make([]SDK.AttributeDefinition, len(r.AttributeDefinitions))
		for i, v := range r.AttributeDefinitions {
			list[i] = v.ToSDK()
		}
		in.AttributeDefinitions = list
	}
	if len(r.KeySchema) != 0 {
		list := make([]SDK.KeySchemaElement, len(r.KeySchema))
		for i, v := range r.KeySchema {
			list[i] = v.ToSDK()
		}
		in.KeySchema = list
	}

	in.BillingMode = SDK.BillingMode(r.BillingMode)

	if len(r.GlobalSecondaryIndexes) != 0 {
		list := make([]SDK.GlobalSecondaryIndex, len(r.GlobalSecondaryIndexes))
		for i, v := range r.GlobalSecondaryIndexes {
			list[i] = v.ToSDK()
		}
		in.GlobalSecondaryIndexes = list
	}
	if len(r.LocalSecondaryIndexes) != 0 {
		list := make([]SDK.LocalSecondaryIndex, len(r.LocalSecondaryIndexes))
		for i, v := range r.LocalSecondaryIndexes {
			list[i] = v.ToSDK()
		}
		in.LocalSecondaryIndexes = list
	}

	if r.ProvisionedThroughput.hasValue() {
		v := r.ProvisionedThroughput.ToSDK()
		in.ProvisionedThroughput = &v
	}
	if r.SSESpecification.hasValue() {
		v := r.SSESpecification.ToSDK()
		in.SSESpecification = &v
	}
	if r.StreamSpecification.hasValue() {
		v := r.StreamSpecification.ToSDK()
		in.StreamSpecification = &v
	}

	if len(r.Tags) != 0 {
		list := make([]SDK.Tag, len(r.Tags))
		for i, v := range r.Tags {
			list[i] = v.ToSDK()
		}
		in.Tags = list
	}
	return in
}

// CreateTableResult contains results from `CreateTable` operation.
type CreateTableResult struct {
	TableDescription TableDescription
}

func NewCreateTableResult(output *SDK.CreateTableResponse) *CreateTableResult {
	r := &CreateTableResult{}
	if output == nil {
		return r
	}

	if output.TableDescription != nil {
		r.TableDescription = newTableDescription(*output.TableDescription)
	}
	return r
}
