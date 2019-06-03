<template>
  <div id="app">
    <NavBar></NavBar>
    <router-view></router-view>
    <ResponseNotification/>
    <FeedbackNotification/>
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";
import ResponseNotification from "./components/ResponseNotification.vue";
import FeedbackNotification from "./components/FeedbackNotification.vue";

export default {
  name: "App",
  components: {
    NavBar,
    ResponseNotification,
    FeedbackNotification
  }
};
</script>


<style>
#app {
  text-align: center;
  color: #000000;
}
</style>
