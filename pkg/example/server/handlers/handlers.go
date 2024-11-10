package handlers

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotate(NewPingHandler),
)
