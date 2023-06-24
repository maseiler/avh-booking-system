import Vue from 'vue'
import VueI18n from 'vue-i18n'
import { pluginOptions } from '../vue.config'

Vue.use(VueI18n)

function loadLocaleMessages () {
  const locales = require.context('./locales', true, /[A-Za-z0-9-_,\s]+\.json$/i)
  const messages = {}
  locales.keys().forEach(key => {
    const matched = key.match(/([A-Za-z0-9-_]+)\./i)
    if (matched && matched.length > 1) {
      const locale = matched[1]
      messages[locale] = locales(key)
    }
  })
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

export default new VueI18n({
  locale: pluginOptions.i18n.locale || 'de',
  fallbackLocale: pluginOptions.i18n.fallbackLocale || 'en',
  messages: loadLocaleMessages(),
  numberFormats: loadLocaleNumbers()
})
