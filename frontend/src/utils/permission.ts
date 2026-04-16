import { usePermissionStore } from '@/stores/permission'

/**
 * 权限控制工具函数
 */

// 权限检查函数
export function usePermission() {
  const permissionStore = usePermissionStore()
  
  return {
    // 检查是否有权限
    can: permissionStore.can,
    
    // 同步检查权限（基于已获取的角色）
    has: permissionStore.hasPermission,
    
    // 检查是否是项目管理员
    isProjectAdmin: permissionStore.isProjectAdmin,
    
    // 检查是否是开发者
    isDeveloper: permissionStore.isDeveloper,
    
    // 检查是否是观察者
    isViewer: permissionStore.isViewer,
    
    // 获取当前角色
    getRole: () => permissionStore.currentRole,
    
    // 获取角色权限
    getPermissions: () => permissionStore.rolePermissions,
    
    // 刷新角色信息
    refreshRole: permissionStore.fetchCurrentRole
  }
}

// 权限指令
export const permissionDirective = {
  mounted(el: HTMLElement, binding: any) {
    const { value } = binding
    const permissionStore = usePermissionStore()
    
    if (!value) return
    
    const { resource, action } = value
    
    if (!permissionStore.hasPermission(resource, action)) {
      el.style.display = 'none'
    }
  },
  
  updated(el: HTMLElement, binding: any) {
    const { value } = binding
    const permissionStore = usePermissionStore()
    
    if (!value) return
    
    const { resource, action } = value
    
    if (permissionStore.hasPermission(resource, action)) {
      el.style.display = ''
    } else {
      el.style.display = 'none'
    }
  }
}

// 角色指令
export const roleDirective = {
  mounted(el: HTMLElement, binding: any) {
    const { value } = binding
    const permissionStore = usePermissionStore()
    
    if (!value) return
    
    const allowedRoles = Array.isArray(value) ? value : [value]
    
    if (!allowedRoles.includes(permissionStore.currentRole)) {
      el.style.display = 'none'
    }
  },
  
  updated(el: HTMLElement, binding: any) {
    const { value } = binding
    const permissionStore = usePermissionStore()
    
    if (!value) return
    
    const allowedRoles = Array.isArray(value) ? value : [value]
    
    if (allowedRoles.includes(permissionStore.currentRole)) {
      el.style.display = ''
    } else {
      el.style.display = 'none'
    }
  }
}

// 权限守卫（用于路由）
export function createPermissionGuard(to: any, from: any, next: any) {
  const permissionStore = usePermissionStore()
  
  // 如果路由需要特定权限
  if (to.meta?.requiresPermission) {
    const { resource, action } = to.meta.requiresPermission
    
    if (!permissionStore.hasPermission(resource, action)) {
      next({ name: 'Forbidden' })
      return
    }
  }
  
  // 如果路由需要特定角色
  if (to.meta?.requiresRole) {
    const requiredRoles = Array.isArray(to.meta.requiresRole) 
      ? to.meta.requiresRole 
      : [to.meta.requiresRole]
    
    if (!requiredRoles.includes(permissionStore.currentRole)) {
      next({ name: 'Forbidden' })
      return
    }
  }
  
  next()
}

// 权限混入
export const permissionMixin = {
  computed: {
    $can() {
      const permissionStore = usePermissionStore()
      return permissionStore.can.bind(permissionStore)
    },
    
    $hasPermission() {
      const permissionStore = usePermissionStore()
      return permissionStore.hasPermission.bind(permissionStore)
    },
    
    $isProjectAdmin() {
      const permissionStore = usePermissionStore()
      return permissionStore.isProjectAdmin
    },
    
    $isDeveloper() {
      const permissionStore = usePermissionStore()
      return permissionStore.isDeveloper
    },
    
    $isViewer() {
      const permissionStore = usePermissionStore()
      return permissionStore.isViewer
    },
    
    $currentRole() {
      const permissionStore = usePermissionStore()
      return permissionStore.currentRole
    }
  },
  
  methods: {
    $checkPermission(resource: string, action: string) {
      const permissionStore = usePermissionStore()
      return permissionStore.can(resource, action)
    },
    
    $refreshRole() {
      const permissionStore = usePermissionStore()
      return permissionStore.fetchCurrentRole()
    }
  }
}