package router

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RegisterRouteParams struct {
	fx.In
}

type RegisterRouteResult struct {
	fx.Out

	Routes gin.IRoutes
}

func RegisterRoute(
	Server *server.Server,
	Method string,
	Pattern string,
	Handler gin.HandlerFunc,
) (gin.IRoutes, error) {
	routes := Server.Gin.Handle(
		Method,
		Pattern,
		Handler,
	)

	return routes, nil
}

var Module = fx.Provide(
	fx.Annotate(RegisterRoute),
)
