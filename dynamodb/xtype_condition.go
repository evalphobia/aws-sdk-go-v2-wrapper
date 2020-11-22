package dynamodb

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/expression"
)

// XConditions is to build Expression Condition for Query/Scan/Update operation.
type XConditions struct {
	KeyConditions []XCondition
	Conditions    []XCondition
	Filters       []XCondition
	Updates       []XUpdateCondition
	Projections   []string
}

func (x XConditions) hasValue() bool {
	switch {
	case len(x.KeyConditions) != 0,
		len(x.Conditions) != 0,
		len(x.Filters) != 0,
		len(x.Updates) != 0,
		len(x.Projections) != 0:
		return true
	}
	return false
}

func (x XConditions) Build() (expression.Expression, error) {
	b := expression.NewBuilder()

	if len(x.KeyConditions) != 0 {
		kc, err := x.KeyConditions[0].KeyCondition()
		if err != nil {
			return expression.Expression{}, err
		}
		// for multiple key conditions
		if len(x.KeyConditions) > 1 {
			for _, v := range x.KeyConditions[1:] {
				kc2, err := v.KeyCondition()
				if err != nil {
					return expression.Expression{}, err
				}
				kc = kc.And(kc2)
			}
		}
		b = b.WithKeyCondition(kc)
	}

	if len(x.Conditions) != 0 {
		cond, err := x.Conditions[0].Condition()
		if err != nil {
			return expression.Expression{}, err
		}

		// for multiple conditions
		if len(x.Conditions) > 1 {
			for _, v := range x.Conditions[1:] {
				cond2, err := v.Condition()
				if err != nil {
					return expression.Expression{}, err
				}
				if v.IsNOT {
					cond2 = cond2.Not()
				}

				switch {
				case v.IsOR:
					cond = cond.Or(cond2)
				default:
					cond = cond.And(cond2)
				}
			}
		}
		b = b.WithCondition(cond)
	}

	if len(x.Filters) != 0 {
		filt, err := x.Filters[0].Condition()
		if err != nil {
			return expression.Expression{}, err
		}

		// for multiple conditions
		if len(x.Filters) > 1 {
			for _, v := range x.Filters[1:] {
				filt2, err := v.Condition()
				if err != nil {
					return expression.Expression{}, err
				}
				if v.IsNOT {
					filt2 = filt2.Not()
				}

				switch {
				case v.IsOR:
					filt = filt.Or(filt2)
				default:
					filt = filt.And(filt2)
				}
			}
		}
		b = b.WithFilter(filt)
	}

	if len(x.Updates) != 0 {
		cond := x.Updates[0].NewCondition()

		// for multiple conditions
		if len(x.Updates) > 1 {
			for _, v := range x.Updates[1:] {
				cond = v.updateCondition(cond)
			}
		}
		b = b.WithUpdate(cond)
	}

	if len(x.Projections) != 0 {
		list := make([]expression.NameBuilder, len(x.Projections))
		for i, v := range x.Projections {
			list[i] = expression.Name(v)
		}
		b = b.WithProjection(expression.ProjectionBuilder{}.AddNames(list...))
	}

	return b.Build()
}

// XCondition contains single condition parameters.
type XCondition struct {
	Name     string
	Value    interface{}
	Operator ComparisonOperator

	// optional
	IsOR  bool
	IsNOT bool
	// - 'BETWEEN': higher value.
	// - 'IN': all of values afre used besides XCondition.Value.
	OtherValues []string
}

func (x XCondition) KeyCondition() (expression.KeyConditionBuilder, error) {
	e := expression.Key(x.Name)
	switch x.Operator {
	case ComparisonOperatorEq:
		return e.Equal(expression.Value(x.Value)), nil
	case ComparisonOperatorLe:
		return e.LessThanEqual(expression.Value(x.Value)), nil
	case ComparisonOperatorLt:
		return e.LessThan(expression.Value(x.Value)), nil
	case ComparisonOperatorGe:
		return e.GreaterThanEqual(expression.Value(x.Value)), nil
	case ComparisonOperatorGt:
		return e.GreaterThan(expression.Value(x.Value)), nil
	case ComparisonOperatorBeginsWith:
		return e.BeginsWith(fmt.Sprint(x.Value)), nil
	case ComparisonOperatorBetween:
		if len(x.OtherValues) == 0 {
			return expression.KeyConditionBuilder{}, errors.New("Condition 'BETWEEN' must set 'OtherValues[0]'")
		}
		return e.Between(expression.Value(x.Value), expression.Value(x.OtherValues[0])), nil
	default:
		return e.Equal(expression.Value(x.Value)), nil
	}
}

