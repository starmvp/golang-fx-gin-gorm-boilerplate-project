package app

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/zap"
)

type App struct {
	Server *server.Server
}

// for multiple app instances
func NewApp(
	Config *config.Config,
	Server *server.Server,
	Logger *zap.Logger,
) (*App, error) {
	app := App{
		Server: Server,
	}

	return &app, nil
}
