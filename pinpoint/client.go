package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "Pinpoint"
)

// Pinpoint has Pinpoint client.
type Pinpoint struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *Pinpoint.
func New(conf config.Config) (*Pinpoint, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasPinpoint() {
		c.EndpointResolver = conf.Endpoints.GetPinpoint()
	}

	svc := &Pinpoint{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *Pinpoint) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *Pinpoint) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *Pinpoint) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *Pinpoint) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
