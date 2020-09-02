package dynamodb

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type BillingMode string

const (
	BillingModeProvisioned   BillingMode = BillingMode(SDK.BillingModeProvisioned)
	BillingModePayPerRequest BillingMode = BillingMode(SDK.BillingModePayPerRequest)
)

func (v BillingMode) IsProvisioned() bool {
	return v == BillingModeProvisioned
}
func (v BillingMode) IsPayPerRequest() bool {
	return v == BillingModePayPerRequest
}

type ComparisonOperator string

const (
	ComparisonOperatorEq          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorEq)
	ComparisonOperatorNe          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorNe)
	ComparisonOperatorIn          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorIn)
	ComparisonOperatorLe          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorLe)
	ComparisonOperatorLt          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorLt)
	ComparisonOperatorGe          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorGe)
	ComparisonOperatorGt          ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorGt)
	ComparisonOperatorBetween     ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorBetween)
	ComparisonOperatorNotNull     ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorNotNull)
	ComparisonOperatorNull        ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorNull)
	ComparisonOperatorContains    ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorContains)
	ComparisonOperatorNotContains ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorNotContains)
	ComparisonOperatorBeginsWith  ComparisonOperator = ComparisonOperator(SDK.ComparisonOperatorBeginsWith)

	// for expression condition
	ComparisonOperatorAttrExists    = ComparisonOperator("ATTR_EXISTS")
	ComparisonOperatorAttrNotExists = ComparisonOperator("ATTR_NOT_EXISTS")
	ComparisonOperatorAttrType      = ComparisonOperator("ATTR_TYPE")
)

func (v ComparisonOperator) IsEq() bool {
	return v == ComparisonOperatorEq
}
func (v ComparisonOperator) IsNe() bool {
	return v == ComparisonOperatorNe
}
func (v ComparisonOperator) IsIn() bool {
	return v == ComparisonOperatorIn
}
func (v ComparisonOperator) IsLe() bool {
	return v == ComparisonOperatorLe
}
func (v ComparisonOperator) IsLt() bool {
	return v == ComparisonOperatorLt
}
func (v ComparisonOperator) IsGe() bool {
	return v == ComparisonOperatorGe
}
func (v ComparisonOperator) IsGt() bool {
	return v == ComparisonOperatorGt
}
func (v ComparisonOperator) IsBetween() bool {
	return v == ComparisonOperatorBetween
}
func (v ComparisonOperator) IsNotNull() bool {
	return v == ComparisonOperatorNotNull
}
func (v ComparisonOperator) IsNull() bool {
	return v == ComparisonOperatorNull
}
func (v ComparisonOperator) IsContains() bool {
	return v == ComparisonOperatorContains
}
func (v ComparisonOperator) IsNotContains() bool {
	return v == ComparisonOperatorNotContains
}
func (v ComparisonOperator) IsBeginsWith() bool {
	return v == ComparisonOperatorBeginsWith
}

type ConditionalOperator string

const (
	ConditionalOperatorAnd ConditionalOperator = ConditionalOperator(SDK.ConditionalOperatorAnd)
	ConditionalOperatorOr  ConditionalOperator = ConditionalOperator(SDK.ConditionalOperatorOr)
)

func (v ConditionalOperator) IsAnd() bool {
	return v == ConditionalOperatorAnd
}
func (v ConditionalOperator) IsOr() bool {
	return v == ConditionalOperatorOr
}

type IndexStatus string

const (
	IndexStatusCreating IndexStatus = IndexStatus(SDK.IndexStatusCreating)
	IndexStatusUpdating IndexStatus = IndexStatus(SDK.IndexStatusUpdating)
	IndexStatusDeleting IndexStatus = IndexStatus(SDK.IndexStatusDeleting)
	IndexStatusActive   IndexStatus = IndexStatus(SDK.IndexStatusActive)
)

func (v IndexStatus) IsCreating() bool {
	return v == IndexStatusCreating
}
func (v IndexStatus) IsUpdating() bool {
	return v == IndexStatusUpdating
}
func (v IndexStatus) IsDeleting() bool {
	return v == IndexStatusDeleting
}
func (v IndexStatus) IsActive() bool {
	return v == IndexStatusActive
}

type KeyType string

const (
	KeyTypeHash  KeyType = KeyType(SDK.KeyTypeHash)
	KeyTypeRange KeyType = KeyType(SDK.KeyTypeRange)
)

func (v KeyType) IsHash() bool {
	return v == KeyTypeHash
}
func (v KeyType) IsRange() bool {
	return v == KeyTypeRange
}

type ProjectionType string

