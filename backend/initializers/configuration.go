package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

const (
	MESSAGE_ERROR_LOADING_ENV   = "Error loading the environment variables: %v"
	MESSAGE_SUCCESS_LOADING_ENV = "Success loading the environment variables"
)

type Configuration struct {
	Port         string
	DbDsn        string
	JwtAlgorithm string
	JwtSecret    string
}

func LoadConfig() *Configuration {
	envFile := filepath.Join(".", ".env")

	err := godotenv.Load(envFile)

	if err != nil {
		fmt.Errorf(MESSAGE_ERROR_LOADING_ENV, err)
		return nil
	}

	fmt.Println(MESSAGE_SUCCESS_LOADING_ENV)

	return &Configuration{
		Port:         os.Getenv("PORT"),
		DbDsn:        os.Getenv("DB_DSN"),
		JwtAlgorithm: os.Getenv("JWT_ALGORITHM"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
	}
}
