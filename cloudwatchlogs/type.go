package cloudwatchlogs

import (
	"strings"

	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// QueryResults contains results of GetQueryREsults.
type QueryResults struct {
	List [][]ResultField
}

func NewQueryResults(list [][]SDK.ResultField) QueryResults {
	result := make([][]ResultField, len(list))
	for i, v := range list {
		result[i] = NewResultFieldList(v)
	}
	return QueryResults{
		List: result,
	}
}

// FilterFieldEqual filters the result whose field equals given string.
func (r QueryResults) FilterFieldEqual(s string) [][]ResultField {
	return r.filterField(func(v string) bool {
		return v == s
	})
}

// FilterFieldContains filters the result whose field equals given string.
func (r QueryResults) FilterFieldContains(s string) [][]ResultField {
	return r.filterField(func(v string) bool {
		return strings.Contains(v, s)
	})
}

// FilterValueEqual filters the result whose value equals given string.
func (r QueryResults) FilterValueEqual(s string) [][]ResultField {
	return r.filterValue(func(v string) bool {
		return v == s
	})
}

// FilterValueContains filters the result whose value equals given string.
func (r QueryResults) FilterValueContains(s string) [][]ResultField {
	return r.filterValue(func(v string) bool {
		return strings.Contains(v, s)
	})
}

func (r QueryResults) filterField(fn func(string) bool) [][]ResultField {
	result := make([][]ResultField, 0, len(r.List))
	for _, vv := range r.List {
		has := false
		for _, v := range vv {
			if fn(v.Field) {
				has = true
				break
			}
		}
		if has {
			result = append(result, vv)
		}
	}
	return result
}

func (r QueryResults) filterValue(fn func(string) bool) [][]ResultField {
	result := make([][]ResultField, 0, len(r.List))
	for _, vv := range r.List {
		has := false
		for _, v := range vv {
			if fn(v.Value) {
				has = true
				break
			}
		}
		if has {
			result = append(result, vv)
		}
	}
	return result
}

type ResultField struct {
	Field string
	Value string
}

func NewResultFieldList(list []SDK.ResultField) []ResultField {
	result := make([]ResultField, len(list))
	for i, v := range list {
		result[i] = NewResultField(v)
	}
	return result
}

func NewResultField(o SDK.ResultField) ResultField {
	r := ResultField{}
	if o.Field != nil {
		r.Field = *o.Field
	}
	if o.Value != nil {
		r.Value = *o.Value
	}
	return r
}
