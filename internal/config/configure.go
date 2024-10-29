package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
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

	fmt.Println("Loading environment file:", ef)

	err := godotenv.Load(ef)
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		// TODO: split in the manner of fx, and make it flexible to extend
		DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
	}
}

var Module = fx.Provide(NewConfig)
