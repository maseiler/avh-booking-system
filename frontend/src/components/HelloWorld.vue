<template>
  <h1>{{ msg }}</h1>
  <h2><icon :icon="['fas', 'user']"/> Icon</h2>
  <div class="card">
    <button type="button" @click="count++">count is {{ count }}</button>
    <button @click="huetteConfetti">Party</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR {{ $t('home.title') }}
    </p>
  </div>
  <div class="box">
    Test
    <button class="button" @click="sendMessage">Send Message</button>
  </div>
  <button class="button" @click="listCategorys">List Categorys</button>
  <p v-for="category in categorys" :key="category.title">{{category.title}} <icon :icon="category.icon" /></p>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
      >create-vue</a
    >, the official Vue + Vite starter
  </p>
  <p>
    Learn more about IDE Support for Vue in the
    <a
      href="https://vuejs.org/guide/scaling-up/tooling.html#ide-support"
      target="_blank"
      >Vue Docs Scaling up Guide</a
    >.
  </p>
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>

  <button class="button" v-for="account in account$.accounts"> {{ account.firstName }} </button>

</template>


<script lang="ts">
import {
  huetteConfetti
} from '../composables/confetti.ts';
import { getTestCategorys, type Category } from '../composables/category.ts';
import { useAccountStore } from '../store/AccountStore.ts';
import { useSocketStore } from '../store/socketStore.ts';

export default {
  data() {
    return{
      count: 0,
      categorys: [] as Category[],
      account$: useAccountStore(),
      socketStore: useSocketStore()
    }
  },
  props: {
    msg: String
  },
  methods:{
    listCategorys(){
      this.categorys = getTestCategorys();
      console.log(getTestCategorys());
    },
    huetteConfetti(){
      huetteConfetti();
    },
    sendMessage(){
      this.socketStore.sendMessage("ABC")
    }
  },
  mounted() {
    console.log(import.meta.env.DEV);
  },
}

</script>