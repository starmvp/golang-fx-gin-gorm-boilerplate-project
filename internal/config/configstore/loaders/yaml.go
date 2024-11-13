package loaders

import (
	"errors"
	"os"
	"regexp"

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

	expandedData := expandEnvWithDefaults(string(data))

	if err := yaml.Unmarshal([]byte(expandedData), &config); err != nil {
		return nil, errors.New("failed to parse config file")
	}
	return config, nil
}

func expandEnvWithDefaults(data string) string {
	re := regexp.MustCompile(`\${(\w+)(:-([^}]*))?}`)

	return re.ReplaceAllStringFunc(data, func(m string) string {
		matches := re.FindStringSubmatch(m)
		varName := matches[1]
		defaultValue := matches[3]

		if val, exists := os.LookupEnv(varName); exists {
			return val
		}
		return defaultValue
	})
}
