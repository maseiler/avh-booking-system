<template>
  <article class="message is-link" v-if="Object.keys(user).length !== 0">
    <div class="message-header">
      <p>{{displayUserName(user)}}</p>
    </div>
    <div class="message-body">
      <div class="columns">
        <div class="column has-text-left">
          <h6 class="subtitle is-6">
            <div v-if="user.BierName !== ''">
              <font-awesome-icon icon="beer" />
              &nbsp;{{user.BierName}}
            </div>
            <div v-if="user.FirstName !== ''">
              <font-awesome-icon icon="user" />
              &nbsp;{{user.FirstName}}
            </div>
            <div v-if="user.LastName !== ''">
              <font-awesome-icon icon="user-circle" />
              &nbsp;{{user.LastName}}
            </div>
            <div v-if="user.BoatName">
              <font-awesome-icon icon="anchor" />
              &nbsp;{{user.BoatName}}
            </div>
            <div v-if="user.Status !== ''">
              <font-awesome-icon icon="info-circle" />
              &nbsp;{{user.Status}}
            </div>
          </h6>
        </div>
        <div class="column is-one-quarter">
          <h5 class="subtitle is-5" v-if="user.Balance !== ''">Balance:</h5>
          <h6 class="title is-5">{{user.Balance}} €</h6>
        </div>
        <div class="column is-one-quarter">
          <button class="button" @click="showPaymentModal = true">Pay</button>
          <PaymentModal v-if="showPaymentModal" @close="showPaymentModal = false" />
        </div>
      </div>
    </div>
  </article>
</template>

<script>
import PaymentModal from "./PaymentModal.vue";

export default {
  components: {
    PaymentModal
  },
  computed: {
    user() {
      return this.$store.state.selectedUser;
    }
  },
  data() {
    return {
      showPaymentModal: false
    };
  }
};
</script>