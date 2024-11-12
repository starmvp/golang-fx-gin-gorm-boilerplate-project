package example

import (
	"boilerplate/internal/app"
	"boilerplate/internal/errors"
	"boilerplate/pkg/example/server"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ExampleApp struct {
	app.App
}

func NewExampleApp(
	app *app.App,
	logger *zap.Logger,
) (*ExampleApp, error) {

	if app == nil {
		return nil, errors.ErrNilApp
	}
	eapp := ExampleApp{App: *app}

	return &eapp, nil
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewExampleApp),
	),
	server.Module,
)
