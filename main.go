package main

import (
	"mylib/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatebaseConfig()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Run(":3000")
}
