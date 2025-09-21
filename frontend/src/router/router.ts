import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import BookingView from '../views/BookingView.vue'
import PaymentView from '../views/PaymentView.vue'
import ProductSettings from '../views/Settings/ProductSettings.vue'
import ProductSettingsSingle from '../views/Settings/ProductSettingsSingle.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/booking', component: BookingView },
  { path: '/payment', component: PaymentView },
  { path: '/settings/products', component: ProductSettings},
  { name: 'ProductSettingsSingle', path: '/settings/products/:productId', component: ProductSettingsSingle},
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router