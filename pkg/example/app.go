package example

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/app"
	e "golang-fx-gin-gorm-boilerplate-project/internal/errors"

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

var Module = fx.Provide(
	fx.Annotate(NewExampleApp),
)
