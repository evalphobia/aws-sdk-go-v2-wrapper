package sqs

import (
	"context"
	"strconv"
)

// XDeleteMessage deletes a message.
func (svc *SQS) XDeleteMessage(ctx context.Context, queueURL, receiptHandle string) error {
	_, err := svc.DeleteMessage(ctx, DeleteMessageRequest{
		QueueURL:      queueURL,
		ReceiptHandle: receiptHandle,
	})
	return err
}

// XReceiveMessage receives messages.
func (svc *SQS) XReceiveMessage(ctx context.Context, queueURL string) ([]Message, error) {
	resp, err := svc.ReceiveMessage(ctx, ReceiveMessageRequest{
		QueueURL: queueURL,
	})
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}

// XReceiveMessageLongPolling receives messages with long polling.
func (svc *SQS) XReceiveMessageLongPolling(ctx context.Context, queueURL string, waitSec int64) ([]Message, error) {
	resp, err := svc.ReceiveMessage(ctx, ReceiveMessageRequest{
		QueueURL:        queueURL,
		WaitTimeSeconds: waitSec,
	})
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}

// XSendMessage sends a message.
func (svc *SQS) XSendMessage(ctx context.Context, queueURL, message string) error {
	_, err := svc.SendMessage(ctx, SendMessageRequest{
		QueueURL:    queueURL,
		MessageBody: message,
	})
	return err
}

// XSendMessageBatch sends messages.
func (svc *SQS) XSendMessageBatch(ctx context.Context, queueURL string, messages []string) (XSendMessageBatchResult, error) {
	const chunkSize = 10
	var chunkList [][]SendMessageBatchRequestEntry
	idNum := 0
	for i, max := 0, len(messages); i < max; i += chunkSize {
		end := i + chunkSize
		if end > max {
			end = max
		}

		chunk := make([]SendMessageBatchRequestEntry, 0, chunkSize)
		for _, v := range messages[i:end] {
			chunk = append(chunk, SendMessageBatchRequestEntry{
				ID:          strconv.Itoa(idNum),
				MessageBody: v,
			})
			idNum++
		}
		chunkList = append(chunkList, chunk)
	}

	result := XSendMessageBatchResult{}
	for _, v := range chunkList {
		resp, err := svc.SendMessageBatch(ctx, SendMessageBatchRequest{
			QueueURL: queueURL,
			Entries:  v,
		})
		if err != nil {
			return result, err
		}

		for _, v := range resp.Successful {
			idNum, _ := strconv.Atoi(v.ID)
			result.Success = append(result.Success, idNum)
		}
		for _, v := range resp.Failed {
			idNum, _ := strconv.Atoi(v.ID)
			result.Failed = append(result.Failed, idNum)
		}
		result.FailEntries = append(result.FailEntries, resp.Failed...)
	}

	return result, nil
}

type XSendMessageBatchResult struct {
	// `[]int` contains slice index of success/fail messages
	Success []int
	Failed  []int

	FailEntries []BatchResultErrorEntry
}
