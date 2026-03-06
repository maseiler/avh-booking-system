<template>
  <ClientSetup v-if="needsSetup" @setup-complete="onSetupComplete" />
  <template v-else>
    <DevModeBar v-if="dev" />
    <MainNavigation />
    <RouterView />
  </template>
</template>

<script lang="ts">
import MainNavigation from './components/MainNavigation.vue';
import DevModeBar from './components/DevModeBar.vue';
import ClientSetup from './views/ClientSetup.vue';
import { useSocketStore } from './store/socketStore';
import { useThemeStore } from './store/themeStore';

export default {
  components: {
    MainNavigation,
    DevModeBar,
    ClientSetup
  },
  data() {
    const clientId = localStorage.getItem('avhbs_client_id');
    return {
      dev: false,
      needsSetup: !clientId,
      socket$: clientId ? useSocketStore() : null,
      theme$: useThemeStore(),
    };
  },
  mounted() {
    this.theme$.init();
    this.dev = import.meta.env.DEV;
    if (this.socket$) {
      this.socket$.getAllFromDb();
    }
  },
  methods: {
    onSetupComplete() {
      location.reload();
    },
  },
};
</script>
