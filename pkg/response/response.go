package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"error": msg,
	})
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
