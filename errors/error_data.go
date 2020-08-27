package errors

import "fmt"

var (
	DefaultErrWrap = func(errData ErrorData) error {
		return fmt.Errorf("error on `%s`: %w", errData.AWSOperation, errData.Err)
	}
)

type ErrorData struct {
	FuncName     string
	AWSOperation string
	Err          error
}
