package dynamodb

import (
	"fmt"
	"strconv"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type ArchivalSummary struct {
	ArchivalBackupARN string
	ArchivalDateTime  time.Time
	ArchivalReason    string
}

func newArchivalSummary(o *SDK.ArchivalSummary) ArchivalSummary {
	result := ArchivalSummary{}
	if o == nil {
		return result
	}

	if o.ArchivalBackupArn != nil {
		result.ArchivalBackupARN = *o.ArchivalBackupArn
	}
	if o.ArchivalDateTime != nil {
		result.ArchivalDateTime = *o.ArchivalDateTime
	}
	if o.ArchivalReason != nil {
		result.ArchivalReason = *o.ArchivalReason
	}
	return result
}

type AttributeDefinition struct {
	AttributeName string
	AttributeType ScalarAttributeType
}

func newAttributeDefinitions(list []SDK.AttributeDefinition) []AttributeDefinition {
	if len(list) == 0 {
		return nil
	}

	results := make([]AttributeDefinition, len(list))
	for i, v := range list {
		results[i] = newAttributeDefinition(v)
	}
	return results
}

func newAttributeDefinition(o SDK.AttributeDefinition) AttributeDefinition {
	result := AttributeDefinition{}

	if o.AttributeName != nil {
		result.AttributeName = *o.AttributeName
	}
	result.AttributeType = ScalarAttributeType(o.AttributeType)
	return result
}

func (r AttributeDefinition) ToSDK() SDK.AttributeDefinition {
	o := SDK.AttributeDefinition{}

	if r.AttributeName != "" {
		o.AttributeName = pointers.String(r.AttributeName)
	}

	o.AttributeType = SDK.ScalarAttributeType(r.AttributeType)
	return o
}

type AttributeValue struct {
	Binary         []byte
	BinarySet      [][]byte
	List           []AttributeValue
	Map            map[string]AttributeValue
	Number         string
	NumberInt      int64
	NumberFloat    float64
	NumberSet      []string
	NumberSetInt   []int64
	NumberSetFloat []float64
	Null           bool
	String         string
	StringSet      []string

	Bool      bool
	HasBool   bool
	HasNumber bool
}

// MustNewAttributeValue creates AttributeValue from given value.
func MustNewAttributeValue(v interface{}) AttributeValue {
	result := AttributeValue{}

	switch vv := v.(type) {
	case []byte:
		result.Binary = vv
	case [][]byte:
		result.BinarySet = vv
	case string:
		result.String = vv
	case []string:
		result.StringSet = vv
	case bool:
		result.Bool = vv
		result.HasBool = true
	case int:
		result.NumberInt = int64(vv)
		result.HasNumber = true
	case int32:
		result.NumberInt = int64(vv)
		result.HasNumber = true
	case int64:
		result.NumberInt = vv
		result.HasNumber = true
	case float32:
		result.NumberFloat = float64(vv)
		result.HasNumber = true
	case float64:
		result.NumberFloat = vv
		result.HasNumber = true
	case []int64:
		result.NumberSetInt = vv
	case []float64:
		result.NumberSetFloat = vv
	case map[string]interface{}:
		m := make(map[string]AttributeValue, len(vv))
		for k, v := range vv {
			m[k] = MustNewAttributeValue(v)
		}
		result.Map = m
	default:
		if v == nil {
			result.Null = true
			return result
		}
		panic(fmt.Sprintf("[MustNewAttributeValue] cannot parse the given value type: [%t]", v))
	}
	return result
}

func newAttributeValue(o SDK.AttributeValue) AttributeValue {
	result := AttributeValue{}

	switch {
	case len(o.B) != 0:
		result.Binary = o.B
	case len(o.BS) != 0:
		result.BinarySet = o.BS
	case len(o.L) != 0:
		result.List = newAttributeValueList(o.L)
	case len(o.M) != 0:
		result.Map = newAttributeValueMap(o.M)
	case o.N != nil:
		result.Number = *o.N
		result.HasNumber = true
	case len(o.NS) != 0:
		result.NumberSet = o.NS
	case o.NULL != nil:
		result.Null = *o.NULL
	case o.S != nil:
		result.String = *o.S
	case len(o.SS) != 0:
		result.StringSet = o.SS
	case o.BOOL != nil:
		result.Bool = *o.BOOL
		result.HasBool = true
	}
	return result
}

func newAttributeValueList(list []SDK.AttributeValue) []AttributeValue {
	if len(list) == 0 {
		return nil
	}

	results := make([]AttributeValue, len(list))
	for i, v := range list {
		results[i] = newAttributeValue(v)
	}
	return results
}

func newAttributeValueMap(o map[string]SDK.AttributeValue) map[string]AttributeValue {
	if len(o) == 0 {
		return nil
	}

	m := make(map[string]AttributeValue, len(o))
	for key, val := range o {
		m[key] = newAttributeValue(val)
	}
	return m
}

func (r AttributeValue) ToSDK() SDK.AttributeValue {
	o := SDK.AttributeValue{}

	switch {
	case len(r.Binary) != 0:
		o.B = r.Binary
	case len(r.BinarySet) != 0:
		o.BS = r.BinarySet
	case len(r.List) != 0:
		list := make([]SDK.AttributeValue, len(r.List))
		for i, v := range r.List {
			list[i] = v.ToSDK()
		}
		o.L = list
	case len(r.Map) != 0:
		m := make(map[string]SDK.AttributeValue, len(r.Map))
		for key, val := range r.Map {
			m[key] = val.ToSDK()
		}
		o.M = m
	case r.Number != "":
		o.N = pointers.String(r.Number)
	case r.NumberInt != 0:
		o.N = pointers.String(strconv.FormatInt(r.NumberInt, 10))
	case r.NumberFloat != 0:
		o.N = pointers.String(strconv.FormatFloat(r.NumberFloat, 'f', -1, 64))
	case r.HasNumber:
		o.N = pointers.String("0")
	case len(r.NumberSet) != 0:
		o.NS = r.NumberSet
	case len(r.NumberSetInt) != 0:
		list := make([]string, len(r.NumberSetInt))
		for i, v := range r.NumberSetInt {
			list[i] = strconv.FormatInt(v, 10)
		}
		o.NS = list
	case len(r.NumberSetFloat) != 0:
		list := make([]string, len(r.NumberSetFloat))
		for i, v := range r.NumberSetFloat {
			list[i] = strconv.FormatFloat(v, 'f', -1, 64)
		}
		o.NS = list
	case r.String != "":
		o.S = pointers.String(r.String)
	case len(r.StringSet) != 0:
		o.SS = r.StringSet
	case r.HasBool,
		r.Bool:
		o.BOOL = pointers.Bool(r.Bool)
	case r.Null:
		o.NULL = pointers.Bool(r.Null)
	}
	return o
}

func (r AttributeValue) GetValue() interface{} {
	switch {
	case len(r.Binary) != 0:
		return r.Binary
	case len(r.BinarySet) != 0:
		return r.BinarySet
	case len(r.List) != 0:
		list := make([]interface{}, len(r.List))
		for i, v := range r.List {
			list[i] = v.GetValue()
		}
		return list
	case len(r.Map) != 0:
		m := make(map[string]interface{}, len(r.Map))
		for key, val := range r.Map {
			m[key] = val.GetValue()
		}
		return m
	case r.Number != "":
		v, _ := strconv.Atoi(r.Number)
		return v
	case r.NumberInt != 0:
		return r.NumberInt
	case r.NumberFloat != 0:
		return r.NumberFloat
	case r.HasNumber:
		return 0
	case len(r.NumberSet) != 0:
		list := make([]int, len(r.NumberSet))
		for i, v := range r.NumberSet {
			list[i], _ = strconv.Atoi(v)
		}
		return list
	case len(r.NumberSetInt) != 0:
		return r.NumberSetInt
	case len(r.NumberSetFloat) != 0:
		return r.NumberSetFloat
	case r.String != "":
		return r.String
	case len(r.StringSet) != 0:
		return r.StringSet
	case r.HasBool,
		r.Bool:
		return r.Bool
	case r.Null:
		return newArchivalSummary
	}
	return nil
}

func (r AttributeValue) hasValue() bool {
	switch {
	case r.HasNumber,
		r.HasBool,
		r.Bool,
		r.Null,
		r.Number != "",
		r.NumberInt != 0,
		r.NumberFloat != 0,
		r.String != "",
		len(r.Binary) != 0,
		len(r.BinarySet) != 0,
		len(r.List) != 0,
		len(r.Map) != 0,
		len(r.NumberSet) != 0,
		len(r.NumberSetInt) != 0,
		len(r.NumberSetFloat) != 0,
		len(r.StringSet) != 0:
		return true
	}
	return false
}

type BillingModeSummary struct {
	BillingMode                       BillingMode
	LastUpdateToPayPerRequestDateTime time.Time
}

func newBillingModeSummary(o *SDK.BillingModeSummary) BillingModeSummary {
	result := BillingModeSummary{}
	if o == nil {
		return result
	}

	if o.LastUpdateToPayPerRequestDateTime != nil {
		result.LastUpdateToPayPerRequestDateTime = *o.LastUpdateToPayPerRequestDateTime
	}

	result.BillingMode = BillingMode(o.BillingMode)
	return result
}

type Capacity struct {
	CapacityUnits      float64
	ReadCapacityUnits  float64
	WriteCapacityUnits float64
}

func newCapacity(o *SDK.Capacity) Capacity {
	result := Capacity{}
	if o == nil {
		return result
	}

	if o.CapacityUnits != nil {
		result.CapacityUnits = *o.CapacityUnits
	}
	if o.ReadCapacityUnits != nil {
		result.ReadCapacityUnits = *o.ReadCapacityUnits
	}
	if o.WriteCapacityUnits != nil {
		result.WriteCapacityUnits = *o.WriteCapacityUnits
	}
	return result
}

func newCapacityMap(m map[string]SDK.Capacity) map[string]Capacity {
	if m == nil {
		return nil
	}

	result := make(map[string]Capacity, len(m))
	for key, val := range m {
		val := val
		result[key] = newCapacity(&val)
	}
	return result
}

type Condition struct {
	ComparisonOperator ComparisonOperator

	// optional
	AttributeValueList []AttributeValue
}

func (r Condition) ToSDK() SDK.Condition {
	o := SDK.Condition{}

	o.ComparisonOperator = SDK.ComparisonOperator(r.ComparisonOperator)

	if len(r.AttributeValueList) != 0 {
		list := make([]SDK.AttributeValue, len(r.AttributeValueList))
		for i, v := range r.AttributeValueList {
			list[i] = v.ToSDK()
		}
		o.AttributeValueList = list
	}

	return o
}

type ConsumedCapacity struct {
	CapacityUnits          float64
	GlobalSecondaryIndexes map[string]Capacity
	LocalSecondaryIndexes  map[string]Capacity
	ReadCapacityUnits      float64
	Table                  Capacity
	TableName              string
	WriteCapacityUnits     float64
}

func newConsumedCapacities(list []SDK.ConsumedCapacity) []ConsumedCapacity {
	if len(list) == 0 {
		return nil
	}

	result := make([]ConsumedCapacity, len(list))
	for i, v := range list {
		result[i] = newConsumedCapacity(v)
	}
	return result
}

func newConsumedCapacity(o SDK.ConsumedCapacity) ConsumedCapacity {
	result := ConsumedCapacity{}

	if o.CapacityUnits != nil {
		result.CapacityUnits = *o.CapacityUnits
	}
	if o.ReadCapacityUnits != nil {
		result.ReadCapacityUnits = *o.ReadCapacityUnits
	}
	if o.TableName != nil {
		result.TableName = *o.TableName
	}
	if o.WriteCapacityUnits != nil {
		result.WriteCapacityUnits = *o.WriteCapacityUnits
	}

	result.GlobalSecondaryIndexes = newCapacityMap(o.GlobalSecondaryIndexes)
	result.LocalSecondaryIndexes = newCapacityMap(o.LocalSecondaryIndexes)
	result.Table = newCapacity(o.Table)
	return result
}

type ExpectedAttributeValue struct {
	AttributeValueList []AttributeValue   `type:"list"`
	ComparisonOperator ComparisonOperator `type:"string" enum:"true"`
	Exists             bool               `type:"boolean"`
	Value              AttributeValue     `type:"structure"`
}

func (r ExpectedAttributeValue) ToSDK() SDK.ExpectedAttributeValue {
	o := SDK.ExpectedAttributeValue{}

	if len(r.AttributeValueList) != 0 {
		list := make([]SDK.AttributeValue, 0, len(r.AttributeValueList))
		for _, v := range r.AttributeValueList {
			if v.hasValue() {
				list = append(list, v.ToSDK())
			}
		}
		o.AttributeValueList = list
	}

	if r.Exists {
		o.Exists = pointers.Bool(r.Exists)
	}

	o.ComparisonOperator = SDK.ComparisonOperator(r.ComparisonOperator)

	if r.Value.hasValue() {
		v := r.Value.ToSDK()
		o.Value = &v
	}

	return o
}

type GlobalSecondaryIndex struct {
	IndexName  string
	KeySchema  []KeySchemaElement
	Projection Projection

	// optional
	ProvisionedThroughput ProvisionedThroughput
}

func (r GlobalSecondaryIndex) ToSDK() SDK.GlobalSecondaryIndex {
	o := SDK.GlobalSecondaryIndex{}

	if r.IndexName != "" {
		o.IndexName = pointers.String(r.IndexName)
	}

	if r.Projection.hasValue() {
		v := r.Projection.ToSDK()
		o.Projection = &v
	}
	if r.ProvisionedThroughput.hasValue() {
		v := r.ProvisionedThroughput.ToSDK()
		o.ProvisionedThroughput = &v
	}

	if len(r.KeySchema) != 0 {
		list := make([]SDK.KeySchemaElement, len(r.KeySchema))
		for i, v := range r.KeySchema {
			list[i] = v.ToSDK()
		}
		o.KeySchema = list
	}
	return o
}

// Represents the properties of a global secondary index.
type GlobalSecondaryIndexDescription struct {
	Backfilling           bool
	IndexARN              string
	IndexName             string
	IndexSizeBytes        int64
	IndexStatus           IndexStatus
	ItemCount             int64
	KeySchema             []KeySchemaElement
	Projection            Projection
	ProvisionedThroughput ProvisionedThroughputDescription
}

func newGlobalSecondaryIndexDescriptions(list []SDK.GlobalSecondaryIndexDescription) []GlobalSecondaryIndexDescription {
	if len(list) == 0 {
		return nil
	}

	result := make([]GlobalSecondaryIndexDescription, len(list))
	for i, v := range list {
		result[i] = newGlobalSecondaryIndexDescription(v)
	}
	return result
}

func newGlobalSecondaryIndexDescription(o SDK.GlobalSecondaryIndexDescription) GlobalSecondaryIndexDescription {
	result := GlobalSecondaryIndexDescription{}

	if o.Backfilling != nil {
		result.Backfilling = *o.Backfilling
	}
	if o.IndexArn != nil {
		result.IndexARN = *o.IndexArn
	}
	if o.IndexName != nil {
		result.IndexName = *o.IndexName
	}
	if o.IndexSizeBytes != nil {
		result.IndexSizeBytes = *o.IndexSizeBytes
	}
	if o.ItemCount != nil {
		result.ItemCount = *o.ItemCount
	}

	result.IndexStatus = IndexStatus(o.IndexStatus)

	result.KeySchema = newKeySchemaElements(o.KeySchema)
	result.Projection = newProjection(o.Projection)
	result.ProvisionedThroughput = newProvisionedThroughputDescription(o.ProvisionedThroughput)
	return result
}

type ItemCollectionMetrics struct {
	ItemCollectionKey   map[string]AttributeValue `type:"map"`
	SizeEstimateRangeGB []float64                 `type:"list"`
}

func newItemCollectionMetrics(o SDK.ItemCollectionMetrics) ItemCollectionMetrics {
	result := ItemCollectionMetrics{}

	m := make(map[string]AttributeValue, len(o.ItemCollectionKey))
	for key, val := range o.ItemCollectionKey {
		m[key] = newAttributeValue(val)
	}
	result.ItemCollectionKey = m

	result.SizeEstimateRangeGB = o.SizeEstimateRangeGB
	return result
}

type KeysAndAttributes struct {
	Keys []map[string]AttributeValue

	// optional
	AttributesToGet          []string
	ConsistentRead           bool
	ExpressionAttributeNames map[string]string
	ProjectionExpression     string
}

func newKeysAndAttributes(o SDK.KeysAndAttributes) KeysAndAttributes {
	result := KeysAndAttributes{}

	result.AttributesToGet = o.AttributesToGet
	result.ExpressionAttributeNames = o.ExpressionAttributeNames

	if o.ConsistentRead != nil {
		result.ConsistentRead = *o.ConsistentRead
	}
	if o.ProjectionExpression != nil {
		result.ProjectionExpression = *o.ProjectionExpression
	}

	if len(o.Keys) != 0 {
		list := make([]map[string]AttributeValue, len(o.Keys))
		for i, v := range o.Keys {
			list[i] = newAttributeValueMap(v)
		}
		result.Keys = list
	}

	return result
}

func (r KeysAndAttributes) ToSDK() SDK.KeysAndAttributes {
	o := SDK.KeysAndAttributes{}

	if len(r.Keys) != 0 {
		list := make([]map[string]SDK.AttributeValue, len(r.Keys))
		for i, v := range r.Keys {
			m := make(map[string]SDK.AttributeValue, len(v))
			for key, val := range v {
				m[key] = val.ToSDK()
			}
			list[i] = m
		}
		o.Keys = list
	}

	if r.ConsistentRead {
		o.ConsistentRead = pointers.Bool(r.ConsistentRead)
	}
	if r.ProjectionExpression != "" {
		o.ProjectionExpression = pointers.String(r.ProjectionExpression)
	}

	o.AttributesToGet = r.AttributesToGet
	o.ExpressionAttributeNames = r.ExpressionAttributeNames
	return o
}

type KeySchemaElement struct {
	AttributeName string
	KeyType       KeyType
}

func newKeySchemaElements(list []SDK.KeySchemaElement) []KeySchemaElement {
	if len(list) == 0 {
		return nil
	}

	results := make([]KeySchemaElement, len(list))
	for i, v := range list {
		results[i] = newKeySchemaElement(v)
	}
	return results
}

func newKeySchemaElement(o SDK.KeySchemaElement) KeySchemaElement {
	result := KeySchemaElement{}

	if o.AttributeName != nil {
		result.AttributeName = *o.AttributeName
	}

	result.KeyType = KeyType(o.KeyType)
	return result
}

func (r KeySchemaElement) ToSDK() SDK.KeySchemaElement {
	o := SDK.KeySchemaElement{}

	if r.AttributeName != "" {
		o.AttributeName = pointers.String(r.AttributeName)
	}
	o.KeyType = SDK.KeyType(r.KeyType)
	return o
}

type LocalSecondaryIndex struct {
	IndexName  string
	KeySchema  []KeySchemaElement
	Projection Projection
}

func (r LocalSecondaryIndex) ToSDK() SDK.LocalSecondaryIndex {
	o := SDK.LocalSecondaryIndex{}

	if r.IndexName != "" {
		o.IndexName = pointers.String(r.IndexName)
	}

	if r.Projection.hasValue() {
		v := r.Projection.ToSDK()
		o.Projection = &v
	}

	if len(r.KeySchema) != 0 {
		list := make([]SDK.KeySchemaElement, len(r.KeySchema))
		for i, v := range r.KeySchema {
			list[i] = v.ToSDK()
		}
		o.KeySchema = list
	}
	return o
}

type LocalSecondaryIndexDescription struct {
	IndexARN       string
	IndexName      string
	IndexSizeBytes int64
	ItemCount      int64
	KeySchema      []KeySchemaElement
	Projection     Projection
}

func newLocalSecondaryIndexDescriptions(list []SDK.LocalSecondaryIndexDescription) []LocalSecondaryIndexDescription {
	if len(list) == 0 {
		return nil
	}

	result := make([]LocalSecondaryIndexDescription, len(list))
	for i, v := range list {
		result[i] = newLocalSecondaryIndexDescription(v)
	}
	return result
}

func newLocalSecondaryIndexDescription(o SDK.LocalSecondaryIndexDescription) LocalSecondaryIndexDescription {
	result := LocalSecondaryIndexDescription{}

	if o.IndexArn != nil {
		result.IndexARN = *o.IndexArn
	}
	if o.IndexName != nil {
		result.IndexName = *o.IndexName
	}
	if o.IndexSizeBytes != nil {
		result.IndexSizeBytes = *o.IndexSizeBytes
	}
	if o.ItemCount != nil {
		result.ItemCount = *o.ItemCount
	}

	result.KeySchema = newKeySchemaElements(o.KeySchema)
	result.Projection = newProjection(o.Projection)
	return result
}

type Projection struct {
	NonKeyAttributes []string
	ProjectionType   ProjectionType
}

func newProjection(o *SDK.Projection) Projection {
	result := Projection{}
	if o == nil {
		return result
	}

	result.NonKeyAttributes = o.NonKeyAttributes
	result.ProjectionType = ProjectionType(o.ProjectionType)
	return result
}

func (r Projection) ToSDK() SDK.Projection {
	o := SDK.Projection{}
	o.NonKeyAttributes = r.NonKeyAttributes
	o.ProjectionType = SDK.ProjectionType(r.ProjectionType)
	return o
}

func (r Projection) hasValue() bool {
	switch {
	case r.ProjectionType != "",
		len(r.NonKeyAttributes) != 0:
		return true
	}
	return false
}

type ProvisionedThroughput struct {
	ReadCapacityUnits  int64
	WriteCapacityUnits int64
}

func (r ProvisionedThroughput) ToSDK() SDK.ProvisionedThroughput {
	return SDK.ProvisionedThroughput{
		ReadCapacityUnits:  pointers.Long64(r.ReadCapacityUnits),
		WriteCapacityUnits: pointers.Long64(r.WriteCapacityUnits),
	}
}

func (r ProvisionedThroughput) hasValue() bool {
	switch {
	case r.ReadCapacityUnits != 0,
		r.WriteCapacityUnits != 0:
		return true
	}
	return false
}

type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime   time.Time
	LastIncreaseDateTime   time.Time
	NumberOfDecreasesToday int64
	ReadCapacityUnits      int64
	WriteCapacityUnits     int64
}

