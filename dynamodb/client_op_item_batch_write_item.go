package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// BatchWriteItem executes `BatchWriteItem` operation.
func (svc *DynamoDB) BatchWriteItem(ctx context.Context, r BatchWriteItemRequest) (*BatchWriteItemResult, error) {
	out, err := svc.RawBatchWriteItem(ctx, r.ToInput())
	if err == nil {
		return NewBatchWriteItemResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "BatchWriteItem",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// BatchWriteItemRequest has parameters for `BatchWriteItem` operation.
type BatchWriteItemRequest struct {
	RequestItems map[string][]WriteRequest

	// optional
	ReturnConsumedCapacity      ReturnConsumedCapacity
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics
}

func (r BatchWriteItemRequest) ToInput() *SDK.BatchWriteItemInput {
	in := &SDK.BatchWriteItemInput{}
	if len(r.RequestItems) != 0 {
		m := make(map[string][]SDK.WriteRequest, len(r.RequestItems))
		for key, val := range r.RequestItems {
			list := make([]SDK.WriteRequest, len(val))
			for i, v := range val {
				list[i] = v.ToSDK()
			}
			m[key] = list
		}
		in.RequestItems = m
	}

	in.ReturnConsumedCapacity = SDK.ReturnConsumedCapacity(r.ReturnConsumedCapacity)
	in.ReturnItemCollectionMetrics = SDK.ReturnItemCollectionMetrics(r.ReturnItemCollectionMetrics)
	return in
}

// BatchWriteItemResult contains results from `BatchWriteItem` operation.
type BatchWriteItemResult struct {
	ConsumedCapacity      []ConsumedCapacity
	ItemCollectionMetrics map[string][]ItemCollectionMetrics
	UnprocessedItems      map[string][]WriteRequest
}

func NewBatchWriteItemResult(output *SDK.BatchWriteItemResponse) *BatchWriteItemResult {
	r := &BatchWriteItemResult{}
	if output == nil {
		return r
	}

	r.ConsumedCapacity = newConsumedCapacities(output.ConsumedCapacity)

	r.ItemCollectionMetrics = make(map[string][]ItemCollectionMetrics, len(output.ItemCollectionMetrics))
	for key, val := range output.ItemCollectionMetrics {
		list := make([]ItemCollectionMetrics, len(val))
		for i, v := range val {
			list[i] = newItemCollectionMetrics(v)
		}
		r.ItemCollectionMetrics[key] = list
	}

	r.UnprocessedItems = make(map[string][]WriteRequest, len(output.UnprocessedItems))
	for key, val := range output.UnprocessedItems {
		list := make([]WriteRequest, len(val))
		for i, v := range val {
			list[i] = newWriteRequest(v)
		}
		r.UnprocessedItems[key] = list
	}
	return r
}
