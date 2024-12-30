package main

import (
	"github.com/gin-gonic/gin"
	"movie-reservation-system/initializers"
)

func main() {
	r := initializers.Init()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my movie reservation system",
		})
	})

	r.Run()
}
