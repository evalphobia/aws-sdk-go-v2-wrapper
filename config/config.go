package config

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const defaultRegion = "us-east-1"

// Config has AWS settings.
type Config struct {
	Region         string
	CommonEndpoint string
	Endpoints      Endpoints
	Logger         log.Logger

	EnableDefaultCreds bool

	// Static Credentials
	AccessKey string
	SecretKey string

	// File Credentials
	Filename string
	Profile  string

	// DefaultPrefix is used for service resource prefix
	// e.g.) DynamoDB table, S3 bucket, SQS Queue
	DefaultPrefix string

	// Specific sevice's options
	S3ForcePathStyle bool

	// Custom Function to wrap errors.
	ErrWrap func(errors.ErrorData) error

	// Custom HTTP Client
	CustomHTTPClient *http.Client
	// Showing req/resp data when `true` (Cannot use with CustomHTTPClient)
	UseDebugRequest bool
}

// AWSConfig creates *aws.Config object from the fields.
func (c Config) AWSConfig() (aws.Config, error) {
	cfg, err := c.loadConfig()
	if err != nil {
		return cfg, err
	}

	region := c.getRegion()
	if region != "" {
		cfg.Region = region
	}
	if c.Logger != nil {
		cfg.Logger = c.Logger
	}
	if c.CommonEndpoint != "" {
		cfg.EndpointResolver = aws.ResolveWithEndpointURL(c.CommonEndpoint)
		cfg.DisableEndpointHostPrefix = true
	}
	switch {
	case c.CustomHTTPClient != nil:
		cfg.HTTPClient = c.CustomHTTPClient
	case c.UseDebugRequest:
		cfg.HTTPClient = &http.Client{
			Transport: DefaultDebugTransport,
		}
	}

	return cfg, nil
}

func (c Config) GetLogger() log.Logger {
	if c.Logger != nil {
		return c.Logger
	}
	return log.DefaultLogger
}

func (c Config) GetErrWrap() func(errors.ErrorData) error {
	if c.ErrWrap != nil {
		return c.ErrWrap
	}
	return errors.DefaultErrWrap
}

func (c Config) loadConfig() (aws.Config, error) {
	if c.EnableDefaultCreds {
		return external.LoadDefaultAWSConfig()
	}

	var opts []external.Config
	if c.AccessKey != "" && c.SecretKey != "" {
		opts = append(opts, external.WithCredentialsProvider{
			CredentialsProvider: aws.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     c.AccessKey,
					SecretAccessKey: c.SecretKey,
				},
			},
		})
	}

	if c.Profile != "" {
		opts = append(opts, external.WithSharedConfigProfile(c.Profile))
	}
	if c.Filename != "" {
		opts = append(opts, external.WithSharedConfigFiles([]string{c.Filename}))
	}
	return external.LoadDefaultAWSConfig(opts...)
}

func (c Config) getRegion() string {
	if c.Region != "" {
		return c.Region
	}
	reg := EnvRegion()
	if reg != "" {
		return reg
	}
	return defaultRegion
}
