package log

import "github.com/sirupsen/logrus"

type logrusLoggerFactory struct {
	application string
	logger      *logrus.Logger
}

func NewLogrusLoggerFactory(appName string, level Level) LoggerFactory {
	logger := logrus.New()
	logger.SetLevel(toLogrusLevel(level))

	return &logrusLoggerFactory{
		application: appName,
		logger:      logger,
	}
}

func (lf *logrusLoggerFactory) NewLogger(name string) Logger {
	logger := logrus.NewEntry(lf.logger).
		WithField("application", lf.application).
		WithField("name", name)

	return NewLogrusAdapter(logger.Logger)
}

func (lf *logrusLoggerFactory) Level() Level {
	return fromLogrusLevel(lf.logger.GetLevel())
}

func (lf *logrusLoggerFactory) SetLevel(level Level) {
	lf.logger.SetLevel(toLogrusLevel(level))
}

func fromLogrusLevel(level logrus.Level) Level {
	switch level {
	case logrus.TraceLevel:
		return Trace
	case logrus.DebugLevel:
		return Debug
	case logrus.InfoLevel:
		return Info
	case logrus.WarnLevel:
		return Warn
	case logrus.ErrorLevel:
		return Error
	case logrus.FatalLevel:
		return Critical
	case logrus.PanicLevel:
		return Critical
	default:
		return Off
	}
}

func toLogrusLevel(level Level) logrus.Level {
	switch level {
	case Trace:
		return logrus.TraceLevel
	case Debug:
		return logrus.DebugLevel
	case Info:
		return logrus.InfoLevel
	case Warn:
		return logrus.WarnLevel
	case Error:
		return logrus.ErrorLevel
	case Critical:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
