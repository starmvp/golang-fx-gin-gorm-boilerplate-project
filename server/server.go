package server

import (
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/utils"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"
	"golang-fx-gin-gorm-boilerplate-project/server/handlers"
	"golang-fx-gin-gorm-boilerplate-project/server/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppServer struct {
	server.Server

	ApiNoAuth   *gin.IRoutes
	ApiNeedAuth *gin.IRoutes
}

func NewAppServer(c *config.Config, d *gorm.DB, l *zap.Logger) *AppServer {
	s := &AppServer{
		Server: *server.NewServer(&c.Config, d, l),
	}
	s.ConfigureRouteGroups()
	return s
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewAppServer),
	),

	services.Module,
	handlers.Module,

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
			s.NoAuth.GET("/api/v1/noauth/ping", handler.Ping())
		},
		func(s *AppServer, handler *handlers.PingHandler) {
			fmt.Println("Server: Configuring routes: need auth ping")
			s.Handlers = append(s.Handlers, handler)
			s.NeedsAuth.GET("/api/v1/needauth/ping", handler.Ping())
		},
	),
)
