package permission

import (
	"errors"
)

type PermissionEngine struct {
	projectMemberService ProjectMemberService
}

type ProjectMemberService interface {
	GetUserRoleInProject(userID, projectID string) (string, error)
	IsUserInProject(userID, projectID string) (bool, error)
}

func NewPermissionEngine(projectMemberService ProjectMemberService) *PermissionEngine {
	return &PermissionEngine{
		projectMemberService: projectMemberService,
	}
}

// HasPermission 检查用户是否有权限
func (e *PermissionEngine) HasPermission(userID, projectID, resource, action string) (bool, error) {
	// 检查用户是否在项目中
	inProject, err := e.projectMemberService.IsUserInProject(userID, projectID)
	if err != nil {
		return false, err
	}

	if !inProject {
		return false, errors.New("user is not a member of this project")
	}

	// 获取用户在项目中的角色
	role, err := e.projectMemberService.GetUserRoleInProject(userID, projectID)
	if err != nil {
		return false, err
	}

	// 检查角色权限
	return HasRolePermission(role, resource, action), nil
}

// CheckPermission 权限检查函数，返回详细结果
func (e *PermissionEngine) CheckPermission(userID, projectID, resource, action string) (*PermissionResult, error) {
	result := &PermissionResult{
		UserID:    userID,
		ProjectID: projectID,
		Resource:  resource,
		Action:    action,
		Allowed:   false,
	}

	// 检查用户是否在项目中
	inProject, err := e.projectMemberService.IsUserInProject(userID, projectID)
	if err != nil {
		result.Message = "Failed to check project membership"
		return result, err
	}

	if !inProject {
		result.Message = "User is not a member of this project"
		return result, nil
	}

	// 获取用户在项目中的角色
	role, err := e.projectMemberService.GetUserRoleInProject(userID, projectID)
	if err != nil {
		result.Message = "Failed to get user role"
		return result, err
	}

	result.Role = role

	// 检查角色权限
	allowed := HasRolePermission(role, resource, action)
	result.Allowed = allowed

	if !allowed {
		result.Message = "Insufficient permissions"
	} else {
		result.Message = "Permission granted"
	}

	return result, nil
}

// CanUserCreateProject 检查用户是否可以创建项目
func (e *PermissionEngine) CanUserCreateProject(userID string) (bool, error) {
	// 只有 super_admin 可以创建项目
	// 这里需要用户服务来检查系统角色
	// 暂时返回 false，实际实现需要依赖用户服务
	return false, nil
}

// CanUserManageProjectMembers 检查用户是否可以管理项目成员
func (e *PermissionEngine) CanUserManageProjectMembers(userID, projectID string) (bool, error) {
	// project_admin 可以管理成员
	role, err := e.projectMemberService.GetUserRoleInProject(userID, projectID)
	if err != nil {
		return false, err
	}

	return role == RoleProjectAdmin, nil
}

type PermissionResult struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
	Resource  string `json:"resource"`
	Action    string `json:"action"`
	Role      string `json:"role"`
	Allowed   bool   `json:"allowed"`
	Message   string `json:"message"`
}
