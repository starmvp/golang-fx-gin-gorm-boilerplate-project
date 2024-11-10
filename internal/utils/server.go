package utils

import (
	"fmt"
	"strings"
)

func GetWebserverAddr() string {
	host := Getenv("HOST", "127.0.0.1")
	port := Getenv("PORT", "18080")
	if strings.Contains(port, ":") {
		port = strings.ReplaceAll(port, ":", "")
	}
	return fmt.Sprintf("%s:%s", host, port)
}
