package controller

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/router"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Controller struct {
	Db     *db.DB
	Logger *logger.Logger

	Handlers []Handler
}

type ControllerParams struct {
	fx.In

	Db       *db.DB
	Logger   *logger.Logger
	Handlers []Handler
}

type ControllerResult struct {
	fx.Out

	Controller *Controller
}

func NewController(params ControllerParams) (ControllerResult, error) {
	l := params.Logger
	if l == nil {
		l = zap.NewNop()
	}
	c := &Controller{
		Db:       params.Db,
		Logger:   l,
		Handlers: params.Handlers,
	}

	return ControllerResult{Controller: c}, nil
}

type RegisterControllerParams struct {
	fx.In

	Server     *server.Server
	Controller *Controller
}

type RegisterControllerResult struct {
	fx.Out

	Controller *Controller
}

func RegisterController(
	params RegisterControllerParams,
) (RegisterControllerResult, error) {
	handlers := params.Controller.Handlers
	for _, handler := range handlers {
		_, err := router.RegisterRoute(router.RegisterRouteParams{
			Server:  params.Server,
			Method:  handler.Method(),
			Pattern: handler.Pattern(),
			Handler: handler.Handler(),
		})
		if err != nil {
			return RegisterControllerResult{Controller: nil}, err
		}
	}

	return RegisterControllerResult{Controller: params.Controller}, nil
}

var Module = fx.Provide(
	NewController,
	RegisterController,
)
