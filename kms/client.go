package kms

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/kms"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "KSM"
)

// KMS has KMS client.
type KMS struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *KMS.
func New(conf config.Config) (*KMS, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasKMS() {
		c.EndpointResolver = conf.Endpoints.GetKMS()
	}

	svc := &KMS{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *KMS) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *KMS) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *KMS) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *KMS) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
