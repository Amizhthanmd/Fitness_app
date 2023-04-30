// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '/Home',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
      },
      {
        path: '/Calculate',
        name: 'Calculate',
        component: () => import('@/views/Calculate.vue'),
      },
      {
        path: '/Exercise',
        name: 'Exercise',
        component: () => import('@/views/Exercise.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
