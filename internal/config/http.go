package config

import (
	"log"
	"os"
)

type HTTPConfig struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	ExposePort string `yaml:"expose_port"`
}

func LoadHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		ExposePort: os.Getenv("EXPOSE_PORT"),
	}
}

// for yaml config

func NewHTTPConfig(config *YamlConfig) *HTTPConfig {
	hc := &HTTPConfig{}
	err := hc.Load(config)
	if err != nil {
		log.Fatalf("FATAL: failed to load http config. err: %+v", err)
	}

	return hc
}

func (hc HTTPConfig) SectionName() string {
	return "http"
}

func (hc *HTTPConfig) Load(c *YamlConfig) error {
	result, err := LoadSection[HTTPConfig](c, hc.SectionName())
	if err != nil {
		return err
	}
	*hc = *result
	c.SetSection(hc.SectionName(), hc)
	return nil
}
