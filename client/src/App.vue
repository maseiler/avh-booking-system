<template>
  <div id="app">
    <NavBar></NavBar>
    <router-view></router-view>
    <ResponseNotification/>
    <FeedbackNotification/>
    <LastBookingNotification/>
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";
import ResponseNotification from "./components/ResponseNotification.vue";
import FeedbackNotification from "./components/FeedbackNotification.vue";
import LastBookingNotification from "./components/LastBookingNotification.vue";

export default {
  name: "App",
  components: {
    NavBar,
    ResponseNotification,
    FeedbackNotification,
    LastBookingNotification
  }
};
</script>


<style>
#app {
  text-align: center;
  color: #000000;
}
</style>
