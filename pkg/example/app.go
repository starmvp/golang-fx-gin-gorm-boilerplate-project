package example

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/app"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"

	"go.uber.org/fx"
)

type ExampleApp struct {
	app.App
}

type ExampleAppParams struct {
	fx.In

	Logger *logger.Logger
}

type ExampleAppResult struct {
	fx.Out

	App *ExampleApp
}

func NewExampleApp(logger *logger.Logger) (ExampleAppResult, error) {

	result, err := app.NewApp(app.AppParams{
		Logger: logger,
	})
	if err != nil {
		return ExampleAppResult{}, err
	}
	e := ExampleApp{App: *result.App}

	return ExampleAppResult{
		App: &e,
	}, nil
}

var Module = fx.Provide(NewExampleApp)
