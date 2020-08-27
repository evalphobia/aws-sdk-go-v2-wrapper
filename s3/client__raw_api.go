package s3

import (
	"context"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"
)

// RawAbortMultipartUpload executes `AbortMultipartUpload` raw operation.
func (svc *S3) RawAbortMultipartUpload(ctx context.Context, in *SDK.AbortMultipartUploadInput) (*SDK.AbortMultipartUploadResponse, error) {
	return svc.client.AbortMultipartUploadRequest(in).Send(ctx)
}

// RawCompleteMultipartUpload executes `CompleteMultipartUpload` raw operation.
func (svc *S3) RawCompleteMultipartUpload(ctx context.Context, in *SDK.CompleteMultipartUploadInput) (*SDK.CompleteMultipartUploadResponse, error) {
	return svc.client.CompleteMultipartUploadRequest(in).Send(ctx)
}

// RawCopyObject executes `CopyObject` raw operation.
func (svc *S3) RawCopyObject(ctx context.Context, in *SDK.CopyObjectInput) (*SDK.CopyObjectResponse, error) {
	return svc.client.CopyObjectRequest(in).Send(ctx)
}

// RawCreateBucket executes `CreateBucket` raw operation.
func (svc *S3) RawCreateBucket(ctx context.Context, in *SDK.CreateBucketInput) (*SDK.CreateBucketResponse, error) {
	return svc.client.CreateBucketRequest(in).Send(ctx)
}

// RawCreateMultipartUpload executes `CreateMultipartUpload` raw operation.
func (svc *S3) RawCreateMultipartUpload(ctx context.Context, in *SDK.CreateMultipartUploadInput) (*SDK.CreateMultipartUploadResponse, error) {
	return svc.client.CreateMultipartUploadRequest(in).Send(ctx)
}

// RawDeleteBucket executes `DeleteBucket` raw operation.
func (svc *S3) RawDeleteBucket(ctx context.Context, in *SDK.DeleteBucketInput) (*SDK.DeleteBucketResponse, error) {
	return svc.client.DeleteBucketRequest(in).Send(ctx)
}

// RawDeleteBucketAnalyticsConfiguration executes `DeleteBucketAnalyticsConfiguration` raw operation.
func (svc *S3) RawDeleteBucketAnalyticsConfiguration(ctx context.Context, in *SDK.DeleteBucketAnalyticsConfigurationInput) (*SDK.DeleteBucketAnalyticsConfigurationResponse, error) {
	return svc.client.DeleteBucketAnalyticsConfigurationRequest(in).Send(ctx)
}

// RawDeleteBucketCors executes `DeleteBucketCors` raw operation.
func (svc *S3) RawDeleteBucketCors(ctx context.Context, in *SDK.DeleteBucketCorsInput) (*SDK.DeleteBucketCorsResponse, error) {
	return svc.client.DeleteBucketCorsRequest(in).Send(ctx)
}

// RawDeleteBucketEncryption executes `DeleteBucketEncryption` raw operation.
func (svc *S3) RawDeleteBucketEncryption(ctx context.Context, in *SDK.DeleteBucketEncryptionInput) (*SDK.DeleteBucketEncryptionResponse, error) {
	return svc.client.DeleteBucketEncryptionRequest(in).Send(ctx)
}

// RawDeleteBucketInventoryConfiguration executes `DeleteBucketInventoryConfiguration` raw operation.
func (svc *S3) RawDeleteBucketInventoryConfiguration(ctx context.Context, in *SDK.DeleteBucketInventoryConfigurationInput) (*SDK.DeleteBucketInventoryConfigurationResponse, error) {
	return svc.client.DeleteBucketInventoryConfigurationRequest(in).Send(ctx)
}

// RawDeleteBucketLifecycle executes `DeleteBucketLifecycle` raw operation.
func (svc *S3) RawDeleteBucketLifecycle(ctx context.Context, in *SDK.DeleteBucketLifecycleInput) (*SDK.DeleteBucketLifecycleResponse, error) {
	return svc.client.DeleteBucketLifecycleRequest(in).Send(ctx)
}

// RawDeleteBucketMetricsConfiguration executes `DeleteBucketMetricsConfiguration` raw operation.
func (svc *S3) RawDeleteBucketMetricsConfiguration(ctx context.Context, in *SDK.DeleteBucketMetricsConfigurationInput) (*SDK.DeleteBucketMetricsConfigurationResponse, error) {
	return svc.client.DeleteBucketMetricsConfigurationRequest(in).Send(ctx)
}

