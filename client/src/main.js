import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import Home from './components/Home.vue'
import Booking from './components/Booking.vue'
import Admin from './components/Admin.vue'
import Statistics from './components/Statistics.vue'
// font awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { faUser} from '@fortawesome/free-solid-svg-icons'
import { faPhone} from '@fortawesome/free-solid-svg-icons'
import { faBeer} from '@fortawesome/free-solid-svg-icons'
import { faEnvelope } from '@fortawesome/free-solid-svg-icons'
import { faUserSecret } from '@fortawesome/free-solid-svg-icons'
import { faEuroSign } from '@fortawesome/free-solid-svg-icons'
import { faSearch } from '@fortawesome/free-solid-svg-icons'
import { faExclamation } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon, FontAwesomeLayers, FontAwesomeLayersText } from '@fortawesome/vue-fontawesome'

library.add(faUser)
library.add(faPhone)
library.add(faBeer)
library.add(faEnvelope)
library.add(faUserSecret)
library.add(faEuroSign)
library.add(faSearch)
library.add(faExclamation)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.component('font-awesome-layers', FontAwesomeLayers)
Vue.component('font-awesome-layers-text', FontAwesomeLayersText)

Vue.use(VueRouter)
const routes = [
  { path: '/', component: Home },
  { path: '/booking', component: Booking },
  { path: '/statistics', component: Statistics},
  { path: '/admin', component: Admin},
]

Vue.use(VueResource)

const router = new VueRouter({
  routes,
  mode: 'history'
})

require("./assets/main.scss")

new Vue({
  render: h => h(App),
  // el: '#app',
  // template: '<App/>',
  // components: {App, Booking},
  router,
}).$mount('#app')
