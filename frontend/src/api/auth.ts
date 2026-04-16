import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json'
  }
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  (response) => response.data,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error.response?.data || error)
  }
)

interface LoginRequest {
  username: string
  password: string
}

interface LoginResponse {
  token: string
  user: {
    id: string
    username: string
    email?: string
    role: string
  }
}

export const login = (data: LoginRequest): Promise<LoginResponse> => {
  return api.post('/auth/login', data)
}

export const logout = (): Promise<void> => {
  return api.post('/auth/logout')
}

export const getCurrentUser = (): Promise<any> => {
  return api.get('/auth/me')
}