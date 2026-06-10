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

func (uc *UserController) CreateUser(c *gin.Context) {
	createUserRequest, createUserRespResult, sentErr := uc.createBasicUser(c)
	if sentErr {
		return
	}

	uc.createUserMeta(createUserRespResult)
	uc.createUserProperty(createUserRequest, createUserRespResult)

	//TODO: удалить мок
	//println(createUserPropertyRespResult)
}

func (uc *UserController) createBasicUser(c *gin.Context) (
	createUserRequest request.CreateUserRequest,
	createUserRespResult map[string]any,
	sentErr bool,
) {
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
		sentErr = true
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

	err = json.NewDecoder(createUserResp.Body).Decode(&createUserRespResult)
	if err != nil {
		slog.Error("invalid json: ", err)
	}

	if service.IsInternalServerError(createUserResp.StatusCode) {
		sentErr = true
		c.JSON(createUserResp.StatusCode, "internal server error")
		return
	}
	if !service.IsStatusCodeOK(createUserResp.StatusCode) {
		sentErr = true
		c.JSON(createUserResp.StatusCode, createUserRespResult)
		return
	}

	return
}

func (uc *UserController) createUserMeta(createUserRespResult map[string]any) {
	userId := createUserRespResult["payload"].(map[string]any)["id"]

	createUMetaReqData := make(map[string]any)
	createUMetaReqData["user_id"] = userId

	createUMetaReqJson, err := json.Marshal(createUMetaReqData)
	if err != nil {
		slog.Error("json marshal error: ", err.Error())
	}

	createUMetaResp, err := uc.Controller.Client.Post(
		viper.GetString(chirikconfig.UserApp)+"/meta",
		"application/json",
		bytes.NewBuffer(createUMetaReqJson),
	)

	var createUserMetaRespResult map[string]any
	err = json.NewDecoder(createUMetaResp.Body).Decode(&createUserMetaRespResult)
	if err != nil {
		slog.Error("invalid json: ", err)
	}
}

func (uc *UserController) createUserProperty(createUserRequest request.CreateUserRequest, createUserRespResult map[string]any) {
	createUPropertyReqData := make(map[string]any)
	createUPropertyReqData["user_id"] = createUserRespResult["payload"].(map[string]any)["id"]
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

	var createUserPropertyRespResult map[string]any
	err = json.NewDecoder(createUserPropertyResp.Body).Decode(&createUserPropertyRespResult)
	if err != nil {
		slog.Error("invalid json: ", err)
	}
}
