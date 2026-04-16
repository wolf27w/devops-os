package common

const (
	// 系统角色
	RoleSuperAdmin = "super_admin"
	RoleUser       = "user"

	// 项目角色
	RoleProjectAdmin = "project_admin"
	RoleDeveloper    = "developer"
	RoleViewer       = "viewer"
)

const (
	// 资源类型
	ResourceProject  = "project"
	ResourceCI       = "ci"
	ResourceRepo     = "repo"
	ResourceWorkflow = "workflow"
	ResourceDB       = "db"
	ResourceWiki     = "wiki"
	ResourceNav      = "nav"
	ResourceMember   = "member"
	ResourceSettings = "settings"
)

const (
	// 操作类型
	ActionView    = "view"
	ActionCreate  = "create"
	ActionUpdate  = "update"
	ActionDelete  = "delete"
	ActionExecute = "execute"
	ActionApprove = "approve"
)

const (
	ContextKeyUserID    = "user_id"
	ContextKeyUsername  = "username"
	ContextKeyUserRole  = "user_role"
	ContextKeyProjectID = "project_id"
)
