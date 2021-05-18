package sqs

import (
	"context"
	"strconv"
)

// XIsEmptyQueue checks if the queue does not have any message.
func (svc *SQS) XIsEmptyQueue(ctx context.Context, queueURL string) (bool, error) {
	resp, err := svc.GetQueueAttributes(ctx, GetQueueAttributesRequest{
		QueueURL: queueURL,
		AttributeNames: []QueueAttributeName{
			QueueAttributeNameApproximateNumberOfMessages,
			QueueAttributeNameApproximateNumberOfMessagesDelayed,
			QueueAttributeNameApproximateNumberOfMessagesNotVisible,
		},
	})
	if err != nil {
		return false, err
	}

	// if v, ok := resp.Attributes[QueueAttributeNameApproximateNumberOfMessages]; ok {

	// }
	for _, v := range resp.Attributes {
		num, err := strconv.Atoi(v)
		if err != nil {
			return false, err
		}
		if num > 0 {
			return false, nil
		}
	}
	return true, nil
}
