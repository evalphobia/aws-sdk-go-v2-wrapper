package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListObjectVersions executes `ListObjectVersions` operation.
func (svc *S3) ListObjectVersions(ctx context.Context, r ListObjectVersionsRequest) (*ListObjectVersionsResult, error) {
	out, err := svc.RawListObjectVersions(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListObjectVersions",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListObjectVersionsResult(out), nil
}

// ListObjectVersionsRequest has parameters for `ListObjectVersions` operation.
type ListObjectVersionsRequest struct {
	Bucket string

	// optional
	Delimiter       string
	EncodingType    EncodingType
	KeyMarker       string
	MaxKeys         int64
	Prefix          string
	VersionIDMarker string
}

func (r ListObjectVersionsRequest) ToInput() *SDK.ListObjectVersionsInput {
	in := &SDK.ListObjectVersionsInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Delimiter != "" {
		in.Delimiter = pointers.String(r.Delimiter)
	}
	if r.KeyMarker != "" {
		in.KeyMarker = pointers.String(r.KeyMarker)
	}
	if r.MaxKeys != 0 {
		in.MaxKeys = pointers.Long64(r.MaxKeys)
	}
	if r.Prefix != "" {
		in.Prefix = pointers.String(r.Prefix)
	}
	if r.VersionIDMarker != "" {
		in.VersionIdMarker = pointers.String(r.VersionIDMarker)
	}

	if r.EncodingType != "" {
		in.EncodingType = SDK.EncodingType(r.EncodingType)
	}
	return in
}

// ListObjectVersionsResult contains results from `ListObjectVersions` operation.
type ListObjectVersionsResult struct {
	CommonPrefixes      []string
	DeleteMarkers       []DeleteMarkerEntry
	Delimiter           string
	EncodingType        EncodingType
	IsTruncated         bool
	KeyMarker           string
	MaxKeys             int64
	Name                string
	NextKeyMarker       string
	NextVersionIDMarker string
	Prefix              string
	VersionIDMarker     string
	Versions            []ObjectVersion
}

func NewListObjectVersionsResult(output *SDK.ListObjectVersionsResponse) *ListObjectVersionsResult {
	r := &ListObjectVersionsResult{}
	if output == nil {
		return r
	}

	if output.Delimiter != nil {
		r.Delimiter = *output.Delimiter
	}
	if output.IsTruncated != nil {
		r.IsTruncated = *output.IsTruncated
	}
	if output.KeyMarker != nil {
		r.KeyMarker = *output.KeyMarker
	}
	if output.MaxKeys != nil {
		r.MaxKeys = *output.MaxKeys
	}
	if output.Name != nil {
		r.Name = *output.Name
	}
	if output.NextKeyMarker != nil {
		r.NextKeyMarker = *output.NextKeyMarker
	}
	if output.NextVersionIdMarker != nil {
		r.NextVersionIDMarker = *output.NextVersionIdMarker
	}
	if output.Prefix != nil {
		r.Prefix = *output.Prefix
	}
	if output.VersionIdMarker != nil {
		r.VersionIDMarker = *output.VersionIdMarker
	}

	if output.EncodingType != "" {
		r.EncodingType = EncodingType(output.EncodingType)
	}

	r.DeleteMarkers = newDeleteMarkerEntries(output.DeleteMarkers)
	r.CommonPrefixes = newCommonPrefixes(output.CommonPrefixes)
	r.Versions = newObjectVersions(output.Versions)
	return r
}
