package kms

import (
	"context"
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

// ScheduleKeyDeletion executes `ScheduleKeyDeletion` operation.
func (svc *KMS) ScheduleKeyDeletion(ctx context.Context, r ScheduleKeyDeletionRequest) (*ScheduleKeyDeletionResult, error) {
	out, err := svc.RawScheduleKeyDeletion(ctx, r.ToInput())
	if err != nil {
		err = svc.errWrap(errors.ErrorData{
			Err:          err,
			AWSOperation: "ScheduleKeyDeletion",
		})
		svc.Errorf(err.Error())
		return nil, err
	}
	return NewScheduleKeyDeletionResult(out), nil
}

// ScheduleKeyDeletionRequest has parameters for `ScheduleKeyDeletion` operation.
type ScheduleKeyDeletionRequest struct {
	KeyID string

	// optional
	PendingWindowInDays int64 // must be [7 ~ 30] days
}

func (r ScheduleKeyDeletionRequest) ToInput() *SDK.ScheduleKeyDeletionInput {
	in := &SDK.ScheduleKeyDeletionInput{}
	if r.KeyID != "" {
		in.KeyId = pointers.String(r.KeyID)
	}
	if r.PendingWindowInDays != 0 {
		in.PendingWindowInDays = pointers.Long64(r.PendingWindowInDays)
	}
	return in
}

type ScheduleKeyDeletionResult struct {
	DeletionDate time.Time
	KeyID        string
}

func NewScheduleKeyDeletionResult(output *SDK.ScheduleKeyDeletionResponse) *ScheduleKeyDeletionResult {
	r := &ScheduleKeyDeletionResult{}
	if output == nil {
		return r
	}

	if output.DeletionDate != nil {
		r.DeletionDate = *output.DeletionDate
	}
	if output.KeyId != nil {
		r.KeyID = *output.KeyId
	}
	return r
}
