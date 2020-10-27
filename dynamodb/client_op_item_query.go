package dynamodb

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// Query executes `Query` operation.
func (svc *DynamoDB) Query(ctx context.Context, r QueryRequest) (*QueryResult, error) {
	in, err := r.ToInput()
	if err != nil {
		return nil, err
	}

	out, err := svc.RawQuery(ctx, in)
	if err == nil {
		return NewQueryResult(out), nil
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "Query",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// QueryRequest has parameters for `Query` operation.
type QueryRequest struct {
	TableName string

	// optional
	ConsistentRead            bool
	ExclusiveStartKey         map[string]AttributeValue
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]AttributeValue
	FilterExpression          string
	IndexName                 string
	KeyConditionExpression    string
	KeyConditions             map[string]Condition
	Limit                     int64
	ProjectionExpression      string
	ReturnConsumedCapacity    ReturnConsumedCapacity
	ScanIndexForward          bool
	Select                    Select

	XConditions XConditions
}

func (r QueryRequest) ToInput() (*SDK.QueryInput, error) {
	in := &SDK.QueryInput{}

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
	if r.KeyConditionExpression != "" {
		in.KeyConditionExpression = pointers.String(r.KeyConditionExpression)
	}

	if len(r.KeyConditions) != 0 {
		m := make(map[string]SDK.Condition, len(r.KeyConditions))
		for key, val := range r.KeyConditions {
			m[key] = val.ToSDK()
		}
		in.KeyConditions = m
	}

	if r.Limit != 0 {
		in.Limit = pointers.Long64(r.Limit)
	}
	if r.ProjectionExpression != "" {
		in.ProjectionExpression = pointers.String(r.ProjectionExpression)
	}

	in.ReturnConsumedCapacity = SDK.ReturnConsumedCapacity(r.ReturnConsumedCapacity)

	if r.ScanIndexForward {
		in.ScanIndexForward = pointers.Bool(r.ScanIndexForward)
	}

	in.Select = SDK.Select(r.Select)

	if r.XConditions.hasValue() {
		expr, err := r.XConditions.Build()
		if err != nil {
			return nil, err
		}
		if v := expr.Names(); v != nil {
			in.ExpressionAttributeNames = v
		}
		if v := expr.Values(); v != nil {
			in.ExpressionAttributeValues = v
		}
		if v := expr.Filter(); v != nil {
			in.FilterExpression = v
		}
		if v := expr.KeyCondition(); v != nil {
			in.KeyConditionExpression = v
		}
		if v := expr.Projection(); v != nil {
			in.ProjectionExpression = v
		}
	}
	return in, nil
}

// QueryResult contains results from `Query` operation.
type QueryResult struct {
	ConsumedCapacity ConsumedCapacity
	Count            int64
	Items            []map[string]SDK.AttributeValue // keep original type to reduce unmarshal cost
	LastEvaluatedKey map[string]AttributeValue
	ScannedCount     int64
}

func NewQueryResult(output *SDK.QueryResponse) *QueryResult {
	r := &QueryResult{}
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

func (r QueryResult) ToSliceMap() ([]map[string]interface{}, error) {
	return ToSliceMapValues(r.Items), nil
}

func (r QueryResult) Unmarshal(out interface{}) error {
	return RawUnmarshalAttributeValues(r.Items, out)
}
