import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'

import './assets/main.scss'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faAnchor, faBeer, faEnvelope, faGear, faPhone, faSearch, faUser } from '@fortawesome/free-solid-svg-icons'

library.add(faAnchor, faBeer, faEnvelope, faGear, faPhone, faSearch, faUser)

const app = createApp(App)
app.use(router).use(createPinia())
app.component("font-awesome-icon", FontAwesomeIcon)
app.mount('#app')
