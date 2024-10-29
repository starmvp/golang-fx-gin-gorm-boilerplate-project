package app

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/fx"
)

type App struct {
	server *server.Server
}

type AppParams struct {
	fx.In

	Logger *logger.Logger
}

type AppResult struct {
	fx.Out

	App *App
}

func NewApp(params AppParams) (AppResult, error) {
	result, err := server.New(server.ServerParams{
		Logger: params.Logger,
	})
	if err != nil {
		return AppResult{}, err
	}
	app := App{
		server: result.Server,
	}

	return AppResult{App: &app}, nil
}
