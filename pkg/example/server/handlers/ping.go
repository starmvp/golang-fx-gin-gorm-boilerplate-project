package handlers

import (
	"boilerplate/pkg/example/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PingHandler struct {
	Service *services.PingService
	Logger  *zap.Logger
}

func (h PingHandler) Name() string {
	return "ping-handler"
}

func NewPingHandler(
	service *services.PingService,
	logger *zap.Logger,
) *PingHandler {
	pc := &PingHandler{
		Logger:  logger,
		Service: service,
	}
	return pc
}

func (e *PingHandler) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
