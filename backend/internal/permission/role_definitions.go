package permission

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

// 角色权限定义
var RolePermissions = map[string]map[string][]string{
	// 系统角色权限
	RoleSuperAdmin: {
		ResourceProject:  {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceCI:       {ActionView, ActionCreate, ActionUpdate, ActionDelete, ActionExecute, ActionApprove},
		ResourceRepo:     {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceWorkflow: {ActionView, ActionCreate, ActionUpdate, ActionDelete, ActionExecute},
		ResourceDB:       {ActionView, ActionCreate, ActionUpdate, ActionDelete, ActionExecute},
		ResourceWiki:     {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceNav:      {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceMember:   {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceSettings: {ActionView, ActionUpdate},
	},

	// 项目角色权限
	RoleProjectAdmin: {
		ResourceProject:  {ActionView, ActionUpdate},
		ResourceCI:       {ActionView, ActionCreate, ActionUpdate, ActionDelete, ActionExecute, ActionApprove},
		ResourceRepo:     {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceWorkflow: {ActionView, ActionCreate, ActionUpdate, ActionDelete, ActionExecute},
		ResourceDB:       {ActionView, ActionCreate, ActionUpdate, ActionDelete, ActionExecute},
		ResourceWiki:     {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceNav:      {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceMember:   {ActionView, ActionCreate, ActionUpdate, ActionDelete},
		ResourceSettings: {ActionView, ActionUpdate},
	},

	RoleDeveloper: {
		ResourceProject:  {ActionView},
		ResourceCI:       {ActionView, ActionCreate, ActionUpdate, ActionExecute},
		ResourceRepo:     {ActionView, ActionCreate, ActionUpdate},
		ResourceWorkflow: {ActionView, ActionCreate, ActionUpdate, ActionExecute},
		ResourceDB:       {ActionView, ActionCreate, ActionUpdate, ActionExecute},
		ResourceWiki:     {ActionView, ActionCreate, ActionUpdate},
		ResourceNav:      {ActionView, ActionCreate, ActionUpdate},
		ResourceMember:   {ActionView},
		ResourceSettings: {ActionView},
	},

	RoleViewer: {
		ResourceProject:  {ActionView},
		ResourceCI:       {ActionView},
		ResourceRepo:     {ActionView},
		ResourceWorkflow: {ActionView},
		ResourceDB:       {ActionView},
		ResourceWiki:     {ActionView},
		ResourceNav:      {ActionView},
		ResourceMember:   {ActionView},
		ResourceSettings: {ActionView},
	},
}

// 检查角色是否有权限
func HasRolePermission(role, resource, action string) bool {
	permissions, ok := RolePermissions[role]
	if !ok {
		return false
	}

	actions, ok := permissions[resource]
	if !ok {
		return false
	}

	for _, a := range actions {
		if a == action {
			return true
		}
	}

	return false
}
