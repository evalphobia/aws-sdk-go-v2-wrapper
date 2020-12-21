package athena

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/athena"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "athena"
)

// Athena has Athena client.
type Athena struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *Athena.
func New(conf config.Config) (*Athena, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasAthena() {
		c.EndpointResolver = conf.Endpoints.GetAthena()
	}

	svc := &Athena{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *Athena) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *Athena) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *Athena) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *Athena) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
