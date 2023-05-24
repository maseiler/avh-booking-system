import Vue from 'vue'
import VueRouter from 'vue-router'

import store from './store.js'
import Home from './pages/Home.vue'
import Booking from './pages/Booking.vue'
import Login from './components/admin/LoginModal.vue'
import Admin from './pages/Admin.vue'
import Statistics from './pages/Statistics.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: Home, name: 'home' },
  {
    path: '/booking', component: Booking, name: 'booking',
    beforeEnter: (to, from, next) => {
      store.commit("getLastNBookEntries", 5);
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
