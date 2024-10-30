package controllers

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/db"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PingController struct {
	controller.Controller
}

func NewPingController(
	Db *db.DB,
	Logger *zap.Logger,
) (*PingController, error) {
	pc := &PingController{}
	handlers := make([]controller.Handler, 0)
	h, err := controller.NewHandler(controller.HandlerParams{
		Method:  controller.Get,
		Pattern: "/ping",
		Handler: pc.handler,
	})
	if err != nil {
		return nil, err
	}
	handlers = append(handlers, h.Handler)

	c, err := controller.NewController(Db, Logger, handlers)

	if err != nil {
		pc = nil
	} else {
		pc.Controller = *c
	}

	return pc, err
}

func (e *PingController) handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
