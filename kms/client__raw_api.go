package kms

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/kms"
)

// RawCancelKeyDeletion executes `CancelKeyDeletion` raw operation.
func (svc *KMS) RawCancelKeyDeletion(ctx context.Context, in *SDK.CancelKeyDeletionInput) (*SDK.CancelKeyDeletionResponse, error) {
	return svc.client.CancelKeyDeletionRequest(in).Send(ctx)
}

// RawConnectCustomKeyStore executes `ConnectCustomKeyStore` raw operation.
func (svc *KMS) RawConnectCustomKeyStore(ctx context.Context, in *SDK.ConnectCustomKeyStoreInput) (*SDK.ConnectCustomKeyStoreResponse, error) {
	return svc.client.ConnectCustomKeyStoreRequest(in).Send(ctx)
}

// RawCreateAlias executes `CreateAlias` raw operation.
func (svc *KMS) RawCreateAlias(ctx context.Context, in *SDK.CreateAliasInput) (*SDK.CreateAliasResponse, error) {
	return svc.client.CreateAliasRequest(in).Send(ctx)
}

// RawCreateCustomKeyStore executes `CreateCustomKeyStore` raw operation.
func (svc *KMS) RawCreateCustomKeyStore(ctx context.Context, in *SDK.CreateCustomKeyStoreInput) (*SDK.CreateCustomKeyStoreResponse, error) {
	return svc.client.CreateCustomKeyStoreRequest(in).Send(ctx)
}

// RawCreateGrant executes `CreateGrant` raw operation.
func (svc *KMS) RawCreateGrant(ctx context.Context, in *SDK.CreateGrantInput) (*SDK.CreateGrantResponse, error) {
	return svc.client.CreateGrantRequest(in).Send(ctx)
}

// RawCreateKey executes `CreateKey` raw operation.
func (svc *KMS) RawCreateKey(ctx context.Context, in *SDK.CreateKeyInput) (*SDK.CreateKeyResponse, error) {
	return svc.client.CreateKeyRequest(in).Send(ctx)
}

// RawDecrypt executes `Decrypt` raw operation.
func (svc *KMS) RawDecrypt(ctx context.Context, in *SDK.DecryptInput) (*SDK.DecryptResponse, error) {
	return svc.client.DecryptRequest(in).Send(ctx)
}

// RawDeleteAlias executes `DeleteAlias` raw operation.
func (svc *KMS) RawDeleteAlias(ctx context.Context, in *SDK.DeleteAliasInput) (*SDK.DeleteAliasResponse, error) {
	return svc.client.DeleteAliasRequest(in).Send(ctx)
}

// RawDeleteCustomKeyStore executes `DeleteCustomKeyStore` raw operation.
func (svc *KMS) RawDeleteCustomKeyStore(ctx context.Context, in *SDK.DeleteCustomKeyStoreInput) (*SDK.DeleteCustomKeyStoreResponse, error) {
	return svc.client.DeleteCustomKeyStoreRequest(in).Send(ctx)
}

// RawDeleteImportedKeyMaterial executes `DeleteImportedKeyMaterial` raw operation.
func (svc *KMS) RawDeleteImportedKeyMaterial(ctx context.Context, in *SDK.DeleteImportedKeyMaterialInput) (*SDK.DeleteImportedKeyMaterialResponse, error) {
	return svc.client.DeleteImportedKeyMaterialRequest(in).Send(ctx)
}

// RawDescribeCustomKeyStores executes `DescribeCustomKeyStores` raw operation.
func (svc *KMS) RawDescribeCustomKeyStores(ctx context.Context, in *SDK.DescribeCustomKeyStoresInput) (*SDK.DescribeCustomKeyStoresResponse, error) {
	return svc.client.DescribeCustomKeyStoresRequest(in).Send(ctx)
}

// RawDescribeKey executes `DescribeKey` raw operation.
func (svc *KMS) RawDescribeKey(ctx context.Context, in *SDK.DescribeKeyInput) (*SDK.DescribeKeyResponse, error) {
	return svc.client.DescribeKeyRequest(in).Send(ctx)
}

// RawDisableKey executes `DisableKey` raw operation.
func (svc *KMS) RawDisableKey(ctx context.Context, in *SDK.DisableKeyInput) (*SDK.DisableKeyResponse, error) {
	return svc.client.DisableKeyRequest(in).Send(ctx)
}

// RawDisableKeyRotation executes `DisableKeyRotation` raw operation.
func (svc *KMS) RawDisableKeyRotation(ctx context.Context, in *SDK.DisableKeyRotationInput) (*SDK.DisableKeyRotationResponse, error) {
	return svc.client.DisableKeyRotationRequest(in).Send(ctx)
}

// RawDisconnectCustomKeyStore executes `DisconnectCustomKeyStore` raw operation.
func (svc *KMS) RawDisconnectCustomKeyStore(ctx context.Context, in *SDK.DisconnectCustomKeyStoreInput) (*SDK.DisconnectCustomKeyStoreResponse, error) {
	return svc.client.DisconnectCustomKeyStoreRequest(in).Send(ctx)
}

