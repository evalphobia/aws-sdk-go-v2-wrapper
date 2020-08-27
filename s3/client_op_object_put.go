package s3

import (
	"bytes"
	"context"
	"io"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutObject executes `PutObject` operation.
func (svc *S3) PutObject(ctx context.Context, r PutObjectRequest) (*PutObjectResult, error) {
	out, err := svc.RawPutObject(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutObject",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewPutObjectResult(out), nil
}

// PutObjectRequest has parameters for `PutObject` operation.
type PutObjectRequest struct {
	Bucket string
	Key    string

	// optional
	ACL                       ObjectCannedACL
	Body                      io.ReadSeeker
	BodyBytes                 []byte
	CacheControl              string
	ContentDisposition        string
	ContentEncoding           string
	ContentLanguage           string
	ContentLength             int64
	ContentMD5                string
	ContentType               string
	Expires                   time.Time
	GrantFullControl          string
	GrantRead                 string
	GrantReadACP              string
	GrantWriteACP             string
	Metadata                  map[string]string
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus
	ObjectLockMode            ObjectLockMode
	ObjectLockRetainUntilDate time.Time
	RequestPayer              RequestPayer
	SSECustomerAlgorithm      string
	SSECustomerKey            string
	SSECustomerKeyMD5         string
	SSEKMSEncryptionContext   string
	SSEKMSKeyId               string
	ServerSideEncryption      ServerSideEncryption
	StorageClass              StorageClass
	Tagging                   string
	WebsiteRedirectLocation   string
}

func (r PutObjectRequest) ToInput() *SDK.PutObjectInput {
	in := &SDK.PutObjectInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}

	if r.CacheControl != "" {
		in.CacheControl = pointers.String(r.CacheControl)
	}
	if r.ContentDisposition != "" {
		in.ContentDisposition = pointers.String(r.ContentDisposition)
	}
	if r.ContentEncoding != "" {
		in.ContentEncoding = pointers.String(r.ContentEncoding)
	}
	if r.ContentLanguage != "" {
		in.ContentLanguage = pointers.String(r.ContentLanguage)
	}
	if r.ContentLength != 0 {
		in.ContentLength = pointers.Long64(r.ContentLength)
	}
	if r.ContentMD5 != "" {
		in.ContentMD5 = pointers.String(r.ContentMD5)
	}
	if r.ContentType != "" {
		in.ContentType = pointers.String(r.ContentType)
	}
	if !r.Expires.IsZero() {
		in.Expires = &r.Expires
	}
	if r.GrantFullControl != "" {
		in.GrantFullControl = pointers.String(r.GrantFullControl)
	}
	if r.GrantRead != "" {
		in.GrantRead = pointers.String(r.GrantRead)
	}
	if r.GrantReadACP != "" {
		in.GrantReadACP = pointers.String(r.GrantReadACP)
	}
	if r.GrantWriteACP != "" {
		in.GrantWriteACP = pointers.String(r.GrantWriteACP)
	}
	if !r.ObjectLockRetainUntilDate.IsZero() {
		in.ObjectLockRetainUntilDate = &r.ObjectLockRetainUntilDate
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
	if r.SSEKMSEncryptionContext != "" {
		in.SSEKMSEncryptionContext = pointers.String(r.SSEKMSEncryptionContext)
	}
	if r.SSEKMSKeyId != "" {
		in.SSEKMSKeyId = pointers.String(r.SSEKMSKeyId)
	}
	if r.Tagging != "" {
		in.Tagging = pointers.String(r.Tagging)
	}
	if r.WebsiteRedirectLocation != "" {
		in.WebsiteRedirectLocation = pointers.String(r.WebsiteRedirectLocation)
	}

	if r.ACL != "" {
		in.ACL = SDK.ObjectCannedACL(r.ACL)
	}
	if r.ObjectLockLegalHoldStatus != "" {
		in.ObjectLockLegalHoldStatus = SDK.ObjectLockLegalHoldStatus(r.ObjectLockLegalHoldStatus)
	}
	if r.ObjectLockMode != "" {
		in.ObjectLockMode = SDK.ObjectLockMode(r.ObjectLockMode)
	}
	if r.RequestPayer != "" {
		in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
	}
	if r.ServerSideEncryption != "" {
		in.ServerSideEncryption = SDK.ServerSideEncryption(r.ServerSideEncryption)
	}
	if r.StorageClass != "" {
		in.StorageClass = SDK.StorageClass(r.StorageClass)
	}

	in.Metadata = r.Metadata

	switch {
	case r.Body != nil:
		in.Body = r.Body
	case len(r.BodyBytes) != 0:
		in.Body = bytes.NewReader(r.BodyBytes)
	}

	return in
}

// PutObjectResult contains results from `PutObject` operation.
type PutObjectResult struct {
	ETag                    string
	Expiration              string
	RequestCharged          RequestCharged
	SSECustomerAlgorithm    string
	SSECustomerKeyMD5       string
	SSEKMSEncryptionContext string
	SSEKMSKeyId             string
	ServerSideEncryption    ServerSideEncryption
	VersionID               string
}

func NewPutObjectResult(output *SDK.PutObjectResponse) *PutObjectResult {
	r := &PutObjectResult{}
	if output == nil {
		return r
	}

	if output.ETag != nil {
		r.ETag = *output.ETag
	}
	if output.Expiration != nil {
		r.Expiration = *output.Expiration
	}
	if output.SSECustomerAlgorithm != nil {
		r.SSECustomerAlgorithm = *output.SSECustomerAlgorithm
	}
	if output.SSECustomerKeyMD5 != nil {
		r.SSECustomerKeyMD5 = *output.SSECustomerKeyMD5
	}
	if output.SSEKMSEncryptionContext != nil {
		r.SSEKMSEncryptionContext = *output.SSEKMSEncryptionContext
	}
	if output.SSEKMSKeyId != nil {
		r.SSEKMSKeyId = *output.SSEKMSKeyId
	}
	if output.VersionId != nil {
		r.VersionID = *output.VersionId
	}

	if output.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	if output.ServerSideEncryption != "" {
		r.ServerSideEncryption = ServerSideEncryption(output.ServerSideEncryption)
	}
	return r
}
