package log

// DefaultLogger is default Logger.
var DefaultLogger Logger

// Logger is logging interface.
type Logger interface {
	// For AWS SDK's log
	Log(v ...interface{})

	// For wrapper's log
	Infof(service, format string, v ...interface{})
	Errorf(service, format string, v ...interface{})
}

func init() {
	v := &DummyLogger{}
	DefaultLogger = v
}

// DummyLogger does not output anything
type DummyLogger struct{}

// Log does nothing.
func (*DummyLogger) Log(...interface{}) {}

// Infof does nothing.
func (*DummyLogger) Infof(service, format string, v ...interface{}) {}

// Errorf does nothing.
func (*DummyLogger) Errorf(serivce, format string, v ...interface{}) {}
