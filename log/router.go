package log

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	h := NewHandler()

	r.POST("/log", h.WriteLog)
	r.GET("/healthz", h.Health)
	r.POST("/registry/updates", h.HandleRegistryUpdate)
}
