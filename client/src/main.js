import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
import store from './store.js'
import mixin from './mixin.js'
import router from './router.js'
import Chart from 'chart.js';

// font awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon, FontAwesomeLayers, FontAwesomeLayersText } from '@fortawesome/vue-fontawesome'

import {
  faUser, faUserCircle, faPhone, faBeer, faEnvelope, faUserSecret, faEuroSign, faSearch, faExclamation,
  faMoneyBill, faFont, faExpandArrowsAlt, faInfoCircle, faTrash, faCreditCard, faBalanceScale, faFingerprint,
  faUtensils, faStar, faAnchor, faGlassWhiskey, faBullhorn, faHeart, faLock, faRedo, faSadTear, faTimes, faCheckCircle, faMoon, faSun, faPalette, faClock, faSpinner, faKey, faCashRegister,
  faServer, faSign,
} from '@fortawesome/free-solid-svg-icons'

import './registerServiceWorker'
import i18n from './i18n'

library.add(faUser, faUserCircle, faPhone, faBeer, faEnvelope, faUserSecret, faEuroSign, faSearch, faExclamation,
  faMoneyBill, faFont, faExpandArrowsAlt, faInfoCircle, faTrash, faCreditCard, faBalanceScale, faFingerprint,
  faUtensils, faStar, faAnchor, faGlassWhiskey, faBullhorn, faHeart, faLock, faRedo, faSadTear, faTimes, faCheckCircle, faMoon, faSun, faPalette, faClock, faSpinner, faKey, faCashRegister,
  faServer, faSign)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.component('font-awesome-layers', FontAwesomeLayers)
Vue.component('font-awesome-layers-text', FontAwesomeLayersText)


Vue.use(VueResource)

require("./assets/main.scss")

Vue.config.devtools = true

new Vue({
  render: h => h(App),
  store,
  router,
  mixin,
  i18n,
  Chart
}).$mount('#app')
