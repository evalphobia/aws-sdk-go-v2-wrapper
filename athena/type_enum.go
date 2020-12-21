package athena

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/athena"
)

type ColumnNullable string

const (
	ColumnNullableNotNull  ColumnNullable = ColumnNullable(SDK.ColumnNullableNotNull)
	ColumnNullableNullable ColumnNullable = ColumnNullable(SDK.ColumnNullableNullable)
	ColumnNullableUnknown  ColumnNullable = ColumnNullable(SDK.ColumnNullableUnknown)
)

func (v ColumnNullable) IsNotNull() bool {
	return v == ColumnNullableNotNull
}
func (v ColumnNullable) IsNullable() bool {
	return v == ColumnNullableNullable
}
func (v ColumnNullable) IsUnknown() bool {
	return v == ColumnNullableUnknown
}

type EncryptionOption string

const (
	EncryptionOptionSSES3  EncryptionOption = EncryptionOption(SDK.EncryptionOptionSseS3)
	EncryptionOptionSSEKMS EncryptionOption = EncryptionOption(SDK.EncryptionOptionSseKms)
	EncryptionOptionCSEKMS EncryptionOption = EncryptionOption(SDK.EncryptionOptionCseKms)
)

func (v EncryptionOption) IsSSES3() bool {
	return v == EncryptionOptionSSES3
}
func (v EncryptionOption) IsSSEKMS() bool {
	return v == EncryptionOptionSSEKMS
}
func (v EncryptionOption) IsCSEKMS() bool {
	return v == EncryptionOptionCSEKMS
}

type StatementType string

const (
	StatementTypeDDL     StatementType = StatementType(SDK.StatementTypeDdl)
	StatementTypeDML     StatementType = StatementType(SDK.StatementTypeDml)
	StatementTypeUtility StatementType = StatementType(SDK.StatementTypeUtility)
)

func (v StatementType) IsDDL() bool {
	return v == StatementTypeDDL
}
func (v StatementType) IsDML() bool {
	return v == StatementTypeDML
}
func (v StatementType) IsUtility() bool {
	return v == StatementTypeUtility
}

type QueryExecutionState string

const (
	QueryExecutionStateQueued    QueryExecutionState = QueryExecutionState(SDK.QueryExecutionStateQueued)
	QueryExecutionStateRunning   QueryExecutionState = QueryExecutionState(SDK.QueryExecutionStateRunning)
	QueryExecutionStateSucceeded QueryExecutionState = QueryExecutionState(SDK.QueryExecutionStateSucceeded)
	QueryExecutionStateFailed    QueryExecutionState = QueryExecutionState(SDK.QueryExecutionStateFailed)
	QueryExecutionStateCancelled QueryExecutionState = QueryExecutionState(SDK.QueryExecutionStateCancelled)
)

func (v QueryExecutionState) IsQueued() bool {
	return v == QueryExecutionStateQueued
}
func (v QueryExecutionState) IsRunning() bool {
	return v == QueryExecutionStateRunning
}
func (v QueryExecutionState) IsSucceeded() bool {
	return v == QueryExecutionStateSucceeded
}
func (v QueryExecutionState) IsFailed() bool {
	return v == QueryExecutionStateFailed
}
func (v QueryExecutionState) IsCancelled() bool {
	return v == QueryExecutionStateCancelled
}
