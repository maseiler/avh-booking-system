<template>
  <div id="response" v-if="content !== undefined && content !== ''">
    <transition>
      <div :class="['notification ' + color]">
        <button class="delete" @click="close"></button>
        {{content}}
      </div>
    </transition>
  </div>
</template>

<script>
import Vue from "vue";

export default {
  props: {
    content: String,
    color: String
  },
  methods: {
    close() {
      this.content = "";
      this.color = "";
      this.$emit("close");
    },
    successMessage(message) {
      this.content = message;
      this.color = "is-success";
      setTimeout(this.close, 3000);
    },
    failureMessage(message) {
      this.content = message;
      this.color = "is-danger";
      setTimeout(this.close, 3000);
    }
  },
  created() {
    this.$responseEventBus.$on("successMessage", this.successMessage);
    this.$responseEventBus.$on("failureMessage", this.failureMessage);
  }
};

var ResponseEventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $responseEventBus: {
    get() {
      return ResponseEventBus;
    }
  }
});
</script>

<style>
#response {
  position: absolute;
  bottom: 0;
  right: 0;
}
</style>
