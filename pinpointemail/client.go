package pinpointemail

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpointemail"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "PinpointEmail"
)

// PinpointEmail has PinpointEmail client.
type PinpointEmail struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *PinpointEmail.
func New(conf config.Config) (*PinpointEmail, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasPinpointEmail() {
		c.EndpointResolver = conf.Endpoints.GetPinpointEmail()
	}

	svc := &PinpointEmail{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *PinpointEmail) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *PinpointEmail) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *PinpointEmail) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *PinpointEmail) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
