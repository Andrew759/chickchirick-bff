package c_controller

import (
	"chickchirick-bff/cmd/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DIContainer struct {
	RedisDecorator *service.RedisDecorator
	*http.Client
}

type Controller struct {
	E  *gin.Engine
	DI *DIContainer
}

type RequestHandler interface {
	RegisterRoutes()
}

type ControllerInterface interface {
	RequestHandler
}
