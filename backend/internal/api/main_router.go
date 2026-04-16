package api

import (
	"devops-os/backend/internal/auth"
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/middleware"
	"devops-os/backend/internal/model"
	"devops-os/backend/internal/permission"
	"devops-os/backend/internal/project"
	"devops-os/backend/internal/user"

	"github.com/gin-gonic/gin"
)

type RouterDependencies struct {
	AuthService          *auth.AuthService
	UserService          *user.UserService
	ProjectService       *project.ProjectService
	ProjectMemberService *project.ProjectMemberService
	PermissionEngine     *permission.PermissionEngine
	AuthHandler          *AuthHandler
	ProjectHandler       *ProjectHandler
}

func SetupMainRouter(deps *RouterDependencies) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	// 创建中间件
	authMiddleware := middleware.AuthMiddleware(deps.AuthService)

	api := router.Group("/api")
	{
		// 认证路由
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", deps.AuthHandler.Login)
			authGroup.POST("/logout", authMiddleware, deps.AuthHandler.Logout)
			authGroup.GET("/me", authMiddleware, deps.AuthHandler.GetCurrentUser)
		}

		// 项目路由（需要认证）
		projectsGroup := api.Group("/projects")
		projectsGroup.Use(authMiddleware)
		{
			projectsGroup.GET("", deps.ProjectHandler.GetProjects)
			projectsGroup.POST("", deps.ProjectHandler.CreateProject)

			projectGroup := projectsGroup.Group("/:project_id")
			projectGroup.Use(middleware.ProjectMiddleware(deps.ProjectService))
			{
				projectGroup.GET("", deps.ProjectHandler.GetProject)

				// 成员管理（需要 project_admin 权限）
				membersGroup := projectGroup.Group("/members")
				membersGroup.Use(middleware.ProjectAdminMiddleware(deps.ProjectMemberService))
				{
					membersGroup.GET("", func(c *gin.Context) {
						projectID := c.Param("project_id")
						members, err := deps.ProjectMemberService.GetMembers(projectID)
						if err != nil {
							common.Error(c, 500, err.Error())
							return
						}
						common.Success(c, members)
					})

					membersGroup.POST("", func(c *gin.Context) {
						projectID := c.Param("project_id")
						var req struct {
							UserID string `json:"user_id" binding:"required"`
							Role   string `json:"role" binding:"required"`
						}

						if err := c.ShouldBindJSON(&req); err != nil {
							common.Error(c, 400, "Invalid request")
							return
						}

						addMemberReq := model.AddMemberRequest{
							UserID: req.UserID,
							Role:   req.Role,
						}

						if err := deps.ProjectMemberService.AddMember(projectID, &addMemberReq); err != nil {
							common.Error(c, 500, err.Error())
							return
						}

						common.Success(c, gin.H{"message": "Member added successfully"})
					})

					membersGroup.DELETE("/:user_id", func(c *gin.Context) {
						projectID := c.Param("project_id")
						userID := c.Param("user_id")

						if err := deps.ProjectMemberService.RemoveMember(projectID, userID); err != nil {
							common.Error(c, 500, err.Error())
							return
						}

						common.Success(c, gin.H{"message": "Member removed successfully"})
					})
				}

				// 模块路由（使用权限引擎）
				setupModuleRoutes(projectGroup, deps.PermissionEngine)
			}
		}

		// 用户管理（仅 super_admin）
		usersGroup := api.Group("/users")
		usersGroup.Use(authMiddleware)
		usersGroup.Use(func(c *gin.Context) {
			userID, _ := c.Get(common.ContextKeyUserID)
			isSuperAdmin, err := deps.UserService.IsSuperAdmin(userID.(string))
			if err != nil || !isSuperAdmin {
				common.Error(c, 403, "Only super admin can access user management")
				c.Abort()
				return
			}
			c.Next()
		})
		{
			usersGroup.GET("", func(c *gin.Context) {
				users, err := deps.UserService.GetAllUsers()
				if err != nil {
					common.Error(c, 500, err.Error())
					return
				}
				common.Success(c, users)
			})

			usersGroup.POST("", func(c *gin.Context) {
				var req struct {
					Username   string `json:"username" binding:"required"`
					Password   string `json:"password" binding:"required"`
					Email      string `json:"email"`
					SystemRole string `json:"system_role"`
				}

				if err := c.ShouldBindJSON(&req); err != nil {
					common.Error(c, 400, "Invalid request")
					return
				}

				createUserReq := model.CreateUserRequest{
					Username:   req.Username,
					Password:   req.Password,
					Email:      req.Email,
					SystemRole: req.SystemRole,
				}
				user, err := deps.UserService.CreateUser(&createUserReq)
				if err != nil {
					common.Error(c, 500, err.Error())
					return
				}

				common.Success(c, user)
			})
		}
	}

	// 设置权限路由
	SetupPermissionRouter(router, deps.PermissionEngine, deps.ProjectMemberService, authMiddleware)

	return router
}

func setupModuleRoutes(projectGroup *gin.RouterGroup, permissionEngine *permission.PermissionEngine) {
	// CI/CD 模块
	ciGroup := projectGroup.Group("/ci")
	ciGroup.Use(middleware.RBACMiddleware(permissionEngine, common.ResourceCI, common.ActionView))
	{
		ciGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		ciGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		ciGroup.POST("", middleware.RBACMiddleware(permissionEngine, common.ResourceCI, common.ActionCreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		ciGroup.PUT("/:id", middleware.RBACMiddleware(permissionEngine, common.ResourceCI, common.ActionUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		ciGroup.DELETE("/:id", middleware.RBACMiddleware(permissionEngine, common.ResourceCI, common.ActionDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	// 其他模块类似实现...
	// 这里只实现 CI 模块作为示例，其他模块可以类似添加
}
