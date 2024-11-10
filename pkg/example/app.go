package example

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/app"
	e "golang-fx-gin-gorm-boilerplate-project/internal/errors"
	"golang-fx-gin-gorm-boilerplate-project/pkg/example/server"

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
		return nil, e.ErrNilApp
	}
	e := ExampleApp{App: *app}

	return &e, nil
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewExampleApp),
	),
	server.Module,
)
