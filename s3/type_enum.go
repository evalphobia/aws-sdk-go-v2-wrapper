package s3

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketCannedACL string

const (
	BucketCannedACLPrivate           BucketCannedACL = BucketCannedACL(SDK.BucketCannedACLPrivate)
	BucketCannedACLPublicRead        BucketCannedACL = BucketCannedACL(SDK.BucketCannedACLPublicRead)
	BucketCannedACLPublicReadWrite   BucketCannedACL = BucketCannedACL(SDK.BucketCannedACLPublicReadWrite)
	BucketCannedACLAuthenticatedRead BucketCannedACL = BucketCannedACL(SDK.BucketCannedACLAuthenticatedRead)
)

func (v BucketCannedACL) IsPrivate() bool {
	return v == BucketCannedACLPrivate
}
func (v BucketCannedACL) IsPublicRead() bool {
	return v == BucketCannedACLPublicRead
}
func (v BucketCannedACL) IsPublicReadWrite() bool {
	return v == BucketCannedACLPublicReadWrite
}
func (v BucketCannedACL) IsAuthenticatedRead() bool {
	return v == BucketCannedACLAuthenticatedRead
}

type BucketLocationConstraint string

const (
	BucketLocationConstraintEu           BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintEu)
	BucketLocationConstraintEuWest1      BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintEuWest1)
	BucketLocationConstraintUsWest1      BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintUsWest1)
	BucketLocationConstraintUsWest2      BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintUsWest2)
	BucketLocationConstraintApSouth1     BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintApSouth1)
	BucketLocationConstraintApSoutheast1 BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintApSoutheast1)
	BucketLocationConstraintApSoutheast2 BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintApSoutheast2)
	BucketLocationConstraintApNortheast1 BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintApNortheast1)
	BucketLocationConstraintSaEast1      BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintSaEast1)
	BucketLocationConstraintCnNorth1     BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintCnNorth1)
	BucketLocationConstraintEuCentral1   BucketLocationConstraint = BucketLocationConstraint(SDK.BucketLocationConstraintEuCentral1)
)

func (v BucketLocationConstraint) IsEu() bool {
	return v == BucketLocationConstraintEu
}
func (v BucketLocationConstraint) IsEuWest1() bool {
	return v == BucketLocationConstraintEuWest1
}
func (v BucketLocationConstraint) IsUsWest1() bool {
	return v == BucketLocationConstraintUsWest1
}
func (v BucketLocationConstraint) IsUsWest2() bool {
	return v == BucketLocationConstraintUsWest2
}
func (v BucketLocationConstraint) IsApSouth1() bool {
	return v == BucketLocationConstraintApSouth1
}
func (v BucketLocationConstraint) IsApSoutheast1() bool {
	return v == BucketLocationConstraintApSoutheast1
}
func (v BucketLocationConstraint) IsApSoutheast2() bool {
	return v == BucketLocationConstraintApSoutheast2
}
func (v BucketLocationConstraint) IsApNortheast1() bool {
	return v == BucketLocationConstraintApNortheast1
}
func (v BucketLocationConstraint) IsSaEast1() bool {
	return v == BucketLocationConstraintSaEast1
}
func (v BucketLocationConstraint) IsCnNorth1() bool {
	return v == BucketLocationConstraintCnNorth1
}
func (v BucketLocationConstraint) IsEuCentral1() bool {
	return v == BucketLocationConstraintEuCentral1
}

type EncodingType string

const (
	EncodingTypeURL EncodingType = EncodingType(SDK.EncodingTypeUrl)
)

func (v EncodingType) IsURL() bool {
	return v == EncodingTypeURL
}

type ExpirationStatus string

const (
	ExpirationStatusEnabled  ExpirationStatus = ExpirationStatus(SDK.ExpirationStatusEnabled)
	ExpirationStatusDisabled ExpirationStatus = ExpirationStatus(SDK.ExpirationStatusDisabled)
)

func (v ExpirationStatus) IsEnabled() bool {
	return v == ExpirationStatusEnabled
}
func (v ExpirationStatus) IsDisabled() bool {
	return v == ExpirationStatusDisabled
}

type MetadataDirective string

const (
	MetadataDirectiveCopy    MetadataDirective = MetadataDirective(SDK.MetadataDirectiveCopy)
	MetadataDirectiveReplace MetadataDirective = MetadataDirective(SDK.MetadataDirectiveReplace)
)

func (v MetadataDirective) IsCopy() bool {
	return v == MetadataDirectiveCopy
}
func (v MetadataDirective) IsReplace() bool {
	return v == MetadataDirectiveReplace
}

type ObjectCannedACL string

