package service

import (
	"bytes"
	"chickchirick-bff/internal/controller/c_controller"
	"chickchirick-bff/internal/request"
	chirikconfig "chickchirick-bff/pkg/chirick_config"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserController struct {
	Controller c_controller.Controller
}

func (uc *UserController) RegisterRouter() {
	e := uc.Controller.E

	group := e.Group("/user")
	group.POST("/create", uc.CreateUser)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var createUserRequest request.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	jsonData, err := json.Marshal(createUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json marshal error"})
		return
	}

	createUserResp, err := uc.Controller.DI.Client.Post(
		viper.GetString(chirikconfig.UserApp)+"/user",
		"application/json",
		bytes.NewBuffer([]byte(jsonData)),
	)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("error closing body: ", err.Error())
		}
	}(createUserResp.Body)

}