func newProvisionedThroughputDescription(o *SDK.ProvisionedThroughputDescription) ProvisionedThroughputDescription {
	result := ProvisionedThroughputDescription{}
	if o == nil {
		return result
	}

	if o.LastDecreaseDateTime != nil {
		result.LastDecreaseDateTime = *o.LastDecreaseDateTime
	}
	if o.LastIncreaseDateTime != nil {
		result.LastIncreaseDateTime = *o.LastIncreaseDateTime
	}
	if o.NumberOfDecreasesToday != nil {
		result.NumberOfDecreasesToday = *o.NumberOfDecreasesToday
	}
	if o.ReadCapacityUnits != nil {
		result.ReadCapacityUnits = *o.ReadCapacityUnits
	}
	if o.WriteCapacityUnits != nil {
		result.WriteCapacityUnits = *o.WriteCapacityUnits
	}
	return result
}

type ReplicaDescription struct {
	GlobalSecondaryIndexes       []ReplicaGlobalSecondaryIndexDescription
	KMSMasterKeyID               string
	RegionName                   string
	ReplicaStatus                ReplicaStatus
	ReplicaStatusDescription     string
	ReplicaStatusPercentProgress string

	ProvisionedThroughputOverrideRCU int64
}

func newReplicaDescriptions(list []SDK.ReplicaDescription) []ReplicaDescription {
	if len(list) == 0 {
		return nil
	}

	result := make([]ReplicaDescription, len(list))
	for i, v := range list {
		result[i] = newReplicaDescription(v)
	}
	return result
}

