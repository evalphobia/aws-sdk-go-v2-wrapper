package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutItem executes `PutItem` operation.
func (svc *DynamoDB) PutItem(ctx context.Context, r PutItemRequest) (*PutItemResult, error) {
	out, err := svc.RawPutItem(ctx, r.ToInput())
	if err == nil {
		return NewPutItemResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "PutItem",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// PutItemRequest has parameters for `PutItem` operation.
type PutItemRequest struct {
	TableName string
	Item      map[string]AttributeValue

	// optional
	ConditionExpression         string
	ConditionalOperator         ConditionalOperator
	Expected                    map[string]ExpectedAttributeValue
	ExpressionAttributeNames    map[string]string
	ExpressionAttributeValues   map[string]AttributeValue
	ReturnConsumedCapacity      ReturnConsumedCapacity
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics
	ReturnValues                ReturnValue
}

func (r PutItemRequest) ToInput() *SDK.PutItemInput {
	in := &SDK.PutItemInput{}

	if r.TableName != "" {
		in.TableName = pointers.String(r.TableName)
	}

	if len(r.Item) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.Item))
		for key, val := range r.Item {
			m[key] = val.ToSDK()
		}
		in.Item = m
	}

	if r.ConditionExpression != "" {
		in.ConditionExpression = pointers.String(r.ConditionExpression)
	}

	in.ConditionalOperator = SDK.ConditionalOperator(r.ConditionalOperator)

	if len(r.Expected) != 0 {
		m := make(map[string]SDK.ExpectedAttributeValue, len(r.Expected))
		for key, val := range r.Expected {
			m[key] = val.ToSDK()
		}
		in.Expected = m
	}

	in.ExpressionAttributeNames = r.ExpressionAttributeNames

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
	return in
}

// PutItemResult contains results from `PutItem` operation.
type PutItemResult struct {
	Attributes            map[string]AttributeValue
	ConsumedCapacity      ConsumedCapacity
	ItemCollectionMetrics ItemCollectionMetrics
}

func NewPutItemResult(output *SDK.PutItemResponse) *PutItemResult {
	r := &PutItemResult{}
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
