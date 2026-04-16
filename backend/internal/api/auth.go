package api

import (
	"devops-os/backend/internal/auth"
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "Invalid request")
		return
	}

	resp, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		common.Error(c, 401, err.Error())
		return
	}

	common.Success(c, resp)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	common.Success(c, gin.H{"message": "Logged out successfully"})
}

func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get(common.ContextKeyUserID)
	if !exists {
		common.Error(c, 401, "User not authenticated")
		return
	}

	username, _ := c.Get(common.ContextKeyUsername)
	role, _ := c.Get(common.ContextKeyUserRole)

	user := gin.H{
		"id":       userID,
		"username": username,
		"role":     role,
	}

	common.Success(c, user)
}
