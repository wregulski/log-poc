package config

import (
	"strings"

	"github.com/spf13/viper"
)

func NewViperConfig() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	setLoggingDefaults()
	return &Config{}
}

func setLoggingDefaults() {
	viper.SetDefault(EnvLoggingLevel, "Debug")
}

func setApplicationName() {
	viper.SetDefault(EnvApplicationName, "4chain-logger-poc")
}
