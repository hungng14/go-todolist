package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var GO_ENV = os.Getenv("GO_ENV")
var DB_DSN = "host=localhost user=postgres password=123456a@ dbname=todolist port=5432 sslmode=disable"
var DB_DSN_TESTING = "host=localhost user=postgres password=123456a@ dbname=todolist_testing port=5432 sslmode=disable"

// Config loads the correct environment file based on the environment (production or test)
func Config(key string) string {
	// Check if the environment is set to "test"
	env := os.Getenv("GO_ENV")

	var envFile string

	if env == "test" {
		envFile = ".env.test"
	} else {
		envFile = ".env"
	}

	// Load the appropriate .env file
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("Error loading %s file: %s\n", envFile, err)
	}

	// Return the value for the requested key
	return os.Getenv(key)
}
