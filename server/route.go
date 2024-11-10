package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *AppServer) ConfigureRouteGroups() {
	fmt.Println("AppServer: Configuring route groups")
	apiNoAuth := s.NoAuth.Group("/api/v1/noauth").Use(func(c *gin.Context) {
	})
	s.ApiNoAuth = &apiNoAuth
	fmt.Println("AppServer: Configuring route groups NoAuth. s.ApiNoAuth: ", s.ApiNoAuth)

	apiNeedAuth := s.NeedsAuth.Group("/api/v1/needauth").Use(func(c *gin.Context) {
	})
	s.ApiNeedAuth = &apiNeedAuth
	fmt.Println("AppServer: Configuring route groups NeedsAuth. s.ApiNeedAuth: ", s.ApiNeedAuth)
}
