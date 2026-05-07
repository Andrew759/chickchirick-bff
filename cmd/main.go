package main

import (
	"chickchirick-bff/cmd/config"
	"chickchirick-bff/cmd/factory"
	"chickchirick-bff/cmd/service"
)

func main() {
	factory.InitViper()

	appConfig := config.AppConfiguration{}.NewAppConfiguration()

	redisDecorator := service.InitRedis(appConfig.RedisConfig)
	defer redisDecorator.RedisClose()

	//TODO: если не потребуется - удалить
	httpClient := factory.InitHttpClient()

	factory.BuildAndServe(redisDecorator, httpClient)
}
