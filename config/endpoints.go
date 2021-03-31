package config

import "github.com/aws/aws-sdk-go-v2/aws"

type Endpoints struct {
	Athena         string
	CloudwatchLogs string
	DynamoDB       string
	EC2            string
	KMS            string
	Pinpoint       string
	PinpointEmail  string
	S3             string
	SES            string
	SSM            string
}

func (e Endpoints) HasAthena() bool {
	if e.Athena != "" {
		return true
	}
	return EnvAthenaEndpoint() != ""
}

func (e Endpoints) GetAthena() aws.ResolveWithEndpoint {
	envvar := EnvAthenaEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.Athena)
}

func (e Endpoints) HasCloudwatchLogs() bool {
	if e.CloudwatchLogs != "" {
		return true
	}
	return EnvCloudwatchLogsEndpoint() != ""
}

func (e Endpoints) GetCloudwatchLogs() aws.ResolveWithEndpoint {
	envvar := EnvCloudwatchLogsEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.CloudwatchLogs)
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

func (e Endpoints) HasPinpoint() bool {
	if e.Pinpoint != "" {
		return true
	}
	return EnvPinpointEndpoint() != ""
}

func (e Endpoints) GetPinpoint() aws.ResolveWithEndpoint {
	envvar := EnvPinpointEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.Pinpoint)
}

func (e Endpoints) HasPinpointEmail() bool {
	if e.PinpointEmail != "" {
		return true
	}
	return EnvPinpointEmailEndpoint() != ""
}

func (e Endpoints) GetPinpointEmail() aws.ResolveWithEndpoint {
	envvar := EnvPinpointEmailEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.PinpointEmail)
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

func (e Endpoints) HasSES() bool {
	if e.SES != "" {
		return true
	}
	return EnvSESEndpoint() != ""
}

func (e Endpoints) GetSES() aws.ResolveWithEndpoint {
	envvar := EnvSESEndpoint()
	if envvar != "" {
		return aws.ResolveWithEndpointURL(envvar)
	}
	return aws.ResolveWithEndpointURL(e.SES)
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
