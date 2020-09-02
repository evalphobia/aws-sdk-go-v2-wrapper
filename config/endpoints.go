package config

import "github.com/aws/aws-sdk-go-v2/aws"

type Endpoints struct {
	EC2      string
	S3       string
	DynamoDB string
}

func (e Endpoints) HasEC2() bool {
	if e.EC2 != "" {
		return true
	}
	return EnvEC2Endpoint() != ""
}

func (e Endpoints) GetEC2() aws.ResolveWithEndpoint {
	envvar := EnvEC2Endpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.EC2)
}

func (e Endpoints) HasS3() bool {
	if e.S3 != "" {
		return true
	}
	return EnvS3Endpoint() != ""
}

func (e Endpoints) GetS3() aws.ResolveWithEndpoint {
	envvar := EnvS3Endpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.S3)
}

func (e Endpoints) HasDynamoDB() bool {
	if e.DynamoDB != "" {
		return true
	}
	return EnvDynamoDBEndpoint() != ""
}

func (e Endpoints) GetDynamoDB() aws.ResolveWithEndpoint {
	envvar := EnvDynamoDBEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.DynamoDB)
}