// RawDeleteBucketPolicy executes `DeleteBucketPolicy` raw operation.
func (svc *S3) RawDeleteBucketPolicy(ctx context.Context, in *SDK.DeleteBucketPolicyInput) (*SDK.DeleteBucketPolicyResponse, error) {
	return svc.client.DeleteBucketPolicyRequest(in).Send(ctx)
}

// RawDeleteBucketReplication executes `DeleteBucketReplication` raw operation.
func (svc *S3) RawDeleteBucketReplication(ctx context.Context, in *SDK.DeleteBucketReplicationInput) (*SDK.DeleteBucketReplicationResponse, error) {
	return svc.client.DeleteBucketReplicationRequest(in).Send(ctx)
}

// RawDeleteBucketTagging executes `DeleteBucketTagging` raw operation.
func (svc *S3) RawDeleteBucketTagging(ctx context.Context, in *SDK.DeleteBucketTaggingInput) (*SDK.DeleteBucketTaggingResponse, error) {
	return svc.client.DeleteBucketTaggingRequest(in).Send(ctx)
}

// RawDeleteBucketWebsite executes `DeleteBucketWebsite` raw operation.
func (svc *S3) RawDeleteBucketWebsite(ctx context.Context, in *SDK.DeleteBucketWebsiteInput) (*SDK.DeleteBucketWebsiteResponse, error) {
	return svc.client.DeleteBucketWebsiteRequest(in).Send(ctx)
}

// RawDeleteObject executes `DeleteObject` raw operation.
func (svc *S3) RawDeleteObject(ctx context.Context, in *SDK.DeleteObjectInput) (*SDK.DeleteObjectResponse, error) {
	return svc.client.DeleteObjectRequest(in).Send(ctx)
}

// RawDeleteObjectTagging executes `DeleteObjectTagging` raw operation.
func (svc *S3) RawDeleteObjectTagging(ctx context.Context, in *SDK.DeleteObjectTaggingInput) (*SDK.DeleteObjectTaggingResponse, error) {
	return svc.client.DeleteObjectTaggingRequest(in).Send(ctx)
}

// RawDeleteObjects executes `DeleteObjects` raw operation.
func (svc *S3) RawDeleteObjects(ctx context.Context, in *SDK.DeleteObjectsInput) (*SDK.DeleteObjectsResponse, error) {
	return svc.client.DeleteObjectsRequest(in).Send(ctx)
}

// RawDeletePublicAccessBlock executes `DeletePublicAccessBlock` raw operation.
func (svc *S3) RawDeletePublicAccessBlock(ctx context.Context, in *SDK.DeletePublicAccessBlockInput) (*SDK.DeletePublicAccessBlockResponse, error) {
	return svc.client.DeletePublicAccessBlockRequest(in).Send(ctx)
}

// RawGetBucketAccelerateConfiguration executes `GetBucketAccelerateConfiguration` raw operation.
func (svc *S3) RawGetBucketAccelerateConfiguration(ctx context.Context, in *SDK.GetBucketAccelerateConfigurationInput) (*SDK.GetBucketAccelerateConfigurationResponse, error) {
	return svc.client.GetBucketAccelerateConfigurationRequest(in).Send(ctx)
}

// RawGetBucketAcl executes `GetBucketAcl` raw operation.
func (svc *S3) RawGetBucketAcl(ctx context.Context, in *SDK.GetBucketAclInput) (*SDK.GetBucketAclResponse, error) { // nolint:golint
	return svc.client.GetBucketAclRequest(in).Send(ctx)
}

// RawGetBucketAnalyticsConfiguration executes `GetBucketAnalyticsConfiguration` raw operation.
func (svc *S3) RawGetBucketAnalyticsConfiguration(ctx context.Context, in *SDK.GetBucketAnalyticsConfigurationInput) (*SDK.GetBucketAnalyticsConfigurationResponse, error) {
	return svc.client.GetBucketAnalyticsConfigurationRequest(in).Send(ctx)
}

// RawGetBucketCors executes `GetBucketCors` raw operation.
func (svc *S3) RawGetBucketCors(ctx context.Context, in *SDK.GetBucketCorsInput) (*SDK.GetBucketCorsResponse, error) {
	return svc.client.GetBucketCorsRequest(in).Send(ctx)
}