func newReplicaDescription(o SDK.ReplicaDescription) ReplicaDescription {
	result := ReplicaDescription{}

	if o.KMSMasterKeyId != nil {
		result.KMSMasterKeyID = *o.KMSMasterKeyId
	}
	if o.RegionName != nil {
		result.RegionName = *o.RegionName
	}
	if o.ReplicaStatusDescription != nil {
		result.ReplicaStatusDescription = *o.ReplicaStatusDescription
	}
	if o.ReplicaStatusPercentProgress != nil {
		result.ReplicaStatusPercentProgress = *o.ReplicaStatusPercentProgress
	}

	result.ReplicaStatus = ReplicaStatus(o.ReplicaStatus)

	result.GlobalSecondaryIndexes = newReplicaGlobalSecondaryIndexDescriptions(o.GlobalSecondaryIndexes)
	if v := o.ProvisionedThroughputOverride; v != nil {
		if v.ReadCapacityUnits != nil {
			result.ProvisionedThroughputOverrideRCU = *v.ReadCapacityUnits
		}
	}
	return result
}

type ReplicaGlobalSecondaryIndexDescription struct {
	IndexName                        string
	ProvisionedThroughputOverrideRCU int64
}

func newReplicaGlobalSecondaryIndexDescriptions(list []SDK.ReplicaGlobalSecondaryIndexDescription) []ReplicaGlobalSecondaryIndexDescription {
	if len(list) == 0 {
		return nil
	}

	result := make([]ReplicaGlobalSecondaryIndexDescription, len(list))
	for i, v := range list {
		result[i] = newReplicaGlobalSecondaryIndexDescription(v)
	}
	return result
}

