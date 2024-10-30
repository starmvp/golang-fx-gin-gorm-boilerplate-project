package controller

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/router"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Controller struct {
	Db     *db.DB
	Logger *zap.Logger

	Handlers []Handler
}

func NewController(
	Db *db.DB,
	Logger *zap.Logger,
	Handlers []Handler,
) (*Controller, error) {
	l := Logger
	if l == nil {
		l = zap.NewNop()
	}
	c := Controller{
		Db:       Db,
		Logger:   l,
		Handlers: Handlers,
	}

	return &c, nil
}

func RegisterController(
	s *server.Server,
	c *Controller,
) (*Controller, error) {
	handlers := c.Handlers
	for _, handler := range handlers {
		_, err := router.RegisterRoute(
			s,
			handler.Method(),
			handler.Pattern(),
			handler.Handler(),
		)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

var Module = fx.Provide(
	fx.Annotate(NewController),
	fx.Annotate(RegisterController),
)