// RawGetBucketEncryption executes `GetBucketEncryption` raw operation.
func (svc *S3) RawGetBucketEncryption(ctx context.Context, in *SDK.GetBucketEncryptionInput) (*SDK.GetBucketEncryptionResponse, error) {
	return svc.client.GetBucketEncryptionRequest(in).Send(ctx)
}

// RawGetBucketInventoryConfiguration executes `GetBucketInventoryConfiguration` raw operation.
func (svc *S3) RawGetBucketInventoryConfiguration(ctx context.Context, in *SDK.GetBucketInventoryConfigurationInput) (*SDK.GetBucketInventoryConfigurationResponse, error) {
	return svc.client.GetBucketInventoryConfigurationRequest(in).Send(ctx)
}

// RawGetBucketLifecycle executes `GetBucketLifecycle` raw operation.
func (svc *S3) RawGetBucketLifecycle(ctx context.Context, in *SDK.GetBucketLifecycleInput) (*SDK.GetBucketLifecycleResponse, error) {
	return svc.client.GetBucketLifecycleRequest(in).Send(ctx)
}

// RawGetBucketLifecycleConfiguration executes `GetBucketLifecycleConfiguration` raw operation.
func (svc *S3) RawGetBucketLifecycleConfiguration(ctx context.Context, in *SDK.GetBucketLifecycleConfigurationInput) (*SDK.GetBucketLifecycleConfigurationResponse, error) {
	return svc.client.GetBucketLifecycleConfigurationRequest(in).Send(ctx)
}

// RawGetBucketLocation executes `GetBucketLocation` raw operation.
func (svc *S3) RawGetBucketLocation(ctx context.Context, in *SDK.GetBucketLocationInput) (*SDK.GetBucketLocationResponse, error) {
	return svc.client.GetBucketLocationRequest(in).Send(ctx)
}

// RawGetBucketLogging executes `GetBucketLogging` raw operation.
func (svc *S3) RawGetBucketLogging(ctx context.Context, in *SDK.GetBucketLoggingInput) (*SDK.GetBucketLoggingResponse, error) {
	return svc.client.GetBucketLoggingRequest(in).Send(ctx)
}

// RawGetBucketMetricsConfiguration executes `GetBucketMetricsConfiguration` raw operation.
func (svc *S3) RawGetBucketMetricsConfiguration(ctx context.Context, in *SDK.GetBucketMetricsConfigurationInput) (*SDK.GetBucketMetricsConfigurationResponse, error) {
	return svc.client.GetBucketMetricsConfigurationRequest(in).Send(ctx)
}

// RawGetBucketNotification executes `GetBucketNotification` raw operation.
func (svc *S3) RawGetBucketNotification(ctx context.Context, in *SDK.GetBucketNotificationInput) (*SDK.GetBucketNotificationResponse, error) {
	return svc.client.GetBucketNotificationRequest(in).Send(ctx)
}

// RawGetBucketNotificationConfiguration executes `GetBucketNotificationConfiguration` raw operation.
func (svc *S3) RawGetBucketNotificationConfiguration(ctx context.Context, in *SDK.GetBucketNotificationConfigurationInput) (*SDK.GetBucketNotificationConfigurationResponse, error) {
	return svc.client.GetBucketNotificationConfigurationRequest(in).Send(ctx)
}

// RawGetBucketPolicy executes `GetBucketPolicy` raw operation.
func (svc *S3) RawGetBucketPolicy(ctx context.Context, in *SDK.GetBucketPolicyInput) (*SDK.GetBucketPolicyResponse, error) {
	return svc.client.GetBucketPolicyRequest(in).Send(ctx)
}

// RawGetBucketPolicyStatus executes `GetBucketPolicyStatus` raw operation.
func (svc *S3) RawGetBucketPolicyStatus(ctx context.Context, in *SDK.GetBucketPolicyStatusInput) (*SDK.GetBucketPolicyStatusResponse, error) {
	return svc.client.GetBucketPolicyStatusRequest(in).Send(ctx)
}

// RawGetBucketReplication executes `GetBucketReplication` raw operation.
func (svc *S3) RawGetBucketReplication(ctx context.Context, in *SDK.GetBucketReplicationInput) (*SDK.GetBucketReplicationResponse, error) {
	return svc.client.GetBucketReplicationRequest(in).Send(ctx)
}

