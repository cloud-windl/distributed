package main

import (
	"distributed/pkg/config"
	"distributed/pkg/middleware"
	"distributed/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	registry.SetupRegistryService()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())

	registry.RegisterRoutes(r)

	port := config.GetEnv("REGISTRY_PORT", "3000")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
