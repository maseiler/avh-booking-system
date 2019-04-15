<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h1>Delete user</h1>
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
                <p>Do you really want to delete</p>
                <p>{{user.BierName}} {{user.FirstName }} {{user.LastName}}?</p>
              </div>
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <div class="field is-grouped">
                <div class="control" v-if="Object.keys(user).length !== 0">
                  <button class="button is-link" @click="deleteUser">Delete</button>
                </div>
                <div class="control">
                  <button class="button" @click="cancelSubmission">Cancel</button>
                </div>
              </div>
            </slot>
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
          .post("/deleteUser", this.user)
          .then(function(response) {
            this.$router.go();
            this.$responseEventBus.$emit(
              "deleteUserSuccess",
              this.user,
              "is-success"
            );
          })
          .catch(function(response) {
            console.log("Error: Couldn't delete user.");
            //TODO
          });
      }
    },
    cancelSubmission() {
      console.log("Canceled submission.");
      this.$emit("close");
    }
  }
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>