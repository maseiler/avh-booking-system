<template>
  <div id="app">
    <section>
      <NavBar></NavBar>
    </section>
    <router-view />
  </div>
</template>

<script setup lang="ts">
import NavBar from "./components/NavBar.vue"
</script>

<style lang="scss">
#app {
  text-align: center;
}
</style>
