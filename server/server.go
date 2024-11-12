package server

import (
	"boilerplate/config"
	"boilerplate/internal/utils"
	"boilerplate/internal/web/server"
	"boilerplate/server/handlers"
	"boilerplate/server/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppServer struct {
	server.Server

	ApiNoAuth    *gin.RouterGroup
	ApiNeedsAuth *gin.RouterGroup
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

	routesModule,

	fx.Invoke(
		func(s *AppServer, logger *zap.Logger) {
			logger.Debug("AppServer module invoked")
			go func() {
				_ = s.Run(utils.GetWebserverAddr())
			}()
		},
	),
)
