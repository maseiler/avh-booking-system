import Vue from 'vue'
import VueRouter from 'vue-router'

import store from './store.js'
import Home from './components/Home.vue'
import Booking from './components/booking/Booking.vue'
import Login from './components/admin/Login.vue'
import Admin from './components/admin/Admin.vue'
import Statistics from './components/statistics/Statistics.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: Home },
  {
    path: '/booking', component: Booking, name: 'booking',
    beforeEnter: (to, from, next) => {
      store.commit("getLast5Bookings");
      store.commit("selectUser", {});
      store.commit("selectSingleItem", {});
      store.commit("selectMultipleItems", []);
      return next();
    }
  },
  { path: '/statistics', component: Statistics },
  {
    path: '/login', component: Login, name: 'login',
  },
  {
    path: '/admin', component: Admin, name: 'admin',
    beforeEnter: (to, from, next) => {
      if (from.name === 'login' || from.name === 'admin') {
        store.commit("selectUser", {});
        store.commit("selectSingleItem", {});
        store.commit("selectMultipleItems", []);
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
