package service

import (
	"chickchirick-bff/internal/controller/c_controller"
	"chickchirick-bff/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Controller c_controller.Controller
}

func (uc *UserController) RegisterRouter() {
	e := uc.Controller.E

	group := e.Group("/user")
	group.POST("/create", uc.GetUser)
	group.POST("")
}

func (uc *UserController) GetUser(c *gin.Context) {
	var createUserRequest request.CreateUserRequest
	if err := c.ShouldBindJSON(&ctr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	uc.Controller.DI.Client.Post()
}
