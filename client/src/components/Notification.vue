<template>
  <div id="response" v-if="content !== undefined && content !== ''">
    <transition name="fade">
      <div class="modal is-active" v-on:click="closeModal()">
        <div
          class="modal-background has-background-black"
          style="opacity: 0.5"
        ></div>
        <div class="modal-content">
          <div class="card has-text-black-bis" v-bind:class="backgroundColor">
            <header class="card-header">
              <div class="card-header-title">
                <p v-if="messageType == 'success'" class="title is-4">
                  {{ $t("notification.success") }}
                </p>
                <p v-if="messageType == 'failure'" class="title is-4">
                  {{ $t("notification.error") }}
                </p>
                <p v-if="messageType == 'paymentProcessing'" class="title is-4">
                  {{ $t("notification.processing") }}
                </p>
              </div>
              <button
                v-if="this.messageType != 'paymentProcessing'"
                class="button is-rounded"
                v-bind:class="textColor"
                v-on:click="close()"
              >
                <font-awesome-icon icon="times" />
              </button>
              <button
              v-if="this.messageType == 'paymentProcessing'"
                class="button is-rounded"
                v-bind:class="textColor"
                v-on:click="cancelAction()"
              >
                <font-awesome-icon icon="times" />
              </button>
            </header>
            <div class="card-content">
              <div class="content">
                {{ content }}
                <br /><br />
                <font-awesome-icon
                  v-if="messageType == 'success'"
                  icon="check-circle"
                  size="3x"
                />
                <font-awesome-icon
                  v-if="messageType == 'failure'"
                  icon="sad-tear"
                  size="3x"
                />
                <font-awesome-icon
                  v-if="messageType == 'paymentProcessing'"
                  icon="spinner"
                  size="3x"
                  class="spinning"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.spinning{
  animation: spin 1000ms ease-in-out infinite forwards;
}
@keyframes spin {
  from{
    transform: rotate(0deg);
  }
  to{
    transform: rotate(359deg);
  }
}

</style>

<script>
import Vue from "vue";

export default {
  props: {
    content: String,
    textColor: String,
    backgroundColor: String,
    messageType: String,
  },
  methods: {
    close() {
      this.content = "";
      this.textColor = "";
      this.backgroundColor = "";
      this.$emit("close");
    },
    cancelAction() {
      this.content = "";
      this.textColor = "";
      this.backgroundColor = "";
      this.$emit("close");
      if(this.messageType == "paymentProcessing"){
        this.$http.post("cancelReaderAction", localStorage.getItem("StripeCardReader"))
      }
    },
    closeModal() {
      if (this.messageType == "success") {
        this.close();
      }
      // in case of error, prevent closing when clicking outside of card
    },
    successMessage(message) {
      this.messageType = "success";
      this.content = message;
      this.textColor = "is-success";
      this.backgroundColor = "has-background-success";
      setTimeout(this.close, 3000);
    },
    failureMessage(message) {
      this.messageType = "failure";
      this.content = message;
      this.textColor = "is-danger";
      this.backgroundColor = "has-background-danger";
    },
    processingMessage(message) {
      this.messageType = "paymentProcessing";
      this.content = message;
      this.textColor = "is-success";
      this.backgroundColor = "has-background-info";
    }
  },
  created() {
    this.$responseEventBus.$on("successMessage", this.successMessage);
    this.$responseEventBus.$on("failureMessage", this.failureMessage);
    this.$responseEventBus.$on("processingMessage", this.processingMessage);
    this.$responseEventBus.$on("close", this.close);
  },
};

var ResponseEventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $responseEventBus: {
    get() {
      return ResponseEventBus;
    },
  },
});
</script>
