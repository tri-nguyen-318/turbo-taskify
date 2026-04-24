package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	authapp "turbo-taskify/backend/internal/application/auth"
	"turbo-taskify/backend/internal/infrastructure/token"
	"turbo-taskify/backend/internal/interfaces/http/handler"
	"turbo-taskify/backend/internal/interfaces/http/middleware"
)

func NewServer(frontendURL string, authSvc *authapp.Service, tokenSvc *token.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors(frontendURL))

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	authH := handler.NewAuthHandler(authSvc)
	authMW := middleware.Auth(tokenSvc)

	api := r.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/signup", authH.SignUp)
		auth.POST("/login", authH.SignIn)
		auth.POST("/google", authH.GoogleSignIn)
		auth.POST("/refresh", authH.Refresh)
		auth.POST("/logout", authMW, authH.SignOut)
		auth.GET("/me", authMW, authH.GetMe)
	}

	return r
}

func ListenAddr(port string) string { return fmt.Sprintf(":%s", port) }

func cors(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
