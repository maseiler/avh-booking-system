<template>
  <div id="app">
    <NavBar></NavBar>
    <router-view></router-view>
    <ResponseNotification/>
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";
import ResponseNotification from "./components/ResponseNotification.vue";

export default {
  name: "App",
  components: {
    NavBar,
    ResponseNotification
  }
};
</script>


<style>
#app {
  text-align: center;
  color: #000000;
}
</style>
