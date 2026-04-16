package middleware

import (
	"devops-os/backend/internal/auth"
	"devops-os/backend/internal/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.Error(c, 401, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			common.Error(c, 401, "Bearer token is required")
			c.Abort()
			return
		}

		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			common.Error(c, 401, "Invalid token")
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			common.Error(c, 401, "Invalid token claims")
			c.Abort()
			return
		}

		username, _ := claims["username"].(string)
		role, _ := claims["role"].(string)

		c.Set(common.ContextKeyUserID, userID)
		c.Set(common.ContextKeyUsername, username)
		c.Set(common.ContextKeyUserRole, role)

		c.Next()
	}
}
