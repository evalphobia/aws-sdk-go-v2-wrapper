package ssm

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "SSM"
)

// SSM has SSM client.
type SSM struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *SSM.
func New(conf config.Config) (*SSM, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasSSM() {
		c.EndpointResolver = conf.Endpoints.GetSSM()
	}

	svc := &SSM{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *SSM) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *SSM) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *SSM) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *SSM) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
