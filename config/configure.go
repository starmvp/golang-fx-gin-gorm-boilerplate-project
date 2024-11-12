package config

import (
	"boilerplate/internal/config"

	"go.uber.org/fx"
)

type Config struct {
	config.Config
}

func NewConfig() *Config {
	return &Config{
		Config: *config.NewConfig(),
	}
}

func GetDBConfig(c *Config) *config.DBConfig {
	return c.Config.DB
}

func GetHTTPConfig(c *Config) *config.HTTPConfig {
	return c.Config.HTTP
}

var Module = fx.Provide(
	fx.Annotate(NewConfig),
	fx.Annotate(GetDBConfig),
	fx.Annotate(GetHTTPConfig),
)
