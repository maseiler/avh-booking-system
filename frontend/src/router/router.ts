import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import BookingView from '../views/BookingView.vue'
import PaymentView from '../views/PaymentView.vue'
import ProductSettings from '../views/Settings/ProductSettings.vue'
import ProductSettingsSingle from '../views/Settings/ProductSettingsSingle.vue'
import AccountSettings from '../views/Settings/AccountSettings.vue'
import AccountSettingsSingle from '../views/Settings/AccountSettingsSingle.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/booking', component: BookingView },
  { path: '/payment', component: PaymentView },
  { path: '/settings/products', component: ProductSettings},
  { name: 'ProductSettingsSingle', path: '/settings/products/:productId', component: ProductSettingsSingle},
  { path: '/settings/accounts', component: AccountSettings},
  { name: 'AccountSettingsSingle', path: '/settings/accounts/:accountId', component: AccountSettingsSingle},
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router