package configstore

import (
	"log"
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
	defaultDBMaxOpenConns = 10
	defaultDBMaxIdleConns = 5
	defaultDBConnMaxLife  = 3600
)

func NewDBConfig(store *ConfigStore) *DBConfig {
	dbc := &DBConfig{}
	err := dbc.Load(store)
	if err != nil {
		log.Fatalf("FATAL: failed to load db config. err: %+v", err)
	}
	if dbc.DBMaxOpenConns == 0 {
		dbc.DBMaxOpenConns = defaultDBMaxOpenConns
	}
	if dbc.DBMaxIdleConns == 0 {
		dbc.DBMaxIdleConns = defaultDBMaxIdleConns
	}
	if dbc.DBConnMaxLife == 0 {
		dbc.DBConnMaxLife = defaultDBConnMaxLife
	}

	return dbc
}

func (dbc DBConfig) SectionName() string {
	return "db"
}

func (dbc *DBConfig) Load(store *ConfigStore) error {
	result, err := LoadSection[DBConfig](store, dbc.SectionName())
	if err != nil {
		return err
	}
	*dbc = *result
	store.SetSection(dbc.SectionName(), dbc)
	return nil
}
