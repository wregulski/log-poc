package log

import (
	"github.com/spf13/viper"
	"github.com/wregulski/log-poc/pkg/config"
)

// DefaultLoggerFactory creates an instance of logger factory with default configuration for this app.
func DefaultLoggerFactory() LoggerFactory {
	lvl := viper.GetString(config.EnvLoggingLevel)
	appName := viper.GetString(config.EnvApplicationName)

	return NewLogrusLoggerFactory(appName, LevelFromString(lvl))
}
