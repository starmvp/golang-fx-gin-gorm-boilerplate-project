package app

import (
	"boilerplate/internal/config"
	"boilerplate/internal/web/server"

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