const (
	ProjectionTypeAll      ProjectionType = ProjectionType(SDK.ProjectionTypeAll)
	ProjectionTypeKeysOnly ProjectionType = ProjectionType(SDK.ProjectionTypeKeysOnly)
	ProjectionTypeInclude  ProjectionType = ProjectionType(SDK.ProjectionTypeInclude)
)

func (v ProjectionType) IsAll() bool {
	return v == ProjectionTypeAll
}
func (v ProjectionType) IsKeysOnly() bool {
	return v == ProjectionTypeKeysOnly
}
func (v ProjectionType) IsInclude() bool {
	return v == ProjectionTypeInclude
}

type ReplicaStatus string

const (
	ReplicaStatusCreating       ReplicaStatus = ReplicaStatus(SDK.ReplicaStatusCreating)
	ReplicaStatusCreationFailed ReplicaStatus = ReplicaStatus(SDK.ReplicaStatusCreationFailed)
	ReplicaStatusUpdating       ReplicaStatus = ReplicaStatus(SDK.ReplicaStatusUpdating)
	ReplicaStatusDeleting       ReplicaStatus = ReplicaStatus(SDK.ReplicaStatusDeleting)
	ReplicaStatusActive         ReplicaStatus = ReplicaStatus(SDK.ReplicaStatusActive)
)

type ReturnConsumedCapacity string

const (
	ReturnConsumedCapacityIndexes ReturnConsumedCapacity = ReturnConsumedCapacity(SDK.ReturnConsumedCapacityIndexes)
	ReturnConsumedCapacityTotal   ReturnConsumedCapacity = ReturnConsumedCapacity(SDK.ReturnConsumedCapacityTotal)
	ReturnConsumedCapacityNone    ReturnConsumedCapacity = ReturnConsumedCapacity(SDK.ReturnConsumedCapacityNone)
)

func (v ReturnConsumedCapacity) IsIndexes() bool {
	return v == ReturnConsumedCapacityIndexes
}
func (v ReturnConsumedCapacity) IsTotal() bool {
	return v == ReturnConsumedCapacityTotal
}
func (v ReturnConsumedCapacity) IsNone() bool {
	return v == ReturnConsumedCapacityNone
}

type ReturnItemCollectionMetrics string

const (
	ReturnItemCollectionMetricsSize ReturnItemCollectionMetrics = ReturnItemCollectionMetrics(SDK.ReturnItemCollectionMetricsSize)
	ReturnItemCollectionMetricsNone ReturnItemCollectionMetrics = ReturnItemCollectionMetrics(SDK.ReturnItemCollectionMetricsNone)
)

func (v ReturnItemCollectionMetrics) IsSize() bool {
	return v == ReturnItemCollectionMetricsSize
}
func (v ReturnItemCollectionMetrics) IsNone() bool {
	return v == ReturnItemCollectionMetricsNone
}

type ReturnValue string

const (
	ReturnValueNone       ReturnValue = ReturnValue(SDK.ReturnValueNone)
	ReturnValueAllOld     ReturnValue = ReturnValue(SDK.ReturnValueAllOld)
	ReturnValueUpdatedOld ReturnValue = ReturnValue(SDK.ReturnValueUpdatedOld)
	ReturnValueAllNew     ReturnValue = ReturnValue(SDK.ReturnValueAllNew)
	ReturnValueUpdatedNew ReturnValue = ReturnValue(SDK.ReturnValueUpdatedNew)
)

func (v ReturnValue) IsNone() bool {
	return v == ReturnValueNone
}

func (v ReturnValue) IsAllOld() bool {
	return v == ReturnValueAllOld
}
func (v ReturnValue) IsUpdatedOld() bool {
	return v == ReturnValueUpdatedOld
}
func (v ReturnValue) IsAllNew() bool {
	return v == ReturnValueAllNew
}
func (v ReturnValue) IsUpdatedNew() bool {
	return v == ReturnValueUpdatedNew
}

type ScalarAttributeType string

const (
	ScalarAttributeTypeS ScalarAttributeType = ScalarAttributeType(SDK.ScalarAttributeTypeS)
	ScalarAttributeTypeN ScalarAttributeType = ScalarAttributeType(SDK.ScalarAttributeTypeN)
	ScalarAttributeTypeB ScalarAttributeType = ScalarAttributeType(SDK.ScalarAttributeTypeB)
)

func (v ScalarAttributeType) IsS() bool {
	return v == ScalarAttributeTypeS
}
func (v ScalarAttributeType) IsN() bool {
	return v == ScalarAttributeTypeN
}
func (v ScalarAttributeType) IsB() bool {
	return v == ScalarAttributeTypeB
}

