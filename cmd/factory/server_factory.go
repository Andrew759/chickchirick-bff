package factory

import (
	"chickchirick-bff/cmd/service"
	"chickchirick-bff/pkg/chirick_config"
	"net/http"

	_ "net/http/pprof"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func BuildAndServe(rDecorator *service.RedisDecorator, client *http.Client) error {
	return BuildServer(rDecorator, client)
}

func BuildServer(rDecorator *service.RedisDecorator, client *http.Client) error {
	e := gin.Default()

	//TODO: доработать CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{viper.GetString(chirik_config.UserApp)}

	e.Use(cors.New(config))

	InitCreateUserService(e, rDecorator, client)
	//InitAuthServer(e, &c_controller.DIContainer{DBDecorator: dbDecorator, RedisDecorator: redisDecorator})

	err := e.Run()
	if err != nil {
		return err
	}

	return nil
}
