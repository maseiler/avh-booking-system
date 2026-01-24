<template>
<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link class="navbar-item" to="/">
      DEVELOPEMENT
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

      <a class="navbar-item" @click="generateTestData()">Generate Local Test Data</a>
      <a class="navbar-item" @click="socket$.queryAccounts()">(Re)Request Accounts from DB/WS</a>
      
    </div>

    <div class="navbar-end">
      <span>This is not visible in production</span>
    </div>
  </div>
</nav>
</template>

<style scoped>
  .navbar{
    border: 3px solid red;
  }
</style>

<script lang="ts">
import { useAccountStore } from '../store/AccountStore';
import { useCategoryStore } from '../store/CategoryStore';
import { useProductStore } from '../store/ProductStore';
import { useSocketStore } from '../store/socketStore';

  export default {
    data() {
      return {
        account$: useAccountStore(),
        category$: useCategoryStore(),
        product$: useProductStore(),
        socket$: useSocketStore(),
        burgerActive: false as Boolean
      }
    },
    methods: {
      toggleBurger(){
        this.burgerActive = !this.burgerActive;
      },
      generateTestData(){
        this.account$.generateTestData();
        this.category$.generateTestData();
        this.product$.generateTestData();
      }
    }
  }
</script>