package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie-reservation-system/initializers"
	"os"
)

func main() {

	db, err := initializers.ConnectDB()

	if err != nil {
		log.Fatal()
	}

	r := initializers.Init(db)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my movie reservation system",
		})
	})

	errEnv := initializers.LoadEnvVariables()

	if errEnv != nil {
		log.Fatal()
	}

	r.Run(":" + os.Getenv("PORT"))
}
