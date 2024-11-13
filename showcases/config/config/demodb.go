package config

import (
	"boilerplate/internal/config"
	"log"
)

type DemoDBConfig struct {
	config.DBConfig
}

func NewDemoDBConfig(config *config.YamlConfig) *DemoDBConfig {
	ddbc := &DemoDBConfig{}
	err := ddbc.Load(config)
	if err != nil {
		log.Fatalf("FATAL: failed to load db config. err: %+v", err)
	}

	return ddbc
}

func (ddbc DemoDBConfig) SectionName() string {
	return "demo-db"
}

func (ddbc *DemoDBConfig) Load(c *config.YamlConfig) error {
	result, err := config.LoadSection[config.DBConfig](c, ddbc.SectionName())
	if err != nil {
		return err
	}
	ddbc.DBConfig = *result
	c.SetSection(ddbc.SectionName(), ddbc)
	return nil
}
