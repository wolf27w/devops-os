<template>
  <el-container class="main-container">
    <el-aside width="200px" class="sidebar">
      <div class="logo">
        <h2>DevOps OS</h2>
      </div>
      <el-menu
        :default-active="activeMenu"
        class="sidebar-menu"
        @select="handleMenuSelect"
      >
        <el-menu-item index="/app/ci">
          <el-icon><Operation /></el-icon>
          <span>CI/CD</span>
        </el-menu-item>
        <el-menu-item index="/app/repo">
          <el-icon><Folder /></el-icon>
          <span>Repo</span>
        </el-menu-item>
        <el-menu-item index="/app/workflow">
          <el-icon><SetUp /></el-icon>
          <span>Workflow</span>
        </el-menu-item>
        <el-menu-item index="/app/db">
          <el-icon><DataBoard /></el-icon>
          <span>Database</span>
        </el-menu-item>
        <el-menu-item index="/app/wiki">
          <el-icon><Document /></el-icon>
          <span>Wiki</span>
        </el-menu-item>
        <el-menu-item index="/app/nav">
          <el-icon><Menu /></el-icon>
          <span>Navigation</span>
        </el-menu-item>
        <el-menu-item index="/app/settings">
          <el-icon><Setting /></el-icon>
          <span>Settings</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    
    <el-container>
      <el-header class="header">
        <div class="header-left">
          <el-select
            v-model="currentProjectId"
            placeholder="Select Project"
            class="project-selector"
            @change="handleProjectChange"
          >
            <el-option
              v-for="project in projectList"
              :key="project.id"
              :label="project.name"
              :value="project.id"
            />
          </el-select>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleUserCommand">
            <span class="user-info">
              <el-avatar :size="32">{{ userInitial }}</el-avatar>
              <span class="username">{{ username }}</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">Profile</el-dropdown-item>
                <el-dropdown-item command="logout">Logout</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      
      <el-main class="content">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProjectStore } from '@/stores/project'
import { useAuthStore } from '@/stores/auth'
import {
  Operation,
  Folder,
  SetUp,
  DataBoard,
  Document,
  Menu,
  Setting
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const projectStore = useProjectStore()
const authStore = useAuthStore()

const activeMenu = computed(() => route.path)
const currentProjectId = computed({
  get: () => projectStore.currentProject?.id || '',
  set: (value) => {
    const project = projectStore.projectList.find(p => p.id === value)
    if (project) {
      projectStore.switchProject(project)
    }
  }
})

const projectList = computed(() => projectStore.projectList)
const username = computed(() => authStore.user?.username || '')
const userInitial = computed(() => username.value.charAt(0).toUpperCase())

const handleMenuSelect = (index: string) => {
  router.push(index)
}

const handleProjectChange = (projectId: string) => {
  const project = projectStore.projectList.find(p => p.id === projectId)
  if (project) {
    projectStore.switchProject(project)
    router.push('/app/ci')
  }
}

const handleUserCommand = (command: string) => {
  if (command === 'logout') {
    authStore.logout()
    router.push('/login')
  } else if (command === 'profile') {
    router.push('/profile')
  }
}

onMounted(async () => {
  await projectStore.fetchProjects()
})
</script>

<style scoped>
.main-container {
  height: 100vh;
}

.sidebar {
  background-color: #304156;
  color: #fff;
}

.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1px solid #2c3e50;
}

.logo h2 {
  margin: 0;
  color: #fff;
  font-size: 18px;
}

.sidebar-menu {
  border-right: none;
  background-color: #304156;
}

.sidebar-menu :deep(.el-menu-item) {
  color: #bfcbd9;
}

.sidebar-menu :deep(.el-menu-item.is-active) {
  background-color: #263445;
  color: #409eff;
}

.sidebar-menu :deep(.el-menu-item:hover) {
  background-color: #263445;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #fff;
  border-bottom: 1px solid #e6e6e6;
  padding: 0 20px;
}

.header-left {
  display: flex;
  align-items: center;
}

.project-selector {
  width: 200px;
}

.header-right {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.username {
  margin-left: 10px;
  font-weight: 500;
}

.content {
  background-color: #f0f2f5;
  padding: 20px;
}
</style>