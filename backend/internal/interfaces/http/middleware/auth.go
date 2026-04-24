package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"turbo-taskify/backend/internal/infrastructure/token"
)

func Auth(ts *token.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil || accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false, "error": "unauthorized",
			})
			return
		}

		claims, err := ts.ValidateAccessToken(accessToken)
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
