package loaders

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlLoader struct {
}

func (yl YamlLoader) Load(source string) (map[string]any, error) {
	var config map[string]any
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, errors.New("failed to read config file")
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.New("failed to parse config file")
	}
	return config, nil
}
