<template>
  <div id="feedbackNotification">
    <button class="button is-primary" :style="{ color: 'black' }" @click="showFeedbackForm = true">
      <span class="icon is-small is-left">
        <font-awesome-icon icon="bullhorn"/>
      </span>
      &nbsp;&nbsp;Feedback
    </button>
    <FeedbackForm v-if="showFeedbackForm" @close="showFeedbackForm = false"/>
  </div>
</template>

<script>
import FeedbackForm from "./FeedbackForm.vue";
export default {
  components: {
    FeedbackForm
  },
  data: function() {
    return {
      showFeedbackForm: false
    };
  }
};
</script>

<style>
#feedbackNotification {
  position: absolute;
  bottom: 0;
  left: 50%;
  right: 50%;
}
</style>
