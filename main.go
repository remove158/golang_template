package main

import (
	"mylib/config"
	"mylib/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB            = config.SetupDatebaseConfig()
	userHandler handler.UserHandler = handler.NewUserHandler()
)

func main() {
	// to close db connection when program executed
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	userRoute := r.Group("users")
	{
		userRoute.GET("", userHandler.GetUser)
		userRoute.POST("", userHandler.PostUser)
	}
	r.Run(":3000")
}
