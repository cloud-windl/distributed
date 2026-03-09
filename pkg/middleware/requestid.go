package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()

		c.Set("request_id", id)

		c.Writer.Header().Set(`X-Request-ID`, id)

		c.Next()
	}
}
