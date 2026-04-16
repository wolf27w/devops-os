import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login as apiLogin, logout as apiLogout, getCurrentUser } from '@/api/auth'

interface User {
  id: string
  username: string
  email?: string
  role: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(null)
  const isAuthenticated = ref(!!token.value)

  const setToken = (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
    isAuthenticated.value = true
  }

  const setUser = (newUser: User) => {
    user.value = newUser
  }

  const clearAuth = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    isAuthenticated.value = false
  }

  const login = async (username: string, password: string) => {
    try {
      const response = await apiLogin({ username, password })
      setToken(response.token)
      setUser(response.user)
      return response
    } catch (error) {
      clearAuth()
      throw error
    }
  }

  const logout = async () => {
    try {
      await apiLogout()
    } finally {
      clearAuth()
    }
  }

  const fetchCurrentUser = async () => {
    if (!token.value) return
    
    try {
      const userData = await getCurrentUser()
      setUser(userData)
    } catch (error) {
      clearAuth()
      throw error
    }
  }

  return {
    token,
    user,
    isAuthenticated,
    login,
    logout,
    fetchCurrentUser,
    clearAuth
  }
})