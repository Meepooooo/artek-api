package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig
	API APIConfig
}

type DBConfig struct {
	Location string
}

type APIConfig struct {
	Port int
}

func Load() (Config, error) {
	config := Config{}

	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	loc, exists := os.LookupEnv("DB_LOCATION")
	if !exists {
		return Config{}, errors.New("environment variable DB_LOCATION does not exist")
	}

	config.DB = DBConfig{Location: loc}

	data, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
