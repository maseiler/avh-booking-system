<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Payment</h1>
            <hr>
          </div>

          <div class="modal-body">
            <p class="subtitle is-4">
              <b>{{displayUserName(user)}}</b> has to pay
              <b>{{user.Balance}}€</b>.
            </p>
            <br>
            <p class="subtitle is-4">Please place the cash in the money box!</p>
          </div>

          <div class="modal-footer">
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="pay">Pay</button>
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
</template>

<script>
export default {
  props: {
    user: {}
  },
  methods: {
    pay: function() {
      this.$http
        .post("/pay", this.user)
        .then(function(response) {
          var message = "".concat(
            this.displayUserName(this.user),
            " payed ",
            this.user.Balance
          );
          this.$router.go();
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch(function(response) {
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    cancel: function() {
      this.$emit("close");
    }
  }
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>