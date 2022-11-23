<template>
  <div>
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

<style>
#feedbackModal {
  display:block;
  margin:0 auto;
  width:fit-content;
  position: fixed;
  bottom: 0px;
  left:50%;
  transform: translateX(-50%);
}
</style>
