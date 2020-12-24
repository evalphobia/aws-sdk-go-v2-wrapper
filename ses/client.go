package ses

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/ses"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "SES"
)

// SES has SES client.
type SES struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *SES.
func New(conf config.Config) (*SES, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasSES() {
		c.EndpointResolver = conf.Endpoints.GetSES()
	}

	svc := &SES{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *SES) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *SES) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *SES) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *SES) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
