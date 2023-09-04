package log

import "github.com/sirupsen/logrus"

type CustomLoggerFactory struct {
	Application string
	Logger      *logrus.Logger
}

func NewCustomLoggerFactory(appName string, level Level) LoggerFactory {
	logger := logrus.New()
	logger.SetLevel(toLogrusLevel(level))
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
		}})

	return &CustomLoggerFactory{
		Application: appName,
		Logger:      logger,
	}
}

func (lf *CustomLoggerFactory) NewLogger(name string) Logger {
	entry := logrus.NewEntry(lf.Logger).
		WithFields(logrus.Fields{
			"application": lf.Application,
			"log.logger":  name,
		})
	return &CustomLogrus{
		entry: entry,
	}
}

func (lf *CustomLoggerFactory) Level() Level {
	return fromLogrusLevel(lf.Logger.GetLevel())
}

func (lf *CustomLoggerFactory) SetLevel(level Level) {
	lf.Logger.SetLevel(toLogrusLevel(level))
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
