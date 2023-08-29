package log

import "github.com/sirupsen/logrus"

type logrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter(logger *logrus.Logger) Logger {
	return &logrusAdapter{
		logger: logger,
	}
}

func (la *logrusAdapter) Trace(msg string) {
	la.logger.Trace(msg)
}

func (la *logrusAdapter) Tracef(format string, params ...interface{}) {
	la.logger.Tracef(format, params...)
}

func (la *logrusAdapter) Debug(msg string) {
	la.logger.Debug(msg)
}

func (la *logrusAdapter) Debugf(format string, params ...interface{}) {
	la.logger.Debugf(format, params...)
}

func (la *logrusAdapter) Info(msg string) {
	la.logger.Info(msg)
}

func (la *logrusAdapter) Infof(format string, params ...interface{}) {
	la.logger.Infof(format, params...)
}

func (la *logrusAdapter) Warn(msg string) {
	la.logger.Warn(msg)
}

func (la *logrusAdapter) Warnf(format string, params ...interface{}) {
	la.logger.Warnf(format, params...)
}

func (la *logrusAdapter) Error(msg string) {
	la.logger.Error(msg)
}

func (la *logrusAdapter) Errorf(format string, params ...interface{}) {
	la.logger.Errorf(format, params...)
}

func (la *logrusAdapter) Critical(msg string) {
	la.logger.Fatal(msg)
}

func (la *logrusAdapter) Criticalf(format string, params ...interface{}) {
	la.logger.Fatalf(format, params...)
}
