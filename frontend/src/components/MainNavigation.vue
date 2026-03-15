<template>
<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link class="navbar-item" to="/">
      LOGO
    </router-link>
    
    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" @click="toggleBurger">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
    </a>
  </div>

  <div class="navbar-menu" :class="{'is-active': burgerActive}">
    <div class="navbar-start">
      <router-link class="navbar-item" to="/" @click="burgerActive = false">
        {{ $t('menu.home') }}
      </router-link>

      <router-link class="navbar-item" to="/booking" @click="burgerActive = false">
        {{ $t('menu.booking') }}
      </router-link>

      <router-link class="navbar-item" to="/payment" @click="burgerActive = false">
        {{ $t('menu.payment') }}
      </router-link>

      <router-link class="navbar-item" to="/orders" @click="burgerActive = false">
        {{ $t('menu.orders') }}
      </router-link>

      <a class="navbar-item" href="#">{{ $t('menu.stats') }}</a>

      <div class="navbar-item has-dropdown is-hoverable">
        <a class="navbar-link" href="#">{{ $t('menu.settings.menuTitle') }}</a>

        <div class="navbar-dropdown">
          <router-link to="/settings/products" class="navbar-item" @click="burgerActive = false">{{ $t('menu.settings.item') }}</router-link>
          <router-link to="/settings/accounts" class="navbar-item" @click="burgerActive = false">{{ $t('menu.settings.account') }}</router-link>
          <router-link to="/settings/payments" class="navbar-item" @click="burgerActive = false">{{ $t('menu.settings.payment') }}</router-link>
          <router-link to="/settings/application" class="navbar-item" @click="burgerActive = false">{{ $t('menu.settings.system') }}</router-link>
        </div>
      </div>
    </div>

    <div class="navbar-end">
      <ThemeSelector />
      <div class="navbar-item">
        <div class="buttons">
          <a class="button is-light">
            {{ $t('menu.login') }}
          </a>
        </div>
      </div>
    </div>
  </div>
</nav>
</template>

<script lang="ts">
  import ThemeSelector from './ThemeSelector.vue'

  export default {
    components: { ThemeSelector },
    data() {
      return {
        burgerActive: false as Boolean
      }
    },
    methods: {
      toggleBurger(){
        this.burgerActive = !this.burgerActive;
      }
    }
  }
</script>