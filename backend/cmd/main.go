package main

import (
	"devops-os/backend/internal/api"
	"devops-os/backend/internal/auth"
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/project"
	"log"
	"os"
)

func main() {
	basePath := "data"
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}

	storage := common.NewFileStorage(basePath)

	authRepo := auth.NewAuthRepository(storage)
	authService := auth.NewAuthService(authRepo, "devops-os-secret-key-change-in-production")

	projectRepo := project.NewProjectRepository(storage)
	projectService := project.NewProjectService(projectRepo)

	authHandler := api.NewAuthHandler(authService)
	projectHandler := api.NewProjectHandler(projectService)

	router := api.SetupRouter(authService, projectService, authHandler, projectHandler)

	log.Println("DevOps OS backend starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
