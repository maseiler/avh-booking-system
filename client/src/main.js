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

function sortUsersByName(array) {
  array.sort(function (a, b) {
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

function sortItemsByName(array) {
  array.sort(function (a, b) {
    var nameA = a["Name"].toLowerCase(),
      nameB = b["Name"].toLowerCase();
    var sizeA = a.Size,
      sizeB = b.Size;
    if (nameA < nameB) return -1;
    if (nameA > nameB) return 1;
    if (sizeA < sizeB) return -1;
    if (sizeA > sizeB) return 1;
    return 0;
  });
}

const store = new Vuex.Store({
  modules: {
    mixin: mixin
  },
  state: {
    users: [],
    items: [],
    bookings: [],
    feedback: []
  },
  mutations: {
    getUsers(state) {
      Vue.http.get("/getUsers").then(response => {
        var users = [].concat.apply([], response.body);
        sortUsersByName(users);
        state.users = users
      })
    },
    getItems(state) {
      Vue.http.get("/getUnreservedItems").then(response => {
        var items = [].concat.apply([], response.body);
        sortItemsByName(items)
        state.items = items
      });
    },
    getLastNBookings(state, n) {
      Vue.http.post("/getLastNBookings", n).then(response => {
        state.bookings = [].concat.apply([], response.body);
      });
    },
    getFeedbackList(state) {
      Vue.http.get("/getFeedback").then(response => {
        state.feedback = [].concat.apply([], response.body);
      });
    },
  },
  getters: {
    usersAH: state => {
      return state.users.filter(user => user["Status"] === "AH");
    },
    usersAktivB: state => {
      return state.users.filter(user => user["Status"] === "Aktiv B");
    },
    usersAktivKA: state => {
      return state.users.filter(user => user["Status"] === "Aktiv KA");
    },
    usersSteganleger: state => {
      return state.users.filter(user => user["Status"] === "Steganleger");
    },
    usersGast: state => {
      return state.users.filter(user => user["Status"] === "Gast");
    },
    itemsAlc: state => {
      return state.items.filter(item => item["Type"] === "alcoholic")
    },
    itemsNonAlc: state => {
      return state.items.filter(item => item["Type"] === "non-alcoholic")
    },
    itemsFood: state => {
      return state.items.filter(item => item["Type"] === "food")
    },
    itemsBoat: state => {
      return state.items.filter(item => item["Type"] === "boat")
    }
  }
})

store.commit("getUsers");
store.commit("getItems");
store.commit("getFeedbackList");
store.commit("getLastNBookings", 50);

new Vue({
  render: h => h(App),
  store,
  router,
  mixin
}).$mount('#app')
