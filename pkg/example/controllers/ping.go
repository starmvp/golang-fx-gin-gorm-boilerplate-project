package controllers

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PingController struct {
	controller.Controller
}

type PingControllerParams struct {
	fx.In

	Db     *db.DB
	Logger *logger.Logger
}

type PingControllerResult struct {
	fx.Out

	Controller *PingController
}

func NewPingController(params PingControllerParams) (PingControllerResult, error) {
	pc := &PingController{}
	handlers := make([]controller.Handler, 0)
	h, err := controller.NewHandler(controller.HandlerParams{
		Method:  controller.Get,
		Pattern: "/ping",
		Handler: pc.handler,
	})
	if err != nil {
		return PingControllerResult{Controller: nil}, err
	}
	handlers = append(handlers, h.Handler)

	result, err := controller.NewController(controller.ControllerParams{
		Db:       params.Db,
		Logger:   params.Logger,
		Handlers: handlers,
	})

	if err != nil {
		pc = nil
	} else {
		pc.Controller = *result.Controller
	}

	return PingControllerResult{Controller: pc}, err
}

func (e *PingController) handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
