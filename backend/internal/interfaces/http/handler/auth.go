package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	authapp "turbo-taskify/backend/internal/application/auth"
)

type AuthHandler struct {
	svc *authapp.Service
}

func NewAuthHandler(svc *authapp.Service) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req authapp.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "validation_error", "details": err.Error()})
		return
	}

	resp, err := h.svc.SignUp(c.Request.Context(), req)
	if err != nil {
		h.handleErr(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req authapp.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "validation_error"})
		return
	}

	resp, err := h.svc.SignIn(c.Request.Context(), req)
	if err != nil {
		h.handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) GoogleSignIn(c *gin.Context) {
	var req authapp.GoogleSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "validation_error"})
		return
	}

	resp, isNew, err := h.svc.GoogleSignIn(c.Request.Context(), req)
	if err != nil {
		h.handleErr(c, err)
		return
	}

	status := http.StatusOK
	if isNew {
		status = http.StatusCreated
	}
	c.JSON(status, resp)
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	// Stateless JWT: client drops the token; extend here to blacklist refresh tokens if needed.
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logged out successfully"})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID := c.GetString("userID")
	u, err := h.svc.GetMe(c.Request.Context(), userID)
	if err != nil {
		h.handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": u})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req authapp.RefreshRequest
	_ = c.ShouldBindJSON(&req)

	if req.RefreshToken == "" {
		req.RefreshToken, _ = c.Cookie("refresh_token")
	}
	if req.RefreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "invalid_refresh_token"})
		return
	}

	accessToken, err := h.svc.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		h.handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "accessToken": accessToken})
}

func (h *AuthHandler) handleErr(c *gin.Context, err error) {
	switch {
	case errors.Is(err, authapp.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "invalid_credentials"})
	case errors.Is(err, authapp.ErrEmailExists):
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "email_already_exists"})
	case errors.Is(err, authapp.ErrUsernameTaken):
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "username_taken"})
	case errors.Is(err, authapp.ErrInvalidToken):
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid_token"})
	case errors.Is(err, authapp.ErrTokenExpired):
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "token_expired"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "internal_error"})
	}
}
