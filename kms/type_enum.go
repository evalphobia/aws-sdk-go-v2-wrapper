package kms

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/kms"
)

type CustomerMasterKeySpec string

const (
	CustomerMasterKeySpecRsa2048          CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecRsa2048)
	CustomerMasterKeySpecRsa3072          CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecRsa3072)
	CustomerMasterKeySpecRsa4096          CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecRsa4096)
	CustomerMasterKeySpecEccNistP256      CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecEccNistP256)
	CustomerMasterKeySpecEccNistP384      CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecEccNistP384)
	CustomerMasterKeySpecEccNistP521      CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecEccNistP521)
	CustomerMasterKeySpecEccSecgP256k1    CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecEccSecgP256k1)
	CustomerMasterKeySpecSymmetricDefault CustomerMasterKeySpec = CustomerMasterKeySpec(SDK.CustomerMasterKeySpecSymmetricDefault)
)

type EncryptionAlgorithmSpec string

const (
	EncryptionAlgorithmSpecSymmetricDefault EncryptionAlgorithmSpec = EncryptionAlgorithmSpec(SDK.EncryptionAlgorithmSpecSymmetricDefault)
	EncryptionAlgorithmSpecRsaesOaepSha1    EncryptionAlgorithmSpec = EncryptionAlgorithmSpec(SDK.EncryptionAlgorithmSpecRsaesOaepSha1)
	EncryptionAlgorithmSpecRsaesOaepSha256  EncryptionAlgorithmSpec = EncryptionAlgorithmSpec(SDK.EncryptionAlgorithmSpecRsaesOaepSha256)
)

type ExpirationModelType string

const (
	ExpirationModelTypeKeyMaterialExpires       ExpirationModelType = ExpirationModelType(SDK.ExpirationModelTypeKeyMaterialExpires)
	ExpirationModelTypeKeyMaterialDoesNotExpire ExpirationModelType = ExpirationModelType(SDK.ExpirationModelTypeKeyMaterialDoesNotExpire)
)

type KeyManagerType string

const (
	KeyManagerTypeAws      KeyManagerType = KeyManagerType(SDK.KeyManagerTypeAws)
	KeyManagerTypeCustomer KeyManagerType = KeyManagerType(SDK.KeyManagerTypeCustomer)
)

type KeyState string

const (
	KeyStateEnabled         KeyState = KeyState(SDK.KeyStateEnabled)
	KeyStateDisabled        KeyState = KeyState(SDK.KeyStateDisabled)
	KeyStatePendingDeletion KeyState = KeyState(SDK.KeyStatePendingDeletion)
	KeyStatePendingImport   KeyState = KeyState(SDK.KeyStatePendingImport)
	KeyStateUnavailable     KeyState = KeyState(SDK.KeyStateUnavailable)
)

type KeyUsageType string

const (
	KeyUsageTypeSignVerify     KeyUsageType = KeyUsageType(SDK.KeyUsageTypeSignVerify)
	KeyUsageTypeEncryptDecrypt KeyUsageType = KeyUsageType(SDK.KeyUsageTypeEncryptDecrypt)
)

type OriginType string

const (
	OriginTypeAwsKms      OriginType = OriginType(SDK.OriginTypeAwsKms)
	OriginTypeExternal    OriginType = OriginType(SDK.OriginTypeExternal)
	OriginTypeAwsCloudhsm OriginType = OriginType(SDK.OriginTypeAwsCloudhsm)
)

type SigningAlgorithmSpec string

const (
	SigningAlgorithmSpecRsassaPssSha256      SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecRsassaPssSha256)
	SigningAlgorithmSpecRsassaPssSha384      SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecRsassaPssSha384)
	SigningAlgorithmSpecRsassaPssSha512      SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecRsassaPssSha512)
	SigningAlgorithmSpecRsassaPkcs1V15Sha256 SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecRsassaPkcs1V15Sha256)
	SigningAlgorithmSpecRsassaPkcs1V15Sha384 SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecRsassaPkcs1V15Sha384)
	SigningAlgorithmSpecRsassaPkcs1V15Sha512 SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecRsassaPkcs1V15Sha512)
	SigningAlgorithmSpecEcdsaSha256          SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecEcdsaSha256)
	SigningAlgorithmSpecEcdsaSha384          SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecEcdsaSha384)
	SigningAlgorithmSpecEcdsaSha512          SigningAlgorithmSpec = SigningAlgorithmSpec(SDK.SigningAlgorithmSpecEcdsaSha512)
)
