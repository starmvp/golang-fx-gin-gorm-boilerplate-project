package config

import (
	"boilerplate/internal/config/configstore"
	"log"
)

type DemoDBConfig struct {
	configstore.DBConfig
}

func NewDemoDBConfig(store *configstore.ConfigStore) *DemoDBConfig {
	ddbc := &DemoDBConfig{}
	err := ddbc.Load(store)
	if err != nil {
		log.Fatalf("FATAL: failed to load db config. err: %+v", err)
	}

	return ddbc
}

func (ddbc DemoDBConfig) SectionName() string {
	return "demo-db"
}

func (ddbc *DemoDBConfig) Load(store *configstore.ConfigStore) error {
	result, err := configstore.LoadSection[configstore.DBConfig](store, ddbc.SectionName())
	if err != nil {
		return err
	}
	ddbc.DBConfig = *result
	store.SetSection(ddbc.SectionName(), ddbc)
	return nil
}
