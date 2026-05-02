<template>
  <!-- ToDo: Maybe a whole 'Welcome Experience' should be introduced.
   It could just be limited to setting the ClientId, when evrything else is done.
   But if this is the first installation, maybe it would be nice to be able to be redirected to the
   Admin Settings directly. -->
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
    // ToDo: move this to a computed value OR
    // Move this functionality 'needsSetup()'/'setupComplete()' to the SettingStore and let it handle it.
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
    // ToDo: either remove this function or use it properly when the ClientSetup actually returns an event.
    // currently the Component doenst emit this Event and the functino will not be called at all.

    //ToDo: Remove this function and move it into SettingStore so the Components and Views try to do as much as possible on their own. 
    onSetupComplete() {
      // At the Top is is asked if 'needsSetup' is true. Why not simply setting this value to false?
      // A reload of the page should never be the solution to fix broken Data/Behaviour.
      location.reload();
    },
  },
};
</script>
