package routers

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func registerRoutes(
	s *server.Server,
	d *db.DB,
	l *zap.Logger,
) ([]*controller.Controller, error) {
	cl := make([]*controller.Controller, 0)
	drc, _ := registerDefaultRoutes(s, d, l)
	cl = append(cl, drc...)
	return cl, nil
}

var Module = fx.Provide(
	fx.Annotate(registerRoutes),
)
