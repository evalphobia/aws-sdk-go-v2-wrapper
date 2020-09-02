package dynamodb

import (
	"context"
	"fmt"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetSingleItem gets single item.
func (svc *DynamoDB) GetSingleItem(ctx context.Context, in GetSingleItemInput) (map[string]AttributeValue, error) {
	req, err := in.ToRequest()
	if err != nil {
		return nil, err
	}

	result, err := svc.GetItem(ctx, req)
	switch {
	case err != nil,
		result == nil:
		return nil, err
	}

	return result.Item, nil
}

type GetSingleItemInput struct {
	TableName string

	HashKeyName  string
	HashKeyValue interface{}

	RangeKeyName  string
	RangeKeyValue interface{}
}

func (in GetSingleItemInput) ToRequest() (GetItemRequest, error) {
	r := GetItemRequest{
		TableName: in.TableName,
		Key:       make(map[string]AttributeValue),
	}

	v, err := RawMarshal(in.HashKeyValue, defaultStructTag)
	if err != nil {
		return r, err
	}
	if v != nil {
		r.Key[in.HashKeyName] = newAttributeValue(*v)
	}

	if in.RangeKeyName != "" {
		v, err := RawMarshal(in.RangeKeyValue, defaultStructTag)
		if err != nil {
			return r, err
		}
		if v != nil {
			r.Key[in.RangeKeyName] = newAttributeValue(*v)
		}
	}

	return r, nil
}

// ForceDeleteAll deltes all data in the table.
// This performs scan all item and delete it via `BatchWriteItem`.
// Consider using `DeleteTable` instead.
func (svc *DynamoDB) ForceDeleteAll(ctx context.Context, tableName string) error {
	descResp, err := svc.DescribeTable(ctx, DescribeTableRequest{
		TableName: tableName,
	})
	switch {
	case err != nil:
		return err
	case descResp == nil:
		return fmt.Errorf("Unexpected response from `DescribeTable` | [%s]", tableName)
	}
	schema := descResp.Table.KeySchema
	hashKey := schema[0].AttributeName
	rangeKey := ""
	if len(schema) > 1 {
		rangeKey = schema[1].AttributeName
	}

	for {
		result, err := svc.RawScan(ctx, &SDK.ScanInput{
			TableName: pointers.String(tableName),
		})
		switch {
		case err != nil:
			return err
		case result == nil:
			return fmt.Errorf("Unexpected response from `Scan` | [%s]", tableName)
		case *result.Count == 0:
			return nil
		}

		const maxDeleteItems = 25
		itemChunks := sliceItemsToChunks(result.Items, 25)
		for _, items := range itemChunks {
			writeReq := make([]WriteRequest, len(items))
			for i, v := range items {
				delKeys := map[string]AttributeValue{
					hashKey: newAttributeValue(v[hashKey]),
				}
				if rangeKey != "" {
					delKeys[rangeKey] = newAttributeValue(v[rangeKey])
				}
				writeReq[i] = WriteRequest{
					DeleteKeys: delKeys,
				}
			}

			_, err = svc.BatchWriteItem(ctx, BatchWriteItemRequest{
				RequestItems: map[string][]WriteRequest{
					tableName: writeReq,
				},
			})
			if err != nil {
				return err
			}
		}
	}
}

func sliceItemsToChunks(nums []map[string]SDK.AttributeValue, size int) [][]map[string]SDK.AttributeValue {
	var parts [][]map[string]SDK.AttributeValue
	maxSize := len(nums)

	for i := 0; i < maxSize; i += size {
		end := i + size
		if end > maxSize {
			end = maxSize
		}

		parts = append(parts, nums[i:end])
	}
	return parts
}
