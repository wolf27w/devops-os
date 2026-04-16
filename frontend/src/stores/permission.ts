import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { checkPermission, getProjectRoles } from '@/api/permission'
import { useProjectStore } from './project'
import { useAuthStore } from './auth'

interface PermissionResult {
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

export const usePermissionStore = defineStore('permission', () => {
  const authStore = useAuthStore()
  const projectStore = useProjectStore()
  
  const currentRole = ref<string>('')
  const rolePermissions = ref<Record<string, string[]>>({})
  const isLoading = ref(false)

  // 当前项目ID
  const currentProjectId = computed(() => projectStore.currentProject?.id || '')
  
  // 当前用户ID
  const currentUserId = computed(() => authStore.user?.id || '')

  // 检查权限
  const can = async (resource: string, action: string): Promise<boolean> => {
    if (!currentProjectId.value || !currentUserId.value) {
      return false
    }

    try {
      const result = await checkPermission(
        currentProjectId.value,
        resource,
        action
      )
      return result.allowed
    } catch (error) {
      console.error('Permission check failed:', error)
      return false
    }
  }

  // 批量检查权限
  const canMultiple = async (permissions: Array<{resource: string, action: string}>): Promise<boolean[]> => {
    if (!currentProjectId.value || !currentUserId.value) {
      return permissions.map(() => false)
    }

    const results = await Promise.all(
      permissions.map(p => can(p.resource, p.action))
    )
    return results
  }

  // 获取当前用户在项目中的角色
  const fetchCurrentRole = async () => {
    if (!currentProjectId.value) {
      currentRole.value = ''
      rolePermissions.value = {}
      return
    }

    isLoading.value = true
    try {
      const roleInfo = await getProjectRoles(currentProjectId.value)
      currentRole.value = roleInfo.role
      rolePermissions.value = roleInfo.permissions
    } catch (error) {
      console.error('Failed to fetch role:', error)
      currentRole.value = ''
      rolePermissions.value = {}
    } finally {
      isLoading.value = false
    }
  }

  // 检查是否有特定权限（同步，基于已获取的角色权限）
  const hasPermission = (resource: string, action: string): boolean => {
    if (!currentRole.value || !rolePermissions.value[resource]) {
      return false
    }
    
    return rolePermissions.value[resource].includes(action)
  }

  // 是否是项目管理员
  const isProjectAdmin = computed(() => currentRole.value === 'project_admin')
  
  // 是否是开发者
  const isDeveloper = computed(() => currentRole.value === 'developer')
  
  // 是否是观察者
  const isViewer = computed(() => currentRole.value === 'viewer')

  // 重置权限状态
  const reset = () => {
    currentRole.value = ''
    rolePermissions.value = {}
  }

  return {
    // state
    currentRole,
    rolePermissions,
    isLoading,
    
    // getters
    currentProjectId,
    currentUserId,
    isProjectAdmin,
    isDeveloper,
    isViewer,
    
    // actions
    can,
    canMultiple,
    fetchCurrentRole,
    hasPermission,
    reset
  }
})