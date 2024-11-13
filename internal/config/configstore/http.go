package configstore

import (
	"log"
)

type HTTPConfig struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	ExposePort string `yaml:"expose_port"`
}

func NewHTTPConfig(store *ConfigStore) *HTTPConfig {
	hc := &HTTPConfig{}
	err := hc.Load(store)
	if err != nil {
		log.Fatalf("FATAL: failed to load http config. err: %+v", err)
	}

	return hc
}

func (hc HTTPConfig) SectionName() string {
	return "http"
}

func (hc *HTTPConfig) Load(store *ConfigStore) error {
	result, err := LoadSection[HTTPConfig](store, hc.SectionName())
	if err != nil {
		return err
	}
	*hc = *result
	store.SetSection(hc.SectionName(), hc)
	return nil
}
