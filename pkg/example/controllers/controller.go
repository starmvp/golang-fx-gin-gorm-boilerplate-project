package controllers

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"

	"go.uber.org/zap"
)

func CreateControllers(
	d *db.DB,
	l *logger.Logger,
) ([]*controller.Controller, error) {
	cl := make([]*controller.Controller, 0)

	c, err := NewPingController(d, l)
	if err != nil {
		l.Error("Failed to create controller", zap.Error(err))
	}

	cl = append(cl, &c.Controller)

	return cl, nil
}
