package service

import (
	"bytes"
	"chickchirick-bff/internal/controller/c_controller"
	"chickchirick-bff/internal/request"
	"chickchirick-bff/internal/service"
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

// TODO: вынести внутреннюю логику в отдельный сервис
func (uc *UserController) CreateUser(c *gin.Context) {
	var createUserRequest request.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	createUserReqData := make(map[string]string)
	createUserReqData["name"] = createUserRequest.Name
	createUserReqData["surname"] = createUserRequest.Surname
	createUserReqData["login"] = createUserRequest.Login
	createUserReqData["phone"] = createUserRequest.Phone

	createUserReqJson, err := json.Marshal(createUserReqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json marshal error"})
		return
	}

	createUserResp, err := uc.Controller.Client.Post(
		viper.GetString(chirikconfig.UserApp)+"/user",
		"application/json",
		bytes.NewBuffer(createUserReqJson),
	)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("error closing body: ", err.Error())
		}
	}(createUserResp.Body)

	var createUserRespResult map[string]interface{}
	err = json.NewDecoder(createUserResp.Body).Decode(&createUserRespResult)
	if err != nil {
		slog.Error("invalid json: ", err)
	}

	if service.IsInternalServerError(createUserResp.StatusCode) {
		c.JSON(createUserResp.StatusCode, "internal server error")
		return
	}
	if !service.IsStatusCodeOK(createUserResp.StatusCode) {
		c.JSON(createUserResp.StatusCode, createUserRespResult)
		return
	}

	createUPropertyReqData := make(map[string]interface{})
	createUPropertyReqData["user_id"] = createUserRespResult["payload"].(map[string]interface{})["id"]
	if createUserRequest.Email != nil {
		createUPropertyReqData["email"] = *createUserRequest.Email
	}
	if createUserRequest.Password != nil {
		createUPropertyReqData["password"] = *createUserRequest.Password
	}

	createUPropertyReqBody, err := json.Marshal(createUPropertyReqData)
	if err != nil {
		slog.Error("error marshaling request data: ", err.Error())
	}

	createUserPropertyResp, err := uc.Controller.Client.Post(
		viper.GetString(chirikconfig.UserApp)+"/property",
		"application/json",
		bytes.NewBuffer(createUPropertyReqBody), // Теперь здесь []byte
	)

	var createUserPropertyRespResult map[string]interface{}
	err = json.NewDecoder(createUserPropertyResp.Body).Decode(&createUserPropertyRespResult)
	if err != nil {
		slog.Error("invalid json: ", err)
	}

	//TODO: удалить мок
	println(createUserPropertyRespResult)
}
