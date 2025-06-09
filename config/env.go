package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:     getENV("DB_HOST", "localhost"),
		Port:     getENV("DB_PORT", "5432"),
		User:     getENV("DB_USER", "postgres"),
		Password: getENV("DB_PASSWORD", "your-password"),
		DBName:   getENV("DB_NAME", "restapitutorial"),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
