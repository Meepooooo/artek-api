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
	User     string
	Password string
	Name     string
}

type APIConfig struct {
	Port int
}

func Load() (Config, error) {
	config := Config{}

	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return Config{}, errors.New("environment variable DB_USER does not exist")
	}

	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return Config{}, errors.New("environment variable DB_PASSWORD does not exist")
	}

	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return Config{}, errors.New("environment variable DB_NAME does not exist")
	}

	config.DB = DBConfig{
		user,
		password,
		dbName,
	}

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
