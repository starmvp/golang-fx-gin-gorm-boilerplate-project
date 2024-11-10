package utils

import "os"

// wrap of os.Getenv to support defaultValue
func Getenv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) == 0 {
			return ""
		}
		return defaultValue[0]
	}

	return value
}