// RawGetBucketRequestPayment executes `GetBucketRequestPayment` raw operation.
func (svc *S3) RawGetBucketRequestPayment(ctx context.Context, in *SDK.GetBucketRequestPaymentInput) (*SDK.GetBucketRequestPaymentResponse, error) {
	return svc.client.GetBucketRequestPaymentRequest(in).Send(ctx)
}

// RawGetBucketTagging executes `GetBucketTagging` raw operation.
func (svc *S3) RawGetBucketTagging(ctx context.Context, in *SDK.GetBucketTaggingInput) (*SDK.GetBucketTaggingResponse, error) {
	return svc.client.GetBucketTaggingRequest(in).Send(ctx)
}

// RawGetBucketVersioning executes `GetBucketVersioning` raw operation.
func (svc *S3) RawGetBucketVersioning(ctx context.Context, in *SDK.GetBucketVersioningInput) (*SDK.GetBucketVersioningResponse, error) {
	return svc.client.GetBucketVersioningRequest(in).Send(ctx)
}

// RawGetBucketWebsite executes `GetBucketWebsite` raw operation.
func (svc *S3) RawGetBucketWebsite(ctx context.Context, in *SDK.GetBucketWebsiteInput) (*SDK.GetBucketWebsiteResponse, error) {
	return svc.client.GetBucketWebsiteRequest(in).Send(ctx)
}

// RawGetObject executes `GetObject` raw operation.
func (svc *S3) RawGetObject(ctx context.Context, in *SDK.GetObjectInput) (*SDK.GetObjectResponse, error) {
	return svc.client.GetObjectRequest(in).Send(ctx)
}

// RawGetObjectAcl executes `GetObjectAcl` raw operation.
func (svc *S3) RawGetObjectAcl(ctx context.Context, in *SDK.GetObjectAclInput) (*SDK.GetObjectAclResponse, error) { // nolint:golint
	return svc.client.GetObjectAclRequest(in).Send(ctx)
}

// RawGetObjectLegalHold executes `GetObjectLegalHold` raw operation.
func (svc *S3) RawGetObjectLegalHold(ctx context.Context, in *SDK.GetObjectLegalHoldInput) (*SDK.GetObjectLegalHoldResponse, error) {
	return svc.client.GetObjectLegalHoldRequest(in).Send(ctx)
}

// RawGetObjectLockConfiguration executes `GetObjectLockConfiguration` raw operation.
func (svc *S3) RawGetObjectLockConfiguration(ctx context.Context, in *SDK.GetObjectLockConfigurationInput) (*SDK.GetObjectLockConfigurationResponse, error) {
	return svc.client.GetObjectLockConfigurationRequest(in).Send(ctx)
}

// RawGetObjectRetention executes `GetObjectRetention` raw operation.
func (svc *S3) RawGetObjectRetention(ctx context.Context, in *SDK.GetObjectRetentionInput) (*SDK.GetObjectRetentionResponse, error) {
	return svc.client.GetObjectRetentionRequest(in).Send(ctx)
}

// RawGetObjectTagging executes `GetObjectTagging` raw operation.
func (svc *S3) RawGetObjectTagging(ctx context.Context, in *SDK.GetObjectTaggingInput) (*SDK.GetObjectTaggingResponse, error) {
	return svc.client.GetObjectTaggingRequest(in).Send(ctx)
}

// RawGetObjectTorrent executes `GetObjectTorrent` raw operation.
func (svc *S3) RawGetObjectTorrent(ctx context.Context, in *SDK.GetObjectTorrentInput) (*SDK.GetObjectTorrentResponse, error) {
	return svc.client.GetObjectTorrentRequest(in).Send(ctx)
}

// RawGetPublicAccessBlock executes `GetPublicAccessBlock` raw operation.
func (svc *S3) RawGetPublicAccessBlock(ctx context.Context, in *SDK.GetPublicAccessBlockInput) (*SDK.GetPublicAccessBlockResponse, error) {
	return svc.client.GetPublicAccessBlockRequest(in).Send(ctx)
}

// RawHeadBucket executes `HeadBucket` raw operation.
func (svc *S3) RawHeadBucket(ctx context.Context, in *SDK.HeadBucketInput) (*SDK.HeadBucketResponse, error) {
	return svc.client.HeadBucketRequest(in).Send(ctx)
}

