package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// HeadObject executes `HeadObject` operation.
func (svc *S3) HeadObject(ctx context.Context, r HeadObjectRequest) (*HeadObjectResult, error) {
	out, err := svc.RawHeadObject(ctx, r.ToInput())
	if err == nil {
		return NewHeadObjectResult(out), nil
	}
	if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == "NotFound" {
			return NewHeadObjectResult(out), nil
		}
	}

	err = svc.errWrap(errors.ErrorData{
		Err:          err,
		AWSOperation: "HeadObject",
	})
	svc.Errorf(err.Error())
	return nil, err
}

// HeadObjectRequest has parameters for `HeadObject` operation.
type HeadObjectRequest struct {
	Bucket string
	Key    string

	// optional
	IfMatch              string
	IfModifiedSince      time.Time
	IfNoneMatch          string
	IfUnmodifiedSince    time.Time
	PartNumber           int64
	Range                string
	RequestPayer         RequestPayer
	SSECustomerAlgorithm string
	SSECustomerKey       string
	SSECustomerKeyMD5    string
	VersionID            string
}

func (r HeadObjectRequest) ToInput() *SDK.HeadObjectInput {
	in := &SDK.HeadObjectInput{}
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
	if r.RequestPayer != "" {
		in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
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

// HeadObjectResult contains results from `HeadObject` operation.
type HeadObjectResult struct {
	IsExist bool

	AcceptRanges              string
	CacheControl              string
	ContentDisposition        string
	ContentEncoding           string
	ContentLanguage           string
	ContentLength             int64
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
	SSEKMSKeyId               string
	ServerSideEncryption      ServerSideEncryption
	StorageClass              StorageClass
	VersionID                 string
	WebsiteRedirectLocation   string
}

func NewHeadObjectResult(output *SDK.HeadObjectResponse) *HeadObjectResult {
	r := &HeadObjectResult{}
	if output == nil {
		return r
	}

	r.IsExist = true

	if output.AcceptRanges != nil {
		r.AcceptRanges = *output.AcceptRanges
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
	if output.MissingMeta != nil {
		r.MissingMeta = *output.MissingMeta
	}
	if output.PartsCount != nil {
		r.PartsCount = *output.PartsCount
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
		r.SSEKMSKeyId = *output.SSEKMSKeyId
	}
	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}
	if output.WebsiteRedirectLocation != nil {
		r.WebsiteRedirectLocation = *output.WebsiteRedirectLocation
	}

	if r.ObjectLockLegalHoldStatus != "" {
		r.ObjectLockLegalHoldStatus = ObjectLockLegalHoldStatus(output.ObjectLockLegalHoldStatus)
	}
	if r.ObjectLockMode != "" {
		r.ObjectLockMode = ObjectLockMode(output.ObjectLockMode)
	}
	if r.ReplicationStatus != "" {
		r.ReplicationStatus = ReplicationStatus(output.ReplicationStatus)
	}
	if r.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	if r.ServerSideEncryption != "" {
		r.ServerSideEncryption = ServerSideEncryption(output.ServerSideEncryption)
	}
	if r.StorageClass != "" {
		r.StorageClass = StorageClass(output.StorageClass)
	}
	r.Metadata = output.Metadata

	return r
}
