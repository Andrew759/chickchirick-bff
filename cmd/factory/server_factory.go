package factory

import (
	"chickchirick-bff/pkg/chirick_config"

	_ "net/http/pprof"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func BuildAndServe() error {
	return BuildServer()
}

func BuildServer() error {
	e := gin.Default()

	//TODO: доработать CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{viper.GetString(chirick_config.UserApp)}

	e.Use(cors.New(config))

	//InitAuthServer(e, &c_controller.DIContainer{DBDecorator: dbDecorator, RedisDecorator: redisDecorator})

	err := e.Run()
	if err != nil {
		return err
	}

	return nil
}
