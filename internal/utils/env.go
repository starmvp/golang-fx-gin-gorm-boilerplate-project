package utils

import "os"

func GetEnv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) == 0 {
			return ""
		}
		return defaultValue[0]
	}

	return value
}