// RawHeadObject executes `HeadObject` raw operation.
func (svc *S3) RawHeadObject(ctx context.Context, in *SDK.HeadObjectInput) (*SDK.HeadObjectResponse, error) {
	return svc.client.HeadObjectRequest(in).Send(ctx)
}

// RawListBucketAnalyticsConfigurations executes `ListBucketAnalyticsConfigurations` raw operation.
func (svc *S3) RawListBucketAnalyticsConfigurations(ctx context.Context, in *SDK.ListBucketAnalyticsConfigurationsInput) (*SDK.ListBucketAnalyticsConfigurationsResponse, error) {
	return svc.client.ListBucketAnalyticsConfigurationsRequest(in).Send(ctx)
}

// RawListBucketInventoryConfigurations executes `ListBucketInventoryConfigurations` raw operation.
func (svc *S3) RawListBucketInventoryConfigurations(ctx context.Context, in *SDK.ListBucketInventoryConfigurationsInput) (*SDK.ListBucketInventoryConfigurationsResponse, error) {
	return svc.client.ListBucketInventoryConfigurationsRequest(in).Send(ctx)
}

// RawListBucketMetricsConfigurations executes `ListBucketMetricsConfigurations` raw operation.
func (svc *S3) RawListBucketMetricsConfigurations(ctx context.Context, in *SDK.ListBucketMetricsConfigurationsInput) (*SDK.ListBucketMetricsConfigurationsResponse, error) {
	return svc.client.ListBucketMetricsConfigurationsRequest(in).Send(ctx)
}

// RawListBuckets executes `ListBuckets` raw operation.
func (svc *S3) RawListBuckets(ctx context.Context, in *SDK.ListBucketsInput) (*SDK.ListBucketsResponse, error) {
	return svc.client.ListBucketsRequest(in).Send(ctx)
}

// RawListMultipartUploads executes `ListMultipartUploads` raw operation.
func (svc *S3) RawListMultipartUploads(ctx context.Context, in *SDK.ListMultipartUploadsInput) (*SDK.ListMultipartUploadsResponse, error) {
	return svc.client.ListMultipartUploadsRequest(in).Send(ctx)
}

// RawListObjectVersions executes `ListObjectVersions` raw operation.
func (svc *S3) RawListObjectVersions(ctx context.Context, in *SDK.ListObjectVersionsInput) (*SDK.ListObjectVersionsResponse, error) {
	return svc.client.ListObjectVersionsRequest(in).Send(ctx)
}

// RawListObjects executes `ListObjects` raw operation.
func (svc *S3) RawListObjects(ctx context.Context, in *SDK.ListObjectsInput) (*SDK.ListObjectsResponse, error) {
	return svc.client.ListObjectsRequest(in).Send(ctx)
}

// RawListObjectsV2 executes `ListObjectsV2` raw operation.
func (svc *S3) RawListObjectsV2(ctx context.Context, in *SDK.ListObjectsV2Input) (*SDK.ListObjectsV2Response, error) {
	return svc.client.ListObjectsV2Request(in).Send(ctx)
}

// RawListParts executes `ListParts` raw operation.
func (svc *S3) RawListParts(ctx context.Context, in *SDK.ListPartsInput) (*SDK.ListPartsResponse, error) {
	return svc.client.ListPartsRequest(in).Send(ctx)
}

// RawPutBucketAccelerateConfiguration executes `PutBucketAccelerateConfiguration` raw operation.
func (svc *S3) RawPutBucketAccelerateConfiguration(ctx context.Context, in *SDK.PutBucketAccelerateConfigurationInput) (*SDK.PutBucketAccelerateConfigurationResponse, error) {
	return svc.client.PutBucketAccelerateConfigurationRequest(in).Send(ctx)
}

// RawPutBucketAcl executes `PutBucketAcl` raw operation.
func (svc *S3) RawPutBucketAcl(ctx context.Context, in *SDK.PutBucketAclInput) (*SDK.PutBucketAclResponse, error) { // nolint:golint
	return svc.client.PutBucketAclRequest(in).Send(ctx)
}

// RawPutBucketAnalyticsConfiguration executes `PutBucketAnalyticsConfiguration` raw operation.
func (svc *S3) RawPutBucketAnalyticsConfiguration(ctx context.Context, in *SDK.PutBucketAnalyticsConfigurationInput) (*SDK.PutBucketAnalyticsConfigurationResponse, error) {
	return svc.client.PutBucketAnalyticsConfigurationRequest(in).Send(ctx)
}

