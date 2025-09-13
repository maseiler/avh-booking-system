import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import BookingView from '../views/BookingView.vue'
import PaymentView from '../views/PaymentView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/booking', component: BookingView },
  { path: '/payment', component: PaymentView },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router