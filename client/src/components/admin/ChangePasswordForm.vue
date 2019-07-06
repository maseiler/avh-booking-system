<template>
  <div>
    <div class="field">
      <div class="control">
        <input class="input" type="password" placeholder="Old password" v-model="oldPassword">
      </div>
    </div>
    <div class="field">
      <div class="control">
        <input class="input" type="password" placeholder="New password" v-model="newPassword1">
      </div>
    </div>
    <div class="field">
      <div class="control">
        <input
          class="input"
          type="password"
          placeholder="Confirm new password"
          v-model="newPassword2"
        >
      </div>
    </div>
    <article v-if="validationError !== ''" class="message is-danger">
      <div class="message-header">
        <div class="field is-grouped">
          <p class="icon is-small is-left">
            <font-awesome-icon icon="exclamation" size="lg"/>
          </p>
          <p>{{validationError}}</p>
        </div>
      </div>
    </article>
    <button class="button is-link" @click="changePassword">Change</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      oldPassword: "",
      newPassword1: "",
      newPassword2: "",
      validationError: ""
    };
  },
  methods: {
    changePassword() {
      if (this.newPassword1 === this.newPassword2) {
        this.$http
          .post("changeAdminPassword", [this.oldPassword, this.newPassword1])
          .then(function(response) {
            this.$responseEventBus.$emit("successMessage", "Changed password.");
          })
          .catch(function(response) {
            this.$responseEventBus.$emit("failureMessage", response.body);
          });
      } else {
        this.validationError = "New passwords do not match!";
      }
    }
  }
};
</script>

