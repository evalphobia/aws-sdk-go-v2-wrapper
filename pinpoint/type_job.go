package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type ImportJobRequest struct {
	Format  Format
	RoleARN string
	S3URL   string

	// optional
	DefineSegment     bool
	RegisterEndpoints bool
	SegmentID         string
	SegmentName       string
}

func (r ImportJobRequest) ToSDK() *SDK.ImportJobRequest {
	o := SDK.ImportJobRequest{}

	o.Format = SDK.Format(r.Format)

	if r.RoleARN != "" {
		o.RoleArn = pointers.String(r.RoleARN)
	}
	if r.S3URL != "" {
		o.S3Url = pointers.String(r.S3URL)
	}

	if r.DefineSegment {
		o.DefineSegment = pointers.Bool(r.DefineSegment)
	}
	if r.RegisterEndpoints {
		o.RegisterEndpoints = pointers.Bool(r.RegisterEndpoints)
	}
	if r.SegmentID != "" {
		o.SegmentId = pointers.String(r.SegmentID)
	}
	if r.SegmentName != "" {
		o.SegmentName = pointers.String(r.SegmentName)
	}
	return &o
}

type ImportJobsResponse struct {
	Item      []ImportJobResponse
	NextToken string
}

func newImportJobsResponse(o *SDK.ImportJobsResponse) ImportJobsResponse {
	result := ImportJobsResponse{}
	if o == nil {
		return result
	}

	if len(o.Item) != 0 {
		list := make([]ImportJobResponse, len(o.Item))
		for i, v := range o.Item {
			v := v
			list[i] = newImportJobResponse(&v)
		}
		result.Item = list
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}

	return result
}

type ImportJobResponse struct {
	ApplicationID string
	CreationDate  string
	Definition    ImportJobResource
	ID            string
	JobStatus     JobStatus
	Type          string

	// optional
	CompletedPieces int64
	CompletionDate  string
	FailedPieces    int64
	Failures        []string
	TotalFailures   int64
	TotalPieces     int64
	TotalProcessed  int64
}

func newImportJobResponse(o *SDK.ImportJobResponse) ImportJobResponse {
	result := ImportJobResponse{}
	if o == nil {
		return result
	}

	if o.ApplicationId != nil {
		result.ApplicationID = *o.ApplicationId
	}
	if o.CreationDate != nil {
		result.CreationDate = *o.CreationDate
	}

	result.Definition = newImportJobResource(o.Definition)

	if o.Id != nil {
		result.ID = *o.Id
	}

	result.JobStatus = JobStatus(o.JobStatus)

	if o.Type != nil {
		result.Type = *o.Type
	}

	if o.CompletedPieces != nil {
		result.CompletedPieces = *o.CompletedPieces
	}
	if o.CompletionDate != nil {
		result.CompletionDate = *o.CompletionDate
	}
	if o.FailedPieces != nil {
		result.FailedPieces = *o.FailedPieces
	}

	result.Failures = o.Failures

	if o.TotalFailures != nil {
		result.TotalFailures = *o.TotalFailures
	}
	if o.TotalPieces != nil {
		result.TotalPieces = *o.TotalPieces
	}
	if o.TotalProcessed != nil {
		result.TotalProcessed = *o.TotalProcessed
	}
	return result
}

type ImportJobResource struct {
	Format  Format
	RoleARN string
	S3URL   string

	// optional
	DefineSegment     bool
	ExternalID        string
	RegisterEndpoints bool
	SegmentID         string
	SegmentName       string
}

func newImportJobResource(o *SDK.ImportJobResource) ImportJobResource {
	result := ImportJobResource{}
	if o == nil {
		return result
	}

	result.Format = Format(o.Format)

	if o.RoleArn != nil {
		result.RoleARN = *o.RoleArn
	}
	if o.S3Url != nil {
		result.S3URL = *o.S3Url
	}

	if o.DefineSegment != nil {
		result.DefineSegment = *o.DefineSegment
	}
	if o.ExternalId != nil {
		result.ExternalID = *o.ExternalId
	}
	if o.RegisterEndpoints != nil {
		result.RegisterEndpoints = *o.RegisterEndpoints
	}
	if o.SegmentId != nil {
		result.SegmentID = *o.SegmentId
	}
	if o.SegmentName != nil {
		result.SegmentName = *o.SegmentName
	}
	return result
}
