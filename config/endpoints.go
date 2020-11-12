package config

import "github.com/aws/aws-sdk-go-v2/aws"

type Endpoints struct {
	DynamoDB string
	EC2      string
	KMS      string
	S3       string
	SSM      string
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

func (e Endpoints) HasKMS() bool {
	if e.KMS != "" {
		return true
	}
	return EnvKMSEndpoint() != ""
}

func (e Endpoints) GetKMS() aws.ResolveWithEndpoint {
	envvar := EnvKMSEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.KMS)
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

func (e Endpoints) HasSSM() bool {
	if e.SSM != "" {
		return true
	}
	return EnvSSMEndpoint() != ""
}

func (e Endpoints) GetSSM() aws.ResolveWithEndpoint {
	envvar := EnvSSMEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.SSM)
}
