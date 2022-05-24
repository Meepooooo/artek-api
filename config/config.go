package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents an application configuration.
type Config struct {
	DB   string
	Port int
}

// Load loads a config from environment variables.
// Returns an empty config and an error if any of the required
// environment variables does not exist.
func Load() (Config, error) {
	godotenv.Load()

	db, exists := os.LookupEnv("DB")
	if !exists {
		return Config{}, errors.New("environment variable DB does not exist")
	}

	portString, exists := os.LookupEnv("PORT")
	if !exists {
		return Config{}, errors.New("environment variable PORT does not exist")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, errors.New("cannot parse environment variable PORT")
	}

	return Config{DB: db, Port: port}, nil
}
