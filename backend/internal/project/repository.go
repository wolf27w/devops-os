package project

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"fmt"
)

type ProjectRepository struct {
	storage *common.FileStorage
}

func NewProjectRepository(storage *common.FileStorage) *ProjectRepository {
	return &ProjectRepository{storage: storage}
}

func (r *ProjectRepository) Create(project *model.Project) error {
	project.ID = common.GenerateID()
	project.CreatedAt = common.GenerateTimestamp()
	project.UpdatedAt = project.CreatedAt

	path := fmt.Sprintf("projects/%s/project.json", project.ID)
	return r.storage.WriteJSON(path, project)
}

func (r *ProjectRepository) GetByID(projectID string) (*model.Project, error) {
	path := fmt.Sprintf("projects/%s/project.json", projectID)

	var project model.Project
	if err := r.storage.ReadJSON(path, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	projectsDir := "projects"

	files, err := r.storage.ListFiles(projectsDir)
	if err != nil {
		return nil, err
	}

	var projects []model.Project
	for _, file := range files {
		projectID := file
		project, err := r.GetByID(projectID)
		if err != nil {
			continue
		}
		projects = append(projects, *project)
	}

	return projects, nil
}

func (r *ProjectRepository) Update(project *model.Project) error {
	project.UpdatedAt = common.GenerateTimestamp()

	path := fmt.Sprintf("projects/%s/project.json", project.ID)
	return r.storage.WriteJSON(path, project)
}

func (r *ProjectRepository) AddMember(projectID, userID, role string) error {
	path := fmt.Sprintf("projects/%s/members.json", projectID)

	var members []model.ProjectMember
	if r.storage.PathExists(path) {
		if err := r.storage.ReadJSON(path, &members); err != nil {
			return err
		}
	}

	member := model.ProjectMember{
		ID:        common.GenerateID(),
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
		JoinedAt:  common.GenerateTimestamp(),
	}

	members = append(members, member)
	return r.storage.WriteJSON(path, members)
}

func (r *ProjectRepository) GetMembers(projectID string) ([]model.ProjectMember, error) {
	path := fmt.Sprintf("projects/%s/members.json", projectID)

	if !r.storage.PathExists(path) {
		return []model.ProjectMember{}, nil
	}

	var members []model.ProjectMember
	if err := r.storage.ReadJSON(path, &members); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *ProjectRepository) GetUserProjects(userID string) ([]model.Project, error) {
	allProjects, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	var userProjects []model.Project
	for _, project := range allProjects {
		members, err := r.GetMembers(project.ID)
		if err != nil {
			continue
		}

		for _, member := range members {
			if member.UserID == userID {
				userProjects = append(userProjects, project)
				break
			}
		}
	}

	return userProjects, nil
}
