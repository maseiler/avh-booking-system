<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Feedback</h1>
            <hr />
          </div>

          <div class="modal-body">
            <p class="subtitle is-5">
              Do you have comments, feature requests or found a bug? Please let
              us know!
              <span class="icon">
                <font-awesome-icon icon="heart" :style="{ color: '#a50303' }" />
              </span>
            </p>
            <div class="field">
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  placeholder="Name (optional)"
                  v-model="name"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="user" />
                </span>
              </div>
            </div>
            <textarea
              class="textarea is-info is-focused"
              placeholder="Comment here..."
              rows="7"
              v-model="text"
            ></textarea>
          </div>

          <div class="modal-footer">
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="submitFeedback">
                    Submit
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button" @click="cancel">Cancel</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data() {
    return {
      name: "",
      text: "",
    };
  },
  methods: {
    submitFeedback() {
      var feedback = { Name: this.name, Content: this.text };
      this.$http
        .post("addFeedback", feedback)
        .then(() => {
          this.$store.commit("getFeedbackList");
          this.$responseEventBus.$emit("successMessage", "Thank you! <3");
        })
        .catch(() => {
          this.$responseEventBus.$emit(
            "failureMessage",
            "Couldn't submit feedback."
          );
        });
      this.resetData();
      this.$emit("close");
    },
    cancel() {
      this.resetData();
      this.$emit("close");
    },
    resetData() {
      this.text = "";
      this.name = "";
    },
  },
};
</script>

<style lang="scss">
@import "../assets/modal.css";
</style>
