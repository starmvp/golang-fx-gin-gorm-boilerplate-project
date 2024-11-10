package handlers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewHealthCheckHandler),
		fx.Annotate(NewPingHandler),
	),
)
