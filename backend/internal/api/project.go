package api

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"devops-os/backend/internal/project"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *project.ProjectService
}

func NewProjectHandler(projectService *project.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req model.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "Invalid request")
		return
	}

	userID, _ := c.Get(common.ContextKeyUserID)
	userRole, _ := c.Get(common.ContextKeyUserRole)

	if userRole != common.RoleSuperAdmin {
		common.Error(c, 403, "Only super admin can create projects")
		return
	}

	proj, err := h.projectService.CreateProject(&req, userID.(string))
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}

	common.Success(c, proj)
}

func (h *ProjectHandler) GetProjects(c *gin.Context) {
	userID, _ := c.Get(common.ContextKeyUserID)
	userRole, _ := c.Get(common.ContextKeyUserRole)

	var projects []model.Project
	var err error

	if userRole == common.RoleSuperAdmin {
		projects, err = h.projectService.GetAllProjects()
	} else {
		projects, err = h.projectService.GetUserProjects(userID.(string))
	}

	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}

	common.Success(c, projects)
}

func (h *ProjectHandler) GetProject(c *gin.Context) {
	projectID := c.Param("project_id")
	if projectID == "" {
		common.Error(c, 400, "project_id is required")
		return
	}

	proj, err := h.projectService.GetProject(projectID)
	if err != nil {
		common.Error(c, 404, "Project not found")
		return
	}

	common.Success(c, proj)
}

func (h *ProjectHandler) AddMember(c *gin.Context) {
	projectID := c.Param("project_id")
	if projectID == "" {
		common.Error(c, 400, "project_id is required")
		return
	}

	var req model.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "Invalid request")
		return
	}

	// 这里需要项目成员服务，暂时返回未实现
	common.Error(c, 501, "Member management not implemented in this handler")
}

func (h *ProjectHandler) GetMembers(c *gin.Context) {
	projectID := c.Param("project_id")
	if projectID == "" {
		common.Error(c, 400, "project_id is required")
		return
	}

	// 这里需要项目成员服务，暂时返回空数组
	common.Success(c, []interface{}{})
}
