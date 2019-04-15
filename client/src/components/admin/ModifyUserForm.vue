<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h1>Modify user</h1>
            </slot>
          </div>

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
            <div class="modal-body">
              <slot name="body">
                <div class="field">
                  <label class="label">Biername</label>
                  <div class="control has-icons-left">
                    <input
                      class="input"
                      type="text"
                      placeholder="Biername"
                      v-model.lazy="user.BierName"
                    >
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="beer"/>
                    </span>
                  </div>
                </div>

                <div class="field">
                  <label class="label">First name</label>
                  <div class="control has-icons-left">
                    <input
                      class="input"
                      type="text"
                      placeholder="First name"
                      v-model.lazy="user.FirstName"
                    >
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="user"/>
                    </span>
                  </div>
                </div>

                <div class="field">
                  <label class="label">Last name</label>
                  <div class="control has-icons-left">
                    <input
                      class="input"
                      type="text"
                      placeholder="Last name"
                      v-model.lazy="user.LastName"
                    >
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="user"/>
                    </span>
                  </div>
                </div>

                <div class="field">
                  <label class="label">Email</label>
                  <div class="control has-icons-left">
                    <input
                      class="input"
                      type="email"
                      placeholder="Email Address"
                      v-model.lazy="user.Email"
                    >
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="envelope"/>
                    </span>
                  </div>
                </div>

                <div class="field">
                  <label class="label">Phone Number</label>
                  <div class="control has-icons-left">
                    <input
                      class="input"
                      type="text"
                      placeholder="Phone Number"
                      v-model.lazy="user.Phone"
                    >
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="phone"/>
                    </span>
                  </div>
                </div>

                <div class="field">
                  <label class="label">Status</label>
                  <div class="control">
                    <div class="select">
                      <select v-model="user.Status">
                        <option disabled value>Status</option>
                        <option>Aktiv B</option>
                        <option>Aktiv KA</option>
                        <option>AH</option>
                        <option>Steganleger</option>
                        <option>Gast</option>
                      </select>
                    </div>
                  </div>
                </div>
              </slot>
            </div>

            <div class="field">
              <label class="label">Balance</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Balance" v-model.lazy="user.Balance">
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="euro-sign"/>
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">Max Debt</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Max Debt" v-model.lazy="user.MaxDebt">
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="money-bill"/>
                </span>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <article v-if="validationError !==''" class="message is-danger">
                <div class="message-header">
                  <div class="field is-grouped">
                    <p class="icon is-small is-left">
                      <font-awesome-icon icon="exclamation" size="lg"/>
                    </p>
                    <p>{{validationError}}</p>
                  </div>
                </div>
              </article>
              <div class="field is-grouped">
                <div class="control">
                  <button class="button is-link" @click="modifyUser">Modify</button>
                </div>
                <div class="control">
                  <button class="button is-text" @click="cancelSubmission">Cancel</button>
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
  data() {
    return {
      validationError: ""
    };
  },
  methods: {
    modifyUser() {
      if (Object.keys(this.user).length === 0) {
        this.validationError = "Please select a user first.";
        console.log("error");
      } else {
        this.$http
          .post("/modifyUser", this.user)
          .then(function(response) {
            this.$router.go();
            this.$responseEventBus.$emit(
              "modifyUserSuccess",
              this.user,
              "is-success"
            );
          })
          .catch(function(response) {
            this.validationError = response.body;
            console.log("Error: Couldn't modify user.");
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