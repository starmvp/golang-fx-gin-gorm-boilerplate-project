package server

import (
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/utils"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"
	"golang-fx-gin-gorm-boilerplate-project/server/handlers"
	"golang-fx-gin-gorm-boilerplate-project/server/services"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppServer struct {
	server.Server
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
		func(s *AppServer, handler *handlers.HealthCheckHandler) {
			fmt.Println("Server: Configuring routes: health")
			s.Handlers = append(s.Handlers, handler)
			s.NoAuth.GET("/health", handler.HealthCheck())
		},
	),
)
