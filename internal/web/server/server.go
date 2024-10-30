package server

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	Gin    *gin.Engine
	Logger *zap.Logger
}

func New(
	Config *config.Config,
	Logger *zap.Logger,
) (*Server, error) {
	// TODO: add configure for server

	g := gin.New()
	_ = g.SetTrustedProxies(nil)

	l := Logger
	if l == nil {
		l = zap.NewNop()
	}

	var s = &Server{
		Gin:    g,
		Logger: l,
	}

	return s, nil
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(New),
	),
)
