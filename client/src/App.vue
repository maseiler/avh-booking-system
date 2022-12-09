<template>
  <div id="app">
    <NavBar></NavBar>
    <router-view></router-view>
    <Notification />
    <FeedbackButton />
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";
import Notification from "./components/Notification.vue";
import FeedbackButton from "./components/FeedbackButton.vue";

export default {
  name: "App",
  components: {
    NavBar,
    Notification,
    FeedbackButton
  }
};
</script>
