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
    checkoutSuccess: function(user, color) {
      this.content = "".concat("Checkout from ", this.displayUserName(user));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    checkoutFailure: function(message, color) {
      console.log(message);
      this.content = message;
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    addUserSuccess: function(user, color) {
      this.content = "".concat("Added new user: ", this.displayUserName(user));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    addItemSuccess: function(item, color) {
      this.content = "".concat("Added new item: ", this.displayItem(item));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    modifyItemSuccess: function(item, color) {
      this.content = "".concat("Modified item: ", this.displayItem(item));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    modifyUserSuccess: function(user, color) {
      this.content = "".concat("Modified user: ", this.displayUserName(user));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    deleteItemSuccess: function(item, color) {
      this.content = "".concat("Deleted item: ", this.displayItem(item));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
    deleteUserSuccess: function(user, color) {
      this.content = "".concat("Deleted user: ", this.displayUserName(user));
      this.color = color;
      var timerId = setTimeout(this.close, 3000);
    },
  },
  created: function() {
    this.$responseEventBus.$on("checkoutSuccess", this.checkoutSuccess);
    this.$responseEventBus.$on("checkoutFailure", this.checkoutFailure);
    this.$responseEventBus.$on("addUserSuccess", this.addUserSuccess);
    this.$responseEventBus.$on("addItemSuccess", this.addItemSuccess);
    this.$responseEventBus.$on("modifyItemSuccess", this.modifyItemSuccess);
    this.$responseEventBus.$on("modifyUserSuccess", this.modifyUserSuccess);
    this.$responseEventBus.$on("deleteItemSuccess", this.deleteItemSuccess);
    this.$responseEventBus.$on("deleteUserSuccess", this.deleteUserSuccess);
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
