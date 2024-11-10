package handlers

import (
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PingHandler struct {
	DB      *gorm.DB
	Logger  *zap.Logger
	Service *services.PingService
}

func (h PingHandler) Name() string {
	return "ping-handler"
}

func NewPingHandler(
	service *services.PingService,
	logger *zap.Logger,
) *PingHandler {
	fmt.Println(">>> NewPingHandler")
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
