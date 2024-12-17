package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"fmt"
    "os"
)

var db *gorm.DB

const (
	MESSAGE_ERROR_CONNECTING_DB = "Error connecting to the database"
	MESSAGE_ERROR_CREATING_DB = "Database already exists or had ocurred an error creating it"
	MESSAGE_SUCCESS_CREATING_DB = "Success creating the database"
)

func ConnectDB() (*gorm.DB, error) {
    err := LoadEnvVariables()
	if err != nil {
		return nil, err
	}

	dsn := os.Getenv("DB_DSN")
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(MESSAGE_ERROR_CONNECTING_DB, err)
    }

	createDatabaseIfNotExist(dsn)

	db, err = gorm.Open(postgres.Open(dsn+" dbname=movie-system-db"), &gorm.Config{})
	if err != nil {
		log.Fatal(MESSAGE_ERROR_CONNECTING_DB, err)
		return nil, err
	}

	return db, nil
}

func createDatabaseIfNotExist(dsn string) {
    conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(MESSAGE_ERROR_CONNECTING_DB, err)
    }

    err = conn.Exec("CREATE DATABASE movie-system-db").Error
    if err != nil {
        fmt.Println(MESSAGE_ERROR_CREATING_DB, err)
    } else {
        fmt.Println(MESSAGE_SUCCESS_CREATING_DB)
    }
}