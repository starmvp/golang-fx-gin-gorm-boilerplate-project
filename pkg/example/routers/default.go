package routers

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"
	"golang-fx-gin-gorm-boilerplate-project/pkg/example/controllers"

	"go.uber.org/fx"
)

type RegisterParams struct {
	fx.In

	Server         *server.Server
	PingController *controllers.PingController
}

func registerDefaultRoutes(params RegisterParams) (controller.RegisterControllerResult, error) {
	result, err := controller.RegisterController(controller.RegisterControllerParams{
		Server:     params.Server,
		Controller: &params.PingController.Controller,
	})

	return result, err
}
