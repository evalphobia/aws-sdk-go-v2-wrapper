package dynamodb

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "DynamoDB"
)

// DynamoDB has DynamoDB client.
type DynamoDB struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *DynamoDB.
func New(conf config.Config) (*DynamoDB, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasDynamoDB() {
		c.EndpointResolver = conf.Endpoints.GetDynamoDB()
	}

	svc := &DynamoDB{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *DynamoDB) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *DynamoDB) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *DynamoDB) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *DynamoDB) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
