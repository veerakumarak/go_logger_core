package logger

type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...interface{})

	Info(msg string)
	Infof(format string, args ...interface{})

	Warn(msg string)
	Warnf(format string, args ...interface{})

	Error(msg string)
	Errorf(format string, args ...interface{})

	With(args ...interface{}) Logger
}
