import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import BookingView from '../views/BookingView.vue'
import PaymentView from '../views/PaymentView.vue'
import ProductSettings from '../views/Settings/ProductSettings.vue'
import ProductSettingsSingle from '../views/Settings/ProductSettingsSingle.vue'
import AccountSettings from '../views/Settings/AccountSettings.vue'
import AccountSettingsSingle from '../views/Settings/AccountSettingsSingle.vue'
import CategorySettingsSingle from '../views/Settings/CategorySettingsSingle.vue'
import PaymentSettings from '../views/Settings/PaymentSettings.vue'
import ClientSettings from '../views/Settings/ClientSettings.vue'
import OrderView from '../views/OrderView.vue'
import ApplicationSettings from '../views/Settings/ApplicationSettings.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/booking', component: BookingView },
  { path: '/payment', component: PaymentView },
  { name: 'ProductSettings', path: '/settings/products', component: ProductSettings},
  { name: 'ProductSettingsSingle', path: '/settings/products/:productId', component: ProductSettingsSingle},
  { name: 'ProductSettingsAdd', path: '/settings/products/add', component: ProductSettingsSingle},
  { name: 'CategorySettingsAdd', path: '/settings/categories/:type/add', component: CategorySettingsSingle },
  { name: 'AccountSettings', path: '/settings/accounts', component: AccountSettings},
  { name: 'AccountSettingsSingle', path: '/settings/accounts/:accountId', component: AccountSettingsSingle},
  { name: 'AccountSettingsAdd', path: '/settings/accounts/add', component: AccountSettingsSingle},
  { path: '/settings/payments', component: PaymentSettings},
  { name: 'ClientSettings', path: '/settings/client', component: ClientSettings},
  { name: 'Orders', path: '/orders/', component: OrderView},
  { name: 'ApplicationSettings', path: '/settings/application', component: ApplicationSettings}
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router