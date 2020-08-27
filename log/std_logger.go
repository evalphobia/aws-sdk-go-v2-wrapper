package log

import "log"

// StdLogger use standard log package.
type StdLogger struct{}

// Log logging the SDK's log.
func (*StdLogger) Log(v ...interface{}) {
	log.Println(v...)
}

// Infof logging information.
func (*StdLogger) Infof(service, format string, v ...interface{}) {
	log.Printf("[INFO] ["+service+"] "+format, v...)
}

// Errorf logging error information.
func (*StdLogger) Errorf(service, format string, v ...interface{}) {
	log.Printf("[ERROR] ["+service+"] "+format, v...)
}
