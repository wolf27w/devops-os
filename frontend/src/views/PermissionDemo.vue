<template>
  <div class="permission-demo">
    <h2>权限系统演示</h2>
    
    <el-card class="demo-card">
      <template #header>
        <div class="card-header">
          <h3>当前权限状态</h3>
          <el-button type="primary" @click="refreshRole" :loading="isLoading">
            刷新权限
          </el-button>
        </div>
      </template>
      
      <div class="status-info">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="当前项目">
            {{ projectStore.currentProject?.name || '未选择项目' }}
          </el-descriptions-item>
          <el-descriptions-item label="当前角色">
            <el-tag :type="roleTagType">
              {{ permissionStore.currentRole || '未知' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="用户ID">
            {{ authStore.user?.id }}
          </el-descriptions-item>
          <el-descriptions-item label="用户名">
            {{ authStore.user?.username }}
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </el-card>
    
    <el-card class="demo-card">
      <template #header>
        <h3>权限检查演示</h3>
      </template>
      
      <div class="permission-checks">
        <div class="check-item">
          <h4>CI/CD 模块权限</h4>
          <div class="check-buttons">
            <el-button 
              :type="ciViewAllowed ? 'success' : 'info'"
              @click="checkCIPermission('view')"
            >
              查看 CI
            </el-button>
            <el-button 
              :type="ciCreateAllowed ? 'success' : 'info'"
              @click="checkCIPermission('create')"
            >
              创建 CI
            </el-button>
            <el-button 
              :type="ciUpdateAllowed ? 'success' : 'info'"
              @click="checkCIPermission('update')"
            >
              更新 CI
            </el-button>
            <el-button 
              :type="ciDeleteAllowed ? 'success' : 'info'"
              @click="checkCIPermission('delete')"
            >
              删除 CI
            </el-button>
          </div>
        </div>
        
        <div class="check-item">
          <h4>成员管理权限</h4>
          <div class="check-buttons">
            <el-button 
              :type="memberViewAllowed ? 'success' : 'info'"
              @click="checkMemberPermission('view')"
            >
              查看成员
            </el-button>
            <el-button 
              :type="memberCreateAllowed ? 'success' : 'info'"
              @click="checkMemberPermission('create')"
            >
              添加成员
            </el-button>
          </div>
        </div>
      </div>
    </el-card>
    
    <el-card class="demo-card">
      <template #header>
        <h3>组件权限控制演示</h3>
      </template>
      
      <div class="component-demo">
        <h4>1. PermissionGuard 组件</h4>
        <PermissionGuard resource="ci" action="create">
          <el-button type="primary" icon="Plus">
            创建 CI 流水线（只有有权限的用户能看到）
          </el-button>
          <template #no-permission>
            <el-alert
              title="权限提示"
              type="info"
              description="您没有创建 CI 流水线的权限"
              show-icon
            />
          </template>
        </PermissionGuard>
        
        <h4>2. RoleGuard 组件</h4>
        <RoleGuard role="project_admin">
          <el-button type="warning" icon="Setting">
            项目设置（只有项目管理员能看到）
          </el-button>
          <template #no-role>
            <el-alert
              title="角色提示"
              type="info"
              description="只有项目管理员可以访问项目设置"
              show-icon
            />
          </template>
        </RoleGuard>
        
        <h4>3. 指令方式</h4>
        <div>
          <el-button 
            v-permission="{ resource: 'member', action: 'create' }"
            type="success"
            icon="User"
          >
            添加成员（使用指令控制）
          </el-button>
          
          <el-button 
            v-role="'project_admin'"
            type="danger"
            icon="Delete"
          >
            删除项目（只有项目管理员能看到）
          </el-button>
        </div>
      </div>
    </el-card>
    
    <el-card class="demo-card">
      <template #header>
        <h3>当前角色权限详情</h3>
      </template>
      
      <div class="permission-details">
        <el-table :data="permissionTableData" style="width: 100%">
          <el-table-column prop="resource" label="资源" width="150" />
          <el-table-column prop="actions" label="可用操作">
            <template #default="{ row }">
              <el-tag
                v-for="action in row.actions"
                :key="action"
                class="action-tag"
                :type="getActionTagType(action)"
              >
                {{ action }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus, Setting, User, Delete } from '@element-plus/icons-vue'
import { usePermissionStore } from '@/stores/permission'
import { useAuthStore } from '@/stores/auth'
import { useProjectStore } from '@/stores/project'
import PermissionGuard from '@/components/PermissionGuard.vue'
import RoleGuard from '@/components/RoleGuard.vue'
import { permissionDirective, roleDirective } from '@/utils/permission'

// 注册指令
const vPermission = permissionDirective
const vRole = roleDirective

const permissionStore = usePermissionStore()
const authStore = useAuthStore()
const projectStore = useProjectStore()

const isLoading = ref(false)
const ciViewAllowed = ref(false)
const ciCreateAllowed = ref(false)
const ciUpdateAllowed = ref(false)
const ciDeleteAllowed = ref(false)
const memberViewAllowed = ref(false)
const memberCreateAllowed = ref(false)

// 角色标签类型
const roleTagType = computed(() => {
  switch (permissionStore.currentRole) {
    case 'project_admin':
      return 'danger'
    case 'developer':
      return 'warning'
    case 'viewer':
      return 'success'
    default:
      return 'info'
  }
})

// 权限表格数据
const permissionTableData = computed(() => {
  const permissions = permissionStore.rolePermissions
  return Object.entries(permissions).map(([resource, actions]) => ({
    resource,
    actions
  }))
})

// 操作标签类型
const getActionTagType = (action: string) => {
  switch (action) {
    case 'view':
      return 'info'
    case 'create':
      return 'success'
    case 'update':
      return 'warning'
    case 'delete':
      return 'danger'
    case 'execute':
      return 'primary'
    case 'approve':
      return ''
    default:
      return 'info'
  }
}

// 刷新角色
const refreshRole = async () => {
  isLoading.value = true
  try {
    await permissionStore.fetchCurrentRole()
    ElMessage.success('权限已刷新')
  } catch (error) {
    ElMessage.error('刷新权限失败')
  } finally {
    isLoading.value = false
  }
}

// 检查 CI 权限
const checkCIPermission = async (action: string) => {
  if (!projectStore.currentProject) {
    ElMessage.warning('请先选择项目')
    return
  }
  
  try {
    const allowed = await permissionStore.can('ci', action)
    
    switch (action) {
      case 'view':
        ciViewAllowed.value = allowed
        break
      case 'create':
        ciCreateAllowed.value = allowed
        break
      case 'update':
        ciUpdateAllowed.value = allowed
        break
      case 'delete':
        ciDeleteAllowed.value = allowed
        break
    }
    
    if (allowed) {
      ElMessage.success(`您有 ${action} CI 的权限`)
    } else {
      ElMessage.warning(`您没有 ${action} CI 的权限`)
    }
  } catch (error) {
    ElMessage.error('权限检查失败')
  }
}

// 检查成员权限
const checkMemberPermission = async (action: string) => {
  if (!projectStore.currentProject) {
    ElMessage.warning('请先选择项目')
    return
  }
  
  try {
    const allowed = await permissionStore.can('member', action)
    
    switch (action) {
      case 'view':
        memberViewAllowed.value = allowed
        break
      case 'create':
        memberCreateAllowed.value = allowed
        break
    }
    
    if (allowed) {
      ElMessage.success(`您有 ${action} 成员的权限`)
    } else {
      ElMessage.warning(`您没有 ${action} 成员的权限`)
    }
  } catch (error) {
    ElMessage.error('权限检查失败')
  }
}

// 初始化时获取角色
onMounted(() => {
  if (projectStore.currentProject) {
    permissionStore.fetchCurrentRole()
  }
})
</script>

<style scoped>
.permission-demo {
  padding: 20px;
}

.demo-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.status-info {
  margin-top: 10px;
}

.permission-checks {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.check-item h4 {
  margin-bottom: 10px;
  color: #606266;
}

.check-buttons {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.component-demo {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.component-demo h4 {
  margin: 10px 0;
  color: #606266;
}

.permission-details {
  margin-top: 10px;
}

.action-tag {
  margin-right: 5px;
  margin-bottom: 5px;
}
</style>