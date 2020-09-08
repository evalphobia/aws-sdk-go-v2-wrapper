package dynamodb

import (
	"context"
	"fmt"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
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

func (svc *DynamoDB) BatchDeleteItems(ctx context.Context, req BatchDeleteItemRequest) error {
	tableName := req.TableName
	hashKey := req.HashKey
	rangeKey := req.RangeKey

	var errs errors.ErrorList
	itemChunks := req.ToChunks()
	for _, items := range itemChunks {
		writeReq := make([]WriteRequest, len(items))
		for i, v := range items {
			av, err := RawMarshal(v.HashKeyValue, defaultStructTag)
			if err != nil {
				return err
			}

			delKeys := map[string]AttributeValue{
				hashKey: newAttributeValue(*av),
			}

			if rangeKey != "" {
				av, err := RawMarshal(v.RangeKeyValue, defaultStructTag)
				if err != nil {
					return err
				}
				delKeys[rangeKey] = newAttributeValue(*av)
			}
			writeReq[i] = WriteRequest{
				DeleteKeys: delKeys,
			}
		}

		_, err := svc.BatchWriteItem(ctx, BatchWriteItemRequest{
			RequestItems: map[string][]WriteRequest{
				tableName: writeReq,
			},
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

type BatchDeleteItemRequest struct {
	TableName string
	HashKey   string
	RangeKey  string
	Items     []BatchDeleteItem
}

func (r BatchDeleteItemRequest) ToChunks() [][]BatchDeleteItem {
	const ddbSize = 25
	var parts [][]BatchDeleteItem

	items := r.Items
	maxSize := len(items)

	for i := 0; i < maxSize; i += ddbSize {
		end := i + ddbSize
		if end > maxSize {
			end = maxSize
		}

		parts = append(parts, items[i:end])
	}
	return parts
}

type BatchDeleteItem struct {
	HashKeyValue  interface{}
	RangeKeyValue interface{}
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

		items := make([]BatchDeleteItem, len(result.Items))
		for i, v := range result.Items {
			item := BatchDeleteItem{
				HashKeyValue: newAttributeValue(v[hashKey]).GetValue(),
			}
			if rangeKey != "" {
				item.RangeKeyValue = newAttributeValue(v[rangeKey]).GetValue()
			}
			items[i] = item
		}

		r := BatchDeleteItemRequest{
			TableName: tableName,
			HashKey:   hashKey,
			RangeKey:  rangeKey,
			Items:     items,
		}
		if err := svc.BatchDeleteItems(ctx, r); err != nil {
			return err
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