func newReplicaGlobalSecondaryIndexDescription(o SDK.ReplicaGlobalSecondaryIndexDescription) ReplicaGlobalSecondaryIndexDescription {
	result := ReplicaGlobalSecondaryIndexDescription{}

	if o.IndexName != nil {
		result.IndexName = *o.IndexName
	}
	if v := o.ProvisionedThroughputOverride; v != nil {
		if v.ReadCapacityUnits != nil {
			result.ProvisionedThroughputOverrideRCU = *v.ReadCapacityUnits
		}
	}
	return result
}

type RestoreSummary struct {
	RestoreDateTime   time.Time
	RestoreInProgress bool

	// optional
	SourceBackupARN string
	SourceTableARN  string
}

func newRestoreSummary(o *SDK.RestoreSummary) RestoreSummary {
	result := RestoreSummary{}
	if o == nil {
		return result
	}

	if o.RestoreDateTime != nil {
		result.RestoreDateTime = *o.RestoreDateTime
	}
	if o.RestoreInProgress != nil {
		result.RestoreInProgress = *o.RestoreInProgress
	}
	if o.SourceBackupArn != nil {
		result.SourceBackupARN = *o.SourceBackupArn
	}
	if o.SourceTableArn != nil {
		result.SourceTableARN = *o.SourceTableArn
	}
	return result
}

