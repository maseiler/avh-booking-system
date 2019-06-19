<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1>Login</h1>
            <input class="input" type="password" v-model="password" placeholder="Password">
            <button class="button is-link" v-on:click="login">Login</button>
            <button class="button" v-on:click="cancel">Cancel</button>
            <!-- TODO: response if wrong -->
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      password: ""
    };
  },
  methods: {
    login() {
      this.$http
        .post("/login", this.password)
        .then(function(response) {
          console.log(response.data);
          this.requireAuth = true;
          this.$router.push("admin");
          this.$emit("close");

          // TODO: prevent access without login: https://stackoverflow.com/questions/52560021/how-to-restrict-page-access-to-unlogged-users-with-vuejs
          // - check going back/forward
          // TODO: reset password by typing in old password
        })
        .catch(function(response) {
          // TODO
        });
    },
    cancel() {
      this.$emit("close");
    }
  }
};
</script>