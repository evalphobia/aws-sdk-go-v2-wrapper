package dynamodb

import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// DeleteTableFromName deletes a table of `name`.
func (svc *DynamoDB) DeleteTableFromName(ctx context.Context, name string) error {
	_, err := svc.DeleteTable(ctx, DeleteTableRequest{
		TableName: name,
	})
	return err
}

// ExistTable checks if the table already exists or not.
func (svc *DynamoDB) ExistTable(ctx context.Context, name string) (bool, error) {
	_, err := svc.DescribeTable(ctx, DescribeTableRequest{
		TableName: name,
	})
	if err == nil {
		return true, nil
	}

	if e, ok := err.(errors.ErrorData); ok {
		if e.GetAWSErrCode() == "ResourceNotFoundException" {
			return false, nil
		}
	}
	return false, err
}
