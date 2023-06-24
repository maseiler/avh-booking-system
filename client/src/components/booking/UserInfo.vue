<template>
  <article class="message is-link" v-if="Object.keys(user).length !== 0">
    <div class="message-header">
      <p>{{ displayUserName(user) }}</p>
    </div>
    <div class="message-body">
      <div class="columns">
        <div class="column has-text-left">
          <h6 class="subtitle is-6">
            <div v-if="user.BierName !== ''">
              <font-awesome-icon icon="beer" />
              &nbsp;{{ user.BierName }}
            </div>
            <div v-if="user.FirstName !== ''">
              <font-awesome-icon icon="user" />
              &nbsp;{{ user.FirstName }}
            </div>
            <div v-if="user.LastName !== ''">
              <font-awesome-icon icon="user-circle" />
              &nbsp;{{ user.LastName }}
            </div>
            <div v-if="user.BoatName">
              <font-awesome-icon icon="anchor" />
              &nbsp;{{ user.BoatName }}
            </div>
            <div v-if="user.Status !== ''">
              <font-awesome-icon icon="info-circle" />
              &nbsp;{{ user.Status }}
            </div>
          </h6>
        </div>
        <div class="column is-one-quarter">
          <h6 class="title is-5" style="color: #1a9000" v-if="user.Balance < 0">
            <!-- ToDo: Fetch Curreny Setting from Database not hard Coded -->
            {{$t('generic.credit')}}:<br />{{ $n(this.invertBalance(user.Balance), "currency", "de-DE") }}
          </h6>
          <h6
            class="title is-5"
            style="color: #c10404"
            v-else-if="user.Balance > 0"
          >
          <!-- ToDo: Fetch Curreny Setting from Database not hard Coded -->
          {{$t('generic.debt')}}:<br />{{ $n(user.Balance, "currency", "de-DE") }}
          </h6>
          <h6 class="title is-5" v-else>Debt:<br />{{ $n(user.Balance, "currency", "de-DE") }}</h6>
        </div>
        <div class="column is-one-quarter">
          <button class="button" @click="showModal()">{{$t('booking.payment.pay')}}</button>
          <PaymentModal
            v-if="showPaymentModal"
            @close="showPaymentModal = false"
          />
          <RechargeModal
            v-if="showRechargeModal"
            @close="showRechargeModal = false"
          />
        </div>
      </div>
    </div>
  </article>
</template>

<script>
import PaymentModal from "./PaymentModal.vue";
import RechargeModal from "./RechargeModal.vue";

export default {
  components: {
    PaymentModal,
    RechargeModal,
  },
  computed: {
    user() {
      return this.$store.state.selectedUser;
    },
  },
  data() {
    return {
      showPaymentModal: false,
      showRechargeModal: false,
    };
  },
  methods: {
    invertBalance(balance) {
      return -1 * balance;
    },
    showModal() {
      if (this.user.Balance > 0) {
        this.showPaymentModal = true;
      } else {
        this.showRechargeModal = true;
      }
    },
  },
};
</script>