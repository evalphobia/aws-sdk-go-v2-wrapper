package ssm

import (
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type Parameter struct {
	ARN              string
	DataType         string
	LastModifiedDate time.Time
	Name             string
	Selector         string
	SourceResult     string
	Type             ParameterType
	Value            string
	Version          int64
}

func newParameter(o *SDK.Parameter) Parameter {
	result := Parameter{}
	if o == nil {
		return result
	}

	if o.ARN != nil {
		result.ARN = *o.ARN
	}
	if o.DataType != nil {
		result.DataType = *o.DataType
	}
	if o.LastModifiedDate != nil {
		result.LastModifiedDate = *o.LastModifiedDate
	}
	if o.Name != nil {
		result.Name = *o.Name
	}
	if o.Selector != nil {
		result.Selector = *o.Selector
	}
	if o.SourceResult != nil {
		result.SourceResult = *o.SourceResult
	}
	result.Type = ParameterType(o.Type)
	if o.Value != nil {
		result.Value = *o.Value
	}
	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}

type ParameterHistory struct {
	AllowedPattern   string
	DataType         string
	Description      string
	KeyID            string
	Labels           []string
	LastModifiedDate time.Time
	LastModifiedUser string
	Name             string
	Policies         []ParameterInlinePolicy
	Tier             ParameterTier
	Type             ParameterType
	Value            string
	Version          int64
}

func newParameterHistoryList(list []*SDK.ParameterHistory) []ParameterHistory {
	if len(list) == 0 {
		return nil
	}

	results := make([]ParameterHistory, len(list))
	for i, v := range list {
		results[i] = newParameterHistory(v)
	}
	return results
}

func newParameterHistory(o *SDK.ParameterHistory) ParameterHistory {
	result := ParameterHistory{}
	if o == nil {
		return result
	}

	if o.AllowedPattern != nil {
		result.AllowedPattern = *o.AllowedPattern
	}
	if o.DataType != nil {
		result.DataType = *o.DataType
	}
	if o.Description != nil {
		result.Description = *o.Description
	}
	if o.KeyId != nil {
		result.KeyID = *o.KeyId
	}
	if o.LastModifiedDate != nil {
		result.LastModifiedDate = *o.LastModifiedDate
	}
	if o.LastModifiedUser != nil {
		result.LastModifiedUser = *o.LastModifiedUser
	}
	if o.Name != nil {
		result.Name = *o.Name
	}
	if o.Value != nil {
		result.Value = *o.Value
	}
	if o.Version != nil {
		result.Version = *o.Version
	}

	result.Labels = o.Labels
	result.Policies = newParameterInlinePolicyList(o.Policies)
	result.Tier = ParameterTier(o.Tier)
	result.Type = ParameterType(o.Type)
	return result
}

type ParameterInlinePolicy struct {
	PolicyStatus string
	PolicyText   string
	PolicyType   string
}

func newParameterInlinePolicyList(list []SDK.ParameterInlinePolicy) []ParameterInlinePolicy {
	if len(list) == 0 {
		return nil
	}

	results := make([]ParameterInlinePolicy, len(list))
	for i := range list {
		results[i] = newParameterInlinePolicy(&list[i])
	}
	return results
}

func newParameterInlinePolicy(o *SDK.ParameterInlinePolicy) ParameterInlinePolicy {
	result := ParameterInlinePolicy{}
	if o == nil {
		return result
	}

	if o.PolicyStatus != nil {
		result.PolicyStatus = *o.PolicyStatus
	}
	if o.PolicyText != nil {
		result.PolicyText = *o.PolicyText
	}
	if o.PolicyType != nil {
		result.PolicyType = *o.PolicyType
	}
	return result
}

type ParameterMetadata struct {
	AllowedPattern   string
	DataType         string
	Description      string
	KeyID            string
	LastModifiedDate time.Time
	LastModifiedUser string
	Name             string
	Policies         []ParameterInlinePolicy
	Tier             ParameterTier
	Type             ParameterType
	Version          int64
}

func newParameterMetadata(o SDK.ParameterMetadata) ParameterMetadata {
	result := ParameterMetadata{}

	if o.AllowedPattern != nil {
		result.AllowedPattern = *o.AllowedPattern
	}
	if o.DataType != nil {
		result.DataType = *o.DataType
	}
	if o.Description != nil {
		result.Description = *o.Description
	}
	if o.KeyId != nil {
		result.KeyID = *o.KeyId
	}
	if o.LastModifiedDate != nil {
		result.LastModifiedDate = *o.LastModifiedDate
	}
	if o.LastModifiedUser != nil {
		result.LastModifiedUser = *o.LastModifiedUser
	}
	if o.Name != nil {
		result.Name = *o.Name
	}
	if o.Version != nil {
		result.Version = *o.Version
	}

	result.Policies = newParameterInlinePolicyList(o.Policies)
	result.Tier = ParameterTier(o.Tier)
	result.Type = ParameterType(o.Type)
	return result
}

type ParameterStringFilter struct {
	Key    string
	Option string
	Values []string
}

func (v ParameterStringFilter) ToSDK() SDK.ParameterStringFilter {
	o := SDK.ParameterStringFilter{}
	if v.Key != "" {
		o.Key = pointers.String(v.Key)
	}
	if v.Option != "" {
		o.Option = pointers.String(v.Option)
	}
	o.Values = v.Values
	return o
}

type Tag struct {
	Key   string
	Value string
}

func (r Tag) ToSDK() SDK.Tag {
	o := SDK.Tag{}

	if r.Key != "" {
		o.Key = pointers.String(r.Key)
	}
	if r.Value != "" {
		o.Value = pointers.String(r.Value)
	}
	return o
}