type Select string

const (
	SelectAllAttributes          Select = Select(SDK.SelectAllAttributes)
	SelectAllProjectedAttributes Select = Select(SDK.SelectAllProjectedAttributes)
	SelectSpecificAttributes     Select = Select(SDK.SelectSpecificAttributes)
	SelectCount                  Select = Select(SDK.SelectCount)
)

func (v Select) IsAllAttributes() bool {
	return v == SelectAllAttributes
}
func (v Select) IsAllProjectedAttributes() bool {
	return v == SelectAllProjectedAttributes
}
func (v Select) IsSpecificAttributes() bool {
	return v == SelectSpecificAttributes
}
func (v Select) IsCount() bool {
	return v == SelectCount
}

type SSEStatus string

const (
	SSEStatusEnabling  SSEStatus = SSEStatus(SDK.SSEStatusEnabling)
	SSEStatusEnabled   SSEStatus = SSEStatus(SDK.SSEStatusEnabled)
	SSEStatusDisabling SSEStatus = SSEStatus(SDK.SSEStatusDisabling)
	SSEStatusDisabled  SSEStatus = SSEStatus(SDK.SSEStatusDisabled)
	SSEStatusUpdating  SSEStatus = SSEStatus(SDK.SSEStatusUpdating)
)

func (v SSEStatus) IsEnabling() bool {
	return v == SSEStatusEnabling
}
func (v SSEStatus) IsEnabled() bool {
	return v == SSEStatusEnabled
}
func (v SSEStatus) IsDisabling() bool {
	return v == SSEStatusDisabling
}
func (v SSEStatus) IsDisabled() bool {
	return v == SSEStatusDisabled
}
func (v SSEStatus) IsUpdating() bool {
	return v == SSEStatusUpdating
}

type SSEType string

const (
	SSETypeAes256 SSEType = SSEType(SDK.SSETypeAes256)
	SSETypeKms    SSEType = SSEType(SDK.SSETypeKms)
)

func (v SSEType) IsAes256() bool {
	return v == SSETypeAes256
}
func (v SSEType) IsKms() bool {
	return v == SSETypeKms
}

type StreamViewType string

const (
	StreamViewTypeNewImage        StreamViewType = StreamViewType(SDK.StreamViewTypeNewImage)
	StreamViewTypeOldImage        StreamViewType = StreamViewType(SDK.StreamViewTypeOldImage)
	StreamViewTypeNewAndOldImages StreamViewType = StreamViewType(SDK.StreamViewTypeNewAndOldImages)
	StreamViewTypeKeysOnly        StreamViewType = StreamViewType(SDK.StreamViewTypeKeysOnly)
)

func (v StreamViewType) IsNewImage() bool {
	return v == StreamViewTypeNewImage
}
func (v StreamViewType) IsOldImage() bool {
	return v == StreamViewTypeOldImage
}
func (v StreamViewType) IsNewAndOldImages() bool {
	return v == StreamViewTypeNewAndOldImages
}
func (v StreamViewType) IsKeysOnly() bool {
	return v == StreamViewTypeKeysOnly
}

type TableStatus string

const (
	TableStatusCreating                          TableStatus = TableStatus(SDK.TableStatusCreating)
	TableStatusUpdating                          TableStatus = TableStatus(SDK.TableStatusUpdating)
	TableStatusDeleting                          TableStatus = TableStatus(SDK.TableStatusDeleting)
	TableStatusActive                            TableStatus = TableStatus(SDK.TableStatusActive)
	TableStatusInaccessibleEncryptionCredentials TableStatus = TableStatus(SDK.TableStatusInaccessibleEncryptionCredentials)
	TableStatusArchiving                         TableStatus = TableStatus(SDK.TableStatusArchiving)
	TableStatusArchived                          TableStatus = TableStatus(SDK.TableStatusArchived)
)

func (v TableStatus) IsCreating() bool {
	return v == TableStatusCreating
}
func (v TableStatus) IsUpdating() bool {
	return v == TableStatusUpdating
}
func (v TableStatus) IsDeleting() bool {
	return v == TableStatusDeleting
}
func (v TableStatus) IsActive() bool {
	return v == TableStatusActive
}
func (v TableStatus) IsInaccessibleEncryptionCredentials() bool {
	return v == TableStatusInaccessibleEncryptionCredentials
}
func (v TableStatus) IsArchiving() bool {
	return v == TableStatusArchiving
}
func (v TableStatus) IsArchived() bool {
	return v == TableStatusArchived
}
