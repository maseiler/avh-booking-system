<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Modify user</h1>
            <hr>
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

              <div v-if="user.Status == 'Steganleger'" class="field">
                <label class="label">Boat Name</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    placeholder="Boat Name"
                    v-model.lazy="user.BoatName"
                  >
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="anchor"/>
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
            </div>

            <div class="field">
              <label class="label">Balance</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  placeholder="Balance"
                  v-model.number="user.Balance"
                >
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="euro-sign"/>
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">Max Debt</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  placeholder="Max Debt"
                  v-model.number="user.MaxDebt"
                >
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="money-bill"/>
                </span>
              </div>
            </div>
          </div>

          <div class="modal-footer">
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
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="modifyUser">Modify</button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="closeForm">Cancel</button>
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
  data() {
    return {
      validationError: ""
    };
  },
  methods: {
    modifyUser() {
      if (Object.keys(this.user).length === 0) {
        this.validationError = "Please select a user first.";
      } else {
        this.$http
          .post("modifyUser", this.user)
          .then(function(response) {
            var message = "".concat(
              "Modified user: ",
              this.displayUserName(this.user)
            );
            this.closeForm()
            this.$store.commit("getUsers");
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch(function(response) {
            this.validationError = response.body;
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
