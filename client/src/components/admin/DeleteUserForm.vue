<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Delete user</h1>
            <hr>
          </div>

          <div class="modal-body">
            <article v-if="Object.keys(user).length === 0" class="message is-danger">
              <div class="message-header">
                <div class="field is-grouped">
                  <p class="icon is-small is-left">
                    <font-awesome-icon icon="exclamation" size="lg"/>
                  </p>
                  <p>Select user first.</p>
                </div>
              </div>
            </article>
            <div v-else>
              <p>
                Do you really want to delete
                <b>{{displayUserName(user)}}</b>?
              </p>
            </div>
          </div>

          <div class="modal-footer">
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <div class="control" v-if="Object.keys(user).length !== 0">
                    <button class="button is-link" @click="deleteUser">Delete</button>
                  </div>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button" @click="closeForm">Cancel</button>
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
  props: {
    user: {}
  },
  methods: {
    deleteUser() {
      if (Object.keys(this.user).length === 0) {
        this.error = "Please select an user first.";
      } else {
        this.$http
          .post("deleteUser", this.user)
          .then(function(response) {
            var message = "".concat(
              "Deleted user: ",
              this.displayUserName(this.user)
            );
            this.closeForm()
            this.$store.commit("getUsers");
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch(function(response) {
            this.$responseEventBus.$emit("failureMessage", "Couldn't delete user.");
          });
      }
    },
    closeForm() {
      this.$emit("close");
    }
  }
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
