<template>
  <!-- ToDo: Maybe a whole 'Welcome Experience' should be introduced.
   It could just be limited to setting the ClientId, when evrything else is done.
   But if this is the first installation, maybe it would be nice to be able to be redirected to the
   Admin Settings directly. -->
  <ClientSetup v-if="setting$.needsSetup" />
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
import { useSettingStore } from './store/SettingStore';
import { useThemeStore } from './store/themeStore';


export default {
  components: {
    MainNavigation,
    DevModeBar,
    ClientSetup,
  },
  data() {
    return {
      dev: false,
      setting$: useSettingStore(),
      socket$: useSocketStore(),
      theme$: useThemeStore(),
    };
  },
  mounted() {
    this.dev = import.meta.env.DEV;
    if (this.socket$ && !this.setting$.needsSetup) {
      this.socket$.getAllFromDb();
    }
    this.theme$.init()
  },
  methods: {
  },
};
</script>
