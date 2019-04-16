<template>
  <article class="message is-link" v-if="Object.keys(user).length !== 0">
    <div class="message-header">
      <p>{{displayUserName(user)}}</p>
    </div>
    <div class="message-body">
      <!-- TODO: favicon -->
      <div class="columns">
        <div class="column has-text-left">
          <h6 class="subtitle is-6">
            <div v-if="user.BierName !== ''">
              <font-awesome-icon icon="beer"/>
              &nbsp;{{user.BierName}}
            </div>
            <div v-if="user.FirstName !== ''">
              <font-awesome-icon icon="user"/>
              &nbsp;{{user.FirstName}}
            </div>
            <div v-if="user.LastName !== ''">
              <font-awesome-icon icon="user-circle"/>
              &nbsp;{{user.LastName}}
            </div>
            <div v-if="user.Status !== ''">
              <font-awesome-icon icon="info-circle"/>
              &nbsp;{{user.Status}}
            </div>
            <div v-if="user.Email !== ''">
              <font-awesome-icon icon="envelope"/>
              &nbsp;{{user.Email}}
            </div>
            <div v-if="user.Phone !== ''">
              <font-awesome-icon icon="phone"/>
              &nbsp;{{user.Phone}}
            </div>
          </h6>
        </div>
        <div class="column is-one-quarter">
          <h5 class="subtitle is-5" v-if="user.Balance !== ''">Balance:</h5>
          <h6 class="title is-5">{{user.Balance}} €</h6>
        </div>
        <div class="column is-one-quarter">
          <button class="button is-link" @click="pay">Pay</button>
        </div>
      </div>
    </div>
  </article>
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
    }
  }
};
</script>