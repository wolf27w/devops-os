<template>
  <slot v-if="hasPermission" />
  <slot v-else name="no-permission">
    <div class="permission-denied">
      <el-alert
        title="权限不足"
        type="warning"
        :description="deniedMessage"
        show-icon
        :closable="false"
      />
    </div>
  </slot>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { usePermissionStore } from '@/stores/permission'

interface Props {
  resource: string
  action: string
  projectId?: string
  showAlert?: boolean
  deniedMessage?: string
}

const props = withDefaults(defineProps<Props>(), {
  showAlert: true,
  deniedMessage: '您没有执行此操作的权限'
})

const permissionStore = usePermissionStore()
const hasPermission = ref(false)
const isLoading = ref(true)

// 计算属性：当前项目ID
const currentProjectId = computed(() => {
  return props.projectId || permissionStore.currentProjectId
})

// 检查权限
const checkPermission = async () => {
  if (!currentProjectId.value) {
    hasPermission.value = false
    isLoading.value = false
    return
  }

  isLoading.value = true
  try {
    hasPermission.value = await permissionStore.can(props.resource, props.action)
  } catch (error) {
    console.error('Permission check failed:', error)
    hasPermission.value = false
  } finally {
    isLoading.value = false
  }
}

// 监听项目变化
watch(() => permissionStore.currentProjectId, () => {
  checkPermission()
})

// 监听权限参数变化
watch(() => [props.resource, props.action, props.projectId], () => {
  checkPermission()
})

onMounted(() => {
  checkPermission()
})

// 暴露方法给父组件
defineExpose({
  checkPermission,
  hasPermission,
  isLoading
})
</script>

<style scoped>
.permission-denied {
  padding: 20px;
}
</style>