package handler

import (
	"github.com/gin-gonic/gin"
	"mylib/common/response"
	"net/http"
)

type UserHandler interface {
	// add userhandler method here
	GetUser(ctx *gin.Context)
	PostUser(ctx *gin.Context)
}

type userHandler struct {
	// add userhanlder service here
}

func NewUserHandler() UserHandler {
	// add contructor here
	return &userHandler{}
}

func (u *userHandler) GetUser(ctx *gin.Context) {
	response := response.BuildResponse(true, "OK!", nil)
	ctx.JSON(http.StatusOK, response)
}
func (u *userHandler) PostUser(ctx *gin.Context) {
	response := response.BuildErrorResponse("Error", "Faile to create user", nil)
	ctx.JSON(http.StatusCreated, response)
}