// RawEnableKey executes `EnableKey` raw operation.
func (svc *KMS) RawEnableKey(ctx context.Context, in *SDK.EnableKeyInput) (*SDK.EnableKeyResponse, error) {
	return svc.client.EnableKeyRequest(in).Send(ctx)
}

// RawEnableKeyRotation executes `EnableKeyRotation` raw operation.
func (svc *KMS) RawEnableKeyRotation(ctx context.Context, in *SDK.EnableKeyRotationInput) (*SDK.EnableKeyRotationResponse, error) {
	return svc.client.EnableKeyRotationRequest(in).Send(ctx)
}

// RawEncrypt executes `Encrypt` raw operation.
func (svc *KMS) RawEncrypt(ctx context.Context, in *SDK.EncryptInput) (*SDK.EncryptResponse, error) {
	return svc.client.EncryptRequest(in).Send(ctx)
}

// RawGenerateDataKey executes `GenerateDataKey` raw operation.
func (svc *KMS) RawGenerateDataKey(ctx context.Context, in *SDK.GenerateDataKeyInput) (*SDK.GenerateDataKeyResponse, error) {
	return svc.client.GenerateDataKeyRequest(in).Send(ctx)
}

// RawGenerateDataKeyPair executes `GenerateDataKeyPair` raw operation.
func (svc *KMS) RawGenerateDataKeyPair(ctx context.Context, in *SDK.GenerateDataKeyPairInput) (*SDK.GenerateDataKeyPairResponse, error) {
	return svc.client.GenerateDataKeyPairRequest(in).Send(ctx)
}

// RawGenerateDataKeyPairWithoutPlaintext executes `GenerateDataKeyPairWithoutPlaintext` raw operation.
func (svc *KMS) RawGenerateDataKeyPairWithoutPlaintext(ctx context.Context, in *SDK.GenerateDataKeyPairWithoutPlaintextInput) (*SDK.GenerateDataKeyPairWithoutPlaintextResponse, error) {
	return svc.client.GenerateDataKeyPairWithoutPlaintextRequest(in).Send(ctx)
}

// RawGenerateDataKeyWithoutPlaintext executes `GenerateDataKeyWithoutPlaintext` raw operation.
func (svc *KMS) RawGenerateDataKeyWithoutPlaintext(ctx context.Context, in *SDK.GenerateDataKeyWithoutPlaintextInput) (*SDK.GenerateDataKeyWithoutPlaintextResponse, error) {
	return svc.client.GenerateDataKeyWithoutPlaintextRequest(in).Send(ctx)
}

// RawGenerateRandom executes `GenerateRandom` raw operation.
func (svc *KMS) RawGenerateRandom(ctx context.Context, in *SDK.GenerateRandomInput) (*SDK.GenerateRandomResponse, error) {
	return svc.client.GenerateRandomRequest(in).Send(ctx)
}

// RawGetKeyPolicy executes `GetKeyPolicy` raw operation.
func (svc *KMS) RawGetKeyPolicy(ctx context.Context, in *SDK.GetKeyPolicyInput) (*SDK.GetKeyPolicyResponse, error) {
	return svc.client.GetKeyPolicyRequest(in).Send(ctx)
}

// RawGetKeyRotationStatus executes `GetKeyRotationStatus` raw operation.
func (svc *KMS) RawGetKeyRotationStatus(ctx context.Context, in *SDK.GetKeyRotationStatusInput) (*SDK.GetKeyRotationStatusResponse, error) {
	return svc.client.GetKeyRotationStatusRequest(in).Send(ctx)
}

// RawGetParametersForImport executes `GetParametersForImport` raw operation.
func (svc *KMS) RawGetParametersForImport(ctx context.Context, in *SDK.GetParametersForImportInput) (*SDK.GetParametersForImportResponse, error) {
	return svc.client.GetParametersForImportRequest(in).Send(ctx)
}

// RawGetPublicKey executes `GetPublicKey` raw operation.
func (svc *KMS) RawGetPublicKey(ctx context.Context, in *SDK.GetPublicKeyInput) (*SDK.GetPublicKeyResponse, error) {
	return svc.client.GetPublicKeyRequest(in).Send(ctx)
}

// RawImportKeyMaterial executes `ImportKeyMaterial` raw operation.
func (svc *KMS) RawImportKeyMaterial(ctx context.Context, in *SDK.ImportKeyMaterialInput) (*SDK.ImportKeyMaterialResponse, error) {
	return svc.client.ImportKeyMaterialRequest(in).Send(ctx)
}

// RawListAliases executes `ListAliases` raw operation.
func (svc *KMS) RawListAliases(ctx context.Context, in *SDK.ListAliasesInput) (*SDK.ListAliasesResponse, error) {
	return svc.client.ListAliasesRequest(in).Send(ctx)
}

// RawListGrants executes `ListGrants` raw operation.
func (svc *KMS) RawListGrants(ctx context.Context, in *SDK.ListGrantsInput) (*SDK.ListGrantsResponse, error) {
	return svc.client.ListGrantsRequest(in).Send(ctx)
}

