package sqs

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type BatchResultErrorEntry struct {
	Code        string
	ID          string
	SenderFault bool

	// optional
	Message string
}

func newBatchResultErrorEntryList(list []SDK.BatchResultErrorEntry) []BatchResultErrorEntry {
	if len(list) == 0 {
		return nil
	}

	results := make([]BatchResultErrorEntry, len(list))
	for i, v := range list {
		results[i] = newBatchResultErrorEntry(v)
	}
	return results
}

func newBatchResultErrorEntry(o SDK.BatchResultErrorEntry) BatchResultErrorEntry {
	result := BatchResultErrorEntry{}

	if o.Code != nil {
		result.Code = *o.Code
	}
	if o.Id != nil {
		result.ID = *o.Id
	}
	if o.SenderFault != nil {
		result.SenderFault = *o.SenderFault
	}

	if o.Message != nil {
		result.Message = *o.Message
	}
	return result
}

type ChangeMessageVisibilityBatchRequestEntry struct {
	ID            string
	ReceiptHandle string

	// optional
	VisibilityTimeout int64
}

func (r ChangeMessageVisibilityBatchRequestEntry) ToSDK() SDK.ChangeMessageVisibilityBatchRequestEntry {
	o := SDK.ChangeMessageVisibilityBatchRequestEntry{}

	if r.ID != "" {
		o.Id = pointers.String(r.ID)
	}
	if r.ReceiptHandle != "" {
		o.ReceiptHandle = pointers.String(r.ReceiptHandle)
	}

	o.VisibilityTimeout = pointers.Long64(r.VisibilityTimeout)
	return o
}

type ChangeMessageVisibilityBatchResultEntry struct {
	ID string
}

func newChangeMessageVisibilityBatchResultEntryList(list []SDK.ChangeMessageVisibilityBatchResultEntry) []ChangeMessageVisibilityBatchResultEntry {
	if len(list) == 0 {
		return nil
	}

	results := make([]ChangeMessageVisibilityBatchResultEntry, len(list))
	for i, v := range list {
		results[i] = newChangeMessageVisibilityBatchResultEntry(v)
	}
	return results
}

func newChangeMessageVisibilityBatchResultEntry(o SDK.ChangeMessageVisibilityBatchResultEntry) ChangeMessageVisibilityBatchResultEntry {
	result := ChangeMessageVisibilityBatchResultEntry{}

	if o.Id != nil {
		result.ID = *o.Id
	}
	return result
}

type DeleteMessageBatchRequestEntry struct {
	ID            string
	ReceiptHandle string
}

func (r DeleteMessageBatchRequestEntry) ToSDK() SDK.DeleteMessageBatchRequestEntry {
	o := SDK.DeleteMessageBatchRequestEntry{}

	if r.ID != "" {
		o.Id = pointers.String(r.ID)
	}
	if r.ReceiptHandle != "" {
		o.ReceiptHandle = pointers.String(r.ReceiptHandle)
	}
	return o
}

type DeleteMessageBatchResultEntry struct {
	ID string
}

func newDeleteMessageBatchResultEntryList(list []SDK.DeleteMessageBatchResultEntry) []DeleteMessageBatchResultEntry {
	if len(list) == 0 {
		return nil
	}

	results := make([]DeleteMessageBatchResultEntry, len(list))
	for i, v := range list {
		results[i] = newDeleteMessageBatchResultEntry(v)
	}
	return results
}

func newDeleteMessageBatchResultEntry(o SDK.DeleteMessageBatchResultEntry) DeleteMessageBatchResultEntry {
	result := DeleteMessageBatchResultEntry{}

	if o.Id != nil {
		result.ID = *o.Id
	}
	return result
}

type Message struct {
	Attributes             map[string]string
	Body                   string
	MD5OfBody              string
	MD5OfMessageAttributes string
	MessageAttributes      map[string]MessageAttributeValue
	MessageID              string
	ReceiptHandle          string
}

func newMessage(o SDK.Message) Message {
	result := Message{}

	result.Attributes = o.Attributes
	if o.Body != nil {
		result.Body = *o.Body
	}
	if o.MD5OfBody != nil {
		result.MD5OfBody = *o.MD5OfBody
	}
	if o.MD5OfMessageAttributes != nil {
		result.MD5OfMessageAttributes = *o.MD5OfMessageAttributes
	}
	result.MessageAttributes = newMessageAttributeValueMap(o.MessageAttributes)
	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}
	if o.ReceiptHandle != nil {
		result.ReceiptHandle = *o.ReceiptHandle
	}
	return result
}

type MessageAttributeValue struct {
	DataType string

	// optional
	BinaryValue []byte
	StringValue string

	// Not implemented. Reserved for future use.
	BinaryListValues [][]byte
	StringListValues []string
}

func newMessageAttributeValueMap(o map[string]SDK.MessageAttributeValue) map[string]MessageAttributeValue {
	if len(o) == 0 {
		return nil
	}

	m := make(map[string]MessageAttributeValue, len(o))
	for key, val := range o {
		m[key] = newMessageAttributeValue(val)
	}
	return m
}

func newMessageAttributeValue(o SDK.MessageAttributeValue) MessageAttributeValue {
	result := MessageAttributeValue{}

	if o.DataType != nil {
		result.DataType = *o.DataType
	}
	result.BinaryValue = o.BinaryValue
	if o.StringValue != nil {
		result.StringValue = *o.StringValue
	}

	result.BinaryListValues = o.BinaryListValues
	result.StringListValues = o.StringListValues
	return result
}

