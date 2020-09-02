package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetItem executes `GetItem` operation.
func (svc *DynamoDB) GetItem(ctx context.Context, r GetItemRequest) (*GetItemResult, error) {
	out, err := svc.RawGetItem(ctx, r.ToInput())
	if err == nil {
		return NewGetItemResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "GetItem",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// GetItemRequest has parameters for `GetItem` operation.
type GetItemRequest struct {
	TableName string
	Key       map[string]AttributeValue

	// optional
	AttributesToGet          []string
	ConsistentRead           bool
	ExpressionAttributeNames map[string]string
	ProjectionExpression     string
	ReturnConsumedCapacity   ReturnConsumedCapacity
}

func (r GetItemRequest) ToInput() *SDK.GetItemInput {
	in := &SDK.GetItemInput{}

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

	in.AttributesToGet = r.AttributesToGet

	if r.ConsistentRead {
		in.ConsistentRead = pointers.Bool(r.ConsistentRead)
	}

	in.ExpressionAttributeNames = r.ExpressionAttributeNames

	if r.ProjectionExpression != "" {
		in.ProjectionExpression = pointers.String(r.ProjectionExpression)
	}

	in.ReturnConsumedCapacity = SDK.ReturnConsumedCapacity(r.ReturnConsumedCapacity)
	return in
}

// GetItemResult contains results from `GetItem` operation.
type GetItemResult struct {
	ConsumedCapacity ConsumedCapacity
	Item             map[string]AttributeValue
}

func NewGetItemResult(output *SDK.GetItemResponse) *GetItemResult {
	r := &GetItemResult{}
	if output == nil {
		return r
	}

	if output.ConsumedCapacity != nil {
		r.ConsumedCapacity = newConsumedCapacity(*output.ConsumedCapacity)
	}

	r.Item = newAttributeValueMap(output.Item)
	return r
}
