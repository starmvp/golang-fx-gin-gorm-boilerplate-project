package loaders

import (
	"encoding/json"
	"errors"
	"os"
)

type JsonLoader struct {
}

func (l JsonLoader) Name() string {
	return "json"
}

func (l JsonLoader) Load(source string) (map[string]any, error) {
	var config map[string]any
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	expandedData := expandEnvWithDefaults(string(data))

	if err := json.Unmarshal([]byte(expandedData), &config); err != nil {
		return nil, errors.New("failed to parse config file")
	}
	return config, nil
}
