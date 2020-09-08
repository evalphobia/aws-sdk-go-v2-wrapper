package dynamodb

import (
	"context"
	"fmt"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// XGetSingleItem gets single item.
func (svc *DynamoDB) XGetSingleItem(ctx context.Context, in XGetSingleItemRequest) (map[string]AttributeValue, error) {
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

type XGetSingleItemRequest struct {
	TableName string

	HashKeyName  string
	HashKeyValue interface{}

	RangeKeyName  string
	RangeKeyValue interface{}
}

func (in XGetSingleItemRequest) ToRequest() (GetItemRequest, error) {
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

// XBatchDeleteItems deletes multiple items using 'BatchWriteItems'.
func (svc *DynamoDB) XBatchDeleteItems(ctx context.Context, req XBatchDeleteItemRequest) error {
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

// XBatchDeleteItemRequest is parameters of 'XBatchDeleteItems'.
type XBatchDeleteItemRequest struct {
	TableName string
	HashKey   string
	RangeKey  string
	Items     []XBatchDeleteItem
}

// ToChunks makes a slice of 25 items slices to avoid the limitation of 'BatchWriteItem'.
func (r XBatchDeleteItemRequest) ToChunks() [][]XBatchDeleteItem {
	const ddbSize = 25
	var parts [][]XBatchDeleteItem

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

// XBatchDeleteItem contains key values to delete and used in 'XBatchDeleteItemRequest'.
type XBatchDeleteItem struct {
	HashKeyValue  interface{}
	RangeKeyValue interface{}
}

// XForceDeleteAll deletes all data in the table.
// This performs scan all item and delete it via 'BatchWriteItem'.
// For big tables, consider using 'DeleteTable' instead.
func (svc *DynamoDB) XForceDeleteAll(ctx context.Context, tableName string) error {
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

		items := make([]XBatchDeleteItem, len(result.Items))
		for i, v := range result.Items {
			item := XBatchDeleteItem{
				HashKeyValue: newAttributeValue(v[hashKey]).GetValue(),
			}
			if rangeKey != "" {
				item.RangeKeyValue = newAttributeValue(v[rangeKey]).GetValue()
			}
			items[i] = item
		}

		r := XBatchDeleteItemRequest{
			TableName: tableName,
			HashKey:   hashKey,
			RangeKey:  rangeKey,
			Items:     items,
		}
		if err := svc.XBatchDeleteItems(ctx, r); err != nil {
			return err
		}
	}
}
