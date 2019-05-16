<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Add new user</h1>
            <hr>
          </div>

          <div class="modal-body">
            <div class="field">
              <label class="label">Biername</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  placeholder="Biername"
                  v-model.lazy="newUser.bierName"
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
                  v-model.lazy="newUser.firstName"
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
                  v-model.lazy="newUser.lastName"
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
                  v-model.lazy="newUser.email"
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
                  v-model.lazy="newUser.phone"
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
                  <select v-model="newUser.status">
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
                  <button class="button is-link" @click="submitUser">Submit</button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="cancelSubmission">Cancel</button>
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
  data: function() {
    return {
      newUser: {
        bierName: "",
        firstName: "",
        lastName: "",
        email: "",
        phone: "",
        status: "",
        balance: "0",
        maxDebt: "0"
      },
      validationError: ""
    };
  },
  methods: {
    submitUser() {
      this.$http
        .post("/addUser", this.newUser)
        .then(function(response) {
          var message = "".concat(
            "Added new user: ",
            this.displayUserName(this.newUser)
          );
          this.resetAndCloseForm();
          this.$router.go();
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch(function(response) {
          this.validationError = response.data;
        });
    },
    cancelSubmission() {
      this.resetAndCloseForm();
    },
    resetAndCloseForm() {
      (this.bierName = ""),
        (this.firstName = ""),
        (this.lastName = ""),
        (this.email = ""),
        (this.phone = ""),
        (this.status = ""),
        (this.validationError = "");
      this.$emit("close");
    }
  }
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>