package config

import "os"

const (
	envRegion = "AWS_REGION"

	envEndpoint               = "AWS_ENDPOINT"
	envAthenaEndpoint         = "AWS_ATHENA_ENDPOINT"
	envCloudwatchLogsEndpoint = "AWS_CLOUDWATCH_LOGS_ENDPOINT"
	envDynamoDBEndpoint       = "AWS_DYNAMODB_ENDPOINT"
	envEC2Endpoint            = "AWS_EC2_ENDPOINT"
	envKMSEndpoint            = "AWS_KMS_ENDPOINT"
	envPinpointEndpoint       = "AWS_PINPOINT_ENDPOINT"
	envPinpointEmailEndpoint  = "AWS_PINPOINT_EMAIL_ENDPOINT"
	envS3Endpoint             = "AWS_S3_ENDPOINT"
	envSESEndpoint            = "AWS_SES_ENDPOINT"
	envSNSEndpoint            = "AWS_SNS_ENDPOINT"
	envSQSEndpoint            = "AWS_SQS_ENDPOINT"
	envSSMEndpoint            = "AWS_SSM_ENDPOINT"
)

// EnvRegion get region from env vars.
func EnvRegion() string {
	return os.Getenv(envRegion)
}

// EnvEndpoint get endpoint from env vars.
func EnvEndpoint() string {
	return os.Getenv(envEndpoint)
}

// EnvAthenaEndpoint get Athena endpoint from env vars.
func EnvAthenaEndpoint() string {
	return os.Getenv(envAthenaEndpoint)
}

// EnvCloudwatchLogsEndpoint get CloudwatchLogs endpoint from env vars.
func EnvCloudwatchLogsEndpoint() string {
	return os.Getenv(envCloudwatchLogsEndpoint)
}

// EnvDynamoDBEndpoint get DynamoDB endpoint from env vars.
func EnvDynamoDBEndpoint() string {
	return os.Getenv(envDynamoDBEndpoint)
}

// EnvEC2Endpoint get EC2 endpoint from env vars.
func EnvEC2Endpoint() string {
	return os.Getenv(envEC2Endpoint)
}

// EnvPinpointEndpoint get Pinpoint endpoint from env vars.
func EnvPinpointEndpoint() string {
	return os.Getenv(envPinpointEndpoint)
}

// EnvPinpointEmailEndpoint get PinpointEmail endpoint from env vars.
func EnvPinpointEmailEndpoint() string {
	return os.Getenv(envPinpointEmailEndpoint)
}

// EnvKMSEndpoint get KMS endpoint from env vars.
func EnvKMSEndpoint() string {
	return os.Getenv(envKMSEndpoint)
}

// EnvS3Endpoint get S3 endpoint from env vars.
func EnvS3Endpoint() string {
	return os.Getenv(envS3Endpoint)
}

// EnvSESEndpoint get SES endpoint from env vars.
func EnvSESEndpoint() string {
	return os.Getenv(envSESEndpoint)
}

// EnvSNSEndpoint get SNS endpoint from env vars.
func EnvSNSEndpoint() string {
	return os.Getenv(envSNSEndpoint)
}

// EnvSQSEndpoint get SQS endpoint from env vars.
func EnvSQSEndpoint() string {
	return os.Getenv(envSQSEndpoint)
}

// EnvSSMEndpoint get SSM endpoint from env vars.
func EnvSSMEndpoint() string {
	return os.Getenv(envSSMEndpoint)
}
