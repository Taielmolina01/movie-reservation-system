package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie-reservation-system/initializers"
	"movie-reservation-system/configuration"
)

func main() {

	config := configuration.LoadConfig()

	db, err := initializers.ConnectDB(config)

	if err != nil {
		log.Fatal()
	}

	r := initializers.Init(db, config)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my movie reservation system",
		})
	})

	r.Run(":" + config.Port)
}
