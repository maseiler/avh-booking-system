<template>
  <DevModeBar v-if="dev"/>
  <MainNavigation />
  <!-- <div v-for="msg in store.notifications"> {{ msg }}</div>
  <button class="button" @click="sendTestMessage">Send Test Message</button> -->
  <RouterView />
</template>

<script lang="ts">
    import MainNavigation from './components/MainNavigation.vue';
  import DevModeBar from './components/DevModeBar.vue';
  
  export default {
    components: {
      MainNavigation,
      DevModeBar
    },
    data() {
      return {
        dev: false
      }
    },
    mounted() {
        this.dev = import.meta.env.DEV;
    }
  }
</script>
