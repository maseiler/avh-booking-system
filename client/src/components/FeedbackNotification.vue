<template>
  <div id="feedbackNotification">
    <button class="button is-primary" :style="{ color: 'black' }" @click="showFeedbackForm = true">
      <span class="icon is-small">
        <font-awesome-icon icon="bullhorn" />
      </span>
      <p>Feedback</p>
    </button>
    <FeedbackForm v-if="showFeedbackForm" @close="showFeedbackForm = false" />
  </div>
</template>

<script>
import FeedbackForm from "./FeedbackForm.vue";
export default {
  components: {
    FeedbackForm
  },
  data() {
    return {
      showFeedbackForm: false
    };
  }
};
</script>

<style>
#feedbackNotification {
  position: fixed;
  bottom: 0px;
  width: 100%;
}
</style>
