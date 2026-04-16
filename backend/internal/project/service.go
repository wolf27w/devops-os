package project

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"errors"
)

type ProjectService struct {
	Repo *ProjectRepository
}

func NewProjectService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) CreateProject(req *model.CreateProjectRequest, ownerID string) (*model.Project, error) {
	project := &model.Project{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     ownerID,
		Status:      "active",
	}

	if err := s.Repo.Create(project); err != nil {
		return nil, err
	}

	// 创建者自动成为 project_admin
	if err := s.Repo.AddMember(project.ID, ownerID, common.RoleProjectAdmin); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetProject(projectID string) (*model.Project, error) {
	return s.Repo.GetByID(projectID)
}

func (s *ProjectService) GetAllProjects() ([]model.Project, error) {
	return s.Repo.GetAll()
}

func (s *ProjectService) GetUserProjects(userID string) ([]model.Project, error) {
	return s.Repo.GetUserProjects(userID)
}

func (s *ProjectService) UpdateProject(projectID string, updates map[string]interface{}) (*model.Project, error) {
	project, err := s.Repo.GetByID(projectID)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, errors.New("project not found")
	}

	// 更新字段
	if name, ok := updates["name"].(string); ok && name != "" {
		project.Name = name
	}

	if description, ok := updates["description"].(string); ok {
		project.Description = description
	}

	if status, ok := updates["status"].(string); ok && status != "" {
		project.Status = status
	}

	project.UpdatedAt = common.GenerateTimestamp()

	if err := s.Repo.Update(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) DeleteProject(projectID string) error {
	return errors.New("project deletion not implemented")
}
