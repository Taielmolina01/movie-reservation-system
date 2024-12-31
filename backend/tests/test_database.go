package tests

import (
	"gorm.io/driver/sqlite"
	"movie-reservation-system/models"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-reservation-system/initializers"
)

func SetUpRouterTest() (*gin.Engine, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connecting to mock database: %w", err)
	}
	
	log.Println("Connected to in-memory SQLite database for testing")
	
	tables := models.GetAllModels()
	for _, t := range tables {
		err = db.AutoMigrate(t)
		if err != nil {
			return nil, fmt.Errorf("Error creating tables in mock database: %w", err)
		}
	}
	
	router := initializers.Init(db)
	return router, nil
}

