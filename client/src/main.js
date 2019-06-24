import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import Vuex from 'vuex'
import Home from './components/Home.vue'
import Booking from './components/booking/Booking.vue'
import Login from './components/admin/Login.vue'
import Admin from './components/admin/Admin.vue'
import Statistics from './components/statistics/Statistics.vue'
// font awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon, FontAwesomeLayers, FontAwesomeLayersText } from '@fortawesome/vue-fontawesome'
import {
  faUser, faUserCircle, faPhone, faBeer, faEnvelope, faUserSecret, faEuroSign, faSearch, faExclamation,
  faMoneyBill, faFont, faExpandArrowsAlt, faInfoCircle, faTrash, faCreditCard, faBalanceScale, faFingerprint, faUtensils, faStar, faAnchor, faGlassWhiskey, faBullhorn, faHeart
} from '@fortawesome/free-solid-svg-icons'

library.add(faUser, faUserCircle, faPhone, faBeer, faEnvelope, faUserSecret, faEuroSign, faSearch, faExclamation,
  faMoneyBill, faFont, faExpandArrowsAlt, faInfoCircle, faTrash, faCreditCard, faBalanceScale, faFingerprint, faUtensils, faStar, faAnchor, faGlassWhiskey, faBullhorn, faHeart)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.component('font-awesome-layers', FontAwesomeLayers)
Vue.component('font-awesome-layers-text', FontAwesomeLayersText)

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

Vue.use(VueResource)

const router = new VueRouter({
  routes,
  mode: 'history'
})

require("./assets/main.scss")

Vue.config.devtools = true

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

// TODO refacore in own file
Vue.use(Vuex)

function sortByName(array) {
  array.sort(function(a, b) {
    var bierNameA = a.BierName.toLowerCase(),
      bierNameB = b.BierName.toLowerCase();
    var firstNameA = a.FirstName.toLowerCase(),
      firstNameB = b.FirstName.toLowerCase();
    var lastNameA = a.LastName.toLowerCase(),
      lastNameB = b.LastName.toLowerCase();
    if (bierNameA < bierNameB) return -1;
    if (bierNameA > bierNameB) return 1;
    if (firstNameA < firstNameB) return -1;
    if (firstNameA > firstNameB) return 1;
    if (lastNameA < lastNameB) return -1;
    if (lastNameA > lastNameB) return 1;
    return 0;
  });
}

const store = new Vuex.Store({
  state: {
    storeUsers: []
  },
  mutations: {
    getUsers(state) {
      console.log("called getUsers in store")
      var users
      Vue.http.get("/getUsers").then(response => {
        users = [].concat.apply([], response.body);
        sortByName(users);
        console.log(users)
        state.storeUsers = users
        console.log(state.storeUsers)
      })
      // console.log(users)
      // state.storeUsers = users
    }
  },
})

store.commit("getUsers");

new Vue({
  render: h => h(App),
  // el: '#app',
  // template: '<App/>',
  // components: {App, Booking},
  store,
  router,
  mixin
}).$mount('#app')
