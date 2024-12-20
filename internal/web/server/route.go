package server

import (
	"fmt"
	"boilerplate/internal/web/server/providers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) ConfigureRouteGroups() {
	fmt.Println("Internal.Server: Configuring routes: enter")
	s.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})).Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			return
		}
		c.Next()
	})

	s.NoAuth = s.Gin.Group("/")

	jwtAuth := providers.NewJwtAuth(s.DB)
	s.NeedsAuth = s.Gin.Group("/")
	s.NeedsAuth.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}).Use(jwtAuth.Middleware().MiddlewareFunc())
}
