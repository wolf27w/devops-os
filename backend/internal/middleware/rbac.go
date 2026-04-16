package middleware

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/permission"
	"devops-os/backend/internal/project"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(
	permissionEngine *permission.PermissionEngine,
	resource, action string,
) gin.HandlerFunc {
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

		hasPermission, err := permissionEngine.HasPermission(
			userID.(string),
			projectID.(string),
			resource,
			action,
		)

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

func ProjectAdminMiddleware(projectMemberService *project.ProjectMemberService) gin.HandlerFunc {
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

		role, err := projectMemberService.GetUserRoleInProject(userID.(string), projectID.(string))
		if err != nil {
			common.Error(c, 500, "Failed to get user role")
			c.Abort()
			return
		}

		if role != common.RoleProjectAdmin {
			common.Error(c, 403, "Only project admin can perform this action")
			c.Abort()
			return
		}

		c.Next()
	}
}
