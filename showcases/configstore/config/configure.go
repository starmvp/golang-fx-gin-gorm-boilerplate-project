package config

import (
	"boilerplate/internal/config/configstore"

	"go.uber.org/fx"
)

type ConfigStore = configstore.ConfigStore

var Module = fx.Provide(
	fx.Annotate(configstore.NewConfigStore),
	fx.Annotate(configstore.NewDBConfig),
	fx.Annotate(configstore.NewHTTPConfig),
	fx.Annotate(NewDemoDBConfig),
)
