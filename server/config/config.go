package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

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
