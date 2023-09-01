package log

import "github.com/sirupsen/logrus"

type CustomLogrus struct {
	entry *logrus.Entry
}

func (cl *CustomLogrus) Trace(msg string) {
	cl.entry.Trace(msg)
}

func (cl *CustomLogrus) Tracef(format string, params ...interface{}) {
	cl.entry.Tracef(format, params...)
}

func (cl *CustomLogrus) Debug(msg string) {
	cl.entry.Debug(msg)
}

func (cl *CustomLogrus) Debugf(format string, params ...interface{}) {
	cl.entry.Debugf(format, params...)
}

func (cl *CustomLogrus) Info(msg string) {
	cl.entry.Info(msg)
}

func (cl *CustomLogrus) Infof(format string, params ...interface{}) {
	cl.entry.Infof(format, params...)
}

func (cl *CustomLogrus) Warn(msg string) {
	cl.entry.Warn(msg)
}

func (cl *CustomLogrus) Warnf(format string, params ...interface{}) {
	cl.entry.Warnf(format, params...)
}

func (cl *CustomLogrus) Error(msg string) {
	cl.entry.Error(msg)
}

func (cl *CustomLogrus) Errorf(format string, params ...interface{}) {
	cl.entry.Errorf(format, params...)
}

func (cl *CustomLogrus) Critical(msg string) {
	cl.entry.Fatal(msg)
}

func (cl *CustomLogrus) Criticalf(format string, params ...interface{}) {
	cl.entry.Fatalf(format, params...)
}
