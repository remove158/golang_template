package handler

import (
	"mylib/common/response"
	"mylib/dto"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var user dto.PostUserDTO
	err := ctx.ShouldBind(&user)
	if err != nil {
		response := response.BuildErrorResponse("Error to post user", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "Created user", user)
	ctx.JSON(http.StatusCreated, response)
}
