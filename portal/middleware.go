package portal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
