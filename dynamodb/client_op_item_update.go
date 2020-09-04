package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// UpdateItem executes `UpdateItem` operation.
func (svc *DynamoDB) UpdateItem(ctx context.Context, r UpdateItemRequest) (*UpdateItemResult, error) {
	in, err := r.ToInput()
	if err != nil {
		return nil, err
	}

	out, err := svc.RawUpdateItem(ctx, in)
	if err == nil {
		return NewUpdateItemResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "UpdateItem",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// UpdateItemRequest has parameters for `UpdateItem` operation.
type UpdateItemRequest struct {
	TableName string
	Key       map[string]AttributeValue

	// optional
	ConditionExpression         string
	ExpressionAttributeNames    map[string]string
	ExpressionAttributeValues   map[string]AttributeValue
	ReturnConsumedCapacity      ReturnConsumedCapacity
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics
	ReturnValues                ReturnValue
	UpdateExpression            string

	XConditions XConditions
}

func (r UpdateItemRequest) ToInput() (*SDK.UpdateItemInput, error) {
	in := &SDK.UpdateItemInput{}

	if r.TableName != "" {
		in.TableName = pointers.String(r.TableName)
	}
	if len(r.Key) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.Key))
		for key, val := range r.Key {
			m[key] = val.ToSDK()
		}
		in.Key = m
	}

	if r.ConditionExpression != "" {
		in.ConditionExpression = pointers.String(r.ConditionExpression)
	}

	if len(r.ExpressionAttributeValues) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.ExpressionAttributeValues))
		for key, val := range r.ExpressionAttributeValues {
			m[key] = val.ToSDK()
		}
		in.ExpressionAttributeValues = m
	}

	in.ReturnConsumedCapacity = SDK.ReturnConsumedCapacity(r.ReturnConsumedCapacity)
	in.ReturnItemCollectionMetrics = SDK.ReturnItemCollectionMetrics(r.ReturnItemCollectionMetrics)
	in.ReturnValues = SDK.ReturnValue(r.ReturnValues)

	if r.UpdateExpression != "" {
		in.UpdateExpression = pointers.String(r.UpdateExpression)
	}

	if r.XConditions.hasValue() {
		expr, err := r.XConditions.Build()
		if err != nil {
			return nil, err
		}
		in.ConditionExpression = expr.Condition()
		in.ExpressionAttributeNames = expr.Names()
		in.ExpressionAttributeValues = expr.Values()
		in.UpdateExpression = expr.Update()
	}
	return in, nil
}

// UpdateItemResult contains results from `UpdateItem` operation.
type UpdateItemResult struct {
	Attributes            map[string]AttributeValue
	ConsumedCapacity      ConsumedCapacity
	ItemCollectionMetrics ItemCollectionMetrics
}

func NewUpdateItemResult(output *SDK.UpdateItemResponse) *UpdateItemResult {
	r := &UpdateItemResult{}
	if output == nil {
		return r
	}

	r.Attributes = newAttributeValueMap(output.Attributes)

	if output.ConsumedCapacity != nil {
		r.ConsumedCapacity = newConsumedCapacity(*output.ConsumedCapacity)
	}
	if output.ItemCollectionMetrics != nil {
		r.ItemCollectionMetrics = newItemCollectionMetrics(*output.ItemCollectionMetrics)
	}
	return r
}
