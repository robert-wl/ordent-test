package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string
}

var config *Config

func Load() *Config {
	if config != nil {
		return config
	}

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	config = &Config{
		ServerPort: getEnv("SERVER_PORT"),
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
