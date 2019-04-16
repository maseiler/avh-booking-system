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
    content: "",
    color: ""
  },
  methods: {
    close: function() {
      this.content = "";
      this.color = "";
      this.$emit("close");
    },
    successMessage: function(message) {
      this.content = message;
      this.color = "is-success";
      var timerId = setTimeout(this.close, 3000);
    },
    failureMessage: function(message) {
      this.content = message;
      this.color = "is-danger";
      var timerId = setTimeout(this.close, 3000);
    }
  },
  created: function() {
    this.$responseEventBus.$on("successMessage", this.successMessage);
    this.$responseEventBus.$on("failureMessage", this.failureMessage);
  }
};

var ResponseEventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $responseEventBus: {
    get: function() {
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
