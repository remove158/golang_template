package handler

import (
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}
func (u *userHandler) PostUser(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