type SSEDescription struct {
	InaccessibleEncryptionDateTime time.Time
	KMSMasterKeyArn                string
	SSEType                        SSEType
	Status                         SSEStatus
}

func newSSEDescription(o *SDK.SSEDescription) SSEDescription {
	result := SSEDescription{}
	if o == nil {
		return result
	}

	if o.InaccessibleEncryptionDateTime != nil {
		result.InaccessibleEncryptionDateTime = *o.InaccessibleEncryptionDateTime
	}
	if o.KMSMasterKeyArn != nil {
		result.KMSMasterKeyArn = *o.KMSMasterKeyArn
	}
	result.SSEType = SSEType(o.SSEType)
	result.Status = SSEStatus(o.Status)
	return result
}

type SSESpecification struct {
	Enabled        bool
	KMSMasterKeyID string
	SSEType        SSEType
}

func (r SSESpecification) ToSDK() SDK.SSESpecification {
	o := SDK.SSESpecification{}

	if r.Enabled {
		o.Enabled = pointers.Bool(r.Enabled)
	}
	if r.KMSMasterKeyID != "" {
		o.KMSMasterKeyId = pointers.String(r.KMSMasterKeyID)
	}
	o.SSEType = SDK.SSEType(r.SSEType)
	return o
}

func (r SSESpecification) hasValue() bool {
	switch {
	case r.Enabled,
		r.KMSMasterKeyID != "",
		r.SSEType != "":
		return true
	}
	return false
}

