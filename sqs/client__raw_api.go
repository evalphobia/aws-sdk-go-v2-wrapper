package sqs

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"
)

// RawAddPermission executes `AddPermission` raw operation.
func (svc *SQS) RawAddPermission(ctx context.Context, in *SDK.AddPermissionInput) (*SDK.AddPermissionResponse, error) {
	return svc.client.AddPermissionRequest(in).Send(ctx)
}

// RawChangeMessageVisibility executes `ChangeMessageVisibility` raw operation.
func (svc *SQS) RawChangeMessageVisibility(ctx context.Context, in *SDK.ChangeMessageVisibilityInput) (*SDK.ChangeMessageVisibilityResponse, error) {
	return svc.client.ChangeMessageVisibilityRequest(in).Send(ctx)
}

// RawChangeMessageVisibilityBatch executes `ChangeMessageVisibilityBatch` raw operation.
func (svc *SQS) RawChangeMessageVisibilityBatch(ctx context.Context, in *SDK.ChangeMessageVisibilityBatchInput) (*SDK.ChangeMessageVisibilityBatchResponse, error) {
	return svc.client.ChangeMessageVisibilityBatchRequest(in).Send(ctx)
}

// RawCreateQueue executes `CreateQueue` raw operation.
func (svc *SQS) RawCreateQueue(ctx context.Context, in *SDK.CreateQueueInput) (*SDK.CreateQueueResponse, error) {
	return svc.client.CreateQueueRequest(in).Send(ctx)
}

// RawDeleteMessage executes `DeleteMessage` raw operation.
func (svc *SQS) RawDeleteMessage(ctx context.Context, in *SDK.DeleteMessageInput) (*SDK.DeleteMessageResponse, error) {
	return svc.client.DeleteMessageRequest(in).Send(ctx)
}

// RawDeleteMessageBatch executes `DeleteMessageBatch` raw operation.
func (svc *SQS) RawDeleteMessageBatch(ctx context.Context, in *SDK.DeleteMessageBatchInput) (*SDK.DeleteMessageBatchResponse, error) {
	return svc.client.DeleteMessageBatchRequest(in).Send(ctx)
}

// RawDeleteQueue executes `DeleteQueue` raw operation.
func (svc *SQS) RawDeleteQueue(ctx context.Context, in *SDK.DeleteQueueInput) (*SDK.DeleteQueueResponse, error) {
	return svc.client.DeleteQueueRequest(in).Send(ctx)
}

// RawGetQueueAttributes executes `GetQueueAttributes` raw operation.
func (svc *SQS) RawGetQueueAttributes(ctx context.Context, in *SDK.GetQueueAttributesInput) (*SDK.GetQueueAttributesResponse, error) {
	return svc.client.GetQueueAttributesRequest(in).Send(ctx)
}

// RawGetQueueUrl executes `GetQueueUrl` raw operation.
//nolint:golint
func (svc *SQS) RawGetQueueUrl(ctx context.Context, in *SDK.GetQueueUrlInput) (*SDK.GetQueueUrlResponse, error) {
	return svc.client.GetQueueUrlRequest(in).Send(ctx)
}

// RawListDeadLetterSourceQueues executes `ListDeadLetterSourceQueues` raw operation.
func (svc *SQS) RawListDeadLetterSourceQueues(ctx context.Context, in *SDK.ListDeadLetterSourceQueuesInput) (*SDK.ListDeadLetterSourceQueuesResponse, error) {
	return svc.client.ListDeadLetterSourceQueuesRequest(in).Send(ctx)
}

// RawListQueueTags executes `ListQueueTags` raw operation.
func (svc *SQS) RawListQueueTags(ctx context.Context, in *SDK.ListQueueTagsInput) (*SDK.ListQueueTagsResponse, error) {
	return svc.client.ListQueueTagsRequest(in).Send(ctx)
}

// RawListQueues executes `ListQueues` raw operation.
func (svc *SQS) RawListQueues(ctx context.Context, in *SDK.ListQueuesInput) (*SDK.ListQueuesResponse, error) {
	return svc.client.ListQueuesRequest(in).Send(ctx)
}

// RawPurgeQueue executes `PurgeQueue` raw operation.
func (svc *SQS) RawPurgeQueue(ctx context.Context, in *SDK.PurgeQueueInput) (*SDK.PurgeQueueResponse, error) {
	return svc.client.PurgeQueueRequest(in).Send(ctx)
}

// RawReceiveMessage executes `ReceiveMessage` raw operation.
func (svc *SQS) RawReceiveMessage(ctx context.Context, in *SDK.ReceiveMessageInput) (*SDK.ReceiveMessageResponse, error) {
	return svc.client.ReceiveMessageRequest(in).Send(ctx)
}

// RawRemovePermission executes `RemovePermission` raw operation.
func (svc *SQS) RawRemovePermission(ctx context.Context, in *SDK.RemovePermissionInput) (*SDK.RemovePermissionResponse, error) {
	return svc.client.RemovePermissionRequest(in).Send(ctx)
}

// RawSendMessage executes `SendMessage` raw operation.
func (svc *SQS) RawSendMessage(ctx context.Context, in *SDK.SendMessageInput) (*SDK.SendMessageResponse, error) {
	return svc.client.SendMessageRequest(in).Send(ctx)
}

// RawSendMessageBatch executes `SendMessageBatch` raw operation.
func (svc *SQS) RawSendMessageBatch(ctx context.Context, in *SDK.SendMessageBatchInput) (*SDK.SendMessageBatchResponse, error) {
	return svc.client.SendMessageBatchRequest(in).Send(ctx)
}

// RawSetQueueAttributes executes `SetQueueAttributes` raw operation.
func (svc *SQS) RawSetQueueAttributes(ctx context.Context, in *SDK.SetQueueAttributesInput) (*SDK.SetQueueAttributesResponse, error) {
	return svc.client.SetQueueAttributesRequest(in).Send(ctx)
}

// RawTagQueue executes `TagQueue` raw operation.
func (svc *SQS) RawTagQueue(ctx context.Context, in *SDK.TagQueueInput) (*SDK.TagQueueResponse, error) {
	return svc.client.TagQueueRequest(in).Send(ctx)
}

// RawUntagQueue executes `UntagQueue` raw operation.
func (svc *SQS) RawUntagQueue(ctx context.Context, in *SDK.UntagQueueInput) (*SDK.UntagQueueResponse, error) {
	return svc.client.UntagQueueRequest(in).Send(ctx)
}
