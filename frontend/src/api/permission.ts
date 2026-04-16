import api from './index'

interface PermissionCheckParams {
  project_id: string
  resource: string
  action: string
}

interface PermissionCheckResult {
  user_id: string
  project_id: string
  resource: string
  action: string
  role: string
  allowed: boolean
  message: string
}

interface RoleInfo {
  role: string
  permissions: Record<string, string[]>
}

export const checkPermission = (
  projectId: string,
  resource: string,
  action: string
): Promise<PermissionCheckResult> => {
  return api.get('/permissions/check', {
    params: {
      project_id: projectId,
      resource,
      action
    }
  })
}

export const getProjectRoles = (projectId: string): Promise<RoleInfo> => {
  return api.get(`/projects/${projectId}/roles`)
}

// 权限常量
export const Resources = {
  PROJECT: 'project',
  CI: 'ci',
  REPO: 'repo',
  WORKFLOW: 'workflow',
  DB: 'db',
  WIKI: 'wiki',
  NAV: 'nav',
  MEMBER: 'member',
  SETTINGS: 'settings'
} as const

export const Actions = {
  VIEW: 'view',
  CREATE: 'create',
  UPDATE: 'update',
  DELETE: 'delete',
  EXECUTE: 'execute',
  APPROVE: 'approve'
} as const

// 预定义的权限检查函数
export const canViewProject = (projectId: string) => 
  checkPermission(projectId, Resources.PROJECT, Actions.VIEW)

export const canCreateCI = (projectId: string) => 
  checkPermission(projectId, Resources.CI, Actions.CREATE)

export const canUpdateCI = (projectId: string) => 
  checkPermission(projectId, Resources.CI, Actions.UPDATE)

export const canDeleteCI = (projectId: string) => 
  checkPermission(projectId, Resources.CI, Actions.DELETE)

export const canExecuteCI = (projectId: string) => 
  checkPermission(projectId, Resources.CI, Actions.EXECUTE)

export const canManageMembers = (projectId: string) => 
  checkPermission(projectId, Resources.MEMBER, Actions.CREATE)

export const canUpdateSettings = (projectId: string) => 
  checkPermission(projectId, Resources.SETTINGS, Actions.UPDATE)