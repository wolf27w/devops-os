import api from './index'

export const getProjects = (): Promise<any[]> => {
  return api.get('/projects')
}

export const getProject = (projectId: string): Promise<any> => {
  return api.get(`/projects/${projectId}`)
}

export const createProject = (data: any): Promise<any> => {
  return api.post('/projects', data)
}

export const getProjectMembers = (projectId: string): Promise<any[]> => {
  return api.get(`/projects/${projectId}/members`)
}

export const addProjectMember = (projectId: string, data: any): Promise<any> => {
  return api.post(`/projects/${projectId}/members`, data)
}