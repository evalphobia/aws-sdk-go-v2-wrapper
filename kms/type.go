package kms

import (
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type KeyMetadata struct {
	KeyID string

	// optional
	AWSAccountID          string
	ARN                   string
	CloudHsmClusterID     string
	CreationDate          time.Time
	CustomKeyStoreID      string
	CustomerMasterKeySpec CustomerMasterKeySpec
	DeletionDate          time.Time
	Description           string
	Enabled               bool
	EncryptionAlgorithms  []EncryptionAlgorithmSpec
	ExpirationModel       ExpirationModelType
	KeyManager            KeyManagerType
	KeyState              KeyState
	KeyUsage              KeyUsageType
	Origin                OriginType
	SigningAlgorithms     []SigningAlgorithmSpec
	ValidTo               time.Time
}

func newKeyMetadata(o *SDK.KeyMetadata) KeyMetadata {
	result := KeyMetadata{}
	if o == nil {
		return result
	}

	if o.KeyId != nil {
		result.KeyID = *o.KeyId
	}

	if o.AWSAccountId != nil {
		result.AWSAccountID = *o.AWSAccountId
	}
	if o.Arn != nil {
		result.ARN = *o.Arn
	}
	if o.CloudHsmClusterId != nil {
		result.CloudHsmClusterID = *o.CloudHsmClusterId
	}
	if o.CreationDate != nil {
		result.CreationDate = *o.CreationDate
	}
	if o.CustomKeyStoreId != nil {
		result.CustomKeyStoreID = *o.CustomKeyStoreId
	}

	o.CustomerMasterKeySpec = SDK.CustomerMasterKeySpec(result.CustomerMasterKeySpec)

	if o.DeletionDate != nil {
		result.DeletionDate = *o.DeletionDate
	}
	if o.Description != nil {
		result.Description = *o.Description
	}
	if o.Enabled != nil {
		result.Enabled = *o.Enabled
	}

	if len(o.EncryptionAlgorithms) != 0 {
		list := make([]EncryptionAlgorithmSpec, len(o.EncryptionAlgorithms))
		for i, v := range o.EncryptionAlgorithms {
			list[i] = EncryptionAlgorithmSpec(v)
		}
		result.EncryptionAlgorithms = list
	}

	o.ExpirationModel = SDK.ExpirationModelType(result.ExpirationModel)
	o.KeyManager = SDK.KeyManagerType(result.KeyManager)
	o.KeyState = SDK.KeyState(result.KeyState)
	o.KeyUsage = SDK.KeyUsageType(result.KeyUsage)
	o.Origin = SDK.OriginType(result.Origin)

	if len(o.SigningAlgorithms) != 0 {
		list := make([]SigningAlgorithmSpec, len(o.SigningAlgorithms))
		for i, v := range o.SigningAlgorithms {
			list[i] = SigningAlgorithmSpec(v)
		}
		result.SigningAlgorithms = list
	}

	if o.ValidTo != nil {
		result.ValidTo = *o.ValidTo
	}
	return result
}

type Tag struct {
	Key   string
	Value string
}

func (r Tag) ToSDK() SDK.Tag {
	return SDK.Tag{
		TagKey:   pointers.String(r.Key),
		TagValue: pointers.String(r.Value),
	}
}