// RawPutBucketCors executes `PutBucketCors` raw operation.
func (svc *S3) RawPutBucketCors(ctx context.Context, in *SDK.PutBucketCorsInput) (*SDK.PutBucketCorsResponse, error) {
	return svc.client.PutBucketCorsRequest(in).Send(ctx)
}

// RawPutBucketEncryption executes `PutBucketEncryption` raw operation.
func (svc *S3) RawPutBucketEncryption(ctx context.Context, in *SDK.PutBucketEncryptionInput) (*SDK.PutBucketEncryptionResponse, error) {
	return svc.client.PutBucketEncryptionRequest(in).Send(ctx)
}

// RawPutBucketInventoryConfiguration executes `PutBucketInventoryConfiguration` raw operation.
func (svc *S3) RawPutBucketInventoryConfiguration(ctx context.Context, in *SDK.PutBucketInventoryConfigurationInput) (*SDK.PutBucketInventoryConfigurationResponse, error) {
	return svc.client.PutBucketInventoryConfigurationRequest(in).Send(ctx)
}

// RawPutBucketLifecycle executes `PutBucketLifecycle` raw operation.
func (svc *S3) RawPutBucketLifecycle(ctx context.Context, in *SDK.PutBucketLifecycleInput) (*SDK.PutBucketLifecycleResponse, error) {
	return svc.client.PutBucketLifecycleRequest(in).Send(ctx)
}

// RawPutBucketLifecycleConfiguration executes `PutBucketLifecycleConfiguration` raw operation.
func (svc *S3) RawPutBucketLifecycleConfiguration(ctx context.Context, in *SDK.PutBucketLifecycleConfigurationInput) (*SDK.PutBucketLifecycleConfigurationResponse, error) {
	return svc.client.PutBucketLifecycleConfigurationRequest(in).Send(ctx)
}

// RawPutBucketLogging executes `PutBucketLogging` raw operation.
func (svc *S3) RawPutBucketLogging(ctx context.Context, in *SDK.PutBucketLoggingInput) (*SDK.PutBucketLoggingResponse, error) {
	return svc.client.PutBucketLoggingRequest(in).Send(ctx)
}

// RawPutBucketMetricsConfiguration executes `PutBucketMetricsConfiguration` raw operation.
func (svc *S3) RawPutBucketMetricsConfiguration(ctx context.Context, in *SDK.PutBucketMetricsConfigurationInput) (*SDK.PutBucketMetricsConfigurationResponse, error) {
	return svc.client.PutBucketMetricsConfigurationRequest(in).Send(ctx)
}

// RawPutBucketNotification executes `PutBucketNotification` raw operation.
func (svc *S3) RawPutBucketNotification(ctx context.Context, in *SDK.PutBucketNotificationInput) (*SDK.PutBucketNotificationResponse, error) {
	return svc.client.PutBucketNotificationRequest(in).Send(ctx)
}

// RawPutBucketNotificationConfiguration executes `PutBucketNotificationConfiguration` raw operation.
func (svc *S3) RawPutBucketNotificationConfiguration(ctx context.Context, in *SDK.PutBucketNotificationConfigurationInput) (*SDK.PutBucketNotificationConfigurationResponse, error) {
	return svc.client.PutBucketNotificationConfigurationRequest(in).Send(ctx)
}

// RawPutBucketPolicy executes `PutBucketPolicy` raw operation.
func (svc *S3) RawPutBucketPolicy(ctx context.Context, in *SDK.PutBucketPolicyInput) (*SDK.PutBucketPolicyResponse, error) {
	return svc.client.PutBucketPolicyRequest(in).Send(ctx)
}

// RawPutBucketReplication executes `PutBucketReplication` raw operation.
func (svc *S3) RawPutBucketReplication(ctx context.Context, in *SDK.PutBucketReplicationInput) (*SDK.PutBucketReplicationResponse, error) {
	return svc.client.PutBucketReplicationRequest(in).Send(ctx)
}

// RawPutBucketRequestPayment executes `PutBucketRequestPayment` raw operation.
func (svc *S3) RawPutBucketRequestPayment(ctx context.Context, in *SDK.PutBucketRequestPaymentInput) (*SDK.PutBucketRequestPaymentResponse, error) {
	return svc.client.PutBucketRequestPaymentRequest(in).Send(ctx)
}

