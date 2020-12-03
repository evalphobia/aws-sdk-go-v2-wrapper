package cloudwatchlogs

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "CloudwatchLogs"
)

// CloudwatchLogs has CloudwatchLogs client.
type CloudwatchLogs struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *CloudwatchLogs.
func New(conf config.Config) (*CloudwatchLogs, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}
	if conf.Endpoints.HasCloudwatchLogs() {
		c.EndpointResolver = conf.Endpoints.GetCloudwatchLogs()
	}

	svc := &CloudwatchLogs{
		client:  SDK.New(c),
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *CloudwatchLogs) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *CloudwatchLogs) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *CloudwatchLogs) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *CloudwatchLogs) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
