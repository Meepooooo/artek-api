package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Name     string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
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

	return Config{
		DBConfig{
			user,
			password,
			dbName,
		},
	}, nil
}
