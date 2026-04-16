package api

import (
	"devops-os/backend/internal/auth"
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/middleware"
	"devops-os/backend/internal/project"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authService *auth.AuthService,
	projectService *project.ProjectService,
	authHandler *AuthHandler,
	projectHandler *ProjectHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/logout", middleware.AuthMiddleware(authService), authHandler.Logout)
			authGroup.GET("/me", middleware.AuthMiddleware(authService), authHandler.GetCurrentUser)
		}

		projectsGroup := api.Group("/projects")
		projectsGroup.Use(middleware.AuthMiddleware(authService))
		{
			projectsGroup.GET("", projectHandler.GetProjects)
			projectsGroup.POST("", projectHandler.CreateProject)

			projectGroup := projectsGroup.Group("/:project_id")
			projectGroup.Use(middleware.ProjectMiddleware(projectService))
			{
				projectGroup.GET("", projectHandler.GetProject)
				projectGroup.GET("/members", projectHandler.GetMembers)
				projectGroup.POST("/members", projectHandler.AddMember)

				setupModuleRoutes(projectGroup, projectService)
			}
		}
	}

	return router
}

func setupModuleRoutes(projectGroup *gin.RouterGroup, projectService *project.ProjectService) {
	ciGroup := projectGroup.Group("/ci")
	ciGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionCIRead))
	{
		ciGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		ciGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		ciGroup.POST("", middleware.RBACMiddleware(projectService, common.PermissionCICreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		ciGroup.PUT("/:id", middleware.RBACMiddleware(projectService, common.PermissionCIUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		ciGroup.DELETE("/:id", middleware.RBACMiddleware(projectService, common.PermissionCIDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	repoGroup := projectGroup.Group("/repo")
	repoGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionRepoRead))
	{
		repoGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		repoGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		repoGroup.POST("", middleware.RBACMiddleware(projectService, common.PermissionRepoCreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		repoGroup.PUT("/:id", middleware.RBACMiddleware(projectService, common.PermissionRepoUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		repoGroup.DELETE("/:id", middleware.RBACMiddleware(projectService, common.PermissionRepoDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	workflowGroup := projectGroup.Group("/workflow")
	workflowGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionWorkflowRead))
	{
		workflowGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		workflowGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		workflowGroup.POST("", middleware.RBACMiddleware(projectService, common.PermissionWorkflowCreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		workflowGroup.PUT("/:id", middleware.RBACMiddleware(projectService, common.PermissionWorkflowUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		workflowGroup.DELETE("/:id", middleware.RBACMiddleware(projectService, common.PermissionWorkflowDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	dbGroup := projectGroup.Group("/db")
	dbGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionDBRead))
	{
		dbGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		dbGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		dbGroup.POST("", middleware.RBACMiddleware(projectService, common.PermissionDBCreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		dbGroup.PUT("/:id", middleware.RBACMiddleware(projectService, common.PermissionDBUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		dbGroup.DELETE("/:id", middleware.RBACMiddleware(projectService, common.PermissionDBDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	wikiGroup := projectGroup.Group("/wiki")
	wikiGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionWikiRead))
	{
		wikiGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		wikiGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		wikiGroup.POST("", middleware.RBACMiddleware(projectService, common.PermissionWikiCreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		wikiGroup.PUT("/:id", middleware.RBACMiddleware(projectService, common.PermissionWikiUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		wikiGroup.DELETE("/:id", middleware.RBACMiddleware(projectService, common.PermissionWikiDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	navGroup := projectGroup.Group("/nav")
	navGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionNavRead))
	{
		navGroup.GET("", func(c *gin.Context) {
			common.Success(c, []interface{}{})
		})
		navGroup.GET("/:id", func(c *gin.Context) {
			common.Success(c, nil)
		})
		navGroup.POST("", middleware.RBACMiddleware(projectService, common.PermissionNavCreate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		navGroup.PUT("/:id", middleware.RBACMiddleware(projectService, common.PermissionNavUpdate), func(c *gin.Context) {
			common.Success(c, nil)
		})
		navGroup.DELETE("/:id", middleware.RBACMiddleware(projectService, common.PermissionNavDelete), func(c *gin.Context) {
			common.Success(c, nil)
		})
	}

	settingsGroup := projectGroup.Group("/settings")
	settingsGroup.Use(middleware.RBACMiddleware(projectService, common.PermissionProjectUpdate))
	{
		settingsGroup.GET("", func(c *gin.Context) {
			common.Success(c, nil)
		})
		settingsGroup.PUT("", func(c *gin.Context) {
			common.Success(c, nil)
		})
	}
}
