import { createApp } from 'vue'
import './assets/style.scss'
import router from './router/router.ts'
import App from './App.vue'
import { createI18n } from 'vue-i18n'
import de  from './locales/de.ts'
import en from './locales/en.ts'
import { createPinia } from 'pinia'

function loadLocaleMessages() {
  //https://vue-i18n.intlify.dev/guide/essentials/started.html
  const messages = {de, en}
  return messages
}

function loadLocaleNumbers(){
  // const locales = require.context('./locales/numbers', true, /[A-Za-z0-9-_,\s]+\.json$/i)
  // const numbers = {}
  // locales.keys().forEach(key => {
  //   const matched = key.match(/([A-Za-z0-9-_]+)\./i)
  //   if (matched && matched.length > 1) {
  //     const locale = matched[1]
  //     numbers[locale] = locales(key)
  //   }
  // })
  // return numbers
  const numberFormats = {
    "de-DE": {
      "currency": {
        "style": "currency",
        "currency": "EUR",
        "notation": "standard",
        "currencyDisplay":"symbol"
      },
      "decimal": {
        "style": "decimal",
        "minimumFractionDigits": 2,
        "maximumFractionDigits": 2
      },
      "percent": {
        "style": "percent",
        "useGrouping": false
      }
    },
    "en-US":{
      "currency":{
        "style": "currency",
        "currency": "USD",
        "notation":"standard",
        "currencyDisplay":"symbol"
      }
    }
  }
  return numberFormats
}

const i18n = createI18n({
  legacy: true,
  locale: 'de',
  fallbackLocale: 'en',
  messages: loadLocaleMessages(),
  numberFormats: loadLocaleNumbers(),
})

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'

library.add( fas )

const pinia = createPinia()
const app = createApp(App)

app.component('icon', FontAwesomeIcon)
app.use(router)
app.use(i18n)
app.use(pinia)
app.mount('#app')
