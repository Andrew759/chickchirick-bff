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
	RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       int
}

func (c AppConfiguration) NewAppConfiguration() AppConfiguration {
	return AppConfiguration{
		ServerURL:   viper.GetString(globalConfig.ServerUrl),
		Environment: viper.GetString(globalConfig.Enviroment),
		RedisConfig: RedisConfig{
			Host:     viper.GetString(globalConfig.RedisHost),
			Port:     viper.GetInt(globalConfig.RedisInternalPort),
			User:     viper.GetString(globalConfig.RedisUser),
			Password: viper.GetString(globalConfig.RedisPassword),
			Db:       viper.GetInt(globalConfig.RedisDB),
		},
	}
}
