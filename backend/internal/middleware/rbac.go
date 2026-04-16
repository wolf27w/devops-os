package middleware

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/project"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(projectService *project.ProjectService, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, exists := c.Get(common.ContextKeyProjectID)
		if !exists {
			common.Error(c, 400, "Project context is required")
			c.Abort()
			return
		}

		userID, exists := c.Get(common.ContextKeyUserID)
		if !exists {
			common.Error(c, 401, "User not authenticated")
			c.Abort()
			return
		}

		hasPermission, err := projectService.HasPermission(projectID.(string), userID.(string), permission)
		if err != nil {
			common.Error(c, 500, "Failed to check permission")
			c.Abort()
			return
		}

		if !hasPermission {
			common.Error(c, 403, "Insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	}
}
