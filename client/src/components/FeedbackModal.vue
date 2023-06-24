<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">{{ $t("feedback.title") }}</h1>
            <hr />
          </div>

          <div class="modal-body">
            <p class="subtitle is-5">
              {{ $t("feedback.subtitle") }}
              <span class="icon">
                <font-awesome-icon icon="heart" :style="{ color: '#a50303' }" />
              </span>
            </p>
            <div class="field">
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('feedback.PHName')"
                  v-model="name"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="user" />
                </span>
              </div>
            </div>
            <textarea
              class="textarea is-info is-focused"
              v-bind:placeholder="$t('feedback.PHComment')"
              rows="7"
              v-model="text"
            ></textarea>
          </div>

          <div class="modal-footer">
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="submitFeedback">
                    {{ $t("generic.submit") }}
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button" @click="cancel">{{ $t("generic.cancel") }}</button>
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
          this.$responseEventBus.$emit("successMessage", this.$t('messages.success.newFeedbackAdded'));
        })
        .catch(() => {
          this.$responseEventBus.$emit(
            "failureMessage",
            this.$t('submitFeedback')
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
