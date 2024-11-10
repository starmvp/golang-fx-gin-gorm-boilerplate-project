package server

import (
	"golang-fx-gin-gorm-boilerplate-project/internal/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger

	Gin       *gin.Engine
	NoAuth    *gin.RouterGroup
	NeedsAuth *gin.RouterGroup

	Handlers []Handler
	Services []Service
}

func NewServer(
	Config *config.Config,
	DB *gorm.DB,
	Logger *zap.Logger,
) *Server {
	// TODO: more configure for server

	g := gin.New()
	_ = g.SetTrustedProxies(nil)

	l := Logger
	if l == nil {
		l = zap.NewNop()
	}

	var s = &Server{
		Config: Config,
		DB:     DB,
		Logger: l,
		Gin:    g,
	}
	s.ConfigureRouteGroups()

	return s
}

func (server *Server) Run(addr string) error {
	s := &http.Server{
		Addr:         addr,
		Handler:      server.Gin,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.ListenAndServe()
	// return server.Gin.Run(":" + addr)
}
