package athena

import (
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type ColumnInfo struct {
	Name string
	Type string

	// optional
	CaseSensitive bool
	CatalogName   string
	Label         string
	Nullable      ColumnNullable
	Precision     int64
	Scale         int64
	SchemaName    string
	TableName     string
}

func NewColumnInfoListFromMetadata(o *SDK.ResultSetMetadata) []ColumnInfo {
	if o == nil {
		return nil
	}
	return NewColumnInfoList(o.ColumnInfo)
}

func NewColumnInfoList(list []SDK.ColumnInfo) []ColumnInfo {
	if len(list) == 0 {
		return nil
	}

	results := make([]ColumnInfo, len(list))
	for i, v := range list {
		results[i] = NewColumnInfo(v)
	}
	return results
}

func NewColumnInfo(o SDK.ColumnInfo) ColumnInfo {
	r := ColumnInfo{}

	if o.Name != nil {
		r.Name = *o.Name
	}
	if o.Type != nil {
		r.Type = *o.Type
	}
	if o.CaseSensitive != nil {
		r.CaseSensitive = *o.CaseSensitive
	}
	if o.CatalogName != nil {
		r.CatalogName = *o.CatalogName
	}
	if o.Label != nil {
		r.Label = *o.Label
	}
	if o.Precision != nil {
		r.Precision = *o.Precision
	}
	if o.Scale != nil {
		r.Scale = *o.Scale
	}
	if o.SchemaName != nil {
		r.SchemaName = *o.SchemaName
	}
	if o.TableName != nil {
		r.TableName = *o.TableName
	}

	r.Nullable = ColumnNullable(o.Nullable)
	return r
}

type Datum struct {
	VarCharValue string

	HasValue bool
}

func NewDatumList(list []SDK.Datum) []Datum {
	if len(list) == 0 {
		return nil
	}

	results := make([]Datum, len(list))
	for i, v := range list {
		results[i] = NewDatum(v)
	}
	return results
}

func NewDatum(o SDK.Datum) Datum {
	r := Datum{}
	if o.VarCharValue != nil {
		r.VarCharValue = *o.VarCharValue
		r.HasValue = true
	}
	return r
}

type QueryExecutionContext struct {
	Catalog  string
	Database string
}

func NewQueryExecutionContext(o *SDK.QueryExecutionContext) QueryExecutionContext {
	r := QueryExecutionContext{}
	if o == nil {
		return r
	}

	if o.Catalog != nil {
		r.Catalog = *o.Catalog
	}
	if o.Database != nil {
		r.Database = *o.Database
	}
	return r
}

func (v QueryExecutionContext) ToSDK() *SDK.QueryExecutionContext {
	switch {
	case v.Catalog != "",
		v.Database != "":
	default:
		return nil
	}

	o := SDK.QueryExecutionContext{}
	if v.Catalog != "" {
		o.Catalog = pointers.String(v.Catalog)
	}
	if v.Database != "" {
		o.Database = pointers.String(v.Database)
	}
	return &o
}

type QueryExecutionStatistics struct {
	DataManifestLocation          string
	DataScannedInBytes            int64
	EngineExecutionTimeInMillis   int64
	QueryPlanningTimeInMillis     int64
	QueryQueueTimeInMillis        int64
	ServiceProcessingTimeInMillis int64
	TotalExecutionTimeInMillis    int64
}

func NewQueryExecutionStatistics(o *SDK.QueryExecutionStatistics) QueryExecutionStatistics {
	r := QueryExecutionStatistics{}
	if o == nil {
		return r
	}

	if o.DataManifestLocation != nil {
		r.DataManifestLocation = *o.DataManifestLocation
	}
	if o.DataScannedInBytes != nil {
		r.DataScannedInBytes = *o.DataScannedInBytes
	}
	if o.EngineExecutionTimeInMillis != nil {
		r.EngineExecutionTimeInMillis = *o.EngineExecutionTimeInMillis
	}
	if o.QueryPlanningTimeInMillis != nil {
		r.QueryPlanningTimeInMillis = *o.QueryPlanningTimeInMillis
	}
	if o.QueryQueueTimeInMillis != nil {
		r.QueryQueueTimeInMillis = *o.QueryQueueTimeInMillis
	}
	if o.ServiceProcessingTimeInMillis != nil {
		r.ServiceProcessingTimeInMillis = *o.ServiceProcessingTimeInMillis
	}
	if o.TotalExecutionTimeInMillis != nil {
		r.TotalExecutionTimeInMillis = *o.TotalExecutionTimeInMillis
	}
	return r
}

type QueryExecutionStatus struct {
	CompletionDateTime time.Time
	State              QueryExecutionState
	StateChangeReason  string
	SubmissionDateTime time.Time
}

func NewQueryExecutionStatus(o *SDK.QueryExecutionStatus) QueryExecutionStatus {
	r := QueryExecutionStatus{}
	if o == nil {
		return r
	}

	if o.CompletionDateTime != nil {
		r.CompletionDateTime = *o.CompletionDateTime
	}
	if o.StateChangeReason != nil {
		r.StateChangeReason = *o.StateChangeReason
	}
	if o.SubmissionDateTime != nil {
		r.SubmissionDateTime = *o.SubmissionDateTime
	}
	r.State = QueryExecutionState(o.State)
	return r
}

type ResultConfiguration struct {
	OutputLocation   string
	EncryptionKMSKey string
	EncryptionOption EncryptionOption
}

func NewResultConfiguration(o *SDK.ResultConfiguration) ResultConfiguration {
	r := ResultConfiguration{}
	if o == nil {
		return r
	}

	if o.OutputLocation != nil {
		r.OutputLocation = *o.OutputLocation
	}
	if v := o.EncryptionConfiguration; v != nil {
		r.EncryptionOption = EncryptionOption(v.EncryptionOption)
		if v.KmsKey != nil {
			r.EncryptionKMSKey = *v.KmsKey
		}
	}
	return r
}

func (v ResultConfiguration) ToSDK() *SDK.ResultConfiguration {
	switch {
	case v.OutputLocation != "",
		v.EncryptionKMSKey != "",
		v.EncryptionOption != "":
	default:
		return nil
	}

	o := SDK.ResultConfiguration{}

	if v.OutputLocation != "" {
		o.OutputLocation = pointers.String(v.OutputLocation)
	}
	switch {
	case v.EncryptionKMSKey != "",
		v.EncryptionOption != "":
		vv := SDK.EncryptionConfiguration{}
		vv.EncryptionOption = SDK.EncryptionOption(v.EncryptionOption)
		if v.EncryptionKMSKey != "" {
			vv.KmsKey = pointers.String(v.EncryptionKMSKey)
		}
		o.EncryptionConfiguration = &vv
	}
	return &o
}

type ResultSet struct {
	ColumnInfo []ColumnInfo
	Rows       []Row
}

func NewResultSet(o *SDK.ResultSet) ResultSet {
	r := ResultSet{}
	if o == nil {
		return r
	}
	r.ColumnInfo = NewColumnInfoListFromMetadata(o.ResultSetMetadata)
	r.Rows = NewRowList(o.Rows)
	return r
}

func (r ResultSet) ToMapString() []map[string]string {
	if len(r.Rows) < 2 {
		return nil
	}

	results := make([]map[string]string, len(r.Rows)-1)
	hd := r.Rows[0].Data
	header := make([]string, len(hd))
	for i, v := range hd {
		header[i] = v.VarCharValue
	}

	for i, row := range r.Rows[1:] {
		m := make(map[string]string, len(row.Data))
		for j, v := range row.Data {
			m[header[j]] = v.VarCharValue
		}
		results[i] = m
	}
	return results
}

func (r ResultSet) ToListString() [][]string {
	if len(r.Rows) < 2 {
		return nil
	}

	results := make([][]string, len(r.Rows))
	for i, row := range r.Rows {
		rr := make([]string, len(r.Rows))
		for j, v := range row.Data {
			rr[j] = v.VarCharValue
		}
		results[i] = rr
	}
	return results
}

type Row struct {
	Data []Datum
}

func NewRowList(list []SDK.Row) []Row {
	if len(list) == 0 {
		return nil
	}

	results := make([]Row, len(list))
	for i, v := range list {
		results[i] = NewRow(v)
	}
	return results
}

func NewRow(o SDK.Row) Row {
	r := Row{}
	r.Data = NewDatumList(o.Data)
	return r
}
