package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"movie-reservation-system/initializers"
	"movie-reservation-system/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
)

func setUpRouterTest() (*gin.Engine, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connecting to mock database: %w", err)
	}

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

func UseRouter(t *testing.T) *gin.Engine {
	router, err := setUpRouterTest()

	if err != nil {
		t.Fatalf("Error setting up test router: %v", err)
	}
	return router
}

func PerformRequest(t *testing.T, router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	var req *http.Request
	var err error

	if body == "" {
		req, err = http.NewRequest(method, path, nil)
	} else {
		req, err = http.NewRequest(method, path, bytes.NewBufferString(body))
	}

	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	return recorder
}