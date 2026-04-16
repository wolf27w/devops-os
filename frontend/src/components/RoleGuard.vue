<template>
  <slot v-if="hasRole" />
  <slot v-else name="no-role">
    <div class="role-denied">
      <el-alert
        title="角色不符"
        type="warning"
        :description="deniedMessage"
        show-icon
        :closable="false"
      />
    </div>
  </slot>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { usePermissionStore } from '@/stores/permission'

interface Props {
  role: string | string[]
  any?: boolean
  deniedMessage?: string
}

const props = withDefaults(defineProps<Props>(), {
  any: false,
  deniedMessage: '您的角色无法访问此内容'
})

const permissionStore = usePermissionStore()
const hasRole = ref(false)

// 检查角色
const checkRole = () => {
  const currentRole = permissionStore.currentRole
  const requiredRoles = Array.isArray(props.role) ? props.role : [props.role]
  
  if (props.any) {
    // 任意一个角色匹配即可
    hasRole.value = requiredRoles.some(role => role === currentRole)
  } else {
    // 需要所有角色都匹配（通常只有一个）
    hasRole.value = requiredRoles.every(role => role === currentRole)
  }
}

// 监听角色变化
watch(() => permissionStore.currentRole, () => {
  checkRole()
})

// 监听角色参数变化
watch(() => props.role, () => {
  checkRole()
})

// 初始检查
checkRole()

// 暴露方法给父组件
defineExpose({
  checkRole,
  hasRole
})
</script>

<style scoped>
.role-denied {
  padding: 20px;
}
</style>