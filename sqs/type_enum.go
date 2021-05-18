package sqs

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"
)

type QueueAttributeName string

const (
	QueueAttributeNameAll                                   QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameAll)
	QueueAttributeNamePolicy                                QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNamePolicy)
	QueueAttributeNameVisibilityTimeout                     QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameVisibilityTimeout)
	QueueAttributeNameMaximumMessageSize                    QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameMaximumMessageSize)
	QueueAttributeNameMessageRetentionPeriod                QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameMessageRetentionPeriod)
	QueueAttributeNameApproximateNumberOfMessages           QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameApproximateNumberOfMessages)
	QueueAttributeNameApproximateNumberOfMessagesNotVisible QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameApproximateNumberOfMessagesNotVisible)
	QueueAttributeNameCreatedTimestamp                      QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameCreatedTimestamp)
	QueueAttributeNameLastModifiedTimestamp                 QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameLastModifiedTimestamp)
	QueueAttributeNameQueueArn                              QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameQueueArn)
	QueueAttributeNameApproximateNumberOfMessagesDelayed    QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameApproximateNumberOfMessagesDelayed)
	QueueAttributeNameDelaySeconds                          QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameDelaySeconds)
	QueueAttributeNameReceiveMessageWaitTimeSeconds         QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameReceiveMessageWaitTimeSeconds)
	QueueAttributeNameRedrivePolicy                         QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameRedrivePolicy)
	QueueAttributeNameFifoQueue                             QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameFifoQueue)
	QueueAttributeNameContentBasedDeduplication             QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameContentBasedDeduplication)
	QueueAttributeNameKmsMasterKeyID                        QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameKmsMasterKeyId)
	QueueAttributeNameKmsDataKeyReusePeriodSeconds          QueueAttributeName = QueueAttributeName(SDK.QueueAttributeNameKmsDataKeyReusePeriodSeconds)
)
