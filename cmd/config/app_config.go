package config

import (
	globalConfig "chickchirick-bff/pkg/chirick_config"

	"github.com/spf13/viper"
)

type AppConfigurationInterface interface {
	NewAppConfiguration() AppConfiguration
}

type AppConfiguration struct {
	ServerURL   string
	Environment string
}

func (c AppConfiguration) NewAppConfiguration() AppConfiguration {
	return AppConfiguration{
		ServerURL:   viper.GetString(globalConfig.ServerUrl),
		Environment: viper.GetString(globalConfig.Enviroment),
	}
}
