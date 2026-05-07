package service

import (
	"chickchirick-bff/internal/controller/c_controller"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Controller c_controller.Controller
}

func (uc *UserController) RegisterRouter() {
	e := uc.Controller.E

	group := e.Group("/user")
	group.POST("/login", uc.GetUser)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
}
