package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"path/filepath"
)

const (
	MESSAGE_ERROR_LOADING_ENV   = "Error loading the environment variables: %v"
	MESSAGE_SUCCESS_LOADING_ENV = "Success loading the environment variables"
)

func LoadEnvVariables() error {
	envFile := filepath.Join(".", ".env") // Should be relative path, idk why doesn't work

	err := godotenv.Load(envFile)

	if err != nil {
		return fmt.Errorf(MESSAGE_ERROR_LOADING_ENV, err)
	}

	fmt.Println(MESSAGE_SUCCESS_LOADING_ENV)
	return nil
}
