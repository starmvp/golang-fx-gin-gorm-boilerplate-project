package handlers

import (
	"golang-fx-gin-gorm-boilerplate-project/server/services"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	service *services.HealthCheckService
}

func (h HealthCheckHandler) Name() string {
	return "health-check-handler"
}

func NewHealthCheckHandler(service *services.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		service: service,
	}
}

func (handler *HealthCheckHandler) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive",
		})
	}
}
