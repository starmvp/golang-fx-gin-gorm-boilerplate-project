package config

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/utils"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	User           string
	Password       string
	Driver         string
	Name           string
	Host           string
	Port           string
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBConnMaxLife  int
}

// TODO: provide default values support
const (
	defaultDBMaxOpenConns = "10"
	defaultDBMaxIdleConns = "5"
	defaultDBConnMaxLife  = "3600"
)

func LoadDBConfig() *DBConfig {
	maxOpenConns, err := strconv.Atoi(
		utils.Getenv(
			"DB_MAX_OPEN_CONNS",
			defaultDBMaxOpenConns,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	maxIdleConns, err := strconv.Atoi(
		utils.Getenv(
			"DB_MAX_IDLE_CONNS",
			defaultDBMaxIdleConns,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	connMaxLife, err := strconv.Atoi(
		utils.Getenv(
			"DB_CONN_MAX_LIFE",
			defaultDBConnMaxLife,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	return &DBConfig{
		User:           os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASSWORD"),
		Driver:         os.Getenv("DB_DRIVER"),
		Name:           os.Getenv("DB_NAME"),
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		DBMaxOpenConns: maxOpenConns,
		DBMaxIdleConns: maxIdleConns,
		DBConnMaxLife:  connMaxLife,
	}
}
