package config

import "os"

const (
	envRegion = "AWS_REGION"

	envEndpoint         = "AWS_ENDPOINT"
	envDynamoDBEndpoint = "AWS_DYNAMODB_ENDPOINT"
	envEC2Endpoint      = "AWS_EC2_ENDPOINT"
	envKMSEndpoint      = "AWS_KMS_ENDPOINT"
	envS3Endpoint       = "AWS_S3_ENDPOINT"
	envSNSEndpoint      = "AWS_SNS_ENDPOINT"
	envSQSEndpoint      = "AWS_SQS_ENDPOINT"
)

// EnvRegion get region from env vars.
func EnvRegion() string {
	return os.Getenv(envRegion)
}

// EnvEndpoint get endpoint from env vars.
func EnvEndpoint() string {
	return os.Getenv(envEndpoint)
}

// EnvDynamoDBEndpoint get DynamoDB endpoint from env vars.
func EnvDynamoDBEndpoint() string {
	return os.Getenv(envDynamoDBEndpoint)
}

// EnvEC2Endpoint get EC2 endpoint from env vars.
func EnvEC2Endpoint() string {
	return os.Getenv(envEC2Endpoint)
}

// EnvKMSEndpoint get KMS endpoint from env vars.
func EnvKMSEndpoint() string {
	return os.Getenv(envKMSEndpoint)
}

// EnvS3Endpoint get S3 endpoint from env vars.
func EnvS3Endpoint() string {
	return os.Getenv(envS3Endpoint)
}

// EnvSNSEndpoint get SNS endpoint from env vars.
func EnvSNSEndpoint() string {
	return os.Getenv(envSNSEndpoint)
}

// EnvSQSEndpoint get SQS endpoint from env vars.
func EnvSQSEndpoint() string {
	return os.Getenv(envSQSEndpoint)
}
