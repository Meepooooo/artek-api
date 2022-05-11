package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBLocation string
	Port       int
}

func Load() (Config, error) {
	godotenv.Load()

	loc, exists := os.LookupEnv("DB_LOCATION")
	if !exists {
		return Config{}, errors.New("environment variable DB_LOCATION does not exist")
	}

	portString, exists := os.LookupEnv("PORT")
	if !exists {
		return Config{}, errors.New("environment variable PORT does not exist")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, errors.New("cannot parse environment variable PORT")
	}

	return Config{DBLocation: loc, Port: port}, nil
}
