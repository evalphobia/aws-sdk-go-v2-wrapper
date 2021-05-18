package sqs

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "SQS"
)

// SQS has SQS client.
type SQS struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *SQS.
func New(conf config.Config) (*SQS, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasSQS() {
		c.EndpointResolver = conf.Endpoints.GetSQS()
	}

	svc := &SQS{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *SQS) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *SQS) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *SQS) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *SQS) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
