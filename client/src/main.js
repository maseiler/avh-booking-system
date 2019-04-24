import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import Home from './components/Home.vue'
import Booking from './components/booking/Booking.vue'
import Admin from './components/admin/Admin.vue'
import Statistics from './components/statistics/Statistics.vue'
// font awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon, FontAwesomeLayers, FontAwesomeLayersText } from '@fortawesome/vue-fontawesome'
import {
  faUser, faUserCircle, faPhone, faBeer, faEnvelope, faUserSecret, faEuroSign, faSearch, faExclamation,
  faMoneyBill, faFont, faExpandArrowsAlt, faInfoCircle, faTrash, faSortNumericDown, faCreditCard
} from '@fortawesome/free-solid-svg-icons'

library.add(faUser, faUserCircle, faPhone, faBeer, faEnvelope, faUserSecret, faEuroSign, faSearch, faExclamation,
  faMoneyBill, faFont, faExpandArrowsAlt, faInfoCircle, faTrash, faSortNumericDown, faCreditCard)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.component('font-awesome-layers', FontAwesomeLayers)
Vue.component('font-awesome-layers-text', FontAwesomeLayersText)

Vue.use(VueRouter)
const routes = [
  { path: '/', component: Home },
  { path: '/booking', component: Booking },
  { path: '/statistics', component: Statistics },
  { path: '/admin', component: Admin }
]

Vue.use(VueResource)

const router = new VueRouter({
  routes,
  mode: 'history'
})

require("./assets/main.scss")

const mixin = Vue.mixin({
  methods: {
    displayUserName: function (user) {
      if (user === undefined) {
        return "???";
      } else
        if (user.BierName !== "") {
          return user.BierName;
        } else if (user.LastName !== "" && user.FirstName !== "") {
          return user.FirstName + " " + user.LastName[0] + ".";
        } else if (user.FirstName !== "") {
          return user.FirstName
        } else {
          return "???";
        }
    },
    displayItem: function (item) {
      if (item === undefined) {
        return "???"
      } else if (item.Type === "alcoholic" || item.Type === "non-alcoholic") {
        return item.Name + " " + item.Size + " " + item.Unit;
      } else if (item.Type === "boat" || item.Type === "food" || item.Type === "Admin") {
        return item.Name;
      } else {
        return "???";
      }
    }
  }
})

new Vue({
  render: h => h(App),
  // el: '#app',
  // template: '<App/>',
  // components: {App, Booking},
  router,
  mixin
}).$mount('#app')
