package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ModuleConfig interface {
	SectionName() string
}

type YamlConfig struct {
	raw     map[string]any
	configs map[string]ModuleConfig
}

func NewYamlConfig() *YamlConfig {
	var config map[string]any
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("FATAL: failed to read config file. err: %+v", err)
		return nil
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		// configure error, MUST exit with fatal error message
		log.Fatalf("FATAL: failed to parse config file. err: %+v", err)
		return nil
	}
	return &YamlConfig{
		raw:     config,
		configs: make(map[string]ModuleConfig),
	}
}

func LoadSection[T any](config *YamlConfig, name string) (*T, error) {
	sectionData, exists := config.raw[name]
	if !exists {
		return nil, nil // Section not found, return nil
	}

	// Marshal the section to YAML and unmarshal into the target struct
	data, err := yaml.Marshal(sectionData)
	if err != nil {
		return nil, err
	}

	var result T
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *YamlConfig) SetSection(name string, section ModuleConfig) {
	fmt.Printf("Setting section: %s\n", name)
	c.configs[name] = section
}

func (c *YamlConfig) GetSection(name string) ModuleConfig {
	section, exists := c.configs[name]
	if !exists {
		return nil
	}
	return section
}
