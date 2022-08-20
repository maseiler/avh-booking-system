import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import './assets/main.scss';

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faBeer, faGear } from '@fortawesome/free-solid-svg-icons'

library.add(faBeer, faGear)

createApp(App).use(router).component("font-awesome-icon", FontAwesomeIcon).mount('#app')
