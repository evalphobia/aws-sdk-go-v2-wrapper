package sqs

import (
	"context"
	"strconv"
)

// XIsEmptyQueue checks if the queue does not have any message.
func (svc *SQS) XIsEmptyQueue(ctx context.Context, queueURL string) (bool, error) {
	attrList := []QueueAttributeName{
		QueueAttributeNameApproximateNumberOfMessages,
		QueueAttributeNameApproximateNumberOfMessagesDelayed,
		QueueAttributeNameApproximateNumberOfMessagesNotVisible,
	}
	resp, err := svc.GetQueueAttributes(ctx, GetQueueAttributesRequest{
		QueueURL:       queueURL,
		AttributeNames: attrList,
	})
	if err != nil {
		return false, err
	}

	result := resp.Attributes
	for _, v := range attrList {
		attr, ok := result[string(v)]
		if !ok {
			return false, nil
		}
		num, err := strconv.Atoi(attr)
		if err != nil {
			return false, err
		}
		if num > 0 {
			return false, nil
		}
	}
	return true, nil
}
