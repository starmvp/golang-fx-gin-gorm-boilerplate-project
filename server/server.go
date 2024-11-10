package server

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"golang-fx-gin-gorm-boilerplate-project/internal/utils"
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppServer struct {
	server.Server
}

func NewAppServer(c *config.Config, d *gorm.DB, l *zap.Logger) *AppServer {
	s := &AppServer{
		Server: *server.NewServer(c, d, l),
	}
	s.Server.ConfigureRouteGroups()
	return s
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewAppServer),
	),
	fx.Invoke(
		func(s *AppServer, logger *zap.Logger) {
			logger.Debug("AppServer module invoked")
			go func() {
				_ = s.Server.Run(utils.GetWebserverAddr())
			}()
		},
		// func(s *AppServer, handler *handlers.HealthCheckHandler) {
		// 	fmt.Println("Server: Configuring routes: health")
		// 	s.Handlers = append(server.Handlers, handler)
		// 	s.Gin.GET("/health", handler.HealthCheck())
		// },
		// func(s *AppServer, handler *handlers.ChatHandler) {
		// 	fmt.Println("Server: Configuring routes: chat")
		// 	s.Handlers = append(server.Handlers, handler)
		// 	s.Gin.GET("/chat", handler.Chat())
		// },
	),
)
