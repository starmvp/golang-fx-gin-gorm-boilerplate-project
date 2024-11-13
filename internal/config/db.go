package config

import (
	"boilerplate/internal/utils"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Driver         string `yaml:"driver"`
	Name           string `yaml:"name"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	DBMaxOpenConns int    `yaml:"max_open_conns"`
	DBMaxIdleConns int    `yaml:"max_idle_conns"`
	DBConnMaxLife  int    `yaml:"conn_max_life"`
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

// for yaml config

func NewDBConfig(c *YamlConfig) *DBConfig {
	dbc := &DBConfig{}
	err := dbc.Load(c)
	if err != nil {
		log.Fatalf("FATAL: failed to load db config. err: %+v", err)
	}

	return dbc
}

func (dbc DBConfig) SectionName() string {
	return "db"
}

func (dbc *DBConfig) Load(c *YamlConfig) error {
	result, err := LoadSection[DBConfig](c, dbc.SectionName())
	if err != nil {
		return err
	}
	*dbc = *result
	c.SetSection(dbc.SectionName(), dbc)
	return nil
}