type StreamSpecification struct {
	StreamEnabled bool

	// optional
	StreamViewType StreamViewType
}

func newStreamSpecification(o *SDK.StreamSpecification) StreamSpecification {
	result := StreamSpecification{}
	if o == nil {
		return result
	}

	if o.StreamEnabled != nil {
		result.StreamEnabled = *o.StreamEnabled
	}
	result.StreamViewType = StreamViewType(o.StreamViewType)
	return result
}

func (r StreamSpecification) ToSDK() SDK.StreamSpecification {
	o := SDK.StreamSpecification{}

	if r.StreamEnabled {
		o.StreamEnabled = pointers.Bool(r.StreamEnabled)
	}

	o.StreamViewType = SDK.StreamViewType(r.StreamViewType)
	return o
}

func (r StreamSpecification) hasValue() bool {
	switch {
	case r.StreamEnabled,
		r.StreamViewType != "":
		return true
	}
	return false
}

type TableDescription struct {
	ArchivalSummary        ArchivalSummary
	AttributeDefinitions   []AttributeDefinition
	BillingModeSummary     BillingModeSummary
	CreationDateTime       time.Time
	GlobalSecondaryIndexes []GlobalSecondaryIndexDescription
	GlobalTableVersion     string
	ItemCount              int64
	KeySchema              []KeySchemaElement
	LatestStreamARN        string
	LatestStreamLabel      string
	LocalSecondaryIndexes  []LocalSecondaryIndexDescription
	ProvisionedThroughput  ProvisionedThroughputDescription
	Replicas               []ReplicaDescription
	RestoreSummary         RestoreSummary
	SSEDescription         SSEDescription
	StreamSpecification    StreamSpecification
	TableARN               string
	TableID                string
	TableName              string
	TableSizeBytes         int64
	TableStatus            TableStatus
}

