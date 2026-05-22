package factory

import (
	"chickchirick-bff/internal/controller/c_controller"
	internalService "chickchirick-bff/internal/controller/service"

	"github.com/gin-gonic/gin"
)

type UserServer struct{}

func InitUserServer(e *gin.Engine, diContainer *c_controller.DIContainer) {
	userServer := UserServer{}

	userServer.initCreateUserService(e, diContainer)

}

func (u UserServer) initCreateUserService(e *gin.Engine, diContainer *c_controller.DIContainer) internalService.UserController {
	createUserService := internalService.UserController{
		Controller: c_controller.Controller{
			E:  e,
			DI: diContainer,
		},
	}
	createUserService.RegisterRouter()

	return createUserService
}
