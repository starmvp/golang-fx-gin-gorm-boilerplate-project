package server

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	Gin    *gin.Engine
	Logger *logger.Logger
}

type ServerParams struct {
	fx.In

	Logger *logger.Logger
}

type ServerResult struct {
	fx.Out

	Server *Server
}

func New(params ServerParams) (ServerResult, error) {
	g := gin.New()
	_ = g.SetTrustedProxies(nil)

	l := params.Logger
	if l == nil {
		l = zap.NewNop()
	}

	var s = &Server{
		Gin:    g,
		Logger: l,
	}

	return ServerResult{Server: s}, nil
}

var Module = fx.Options(fx.Provide(New))
