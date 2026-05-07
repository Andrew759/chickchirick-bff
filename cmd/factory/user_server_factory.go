package factory

import (
	"chickchirick-bff/cmd/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserServer struct{}

func InitUserServer(e *gin.Engine, rDecorator *service.RedisDecorator, client *http.Client) {

}
