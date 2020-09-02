package dynamodb

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

const defaultStructTag = "dynamodb"

// MarshalFromMap converts result AttributeValue to slice of map values.
func MarshalFromMap(o map[string]interface{}) (map[string]AttributeValue, error) {
	m, err := dynamodbattribute.MarshalMap(o)
	if err != nil {
		return nil, err
	}

	return newAttributeValueMap(m), nil
}

// MarshalToMap converts result AttributeValue to map values.
func MarshalToMap(in interface{}) (map[string]AttributeValue, error) {
	out, err := RawMarshal(in, defaultStructTag)
	if err != nil {
		return nil, err
	}

	if out != nil {
		return newAttributeValueMap(out.M), nil
	}
	return nil, nil
}

// MarshalToMap converts result AttributeValue to list values.
func MarshalToList(in interface{}) ([]AttributeValue, error) {
	out, err := RawMarshal(in, defaultStructTag)
	if err != nil {
		return nil, err
	}

	if out != nil {
		return newAttributeValueList(out.L), nil
	}
	return nil, nil
}

func RawMarshal(in interface{}, structTag string) (*SDK.AttributeValue, error) {
	encoder := dynamodbattribute.NewEncoder()
	encoder.TagKey = structTag

	return encoder.Encode(in)
}

// ToSliceMapValues converts result AttributeValue to slice of map values.
func ToSliceMapValues(o []map[string]SDK.AttributeValue) []map[string]interface{} {
	list := make([]map[string]interface{}, len(o))
	for i, item := range o {
		var v map[string]interface{}
		_ = dynamodbattribute.UnmarshalMap(item, &v)
		list[i] = v
	}
	return list
}

// ToMapValue converts result AttributeValue to map value.
func ToMapValue(o map[string]AttributeValue) map[string]interface{} {
	m := make(map[string]SDK.AttributeValue, len(o))
	for key, val := range o {
		m[key] = val.ToSDK()
	}

	var v map[string]interface{}
	_ = dynamodbattribute.UnmarshalMap(m, &v)
	return v
}

// Unmarshal unmarshals given slice pointer sturct from DynamoDB item result to mapping.
//     e.g. err = Unmarshal(&[]*yourStruct)
// The struct tag `dynamodb:""` is used to unmarshal.
func RawUnmarshalAttributeValues(items []map[string]SDK.AttributeValue, out interface{}) error {
	return RawUnmarshalWithTagName(items, out, defaultStructTag)
}

// RawUnmarshalWithTagName unmarshals given slice pointer sturct and tag name from DynamoDB item result to mapping.
func RawUnmarshalWithTagName(items []map[string]SDK.AttributeValue, out interface{}, structTag string) error {
	if len(items) == 0 {
		return nil
	}

	decoder := dynamodbattribute.NewDecoder()
	decoder.TagKey = structTag

	list := make([]SDK.AttributeValue, len(items))
	for i, m := range items {
		list[i] = SDK.AttributeValue{M: m}
	}
	err := decoder.Decode(&SDK.AttributeValue{L: list}, out)
	return err
}
