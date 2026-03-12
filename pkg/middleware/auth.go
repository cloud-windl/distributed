package middleware

import (
	"strings"

	myjwt "distributed/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"code": 8001, "msg": "missing authorization header"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"code": 8002, "msg": "invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := myjwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"code": 8003, "msg": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
