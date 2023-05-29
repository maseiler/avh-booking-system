<template>
  <div>
    <div id="feedbackModal">
      <button
        class="button is-primary"
        @click="showFeedbackModal = true"
      >
        <span class="icon is-small">
          <font-awesome-icon icon="bullhorn" />
        </span>
        <p>{{ $t("feedback.button") }}</p>
      </button>
    </div>
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
