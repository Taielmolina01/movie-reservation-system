package main

import (
	"github.com/gin-gonic/gin"
	"movie-reservation-system/initializers"
	"log"
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

	r.Run()
}