const (
	ObjectCannedACLPrivate                ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLPrivate)
	ObjectCannedACLPublicRead             ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLPublicRead)
	ObjectCannedACLPublicReadWrite        ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLPublicReadWrite)
	ObjectCannedACLAuthenticatedRead      ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLAuthenticatedRead)
	ObjectCannedACLAwsExecRead            ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLAwsExecRead)
	ObjectCannedACLBucketOwnerRead        ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLBucketOwnerRead)
	ObjectCannedACLBucketOwnerFullControl ObjectCannedACL = ObjectCannedACL(SDK.ObjectCannedACLBucketOwnerFullControl)
)

func (v ObjectCannedACL) IsPrivate() bool {
	return v == ObjectCannedACLPrivate
}
func (v ObjectCannedACL) IsPublicRead() bool {
	return v == ObjectCannedACLPublicRead
}
func (v ObjectCannedACL) IsPublicReadWrite() bool {
	return v == ObjectCannedACLPublicReadWrite
}
func (v ObjectCannedACL) IsAuthenticatedRead() bool {
	return v == ObjectCannedACLAuthenticatedRead
}
func (v ObjectCannedACL) IsAwsExecRead() bool {
	return v == ObjectCannedACLAwsExecRead
}
func (v ObjectCannedACL) IsBucketOwnerRead() bool {
	return v == ObjectCannedACLBucketOwnerRead
}
func (v ObjectCannedACL) IsBucketOwnerFullControl() bool {
	return v == ObjectCannedACLBucketOwnerFullControl
}

type ObjectLockLegalHoldStatus string

const (
	ObjectLockLegalHoldStatusOn  ObjectLockLegalHoldStatus = ObjectLockLegalHoldStatus(SDK.ObjectLockLegalHoldStatusOn)
	ObjectLockLegalHoldStatusOff ObjectLockLegalHoldStatus = ObjectLockLegalHoldStatus(SDK.ObjectLockLegalHoldStatusOff)
)

func (v ObjectLockLegalHoldStatus) IsOn() bool {
	return v == ObjectLockLegalHoldStatusOn
}
func (v ObjectLockLegalHoldStatus) IsOff() bool {
	return v == ObjectLockLegalHoldStatusOff
}

type ObjectLockMode string

const (
	ObjectLockModeGovernance ObjectLockMode = ObjectLockMode(SDK.ObjectLockModeGovernance)
	ObjectLockModeCompliance ObjectLockMode = ObjectLockMode(SDK.ObjectLockModeCompliance)
)

func (v ObjectLockMode) IsGovernance() bool {
	return v == ObjectLockModeGovernance
}
func (v ObjectLockMode) IsCompliance() bool {
	return v == ObjectLockModeCompliance
}

type ObjectLockRetentionMode string

const (
	ObjectLockRetentionModeGovernance ObjectLockRetentionMode = ObjectLockRetentionMode(SDK.ObjectLockRetentionModeGovernance)
	ObjectLockRetentionModeCompliance ObjectLockRetentionMode = ObjectLockRetentionMode(SDK.ObjectLockRetentionModeCompliance)
)

func (v ObjectLockRetentionMode) IsGovernance() bool {
	return v == ObjectLockRetentionModeGovernance
}
func (v ObjectLockRetentionMode) IsCompliance() bool {
	return v == ObjectLockRetentionModeCompliance
}

type ObjectVersionStorageClass string

const (
	ObjectVersionStorageClassStandard ObjectVersionStorageClass = ObjectVersionStorageClass(SDK.ObjectVersionStorageClassStandard)
)

func (v ObjectVersionStorageClass) IsStandard() bool {
	return v == ObjectVersionStorageClassStandard
}

type Permission string

const (
	PermissionFullControl Permission = Permission(SDK.PermissionFullControl)
	PermissionWrite       Permission = Permission(SDK.PermissionWrite)
	PermissionWriteAcp    Permission = Permission(SDK.PermissionWriteAcp)
	PermissionRead        Permission = Permission(SDK.PermissionRead)
	PermissionReadAcp     Permission = Permission(SDK.PermissionReadAcp)
)

type ReplicationStatus string

const (
	ReplicationStatusComplete ReplicationStatus = ReplicationStatus(SDK.ReplicationStatusComplete)
	ReplicationStatusPending  ReplicationStatus = ReplicationStatus(SDK.ReplicationStatusPending)
	ReplicationStatusFailed   ReplicationStatus = ReplicationStatus(SDK.ReplicationStatusFailed)
	ReplicationStatusReplica  ReplicationStatus = ReplicationStatus(SDK.ReplicationStatusReplica)
)

func (v ReplicationStatus) IsComplete() bool {
	return v == ReplicationStatusComplete
}
func (v ReplicationStatus) IsPending() bool {
	return v == ReplicationStatusPending
}
func (v ReplicationStatus) IsFailed() bool {
	return v == ReplicationStatusFailed
}
func (v ReplicationStatus) IsReplica() bool {
	return v == ReplicationStatusReplica
}

