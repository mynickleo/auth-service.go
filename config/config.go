package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DB         string
	SecretKey  string
}

var AppConfig *Config

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error initializaition config")
	}

	AppConfig = &Config{
		Port:       getEnv("PORT", "3000"),
		DBUser:     getEnv("POSTGRES_USER", ""),
		DBPassword: getEnv("POSTGRES_PASSWORD", ""),
		DBHost:     getEnv("POSTGRES_HOST", ""),
		DBPort:     getEnv("POSTGRES_PORT", ""),
		DB:         getEnv("POSTGRES_DB", ""),
		SecretKey:  getEnv("SECRET_KEY", ""),
	}

	return nil
}