// RawPutBucketTagging executes `PutBucketTagging` raw operation.
func (svc *S3) RawPutBucketTagging(ctx context.Context, in *SDK.PutBucketTaggingInput) (*SDK.PutBucketTaggingResponse, error) {
	return svc.client.PutBucketTaggingRequest(in).Send(ctx)
}

// RawPutBucketVersioning executes `PutBucketVersioning` raw operation.
func (svc *S3) RawPutBucketVersioning(ctx context.Context, in *SDK.PutBucketVersioningInput) (*SDK.PutBucketVersioningResponse, error) {
	return svc.client.PutBucketVersioningRequest(in).Send(ctx)
}

// RawPutBucketWebsite executes `PutBucketWebsite` raw operation.
func (svc *S3) RawPutBucketWebsite(ctx context.Context, in *SDK.PutBucketWebsiteInput) (*SDK.PutBucketWebsiteResponse, error) {
	return svc.client.PutBucketWebsiteRequest(in).Send(ctx)
}

// RawPutObject executes `PutObject` raw operation.
func (svc *S3) RawPutObject(ctx context.Context, in *SDK.PutObjectInput) (*SDK.PutObjectResponse, error) {
	return svc.client.PutObjectRequest(in).Send(ctx)
}

// RawPutObjectAcl executes `PutObjectAcl` raw operation.
func (svc *S3) RawPutObjectAcl(ctx context.Context, in *SDK.PutObjectAclInput) (*SDK.PutObjectAclResponse, error) { // nolint:golint
	return svc.client.PutObjectAclRequest(in).Send(ctx)
}

// RawPutObjectLegalHold executes `PutObjectLegalHold` raw operation.
func (svc *S3) RawPutObjectLegalHold(ctx context.Context, in *SDK.PutObjectLegalHoldInput) (*SDK.PutObjectLegalHoldResponse, error) {
	return svc.client.PutObjectLegalHoldRequest(in).Send(ctx)
}

// RawPutObjectLockConfiguration executes `PutObjectLockConfiguration` raw operation.
func (svc *S3) RawPutObjectLockConfiguration(ctx context.Context, in *SDK.PutObjectLockConfigurationInput) (*SDK.PutObjectLockConfigurationResponse, error) {
	return svc.client.PutObjectLockConfigurationRequest(in).Send(ctx)
}

// RawPutObjectRetention executes `PutObjectRetention` raw operation.
func (svc *S3) RawPutObjectRetention(ctx context.Context, in *SDK.PutObjectRetentionInput) (*SDK.PutObjectRetentionResponse, error) {
	return svc.client.PutObjectRetentionRequest(in).Send(ctx)
}

// RawPutObjectTagging executes `PutObjectTagging` raw operation.
func (svc *S3) RawPutObjectTagging(ctx context.Context, in *SDK.PutObjectTaggingInput) (*SDK.PutObjectTaggingResponse, error) {
	return svc.client.PutObjectTaggingRequest(in).Send(ctx)
}

// RawPutPublicAccessBlock executes `PutPublicAccessBlock` raw operation.
func (svc *S3) RawPutPublicAccessBlock(ctx context.Context, in *SDK.PutPublicAccessBlockInput) (*SDK.PutPublicAccessBlockResponse, error) {
	return svc.client.PutPublicAccessBlockRequest(in).Send(ctx)
}

// RawRestoreObject executes `RestoreObject` raw operation.
func (svc *S3) RawRestoreObject(ctx context.Context, in *SDK.RestoreObjectInput) (*SDK.RestoreObjectResponse, error) {
	return svc.client.RestoreObjectRequest(in).Send(ctx)
}

// RawUploadPart executes `UploadPart` raw operation.
func (svc *S3) RawUploadPart(ctx context.Context, in *SDK.UploadPartInput) (*SDK.UploadPartResponse, error) {
	return svc.client.UploadPartRequest(in).Send(ctx)
}

// RawUploadPartCopy executes `UploadPartCopy` raw operation.
func (svc *S3) RawUploadPartCopy(ctx context.Context, in *SDK.UploadPartCopyInput) (*SDK.UploadPartCopyResponse, error) {
	return svc.client.UploadPartCopyRequest(in).Send(ctx)
}
