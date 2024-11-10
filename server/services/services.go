package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewHealthCheckService),
		fx.Annotate(NewPingService),
	),
)
