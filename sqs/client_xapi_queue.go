package sqs

import (
	"context"
	"fmt"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// XCreateQueue creates a SQS queue.
func (svc *SQS) XCreateQueue(ctx context.Context, queueName string) (queueURL string, err error) {
	resp, err := svc.CreateQueue(ctx, CreateQueueRequest{
		QueueName: queueName,
	})
	if err != nil {
		return "", err
	}
	return resp.QueueURL, nil
}

// XDeleteQueue deletes a SQS queue.
func (svc *SQS) XDeleteQueue(ctx context.Context, queueURL string) error {
	_, err := svc.DeleteQueue(ctx, DeleteQueueRequest{
		QueueURL: queueURL,
	})
	return err
}

// XDeleteQueueByName delete a SQS queue by queue name.
func (svc *SQS) XDeleteQueueByName(ctx context.Context, queueName string) error {
	queueURL, err := svc.XGetQueueURL(ctx, queueName)
	switch {
	case err != nil:
		return err
	case queueURL == "":
		return fmt.Errorf("Cannot find the queue name:[%s]", queueName)
	}

	_, err = svc.DeleteQueue(ctx, DeleteQueueRequest{
		QueueURL: queueURL,
	})
	return err
}

// XExistsQueue checks if the SQS queue exists or not.
func (svc *SQS) XExistsQueue(ctx context.Context, queueName string) (bool, error) {
	_, err := svc.GetQueueURL(ctx, GetQueueURLRequest{
		QueueName: queueName,
	})
	if err == nil {
		return true, nil
	}

	if e, ok := err.(errors.ErrorData); ok {
		if e.GetAWSErrCode() == "AWS.SimpleQueueService.NonExistentQueue" {
			return false, nil
		}
	}
	return false, err
}

// XGetQueueURL gets a SQS queue url.
func (svc *SQS) XGetQueueURL(ctx context.Context, queueName string) (queueURL string, err error) {
	resp, err := svc.GetQueueURL(ctx, GetQueueURLRequest{
		QueueName: queueName,
	})
	if err != nil {
		return "", err
	}
	return resp.QueueURL, nil
}

// XListAllQueueURLs lists all of the SQS queue url.
func (svc *SQS) XListAllQueueURLs(ctx context.Context, queuePrefix string) (queueURLs []string, err error) {
	resp, err := svc.ListQueues(ctx, ListQueuesRequest{
		QueueNamePrefix: queuePrefix,
	})
	if err != nil {
		return nil, err
	}

	results := resp.QueueURLs
	nextToken := resp.NextToken
	for nextToken != "" {
		resp, err := svc.ListQueues(ctx, ListQueuesRequest{
			QueueNamePrefix: queuePrefix,
			NextToken:       nextToken,
		})
		if err != nil {
			return nil, err
		}

		results = append(results, resp.QueueURLs...)
		nextToken = resp.NextToken
	}

	return results, nil
}

// XPurgeQueue purges messages of the queue.
func (svc *SQS) XPurgeQueue(ctx context.Context, queueURL string) error {
	_, err := svc.PurgeQueue(ctx, PurgeQueueRequest{
		QueueURL: queueURL,
	})
	return err
}

// XPurgeQueueByName purges messages of the queue by queue name.
func (svc *SQS) XPurgeQueueByName(ctx context.Context, queueName string) error {
	queueURL, err := svc.XGetQueueURL(ctx, queueName)
	switch {
	case err != nil:
		return err
	case queueURL == "":
		return fmt.Errorf("Cannot find the queue name:[%s]", queueName)
	}

	_, err = svc.PurgeQueue(ctx, PurgeQueueRequest{
		QueueURL: queueURL,
	})
	return err
}
