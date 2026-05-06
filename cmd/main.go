package main

import (
	"chickchirick-auth/cmd/config"
	"chickchirick-auth/cmd/factory"
)

func main() {
	factory.InitViper()

	appConfig := config.AppConfiguration{}.NewAppConfiguration()

	//TODO: если не потребуется - удалить
	//httpClient := factory.InitHttpClient()

	factory.BuildAndServe(dbDecorator, redisDecorator)
}
