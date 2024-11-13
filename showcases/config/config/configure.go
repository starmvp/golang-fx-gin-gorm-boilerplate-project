package config

import (
	"boilerplate/internal/config"

	"go.uber.org/fx"
)

type YamlConfig = config.YamlConfig

var Module = fx.Provide(
	fx.Annotate(config.NewYamlConfig),
	fx.Annotate(config.NewDBConfig),
	fx.Annotate(config.NewHTTPConfig),
	fx.Annotate(NewDemoDBConfig),
)
