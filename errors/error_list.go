package errors

import "strings"

type ErrorList []error

func (ee ErrorList) Error() string {
	var buf strings.Builder
	for _, e := range ee {
		buf.WriteString(e.Error())
	}

	return buf.String()
}
