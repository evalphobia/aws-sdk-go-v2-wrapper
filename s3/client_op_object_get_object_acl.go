package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// GetObjectACL executes `GetObjectACL` operation.
func (svc *S3) GetObjectACL(ctx context.Context, r GetObjectACLRequest) (*GetObjectACLResult, error) {
	out, err := svc.RawGetObjectAcl(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "GetObjectACL",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewGetObjectACLResult(out), nil
}

// GetObjectACLRequest has parameters for `GetObjectACL` operation.
type GetObjectACLRequest struct {
	Bucket string
	Key    string

	// optional
	RequestPayer RequestPayer
	VersionID    string
}

func (r GetObjectACLRequest) ToInput() *SDK.GetObjectAclInput {
	in := &SDK.GetObjectAclInput{}
	if r.Bucket != "" {
		in.Bucket = pointers.String(r.Bucket)
	}
	if r.Key != "" {
		in.Key = pointers.String(r.Key)
	}
	if r.RequestPayer != "" {
		in.RequestPayer = SDK.RequestPayer(r.RequestPayer)
	}
	if r.VersionID != "" {
		in.VersionId = pointers.String(r.VersionID)
	}
	return in
}

// GetObjectACLResult contains results from `GetObjectACL` operation.
type GetObjectACLResult struct {
	Grants         []Grant
	Owner          Owner
	RequestCharged RequestCharged
}

func NewGetObjectACLResult(output *SDK.GetObjectAclResponse) *GetObjectACLResult {
	r := &GetObjectACLResult{}
	if output == nil {
		return r
	}

	if r.RequestCharged != "" {
		r.RequestCharged = RequestCharged(output.RequestCharged)
	}

	r.Grants = newGrants(output.Grants)
	r.Owner = newOwner(output.Owner)
	return r
}
