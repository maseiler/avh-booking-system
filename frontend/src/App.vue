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
  import { useSocketStore } from './store/socketStore';
  import { useThemeStore } from './store/themeStore';

  export default {
    components: {
      MainNavigation,
      DevModeBar
    },
    data() {
      return {
        dev: false,
        socket$: useSocketStore(),
        theme$: useThemeStore(),
      }
    },
    mounted() {
        this.theme$.init();
        this.dev = import.meta.env.DEV;
        this.socket$.getAllFromDb();
    }
  }
</script>