type RequestCharged string

const (
	RequestChargedRequester RequestCharged = RequestCharged(SDK.RequestChargedRequester)
)

func (v RequestCharged) IsRequester() bool {
	return v == RequestChargedRequester
}

type RequestPayer string

const (
	RequestPayerRequester RequestPayer = "requester"
)

func (v RequestPayer) IsRequester() bool {
	return v == RequestPayerRequester
}

type ServerSideEncryption string

const (
	ServerSideEncryptionAes256 ServerSideEncryption = "AES256"
	ServerSideEncryptionAwsKms ServerSideEncryption = "aws:kms"
)

func (v ServerSideEncryption) IsAes256() bool {
	return v == ServerSideEncryptionAes256
}
func (v ServerSideEncryption) IsAwsKms() bool {
	return v == ServerSideEncryptionAwsKms
}

type StorageClass string

const (
	StorageClassStandard           StorageClass = StorageClass(SDK.StorageClassStandard)
	StorageClassReducedRedundancy  StorageClass = StorageClass(SDK.StorageClassReducedRedundancy)
	StorageClassStandardIa         StorageClass = StorageClass(SDK.StorageClassStandardIa)
	StorageClassOnezoneIa          StorageClass = StorageClass(SDK.StorageClassOnezoneIa)
	StorageClassIntelligentTiering StorageClass = StorageClass(SDK.StorageClassIntelligentTiering)
	StorageClassGlacier            StorageClass = StorageClass(SDK.StorageClassGlacier)
	StorageClassDeepArchive        StorageClass = StorageClass(SDK.StorageClassDeepArchive)
)

func (v StorageClass) IsStandard() bool {
	return v == StorageClassStandard
}
func (v StorageClass) IsReducedRedundancy() bool {
	return v == StorageClassReducedRedundancy
}
func (v StorageClass) IsStandardIa() bool {
	return v == StorageClassStandardIa
}
func (v StorageClass) IsOnezoneIa() bool {
	return v == StorageClassOnezoneIa
}
func (v StorageClass) IsIntelligentTiering() bool {
	return v == StorageClassIntelligentTiering
}
func (v StorageClass) IsGlacier() bool {
	return v == StorageClassGlacier
}
func (v StorageClass) IsDeepArchive() bool {
	return v == StorageClassDeepArchive
}

type TaggingDirective string

const (
	TaggingDirectiveCopy    TaggingDirective = TaggingDirective(SDK.TaggingDirectiveCopy)
	TaggingDirectiveReplace TaggingDirective = TaggingDirective(SDK.TaggingDirectiveReplace)
)

func (v TaggingDirective) IsCopy() bool {
	return v == TaggingDirectiveCopy
}
func (v TaggingDirective) IsReplace() bool {
	return v == TaggingDirectiveReplace
}

type TransitionStorageClass string

const (
	TransitionStorageClassGlacier            TransitionStorageClass = TransitionStorageClass(SDK.TransitionStorageClassGlacier)
	TransitionStorageClassStandardIa         TransitionStorageClass = TransitionStorageClass(SDK.TransitionStorageClassStandardIa)
	TransitionStorageClassOnezoneIa          TransitionStorageClass = TransitionStorageClass(SDK.TransitionStorageClassOnezoneIa)
	TransitionStorageClassIntelligentTiering TransitionStorageClass = TransitionStorageClass(SDK.TransitionStorageClassIntelligentTiering)
	TransitionStorageClassDeepArchive        TransitionStorageClass = TransitionStorageClass(SDK.TransitionStorageClassDeepArchive)
)

func (v TransitionStorageClass) IsGlacier() bool {
	return v == TransitionStorageClassGlacier
}
func (v TransitionStorageClass) IsStandardIa() bool {
	return v == TransitionStorageClassStandardIa
}
func (v TransitionStorageClass) IsOnezoneIa() bool {
	return v == TransitionStorageClassOnezoneIa
}
func (v TransitionStorageClass) IsIntelligentTiering() bool {
	return v == TransitionStorageClassIntelligentTiering
}
func (v TransitionStorageClass) IsDeepArchive() bool {
	return v == TransitionStorageClassDeepArchive
}

type Type string

const (
	TypeCanonicalUser         Type = Type(SDK.TypeCanonicalUser)
	TypeAmazonCustomerByEmail Type = Type(SDK.TypeAmazonCustomerByEmail)
	TypeGroup                 Type = Type(SDK.TypeGroup)
)

func (v Type) IsCanonicalUser() bool {
	return v == TypeCanonicalUser
}
func (v Type) IsAmazonCustomerByEmail() bool {
	return v == TypeAmazonCustomerByEmail
}
func (v Type) IsGroup() bool {
	return v == TypeGroup
}
