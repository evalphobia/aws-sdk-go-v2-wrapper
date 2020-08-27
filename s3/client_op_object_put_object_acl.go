package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// PutObjectACL executes `PutObjectACL` operation.
func (svc *S3) PutObjectACL(ctx context.Context, r PutObjectACLRequest) (*PutObjectACLResult, error) {
	out, err := svc.RawPutObjectAcl(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "PutObjectACL",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewPutObjectACLResult(out), nil
}

// PutObjectACLRequest has parameters for `PutObjectACL` operation.
type PutObjectACLRequest struct {
	Bucket string
	Key    string

	// optional
	ACL                 ObjectCannedACL
	AccessControlPolicy AccessControlPolicy
	GrantFullControl    string
	GrantRead           string
	GrantReadACP        string
	GrantWrite          string
	GrantWriteACP       string
	RequestPayer        RequestPayer
	VersionID           string
}

func (r PutObjectACLRequest) ToInput() *SDK.PutObjectAclInput {
	in := &SDK.PutObjectAclInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
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
	if r.GrantWrite != "" {
		in.GrantWrite = pointers.String(r.GrantWrite)
	}
	if r.GrantWriteACP != "" {
		in.GrantWriteACP = pointers.String(r.GrantWriteACP)
	}
	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}

	if r.ACL != "" {
		in.ACL = SDK.ObjectCannedACL(r.ACL)
	}
	if r.RequestPayer != "" {
		in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
	}
	in.AccessControlPolicy = r.AccessControlPolicy.ToSDK()
	return in
}

// PutObjectACLResult contains results from `PutObjectACL` operation.
type PutObjectACLResult struct {
	RequestCharged RequestCharged
}

func NewPutObjectACLResult(output *SDK.PutObjectAclResponse) *PutObjectACLResult {
	r := &PutObjectACLResult{}
	if output == nil {
		return r
	}

	if output.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}
	return r
}