func (r MessageAttributeValue) ToSDK() SDK.MessageAttributeValue {
	o := SDK.MessageAttributeValue{}

	if r.DataType != "" {
		o.DataType = pointers.String(r.DataType)
	}
	if len(r.BinaryValue) != 0 {
		o.BinaryValue = r.BinaryValue
	}
	if r.StringValue != "" {
		o.StringValue = pointers.String(r.StringValue)
	}

	if len(r.BinaryListValues) != 0 {
		o.BinaryListValues = r.BinaryListValues
	}
	if len(r.StringListValues) != 0 {
		o.StringListValues = r.StringListValues
	}
	return o
}

type MessageSystemAttributeValue struct {
	DataType string

	// optional
	BinaryValue []byte
	StringValue string

	// Not implemented. Reserved for future use.
	BinaryListValues [][]byte
	StringListValues []string
}

func newMessageSystemAttributeValueMap(o map[string]SDK.MessageSystemAttributeValue) map[string]MessageSystemAttributeValue {
	if len(o) == 0 {
		return nil
	}

	m := make(map[string]MessageSystemAttributeValue, len(o))
	for key, val := range o {
		m[key] = newMessageSystemAttributeValue(val)
	}
	return m
}

func newMessageSystemAttributeValue(o SDK.MessageSystemAttributeValue) MessageSystemAttributeValue {
	result := MessageSystemAttributeValue{}

	if o.DataType != nil {
		result.DataType = *o.DataType
	}
	result.BinaryValue = o.BinaryValue
	if o.StringValue != nil {
		result.StringValue = *o.StringValue
	}

	result.BinaryListValues = o.BinaryListValues
	result.StringListValues = o.StringListValues
	return result
}

func (r MessageSystemAttributeValue) ToSDK() SDK.MessageSystemAttributeValue {
	o := SDK.MessageSystemAttributeValue{}

	if r.DataType != "" {
		o.DataType = pointers.String(r.DataType)
	}
	if len(r.BinaryValue) != 0 {
		o.BinaryValue = r.BinaryValue
	}
	if r.StringValue != "" {
		o.StringValue = pointers.String(r.StringValue)
	}

	if len(r.BinaryListValues) != 0 {
		o.BinaryListValues = r.BinaryListValues
	}
	if len(r.StringListValues) != 0 {
		o.StringListValues = r.StringListValues
	}
	return o
}

type SendMessageBatchResultEntry struct {
	ID               string
	MD5OfMessageBody string
	MessageID        string

	// for FIFO
	SequenceNumber string

	MD5OfMessageAttributes       string
	MD5OfMessageSystemAttributes string
}

func newSendMessageBatchResultEntryList(list []SDK.SendMessageBatchResultEntry) []SendMessageBatchResultEntry {
	if len(list) == 0 {
		return nil
	}

	results := make([]SendMessageBatchResultEntry, len(list))
	for i, v := range list {
		results[i] = newSendMessageBatchResultEntry(v)
	}
	return results
}

func newSendMessageBatchResultEntry(o SDK.SendMessageBatchResultEntry) SendMessageBatchResultEntry {
	result := SendMessageBatchResultEntry{}

	if o.Id != nil {
		result.ID = *o.Id
	}
	if o.MD5OfMessageBody != nil {
		result.MD5OfMessageBody = *o.MD5OfMessageBody
	}
	if o.MessageId != nil {
		result.MessageID = *o.MessageId
	}

	if o.SequenceNumber != nil {
		result.SequenceNumber = *o.SequenceNumber
	}

	if o.MD5OfMessageAttributes != nil {
		result.MD5OfMessageAttributes = *o.MD5OfMessageAttributes
	}
	if o.MD5OfMessageSystemAttributes != nil {
		result.MD5OfMessageSystemAttributes = *o.MD5OfMessageSystemAttributes
	}
	return result
}

type SendMessageBatchRequestEntry struct {
	ID          string
	MessageBody string

	// for FIFO
	MessageDeduplicationID string
	MessageGroupID         string

	DelaySeconds            int64
	MessageAttributes       map[string]MessageAttributeValue
	MessageSystemAttributes map[string]MessageSystemAttributeValue
}

func (r SendMessageBatchRequestEntry) ToSDK() SDK.SendMessageBatchRequestEntry {
	o := SDK.SendMessageBatchRequestEntry{}

	if r.ID != "" {
		o.Id = pointers.String(r.ID)
	}
	if r.MessageBody != "" {
		o.MessageBody = pointers.String(r.MessageBody)
	}

	if r.MessageDeduplicationID != "" {
		o.MessageDeduplicationId = pointers.String(r.MessageDeduplicationID)
	}
	if r.MessageGroupID != "" {
		o.MessageGroupId = pointers.String(r.MessageGroupID)
	}

	if r.DelaySeconds != 0 {
		o.DelaySeconds = pointers.Long64(r.DelaySeconds)
	}
	r.MessageAttributes = newMessageAttributeValueMap(o.MessageAttributes)
	r.MessageSystemAttributes = newMessageSystemAttributeValueMap(o.MessageSystemAttributes)
	return o
}
