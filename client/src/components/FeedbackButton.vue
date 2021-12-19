<template>
  <div id="feedbackModal">
    <button
      class="button is-primary"
      :style="{ color: 'black' }"
      @click="showFeedbackModal = true"
    >
      <span class="icon is-small">
        <font-awesome-icon icon="bullhorn" />
      </span>
      <p>Feedback</p>
    </button>
    <FeedbackModal v-if="showFeedbackModal" @close="showFeedbackModal = false" />
  </div>
</template>

<script>
import FeedbackModal from "./FeedbackModal.vue";
export default {
  components: {
    FeedbackModal,
  },
  data() {
    return {
      showFeedbackModal: false,
    };
  },
};
</script>

<style>
#feedbackModal {
  position: fixed;
  bottom: 0px;
  width: 100%;
}
</style>
