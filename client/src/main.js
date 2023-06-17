import { createApp } from 'vue'
import App from './App.vue'

// Router
import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './pages/HomePage.vue'
import Booking from './pages/BookingPage.vue'

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/booking', component: Booking, name: 'booking' },
]
const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

// Stores
import { createPinia } from 'pinia'

// Icons
// TODO refactor to dedicated file
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faUser,
  faUserCircle,
  faPhone,
  faBeer,
  faEnvelope,
  faUserSecret,
  faEuroSign,
  faSearch,
  faExclamation,
  faMoneyBill,
  faFont,
  faExpandArrowsAlt,
  faInfoCircle,
  faTrash,
  faCreditCard,
  faBalanceScale,
  faFingerprint,
  faUtensils,
  faStar,
  faAnchor,
  faGlassWhiskey,
  faBullhorn,
  faHeart,
  faLock,
  faRedo,
  faSadTear,
  faTimes,
  faCheckCircle,
  faMoon,
  faSun,
  faPalette,
  faClock,
} from '@fortawesome/free-solid-svg-icons'

library.add(
  faUser,
  faUserCircle,
  faPhone,
  faBeer,
  faEnvelope,
  faUserSecret,
  faEuroSign,
  faSearch,
  faExclamation,
  faMoneyBill,
  faFont,
  faExpandArrowsAlt,
  faInfoCircle,
  faTrash,
  faCreditCard,
  faBalanceScale,
  faFingerprint,
  faUtensils,
  faStar,
  faAnchor,
  faGlassWhiskey,
  faBullhorn,
  faHeart,
  faLock,
  faRedo,
  faSadTear,
  faTimes,
  faCheckCircle,
  faMoon,
  faSun,
  faPalette,
  faClock
)

// import CSS styles
import './styles/main.scss'

createApp(App)
  .use(router)
  .use(createPinia())
  .component('font-awesome-icon', FontAwesomeIcon)
  .mount('#app')
