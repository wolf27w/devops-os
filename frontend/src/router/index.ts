import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/LoginView.vue')
    },
    {
      path: '/',
      redirect: '/app/ci'
    },
    {
      path: '/app',
      component: () => import('@/layout/MainLayout.vue'),
      meta: { requiresAuth: true },
      children: [
        {
          path: 'ci',
          name: 'CI',
          component: () => import('@/views/CIView.vue')
        },
        {
          path: 'repo',
          name: 'Repo',
          component: () => import('@/views/RepoView.vue')
        },
        {
          path: 'workflow',
          name: 'Workflow',
          component: () => import('@/views/WorkflowView.vue')
        },
        {
          path: 'db',
          name: 'Database',
          component: () => import('@/views/DBView.vue')
        },
        {
          path: 'wiki',
          name: 'Wiki',
          component: () => import('@/views/WikiView.vue')
        },
        {
          path: 'nav',
          name: 'Navigation',
          component: () => import('@/views/NavView.vue')
        },
        {
          path: 'settings',
          name: 'Settings',
          component: () => import('@/views/SettingsView.vue')
        }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.path === '/login' && authStore.isAuthenticated) {
    next('/app/ci')
  } else {
    next()
  }
})

export default router