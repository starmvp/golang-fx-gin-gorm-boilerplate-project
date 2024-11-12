package server

import (
	"fmt"
	"golang-fx-gin-gorm-boilerplate-project/server/handlers"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func (s *AppServer) ConfigureRouteGroups() {
	fmt.Println("AppServer: Configuring route groups")

	s.Gin.Static("/static/assets", "./webroot/assets")

	s.ApiNoAuth = s.NoAuth.Group("/api/v1/noauth")
	s.ApiNoAuth.Use(func(c *gin.Context) {
	})
	fmt.Println("AppServer: Configuring route groups ApiNoAuth. s.ApiNoAuth: ", s.ApiNoAuth)

	s.ApiNeedsAuth = s.NeedsAuth.Group("/api/v1/needauth")
	s.ApiNeedsAuth.Use(func(c *gin.Context) {
	})
	fmt.Println("AppServer: Configuring route groups ApiNeedsAuth. s.ApiNeedAuth: ", s.ApiNeedsAuth)
}

func AppendFallbackRoute() fx.Option {
	return fx.Invoke(func(s *AppServer) {
		fmt.Println("Register fallback route")
		s.Gin.NoRoute(func(c *gin.Context) {
			if !strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.File("./webroot/index.html")
			}
		})
	})
}

var routesModule = fx.Options(
	fx.Invoke(
		func(s *AppServer, handler *handlers.HealthCheckHandler) {
			fmt.Println("Server: Configuring routes: health")
			s.Handlers = append(s.Handlers, handler)
			s.NoAuth.GET("/health", handler.HealthCheck())
		},
		func(s *AppServer, handler *handlers.PingHandler) {
			fmt.Println("Server: Configuring routes: no auth ping")
			s.Handlers = append(s.Handlers, handler)
			s.ApiNoAuth.GET("/ping", handler.Ping())
		},
		func(s *AppServer, handler *handlers.PingHandler) {
			fmt.Println("Server: Configuring routes: need auth ping")
			s.Handlers = append(s.Handlers, handler)
			s.ApiNeedsAuth.GET("/ping", handler.Ping())
		},
	),
	fx.Invoke(func() { fmt.Println("All primary routes have been registered.") }),
	AppendFallbackRoute(),
)
