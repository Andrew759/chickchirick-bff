package main

import (
	"chickchirick-auth/cmd/config"
	"chickchirick-auth/cmd/factory"
	"chickchirick-auth/cmd/service"
)

func main() {
	factory.InitViper()

	appConfig := config.AppConfiguration{}.NewAppConfiguration()

	redisDecorator := service.InitRedis(appConfig.RedisConfig)
	defer redisDecorator.RedisClose()

	//TODO: если не потребуется - удалить
	httpClient := factory.InitHttpClient()

	factory.BuildAndServe(redisDecorator)
}
