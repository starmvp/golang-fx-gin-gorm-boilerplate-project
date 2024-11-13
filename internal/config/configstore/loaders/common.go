package loaders

import (
	"os"
	"regexp"
)

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
