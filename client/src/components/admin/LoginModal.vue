<template>
  <div>
    <transition name="modal">
      <div class="modal-mask">
        <div class="modal-wrapper">
          <div class="modal-container">
            <div class="modal-header">
              <h1 class="title is-4">Login to Admin</h1>
            </div>
            <div class="modal-body">
              <div class="field">
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="password"
                    placeholder="Password"
                    v-model="password"
                    @keyup.enter="login"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="lock" />
                  </span>
                </div>
              </div>
              <article v-if="validationError !== ''" class="message is-danger">
                <div class="message-header">
                  <div class="field is-grouped">
                    <p class="icon is-small is-left">
                      <font-awesome-icon icon="exclamation" size="lg" />
                    </p>
                    <p>{{ validationError }}</p>
                  </div>
                </div>
              </article>
            </div>

            <div class="modal-footer">
              <div class="level">
                <div class="level-left">
                  <div class="level-item">
                    <button class="button is-link" @click="login">
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
  </div>
</template>

<script>
export default {
  data() {
    return {
      password: "",
      validationError: "",
    };
  },
  methods: {
    login() {
      this.$http
        .post("login", this.password)
        .then(() => {
          this.$router.push("/admin");
        })
        .catch(() => {
          this.validationError = "Error. Wrong password?";
        });
    },
    cancel() {
      this.$router.push("/");
    },
  },
};
</script>

<style lang="scss">
@import "../../assets/modal.css";
</style>
