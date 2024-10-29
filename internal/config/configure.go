package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB   *DBConfig
	HTTP *HTTPConfig
}

func NewConfig() *Config {
	var ef string
	en := os.Getenv("ENVIRONMENT_ORIGINAL")
	if "" == en {
		en = os.Getenv("ENVIRONMENT")
		if "" == en {
			ef = ".env"
		} else {
			ef = ".env." + en + ".local"
		}
	} else {
		ef = ".env." + en
	}

	err := godotenv.Load(ef)
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
	}
}
