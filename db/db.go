package db

import (
	"boilerplate/internal/db"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(db.New),
	),
)
