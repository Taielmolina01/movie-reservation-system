package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-reservation-system/initializers"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	DB, _ = initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	userController := initializers.Init(DB)

	initializers.SetUpRoutes(r, userController)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my movie reservation system",
		})
	})

	r.Run()
}
