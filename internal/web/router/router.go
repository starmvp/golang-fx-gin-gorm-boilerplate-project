package router

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/web/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RegisterRouteParams struct {
	fx.In

	Server  *server.Server
	Method  string
	Pattern string
	Handler gin.HandlerFunc
}

type RegisterRouteResult struct {
	fx.Out

	Routes gin.IRoutes
}

func RegisterRoute(params RegisterRouteParams) (RegisterRouteResult, error) {
	routes := params.Server.Gin.Handle(
		params.Method,
		params.Pattern,
		params.Handler,
	)

	return RegisterRouteResult{Routes: routes}, nil
}

var Module = fx.Provide(RegisterRoute)
