package errors

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
)

var (
	DefaultErrWrap = func(errData ErrorData) error {
		errData.PrintErr = func() string {
			return fmt.Sprintf("error on `%s`: [%s]", errData.AWSOperation, errData.Err.Error())
		}
		return errData
	}
)

type ErrorData struct {
	PrintErr     func() string
	FuncName     string
	AWSOperation string
	Err          error
}

func (e ErrorData) Error() string {
	return e.PrintErr()
}

func (e ErrorData) GetAWSErrCode() string {
	return GetAWSErrorCode(e.Err)
}

func GetAWSErrorCode(err error) string {
	if aerr, ok := err.(awserr.Error); ok {
		return aerr.Code()
	}
	return ""
}