// RawListKeyPolicies executes `ListKeyPolicies` raw operation.
func (svc *KMS) RawListKeyPolicies(ctx context.Context, in *SDK.ListKeyPoliciesInput) (*SDK.ListKeyPoliciesResponse, error) {
	return svc.client.ListKeyPoliciesRequest(in).Send(ctx)
}

// RawListKeys executes `ListKeys` raw operation.
func (svc *KMS) RawListKeys(ctx context.Context, in *SDK.ListKeysInput) (*SDK.ListKeysResponse, error) {
	return svc.client.ListKeysRequest(in).Send(ctx)
}

// RawListResourceTags executes `ListResourceTags` raw operation.
func (svc *KMS) RawListResourceTags(ctx context.Context, in *SDK.ListResourceTagsInput) (*SDK.ListResourceTagsResponse, error) {
	return svc.client.ListResourceTagsRequest(in).Send(ctx)
}

// RawListRetirableGrants executes `ListRetirableGrants` raw operation.
func (svc *KMS) RawListRetirableGrants(ctx context.Context, in *SDK.ListRetirableGrantsInput) (*SDK.ListRetirableGrantsResponse, error) {
	return svc.client.ListRetirableGrantsRequest(in).Send(ctx)
}

// RawPutKeyPolicy executes `PutKeyPolicy` raw operation.
func (svc *KMS) RawPutKeyPolicy(ctx context.Context, in *SDK.PutKeyPolicyInput) (*SDK.PutKeyPolicyResponse, error) {
	return svc.client.PutKeyPolicyRequest(in).Send(ctx)
}

// RawReEncrypt executes `ReEncrypt` raw operation.
func (svc *KMS) RawReEncrypt(ctx context.Context, in *SDK.ReEncryptInput) (*SDK.ReEncryptResponse, error) {
	return svc.client.ReEncryptRequest(in).Send(ctx)
}

// RawRetireGrant executes `RetireGrant` raw operation.
func (svc *KMS) RawRetireGrant(ctx context.Context, in *SDK.RetireGrantInput) (*SDK.RetireGrantResponse, error) {
	return svc.client.RetireGrantRequest(in).Send(ctx)
}

// RawRevokeGrant executes `RevokeGrant` raw operation.
func (svc *KMS) RawRevokeGrant(ctx context.Context, in *SDK.RevokeGrantInput) (*SDK.RevokeGrantResponse, error) {
	return svc.client.RevokeGrantRequest(in).Send(ctx)
}

// RawScheduleKeyDeletion executes `ScheduleKeyDeletion` raw operation.
func (svc *KMS) RawScheduleKeyDeletion(ctx context.Context, in *SDK.ScheduleKeyDeletionInput) (*SDK.ScheduleKeyDeletionResponse, error) {
	return svc.client.ScheduleKeyDeletionRequest(in).Send(ctx)
}

// RawSign executes `Sign` raw operation.
func (svc *KMS) RawSign(ctx context.Context, in *SDK.SignInput) (*SDK.SignResponse, error) {
	return svc.client.SignRequest(in).Send(ctx)
}

// RawTagResource executes `TagResource` raw operation.
func (svc *KMS) RawTagResource(ctx context.Context, in *SDK.TagResourceInput) (*SDK.TagResourceResponse, error) {
	return svc.client.TagResourceRequest(in).Send(ctx)
}

// RawUntagResource executes `UntagResource` raw operation.
func (svc *KMS) RawUntagResource(ctx context.Context, in *SDK.UntagResourceInput) (*SDK.UntagResourceResponse, error) {
	return svc.client.UntagResourceRequest(in).Send(ctx)
}

// RawUpdateAlias executes `UpdateAlias` raw operation.
func (svc *KMS) RawUpdateAlias(ctx context.Context, in *SDK.UpdateAliasInput) (*SDK.UpdateAliasResponse, error) {
	return svc.client.UpdateAliasRequest(in).Send(ctx)
}

// RawUpdateCustomKeyStore executes `UpdateCustomKeyStore` raw operation.
func (svc *KMS) RawUpdateCustomKeyStore(ctx context.Context, in *SDK.UpdateCustomKeyStoreInput) (*SDK.UpdateCustomKeyStoreResponse, error) {
	return svc.client.UpdateCustomKeyStoreRequest(in).Send(ctx)
}

// RawUpdateKeyDescription executes `UpdateKeyDescription` raw operation.
func (svc *KMS) RawUpdateKeyDescription(ctx context.Context, in *SDK.UpdateKeyDescriptionInput) (*SDK.UpdateKeyDescriptionResponse, error) {
	return svc.client.UpdateKeyDescriptionRequest(in).Send(ctx)
}

// RawVerify executes `Verify` raw operation.
func (svc *KMS) RawVerify(ctx context.Context, in *SDK.VerifyInput) (*SDK.VerifyResponse, error) {
	return svc.client.VerifyRequest(in).Send(ctx)
}
