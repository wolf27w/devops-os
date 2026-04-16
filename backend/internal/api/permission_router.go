package api

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/permission"
	"devops-os/backend/internal/project"

	"github.com/gin-gonic/gin"
)

func SetupPermissionRouter(
	router *gin.Engine,
	permissionEngine *permission.PermissionEngine,
	projectMemberService *project.ProjectMemberService,
) {
	api := router.Group("/api")
	api.Use(AuthMiddleware)

	// 权限检查接口
	api.GET("/permissions/check", func(c *gin.Context) {
		userID, _ := c.Get(common.ContextKeyUserID)
		projectID := c.Query("project_id")
		resource := c.Query("resource")
		action := c.Query("action")

		if projectID == "" || resource == "" || action == "" {
			common.Error(c, 400, "project_id, resource, and action are required")
			return
		}

		result, err := permissionEngine.CheckPermission(userID.(string), projectID, resource, action)
		if err != nil {
			common.Error(c, 500, err.Error())
			return
		}

		common.Success(c, result)
	})

	// 获取项目角色
	api.GET("/projects/:project_id/roles", func(c *gin.Context) {
		projectID := c.Param("project_id")
		userID, _ := c.Get(common.ContextKeyUserID)

		// 检查用户是否在项目中
		inProject, err := projectMemberService.IsUserInProject(userID.(string), projectID)
		if err != nil || !inProject {
			common.Error(c, 403, "Access denied")
			return
		}

		// 获取用户在项目中的角色
		role, err := projectMemberService.GetUserRoleInProject(userID.(string), projectID)
		if err != nil {
			common.Error(c, 500, err.Error())
			return
		}

		common.Success(c, gin.H{
			"role":        role,
			"permissions": permission.GetRolePermissions(role),
		})
	})
}

// 获取角色权限（辅助函数）
func GetRolePermissions(role string) map[string][]string {
	return permission.RolePermissions[role]
}
