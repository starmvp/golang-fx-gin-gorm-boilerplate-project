package configstore

import (
	"fmt"
	"log"
	"os"

	"boilerplate/internal/config/configstore/loaders"

	"gopkg.in/yaml.v3"
)

type moduleConfig interface {
	SectionName() string
}

type ConfigLoader interface {
	Load(source string) (map[string]any, error)
}

type ConfigStore struct {
	raw     map[string]any
	configs map[string]moduleConfig
}

func NewConfigStore() *ConfigStore {
	loaderName := os.Getenv("ENV_LOADER")
	if loaderName == "" {
		loaderName = "yaml"
	}
	envSource := os.Getenv("ENV_SOURCE")
	if envSource == "" {
		envSource = "config.yml"
	}
	var loader ConfigLoader
	switch loaderName {
	case "yaml":
		loader = loaders.YamlLoader{}
	default:
		loader = loaders.YamlLoader{}
	}
	raw, err := loader.Load(envSource)
	if err != nil {
		log.Fatalf("FATAL: failed to load config. err: %+v", err)
	}
	return &ConfigStore{
		raw:     raw,
		configs: make(map[string]moduleConfig),
	}
}

func LoadSection[T any](store *ConfigStore, name string) (*T, error) {
	sectionData, exists := store.raw[name]
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

func (c *ConfigStore) SetSection(name string, section moduleConfig) {
	fmt.Printf("Setting section: %s\n", name)
	c.configs[name] = section
}

func (c *ConfigStore) GetSection(name string) moduleConfig {
	section, exists := c.configs[name]
	if !exists {
		return nil
	}
	return section
}
