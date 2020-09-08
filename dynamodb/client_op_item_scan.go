package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// Scan executes `Scan` operation.
func (svc *DynamoDB) Scan(ctx context.Context, r ScanRequest) (*ScanResult, error) {
	in, err := r.ToInput()
	if err != nil {
		return nil, err
	}

	out, err := svc.RawScan(ctx, in)
	if err == nil {
		return NewScanResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "Scan",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// ScanRequest has parameters for `Scan` operation.
type ScanRequest struct {
	TableName string

	// optional
	ConsistentRead            bool
	ExclusiveStartKey         map[string]AttributeValue
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]AttributeValue
	FilterExpression          string
	IndexName                 string
	Limit                     int64
	ProjectionExpression      string
	ReturnConsumedCapacity    ReturnConsumedCapacity
	Segment                   int64
	Select                    Select
	TotalSegments             int64

	XConditions XConditions
}

func (r ScanRequest) ToInput() (*SDK.ScanInput, error) {
	in := &SDK.ScanInput{}

	if r.TableName != "" {
		in.TableName = pointers.String(r.TableName)
	}

	if r.ConsistentRead {
		in.ConsistentRead = pointers.Bool(r.ConsistentRead)
	}

	if len(r.ExclusiveStartKey) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.ExclusiveStartKey))
		for key, val := range r.ExclusiveStartKey {
			m[key] = val.ToSDK()
		}
		in.ExclusiveStartKey = m
	}

	in.ExpressionAttributeNames = r.ExpressionAttributeNames

	if len(r.ExpressionAttributeValues) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.ExpressionAttributeValues))
		for key, val := range r.ExpressionAttributeValues {
			m[key] = val.ToSDK()
		}
		in.ExpressionAttributeValues = m
	}

	if r.FilterExpression != "" {
		in.FilterExpression = pointers.String(r.FilterExpression)
	}
	if r.IndexName != "" {
		in.IndexName = pointers.String(r.IndexName)
	}

	if r.Limit != 0 {
		in.Limit = pointers.Long64(r.Limit)
	}
	if r.ProjectionExpression != "" {
		in.ProjectionExpression = pointers.String(r.ProjectionExpression)
	}

	in.ReturnConsumedCapacity = SDK.ReturnConsumedCapacity(r.ReturnConsumedCapacity)

	in.Select = SDK.Select(r.Select)

	if r.XConditions.hasValue() {
		expr, err := r.XConditions.Build()
		if err != nil {
			return nil, err
		}
		in.ExpressionAttributeNames = expr.Names()
		in.ExpressionAttributeValues = expr.Values()
		in.FilterExpression = expr.Filter()
		in.ProjectionExpression = expr.Projection()
	}
	return in, nil
}

// ScanResult contains results from `Scan` operation.
type ScanResult struct {
	ConsumedCapacity ConsumedCapacity
	Count            int64
	Items            []map[string]SDK.AttributeValue // keep original type to reduce unmarshal cost
	LastEvaluatedKey map[string]AttributeValue
	ScannedCount     int64
}

func NewScanResult(output *SDK.ScanResponse) *ScanResult {
	r := &ScanResult{}
	if output == nil {
		return r
	}

	if output.ConsumedCapacity != nil {
		r.ConsumedCapacity = newConsumedCapacity(*output.ConsumedCapacity)
	}

	if output.Count != nil {
		r.Count = *output.Count
	}

	r.Items = output.Items
	r.LastEvaluatedKey = newAttributeValueMap(output.LastEvaluatedKey)

	if output.ScannedCount != nil {
		r.ScannedCount = *output.ScannedCount
	}
	return r
}

func (r ScanResult) ToSliceMap() ([]map[string]interface{}, error) {
	return ToSliceMapValues(r.Items), nil
}

func (r ScanResult) Unmarshal(out interface{}) error {
	return RawUnmarshalAttributeValues(r.Items, out)
}
