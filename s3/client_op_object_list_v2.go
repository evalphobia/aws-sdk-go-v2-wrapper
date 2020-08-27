package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ListObjectsV2 executes `ListObjectsV2` operation.
func (svc *S3) ListObjectsV2(ctx context.Context, r ListObjectsRequest) (*ListObjectsResult, error) {
	out, err := svc.RawListObjectsV2(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ListObjectsV2",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewListObjectsResultV2(out), nil
}

// ListObjectsRequest has parameters for `ListObjectsV2` operation.
type ListObjectsRequest struct {
	Bucket string

	// optional
	ContinuationToken string
	Delimiter         string
	EncodingType      string
	FetchOwner        bool
	MaxKeys           int64
	Prefix            string
	RequestPayer      string
	StartAfter        string
}

func (r ListObjectsRequest) ToInput() *SDK.ListObjectsV2Input {
	in := &SDK.ListObjectsV2Input{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.ContinuationToken != "" {
		in.ContinuationToken = pointers.String(r.ContinuationToken)
	}
	if r.Delimiter != "" {
		in.Delimiter = pointers.String(r.Delimiter)
	}
	if r.EncodingType != "" {
		in.EncodingType = SDK.EncodingType(r.EncodingType)
	}
	if r.FetchOwner {
		in.FetchOwner = pointers.Bool(r.FetchOwner)
	}
	if r.MaxKeys != 0 {
		in.MaxKeys = pointers.Long64(r.MaxKeys)
	}
	if r.Prefix != "" {
		in.Prefix = pointers.String(r.Prefix)
	}
	if r.RequestPayer != "" {
		in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
	}
	if r.StartAfter != "" {
		in.StartAfter = pointers.String(r.StartAfter)
	}
	return in
}

// ListObjectsResult contains data from ListObjectsV2.
type ListObjectsResult struct {
	CommonPrefixes        []string
	Contents              []Object
	ContinuationToken     string
	Delimiter             string
	EncodingType          EncodingType
	IsTruncated           bool
	KeyCount              int64
	MaxKeys               int64
	Name                  string
	NextContinuationToken string
	Prefix                string
	StartAfter            string
}

func NewListObjectsResultV2(output *SDK.ListObjectsV2Response) *ListObjectsResult {
	r := &ListObjectsResult{}
	if output == nil {
		return r
	}

	if output.ContinuationToken != nil {
		r.ContinuationToken = *output.ContinuationToken
	}
	if output.Delimiter != nil {
		r.Delimiter = *output.Delimiter
	}
	if output.EncodingType != "" {
		r.EncodingType = EncodingType(output.EncodingType)
	}
	if output.IsTruncated != nil {
		r.IsTruncated = *output.IsTruncated
	}
	if output.KeyCount != nil {
		r.KeyCount = *output.KeyCount
	}
	if output.MaxKeys != nil {
		r.MaxKeys = *output.MaxKeys
	}
	if output.Name != nil {
		r.Name = *output.Name
	}
	if output.NextContinuationToken != nil {
		r.NextContinuationToken = *output.NextContinuationToken
	}
	if output.Prefix != nil {
		r.Prefix = *output.Prefix
	}
	if output.StartAfter != nil {
		r.StartAfter = *output.StartAfter
	}

	if len(output.Contents) != 0 {
		list := make([]Object, len(output.Contents))
		for i, v := range output.Contents {
			list[i] = NewObject(v)
		}
		r.Contents = list
	}

	r.CommonPrefixes = newCommonPrefixes(output.CommonPrefixes)
	return r
}
