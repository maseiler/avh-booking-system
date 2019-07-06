import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from './components/Home.vue'
import Booking from './components/booking/Booking.vue'
import Login from './components/admin/Login.vue'
import Admin from './components/admin/Admin.vue'
import Statistics from './components/statistics/Statistics.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: Home },
  { path: '/booking', component: Booking },
  { path: '/statistics', component: Statistics },
  {
    path: '/login', component: Login, name: 'login',
  },
  {
    path: '/admin', component: Admin, name: 'admin',
    beforeEnter: (to, from, next) => {
      if (from.name === 'login' || from.name === 'admin') {
        return next();
      }
      return next({ path: '/' });
    }
  }
]

const router = new VueRouter({
  routes,
  mode: 'hash'
})

export default router
