package ses

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/ses"
)

type BulkEmailStatus string

const (
	BulkEmailStatusSuccess                       BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusSuccess)
	BulkEmailStatusMessageRejected               BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusMessageRejected)
	BulkEmailStatusMailFromDomainNotVerified     BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusMailFromDomainNotVerified)
	BulkEmailStatusConfigurationSetDoesNotExist  BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusConfigurationSetDoesNotExist)
	BulkEmailStatusTemplateDoesNotExist          BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusTemplateDoesNotExist)
	BulkEmailStatusAccountSuspended              BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusAccountSuspended)
	BulkEmailStatusAccountThrottled              BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusAccountThrottled)
	BulkEmailStatusAccountDailyQuotaExceeded     BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusAccountDailyQuotaExceeded)
	BulkEmailStatusInvalidSendingPoolName        BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusInvalidSendingPoolName)
	BulkEmailStatusAccountSendingPaused          BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusAccountSendingPaused)
	BulkEmailStatusConfigurationSetSendingPaused BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusConfigurationSetSendingPaused)
	BulkEmailStatusInvalidParameterValue         BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusInvalidParameterValue)
	BulkEmailStatusTransientFailure              BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusTransientFailure)
	BulkEmailStatusFailed                        BulkEmailStatus = BulkEmailStatus(SDK.BulkEmailStatusFailed)
)
