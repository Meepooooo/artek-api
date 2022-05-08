package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN  string
	Port int
}

type DBConfig struct {
	DSN string
}

type APIConfig struct {
	Port int
}

func Load() (Config, error) {
	godotenv.Load()

	dsn, exists := os.LookupEnv("DSN")
	if !exists {
		return Config{}, errors.New("environment variable DSN does not exist")
	}

	portString, exists := os.LookupEnv("PORT")
	if !exists {
		return Config{}, errors.New("environment variable PORT does not exist")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, errors.New("cannot parse environment variable PORT")
	}

	return Config{DSN: dsn, Port: port}, nil
}
