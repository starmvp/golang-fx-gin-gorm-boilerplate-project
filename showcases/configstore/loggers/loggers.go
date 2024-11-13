package loggers

import (
	"boilerplate/internal/logger"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(logger.NewGormLogger),
)
