package routers

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"
	"golang-fx-gin-gorm-boilerplate-project/pkg/example/controllers"

	"go.uber.org/zap"
)

func registerDefaultRoutes(
	s *server.Server,
	d *db.DB,
	l *logger.Logger,
) ([]*controller.Controller, error) {
	cl, err := controllers.CreateControllers(d, l)
	if err != nil {
		return nil, err
	}

	rcl := make([]*controller.Controller, 0)
	for _, c := range cl {
		c, err := controller.RegisterController(s, c)
		if err != nil {
			rcl = append(rcl, c)
		} else {
			l.Warn("Failed to register controller", zap.Error(err))
		}
	}

	return rcl, nil
}
