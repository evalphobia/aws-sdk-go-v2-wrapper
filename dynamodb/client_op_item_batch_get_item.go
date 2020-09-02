package dynamodb

import (
	"context"
	"fmt"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// BatchGetItem executes `BatchGetItem` operation.
func (svc *DynamoDB) BatchGetItem(ctx context.Context, r BatchGetItemRequest) (*BatchGetItemResult, error) {
	out, err := svc.RawBatchGetItem(ctx, r.ToInput())
	if err == nil {
		return NewBatchGetItemResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "BatchGetItem",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// BatchGetItemRequest has parameters for `BatchGetItem` operation.
type BatchGetItemRequest struct {
	RequestItems map[string]KeysAndAttributes

	// optional
	ReturnConsumedCapacity ReturnConsumedCapacity
}

func (r BatchGetItemRequest) ToInput() *SDK.BatchGetItemInput {
	in := &SDK.BatchGetItemInput{}
	if len(r.RequestItems) != 0 {
		m := make(map[string]SDK.KeysAndAttributes)
		for key, val := range r.RequestItems {
			m[key] = val.ToSDK()
		}
		in.RequestItems = m
	}

	in.ReturnConsumedCapacity = SDK.ReturnConsumedCapacity(r.ReturnConsumedCapacity)
	return in
}

// BatchGetItemResult contains results from `BatchGetItem` operation.
type BatchGetItemResult struct {
	ConsumedCapacity []ConsumedCapacity
	Responses        map[string][]map[string]SDK.AttributeValue // keep original type to reduce unmarshal cost
	UnprocessedKeys  map[string]KeysAndAttributes
}

func NewBatchGetItemResult(output *SDK.BatchGetItemResponse) *BatchGetItemResult {
	r := &BatchGetItemResult{}
	if output == nil {
		return r
	}

	r.ConsumedCapacity = newConsumedCapacities(output.ConsumedCapacity)
	r.Responses = output.Responses

	r.UnprocessedKeys = make(map[string]KeysAndAttributes, len(output.UnprocessedKeys))
	for key, val := range output.UnprocessedKeys {
		r.UnprocessedKeys[key] = newKeysAndAttributes(val)
	}
	return r
}

func (r BatchGetItemResult) ToSliceMap(tableName string) ([]map[string]interface{}, error) {
	resp, ok := r.Responses[tableName]
	if !ok {
		return nil, fmt.Errorf("table name [%s] does not exists in the response", tableName)
	}

	return ToSliceMapValues(resp), nil
}

func (r BatchGetItemResult) Unmarshal(tableName string, out interface{}) error {
	resp, ok := r.Responses[tableName]
	if !ok {
		return fmt.Errorf("table name [%s] does not exists in the response", tableName)
	}

	return RawUnmarshalAttributeValues(resp, out)
}
