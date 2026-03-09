package registry

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	h := NewHandler()

	r.POST("/services", h.RegisterService)
	r.DELETE("/services", h.DeregisterService)
}
