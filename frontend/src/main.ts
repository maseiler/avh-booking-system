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


/**
 * Communication between Websocket and Pinia Stores
 * Could be moved into the Stores by writing a subscribe method with this.$subscribe
 * this needs to be called from within the vue app context tho
 * so it could be in App.vue and called from the onMounted() function
 */
import { useAccountStore } from './store/AccountStore.ts'
import { useSocketStore } from './store/socketStore.ts'
const accountStore = useAccountStore();
const socketStore = useSocketStore();

accountStore.$subscribe((mutation, state) => {
  console.log("mutation: ", mutation)

  // Called when an Account is added in the Store
  if(mutation.storeId === "account" && mutation.events.type === "add"){
    const table = "account";
    const operation = "insert";
    const values = mutation.events.newValue;
    let payload = {
      "operation": operation,
      "table": table,
      "values": values
    }
    let msg = {type: "mutation", payload: payload};
    socketStore.wsClient.send(msg)
  }

  // Called when an Account gets Modified
  if(mutation.storeId === "account" && mutation.events.type === "set" && mutation.events.key != "accounts" && mutation.events.key != "selected"){
    const table = "account";
    const operation = "update";
    const values = mutation.events.target;
    const where = {"account_id": mutation.events.target.id.toString()};
    let payload = {
      "operation": operation,
      "table": table,
      "where": where,
      "values": values
    }
    let msg = {type: "mutation", payload: payload};
    socketStore.wsClient.send(JSON.stringify(msg));
  }
})
