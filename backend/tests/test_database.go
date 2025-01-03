package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"movie-reservation-system/initializers"
	"movie-reservation-system/models"
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

	config := &initializers.Configuration{
		Port:         "3000",
		DbDsn:        "host=localhost user=postgres password=taiel0101 port=5432 sslmode=disable dbname=movie-system-db",
		JwtAlgorithm: "HS256",
		JwtSecret:    "ASDADASD",
	}

	router := initializers.Init(db, config)

	gin.SetMode(gin.ReleaseMode)

	return router, nil
}