func (x XCondition) Condition() (expression.ConditionBuilder, error) {
	e := expression.Name(x.Name)
	switch x.Operator {
	case ComparisonOperatorEq:
		return e.Equal(expression.Value(x.Value)), nil
	case ComparisonOperatorLe:
		return e.LessThanEqual(expression.Value(x.Value)), nil
	case ComparisonOperatorLt:
		return e.LessThan(expression.Value(x.Value)), nil
	case ComparisonOperatorGe:
		return e.GreaterThanEqual(expression.Value(x.Value)), nil
	case ComparisonOperatorGt:
		return e.GreaterThan(expression.Value(x.Value)), nil
	case ComparisonOperatorBeginsWith:
		return e.BeginsWith(fmt.Sprint(x.Value)), nil
	case ComparisonOperatorBetween:
		if len(x.OtherValues) == 0 {
			return expression.ConditionBuilder{}, errors.New("Condition 'BETWEEN' must set 'OtherValues[0]'")
		}
		return e.Between(expression.Value(x.Value), expression.Value(x.OtherValues[0])), nil
	case ComparisonOperatorIn:
		list := make([]expression.OperandBuilder, len(x.OtherValues))
		for i, v := range x.OtherValues {
			list[i] = expression.Value(v)
		}
		return e.In(expression.Value(x.Value), list...), nil
	case ComparisonOperatorNe:
		return e.NotEqual(expression.Value(x.Value)), nil
	case ComparisonOperatorContains:
		return e.Contains(fmt.Sprint(x.Value)), nil
	case ComparisonOperatorAttrExists:
		return e.AttributeExists(), nil
	case ComparisonOperatorAttrNotExists:
		return e.AttributeNotExists(), nil
	case ComparisonOperatorAttrType:
		return e.AttributeType(expression.DynamoDBAttributeType(fmt.Sprint(x.Value))), nil
	default:
		return e.Equal(expression.Value(x.Value)), nil
	}
}

type XUpdateCondition struct {
	Name      string
	Value     interface{}
	Operation OperationMode

	SetType       SetType
	SetTypeKey    string
	SetTypeValue2 interface{}
}

func (x XUpdateCondition) NewCondition() expression.UpdateBuilder {
	return x.updateCondition(expression.UpdateBuilder{})
}

func (x XUpdateCondition) updateCondition(b expression.UpdateBuilder) expression.UpdateBuilder {
	switch x.Operation {
	case OperationModeSET:
		return x.updateConditionSet(b)
	case OperationModeREMOVE:
		return b.Remove(expression.Name(x.Name))
	case OperationModeADD:
		return b.Add(expression.Name(x.Name), expression.Value(x.Value))
	default:
		return x.updateConditionSet(b)
	}
}

func (x XUpdateCondition) updateConditionSet(b expression.UpdateBuilder) expression.UpdateBuilder {
	name := expression.Name(x.Name)
	value := expression.Value(x.Value)
	switch x.SetType {
	case SetTypePlus:
		return b.Set(name, expression.Plus(value, expression.Value(x.SetTypeValue2)))
	case SetTypeMinus:
		return b.Set(name, expression.Minus(value, expression.Value(x.SetTypeValue2)))
	case SetTypeListAppend:
		return b.Set(name, expression.ListAppend(value, expression.Value(x.SetTypeValue2)))
	case SetTypeIfNotExists:
		return b.Set(name, expression.IfNotExists(expression.Name(x.SetTypeKey), value))
	default:
		return b.Set(name, value)
	}
}

type OperationMode string

const (
	OperationModeSET    OperationMode = "SET"
	OperationModeREMOVE OperationMode = "REMOVE"
	OperationModeADD    OperationMode = "ADD"
	OperationModeDELETE OperationMode = "DELETE"
)

type SetType string

const (
	SetTypePlus        SetType = "PLUS"
	SetTypeMinus       SetType = "MINUS"
	SetTypeListAppend  SetType = "LIST_APPEND"
	SetTypeIfNotExists SetType = "IF_NOT_EXISTS"
)
