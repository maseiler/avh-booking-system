import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('./components/Home.vue'),
  },
  {
    path: '/booking',
    name: 'Booking',
    component: () => import('./components/Booking.vue'),
  },
  {
    path: '/statistics',
    name: 'Statistics',
    component: () => import('./components/Statistics.vue')
  },
  {
    path: '/feedback',
    name: 'Feedback',
    component: () => import('./components/Feedback.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
