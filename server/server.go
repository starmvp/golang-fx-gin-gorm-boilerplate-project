package server

import (
	"golang-fx-gin-gorm-boilerplate-project/config"
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
)
