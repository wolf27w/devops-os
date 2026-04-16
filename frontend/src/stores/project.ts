import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getProjects, getProject } from '@/api/project'

interface Project {
  id: string
  name: string
  description: string
  owner_id: string
  status: string
  created_at: string
  updated_at: string
}

export const useProjectStore = defineStore('project', () => {
  const currentProject = ref<Project | null>(null)
  const projectList = ref<Project[]>([])
  const isLoading = ref(false)

  const fetchProjects = async () => {
    isLoading.value = true
    try {
      const projects = await getProjects()
      projectList.value = projects
      
      if (projects.length > 0 && !currentProject.value) {
        currentProject.value = projects[0]
      }
    } finally {
      isLoading.value = false
    }
  }

  const switchProject = (project: Project) => {
    currentProject.value = project
  }

  const fetchProject = async (projectId: string) => {
    isLoading.value = true
    try {
      const project = await getProject(projectId)
      currentProject.value = project
    } finally {
      isLoading.value = false
    }
  }

  return {
    currentProject,
    projectList,
    isLoading,
    fetchProjects,
    switchProject,
    fetchProject
  }
})