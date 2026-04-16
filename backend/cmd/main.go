package main

import (
	"devops-os/backend/internal/api"
	"devops-os/backend/internal/auth"
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/permission"
	"devops-os/backend/internal/project"
	"devops-os/backend/internal/user"
	"log"
	"os"
)

func main() {
	basePath := "data"
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}

	storage := common.NewFileStorage(basePath)

	// 初始化服务
	authRepo := auth.NewAuthRepository(storage)
	authService := auth.NewAuthService(authRepo, "devops-os-secret-key-change-in-production")

	userRepo := user.NewUserRepository(storage)
	userService := user.NewUserService(userRepo)

	projectRepo := project.NewProjectRepository(storage)
	projectService := project.NewProjectService(projectRepo)
	projectMemberService := project.NewProjectMemberService(projectRepo)

	permissionEngine := permission.NewPermissionEngine(projectMemberService)

	// 初始化处理器
	authHandler := api.NewAuthHandler(authService)
	projectHandler := api.NewProjectHandler(projectService)

	// 设置依赖
	deps := &api.RouterDependencies{
		AuthService:          authService,
		UserService:          userService,
		ProjectService:       projectService,
		ProjectMemberService: projectMemberService,
		PermissionEngine:     permissionEngine,
		AuthHandler:          authHandler,
		ProjectHandler:       projectHandler,
	}

	// 创建路由
	router := api.SetupMainRouter(deps)

	log.Println("DevOps OS backend starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
