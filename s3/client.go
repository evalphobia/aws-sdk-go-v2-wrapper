package s3

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	serviceName = "S3"
)

// S3 has S3 client.
type S3 struct {
	client *SDK.Client

	logger  log.Logger
	errWrap func(errors.ErrorData) error
}

// New returns initialized *S3.
func New(conf config.Config) (*S3, error) {
	c, err := conf.AWSConfig()
	if err != nil {
		return nil, err
	}

	cli := SDK.New(c)
	switch {
	case conf.Endpoints.HasS3():
		c.EndpointResolver = conf.Endpoints.GetS3()
		// c.DisableEndpointHostPrefix = true
		cli.ForcePathStyle = true
	case conf.CommonEndpoint != "":
		cli.ForcePathStyle = true
	}

	svc := &S3{
		client:  cli,
		logger:  conf.GetLogger(),
		errWrap: conf.GetErrWrap(),
	}
	return svc, nil
}

// GetClient gets original SDK client.
func (svc *S3) GetClient() *SDK.Client {
	return svc.client
}

// SetLogger sets logger.
func (svc *S3) SetLogger(logger log.Logger) {
	svc.logger = logger
}

// Infof logging information.
func (svc *S3) Infof(format string, v ...interface{}) {
	svc.logger.Infof(serviceName, format, v...)
}

// Errorf logging error information.
func (svc *S3) Errorf(format string, v ...interface{}) {
	svc.logger.Errorf(serviceName, format, v...)
}
