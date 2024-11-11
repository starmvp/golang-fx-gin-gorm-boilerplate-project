package server

import (
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/internal/utils"
	"golang-fx-gin-gorm-boilerplate-project/server/handlers"
	"golang-fx-gin-gorm-boilerplate-project/server/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func (s *AppServer) ConfigureRouteGroups() {
	fmt.Println("AppServer: Configuring route groups")

	s.ApiNoAuth = s.NoAuth.Group("/api/v1/noauth")
	s.ApiNoAuth.Use(func(c *gin.Context) {
	})
	fmt.Println("AppServer: Configuring route groups NoAuth. s.ApiNoAuth: ", s.ApiNoAuth)

	s.ApiNeedsAuth = s.NeedsAuth.Group("/api/v1/needauth")
	s.ApiNeedsAuth.Use(func(c *gin.Context) {
	})
	fmt.Println("AppServer: Configuring route groups NeedsAuth. s.ApiNeedAuth: ", s.ApiNeedsAuth)
}

var routesModule = fx.Options(
	fx.Invoke(
		func(s *AppServer, logger *zap.Logger) {
			logger.Debug("AppServer module invoked")
			go func() {
				_ = s.Run(utils.GetWebserverAddr())
			}()
		},
		func(s *services.PingService) {
			fmt.Println("Server: Configuring services: ping: ", s)
		},
		func(s *AppServer, handler *handlers.HealthCheckHandler) {
			fmt.Println("Server: Configuring routes: health")
			s.Handlers = append(s.Handlers, handler)
			s.NoAuth.GET("/health", handler.HealthCheck())
		},
		func(s *AppServer, handler *handlers.PingHandler) {
			fmt.Println("Server: Configuring routes: no auth ping")
			s.Handlers = append(s.Handlers, handler)
			s.ApiNoAuth.GET("/ping", handler.Ping())
		},
		func(s *AppServer, handler *handlers.PingHandler) {
			fmt.Println("Server: Configuring routes: need auth ping")
			s.Handlers = append(s.Handlers, handler)
			s.ApiNeedsAuth.GET("/ping", handler.Ping())
		},
	),
)
