package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string

	PostgresUsername string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
}

var config *Config

func Get() *Config {
	if config != nil {
		return config
	}

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	config = &Config{
		ServerPort: getEnv("SERVER_PORT"),

		PostgresUsername: getEnv("POSTGRES_USERNAME"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD"),
		PostgresHost:     getEnv("POSTGRES_HOST"),
		PostgresPort:     getEnv("POSTGRES_PORT"),
		PostgresDB:       getEnv("POSTGRES_DB"),
	}

	return config
}
func getEnv(key string) string {
	v := os.Getenv(key)

	if v == "" {
		panic("missing env var: " + key)
	}

	return v
}
