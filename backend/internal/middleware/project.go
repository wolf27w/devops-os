package middleware

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/project"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProjectMiddleware(projectService *project.ProjectService) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := c.Param("project_id")
		if projectID == "" {
			projectID = c.Query("project_id")
		}

		if projectID == "" {
			common.Error(c, 400, "project_id is required")
			c.Abort()
			return
		}

		userID, exists := c.Get(common.ContextKeyUserID)
		if !exists {
			common.Error(c, 401, "User not authenticated")
			c.Abort()
			return
		}

		// 检查用户是否在项目中
		memberService := project.NewProjectMemberService(projectService.repo)
		inProject, err := memberService.IsUserInProject(userID.(string), projectID)
		if err != nil || !inProject {
			common.Error(c, 403, "Access denied to project")
			c.Abort()
			return
		}

		c.Set(common.ContextKeyProjectID, projectID)
		c.Next()
	}
}

func ProjectContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		parts := strings.Split(path, "/")

		for i, part := range parts {
			if part == "projects" && i+1 < len(parts) {
				projectID := parts[i+1]
				c.Set(common.ContextKeyProjectID, projectID)
				break
			}
		}

		c.Next()
	}
}