func newTableDescription(o SDK.TableDescription) TableDescription {
	result := TableDescription{}

	if o.CreationDateTime != nil {
		result.CreationDateTime = *o.CreationDateTime
	}
	if o.GlobalTableVersion != nil {
		result.GlobalTableVersion = *o.GlobalTableVersion
	}
	if o.ItemCount != nil {
		result.ItemCount = *o.ItemCount
	}
	if o.LatestStreamArn != nil {
		result.LatestStreamARN = *o.LatestStreamArn
	}
	if o.LatestStreamLabel != nil {
		result.LatestStreamLabel = *o.LatestStreamLabel
	}
	if o.TableArn != nil {
		result.TableARN = *o.TableArn
	}
	if o.TableId != nil {
		result.TableID = *o.TableId
	}
	if o.TableName != nil {
		result.TableName = *o.TableName
	}
	if o.TableSizeBytes != nil {
		result.TableSizeBytes = *o.TableSizeBytes
	}

	result.TableStatus = TableStatus(o.TableStatus)

	result.ArchivalSummary = newArchivalSummary(o.ArchivalSummary)
	result.BillingModeSummary = newBillingModeSummary(o.BillingModeSummary)
	result.ProvisionedThroughput = newProvisionedThroughputDescription(o.ProvisionedThroughput)
	result.RestoreSummary = newRestoreSummary(o.RestoreSummary)
	result.SSEDescription = newSSEDescription(o.SSEDescription)
	result.StreamSpecification = newStreamSpecification(o.StreamSpecification)

	result.AttributeDefinitions = newAttributeDefinitions(o.AttributeDefinitions)
	result.GlobalSecondaryIndexes = newGlobalSecondaryIndexDescriptions(o.GlobalSecondaryIndexes)
	result.LocalSecondaryIndexes = newLocalSecondaryIndexDescriptions(o.LocalSecondaryIndexes)
	result.KeySchema = newKeySchemaElements(o.KeySchema)
	result.Replicas = newReplicaDescriptions(o.Replicas)
	return result
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

type WriteRequest struct {
	DeleteKeys map[string]AttributeValue
	PutItems   map[string]AttributeValue
}

func newWriteRequest(o SDK.WriteRequest) WriteRequest {
	result := WriteRequest{
		DeleteKeys: make(map[string]AttributeValue),
		PutItems:   make(map[string]AttributeValue),
	}

	if o.DeleteRequest != nil {
		m := make(map[string]AttributeValue, len(o.DeleteRequest.Key))
		for key, val := range o.DeleteRequest.Key {
			m[key] = newAttributeValue(val)
		}
		result.DeleteKeys = m
	}

	if o.PutRequest != nil {
		m := make(map[string]AttributeValue, len(o.PutRequest.Item))
		for key, val := range o.PutRequest.Item {
			m[key] = newAttributeValue(val)
		}
		result.DeleteKeys = m
	}

	return result
}

func (r WriteRequest) ToSDK() SDK.WriteRequest {
	o := SDK.WriteRequest{}

	if len(r.DeleteKeys) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.DeleteKeys))
		for key, val := range r.DeleteKeys {
			m[key] = val.ToSDK()
		}
		o.DeleteRequest = &SDK.DeleteRequest{
			Key: m,
		}
	}

	if len(r.PutItems) != 0 {
		m := make(map[string]SDK.AttributeValue, len(r.PutItems))
		for key, val := range r.PutItems {
			m[key] = val.ToSDK()
		}
		o.PutRequest = &SDK.PutRequest{
			Item: m,
		}
	}

	return o
}
