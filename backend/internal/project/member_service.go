package project

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"errors"
	"fmt"
)

type ProjectMemberService struct {
	repo *ProjectRepository
}

func NewProjectMemberService(repo *ProjectRepository) *ProjectMemberService {
	return &ProjectMemberService{repo: repo}
}

func (s *ProjectMemberService) AddMember(projectID string, req *model.AddMemberRequest) error {
	// 检查项目是否存在
	project, err := s.repo.GetByID(projectID)
	if err != nil {
		return err
	}

	if project == nil {
		return errors.New("project not found")
	}

	// 检查用户是否已经是成员
	members, err := s.repo.GetMembers(projectID)
	if err != nil {
		return err
	}

	for _, member := range members {
		if member.UserID == req.UserID {
			return errors.New("user is already a member of this project")
		}
	}

	// 添加成员
	return s.repo.AddMember(projectID, req.UserID, req.Role)
}

func (s *ProjectMemberService) RemoveMember(projectID, userID string) error {
	// 检查项目是否存在
	project, err := s.repo.GetByID(projectID)
	if err != nil {
		return err
	}

	if project == nil {
		return errors.New("project not found")
	}

	// 获取所有成员
	members, err := s.repo.GetMembers(projectID)
	if err != nil {
		return err
	}

	// 检查是否是最后一个 project_admin
	var projectAdminCount int
	var memberIndex = -1

	for i, member := range members {
		if member.UserID == userID {
			memberIndex = i
		}
		if member.Role == common.RoleProjectAdmin {
			projectAdminCount++
		}
	}

	if memberIndex == -1 {
		return errors.New("user is not a member of this project")
	}

	// 如果要删除的是 project_admin，检查是否是最后一个
	if members[memberIndex].Role == common.RoleProjectAdmin && projectAdminCount <= 1 {
		return errors.New("cannot remove the last project admin")
	}

	// 删除成员
	members = append(members[:memberIndex], members[memberIndex+1:]...)

	path := fmt.Sprintf("projects/%s/members.json", projectID)
	return s.repo.storage.WriteJSON(path, members)
}

func (s *ProjectMemberService) UpdateMemberRole(projectID, userID, newRole string) error {
	// 检查项目是否存在
	project, err := s.repo.GetByID(projectID)
	if err != nil {
		return err
	}

	if project == nil {
		return errors.New("project not found")
	}

	// 获取所有成员
	members, err := s.repo.GetMembers(projectID)
	if err != nil {
		return err
	}

	// 查找成员并更新角色
	var memberFound bool
	var projectAdminCount int

	for i, member := range members {
		if member.Role == common.RoleProjectAdmin {
			projectAdminCount++
		}

		if member.UserID == userID {
			memberFound = true

			// 如果要降级最后一个 project_admin，不允许
			if member.Role == common.RoleProjectAdmin &&
				newRole != common.RoleProjectAdmin &&
				projectAdminCount <= 1 {
				return errors.New("cannot demote the last project admin")
			}

			members[i].Role = newRole
			break
		}
	}

	if !memberFound {
		return errors.New("user is not a member of this project")
	}

	path := fmt.Sprintf("projects/%s/members.json", projectID)
	return s.repo.storage.WriteJSON(path, members)
}

func (s *ProjectMemberService) GetMembers(projectID string) ([]model.ProjectMember, error) {
	return s.repo.GetMembers(projectID)
}

func (s *ProjectMemberService) GetUserRoleInProject(userID, projectID string) (string, error) {
	members, err := s.repo.GetMembers(projectID)
	if err != nil {
		return "", err
	}

	for _, member := range members {
		if member.UserID == userID {
			return member.Role, nil
		}
	}

	return "", errors.New("user is not a member of this project")
}

func (s *ProjectMemberService) IsUserInProject(userID, projectID string) (bool, error) {
	members, err := s.repo.GetMembers(projectID)
	if err != nil {
		return false, err
	}

	for _, member := range members {
		if member.UserID == userID {
			return true, nil
		}
	}

	return false, nil
}

func (s *ProjectMemberService) GetUserProjects(userID string) ([]model.Project, error) {
	return s.repo.GetUserProjects(userID)
}

func (s *ProjectMemberService) CountProjectAdmins(projectID string) (int, error) {
	members, err := s.repo.GetMembers(projectID)
	if err != nil {
		return 0, err
	}

	var count int
	for _, member := range members {
		if member.Role == common.RoleProjectAdmin {
			count++
		}
	}

	return count, nil
}
