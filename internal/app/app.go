package app

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/fx"
)

type App struct {
	Server *server.Server
}

func NewApp(
	Server *server.Server,
	Config *config.Config,
	Logger *logger.Logger,
) (*App, error) {
	app := App{
		Server: Server,
	}

	return &app, nil
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewApp),
	),
)
