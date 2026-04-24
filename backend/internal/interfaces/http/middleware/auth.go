package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"turbo-taskify/backend/internal/infrastructure/token"
)

func Auth(ts *token.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false, "error": "unauthorized",
			})
			return
		}

		claims, err := ts.ValidateAccessToken(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false, "error": "invalid_token",
			})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
