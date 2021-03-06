package s3

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetObject executes `GetObject` operation.
func (svc *S3) GetObject(ctx context.Context, r GetObjectRequest) (*GetObjectResult, error) {
	out, err := svc.RawGetObject(ctx, r.ToInput())
	if err == nil {
		return NewGetObjectResult(out), nil
	}
	if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == "NoSuchKey" {
			return NewGetObjectResult(out), nil
		}
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "GetObject",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// GetObjectRequest has parameters for `GetObject` operation.
type GetObjectRequest struct {
	Bucket string
	Key    string

	// optional
	IfMatch                    string
	IfModifiedSince            time.Time
	IfNoneMatch                string
	IfUnmodifiedSince          time.Time
	PartNumber                 int64
	Range                      string
	RequestPayer               RequestPayer
	ResponseCacheControl       string
	ResponseContentDisposition string
	ResponseContentEncoding    string
	ResponseContentLanguage    string
	ResponseContentType        string
	ResponseExpires            time.Time
	SSECustomerAlgorithm       string
	SSECustomerKey             string
	SSECustomerKeyMD5          string
	VersionID                  string
}

func (r GetObjectRequest) ToInput() *SDK.GetObjectInput {
	in := &SDK.GetObjectInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}

	if r.IfMatch != "" {
		in.IfMatch = pointers.String(r.IfMatch)
	}
	if !r.IfModifiedSince.IsZero() {
		in.IfModifiedSince = &r.IfModifiedSince
	}
	if r.IfNoneMatch != "" {
		in.IfNoneMatch = pointers.String(r.IfNoneMatch)
	}
	if !r.IfUnmodifiedSince.IsZero() {
		in.IfUnmodifiedSince = &r.IfUnmodifiedSince
	}
	if r.PartNumber != 0 {
		in.PartNumber = pointers.Long64(r.PartNumber)
	}
	if r.Range != "" {
		in.Range = pointers.String(r.Range)
	}

	in.RequestPayer = SDK.RequestPayer(r.RequestPayer)

	if r.ResponseCacheControl != "" {
		in.ResponseCacheControl = pointers.String(r.ResponseCacheControl)
	}
	if r.ResponseContentDisposition != "" {
		in.ResponseContentDisposition = pointers.String(r.ResponseContentDisposition)
	}
	if r.ResponseContentEncoding != "" {
		in.ResponseContentEncoding = pointers.String(r.ResponseContentEncoding)
	}
	if r.ResponseContentLanguage != "" {
		in.ResponseContentLanguage = pointers.String(r.ResponseContentLanguage)
	}
	if r.ResponseContentType != "" {
		in.ResponseContentType = pointers.String(r.ResponseContentType)
	}
	if !r.ResponseExpires.IsZero() {
		in.ResponseExpires = &r.ResponseExpires
	}
	if r.SSECustomerAlgorithm != "" {
		in.SSECustomerAlgorithm = pointers.String(r.SSECustomerAlgorithm)
	}
	if r.SSECustomerKey != "" {
		in.SSECustomerKey = pointers.String(r.SSECustomerKey)
	}
	if r.SSECustomerKeyMD5 != "" {
		in.SSECustomerKeyMD5 = pointers.String(r.SSECustomerKeyMD5)
	}
	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}
	return in
}

// GetObjectResult contains results from `GetObject` operation.
type GetObjectResult struct {
	Exists bool

	AcceptRanges              string
	Body                      io.ReadCloser
	CacheControl              string
	ContentDisposition        string
	ContentEncoding           string
	ContentLanguage           string
	ContentLength             int64
	ContentRange              string
	ContentType               string
	DeleteMarker              bool
	ETag                      string
	Expiration                string
	Expires                   string
	LastModified              time.Time
	Metadata                  map[string]string
	MissingMeta               int64
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus
	ObjectLockMode            ObjectLockMode
	ObjectLockRetainUntilDate time.Time
	PartsCount                int64
	ReplicationStatus         ReplicationStatus
	RequestCharged            RequestCharged
	Restore                   string
	SSECustomerAlgorithm      string
	SSECustomerKeyMD5         string
	SSEKMSKeyID               string
	ServerSideEncryption      ServerSideEncryption
	StorageClass              StorageClass
	TagCount                  int64
	VersionID                 string
	WebsiteRedirectLocation   string
}

func NewGetObjectResult(output *SDK.GetObjectResponse) *GetObjectResult {
	r := &GetObjectResult{}
	if output == nil {
		return r
	}

	r.Exists = true

	if output.AcceptRanges != nil {
		r.AcceptRanges = *output.AcceptRanges
	}
	if output.Body != nil {
		r.Body = output.Body
	}
	if output.CacheControl != nil {
		r.CacheControl = *output.CacheControl
	}
	if output.ContentDisposition != nil {
		r.ContentDisposition = *output.ContentDisposition
	}
	if output.ContentEncoding != nil {
		r.ContentEncoding = *output.ContentEncoding
	}
	if output.ContentLanguage != nil {
		r.ContentLanguage = *output.ContentLanguage
	}
	if output.ContentLength != nil {
		r.ContentLength = *output.ContentLength
	}
	if output.ContentRange != nil {
		r.ContentRange = *output.ContentRange
	}
	if output.ContentType != nil {
		r.ContentType = *output.ContentType
	}
	if output.DeleteMarker != nil {
		r.DeleteMarker = *output.DeleteMarker
	}
	if output.ETag != nil {
		r.ETag = *output.ETag
	}
	if output.Expiration != nil {
		r.Expiration = *output.Expiration
	}
	if output.Expires != nil {
		r.Expires = *output.Expires
	}
	if output.LastModified != nil {
		r.LastModified = *output.LastModified
	}

	r.Metadata = output.Metadata

	if output.MissingMeta != nil {
		r.MissingMeta = *output.MissingMeta
	}
	if r.ObjectLockLegalHoldStatus != "" {
		r.ObjectLockLegalHoldStatus = ObjectLockLegalHoldStatus(output.ObjectLockLegalHoldStatus)
	}
	if r.ObjectLockMode != "" {
		r.ObjectLockMode = ObjectLockMode(output.ObjectLockMode)
	}
	if output.PartsCount != nil {
		r.PartsCount = *output.PartsCount
	}
	if r.ReplicationStatus != "" {
		r.ReplicationStatus = ReplicationStatus(output.ReplicationStatus)
	}
	if r.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	if output.Restore != nil {
		r.Restore = *output.Restore
	}
	if output.SSECustomerAlgorithm != nil {
		r.SSECustomerAlgorithm = *output.SSECustomerAlgorithm
	}
	if output.SSECustomerKeyMD5 != nil {
		r.SSECustomerKeyMD5 = *output.SSECustomerKeyMD5
	}
	if output.SSEKMSKeyId != nil {
		r.SSEKMSKeyID = *output.SSEKMSKeyId
	}
	if r.ServerSideEncryption != "" {
		r.ServerSideEncryption = ServerSideEncryption(output.ServerSideEncryption)
	}
	if r.StorageClass != "" {
		r.StorageClass = StorageClass(output.StorageClass)
	}
	if output.TagCount != nil {
		r.TagCount = *output.TagCount
	}
	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}
	if output.WebsiteRedirectLocation != nil {
		r.WebsiteRedirectLocation = *output.WebsiteRedirectLocation
	}
	return r
}

func (r GetObjectResult) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
