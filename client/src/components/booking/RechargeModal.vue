<template>
  <transition name="paymentModal">
    <div class="paymentModal-mask big">
      <div class="paymentModal-wrapper big">
        <div class="paymentModal-container big">
          <div class="paymentModal-header big">
            <h1 class="title is-4">{{$t('booking.payment.recharge.title')}}</h1>
            <hr />
          </div>

          <div class="paymentModal-body big">
            <p class="subtitle is-4">
              <b>{{ displayUserName(user) }}</b> {{$t('booking.payment.recharge.sub1')}}<br />
              {{$t('booking.payment.recharge.sub2')}}
              <b>{{ this.invertBalance(user.Balance) }}{{$t('booking.payment.recharge.sub3')}}</b>.
            </p>
            <br />
            <p class="subtitle is-5">
              {{$t('booking.payment.recharge.addCreditNotice')}}
            </p>
            <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-3">
                <div class="field">
                  <p class="control has-icons-right">
                    <input
                      class="input no-controls"
                      type="number"
                      placeholder="00.00"
                      v-model.number="newCredit"
                      style="text-align: right; font-weight: bold"
                    />
                    <span class="icon is-right">
                      <font-awesome-icon
                        icon="euro-sign"
                        :style="{ color: 'black' }"
                      />
                    </span>
                  </p>
                </div>
              </div>
              <div class="column">
                <!-- <button class="button is-success is-fullwidth" @click="payCash">
                  {{$t('booking.payment.cash')}}
                </button> -->
              </div>
              <div class="column">
                <button class="button is-link is-fullwidth" @click="payEC">
                  {{$t('booking.payment.card')}}
                </button>
              </div>
              <div class="column">
                <button
                  class="button is-danger is-outlined is-fullwidth"
                  @click="cancel"
                >
                {{$t('generic.cancel')}}
                </button>
              </div>
              <div class="column is-1"></div>
            </div>

            <div class="level">
              <div class="level-item is-centered">
                <div v-if="showPasswordForm" class="field">
                  <div class="control has-icons-left">
                    <input
                      class="input"
                      type="password"
                      v-bind:placeholder="$t('user.password')"
                      v-model="password"
                      @keyup.enter="loginAndPay"
                    />
                    <span class="icon is-small is-left">
                      <font-awesome-icon icon="lock" />
                    </span>
                  </div>
                </div>
              </div>
            </div>
            <p
              v-if="userBookEntries.LastPayment !== '0001-01-01T00:00:00Z'"
              class="subtitle is-6"
            >
            {{$t('booking.payment.lastPayment')}}: {{ printDateTime(userBookEntries.LastPayment) }}
            </p>
          </div>

          <div class="paymentModal-footer big">
            <table class="table is-hoverable is-striped">
              <thead>
                <tr>
                  <th>{{$t('booking.payment.timestamp')}}</th>
                  <th>{{$t('item.item')}}</th>
                  <th>{{$t('cart.amount')}}</th>
                  <th>{{$t('cart.total')}}</th>
                  <th>{{$t('generic.comment')}}</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="entry in userBookEntries.Debts" :key="entry">
                  <td>{{ printDateTime(entry.TimeStamp) }}</td>
                  <td>{{ displayItem(getItemByID(items, entry.ItemID)) }}</td>
                  <td>{{ entry.Amount }}</td>
                  <td>{{ $n(entry.TotalPrice, "currency", "de-DE") }}</td>
                  <td>{{ entry.Comment }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  computed: {
    user() {
      return this.$store.state.selectedUser;
    },
    items() {
      return this.$store.state.items;
    },
  },
  data() {
    return {
      userBookEntries: {},
      showPasswordForm: false,
      password: "",
      newCredit: 0,
      paymentMethod: "",
    };
  },
  methods: {
    invertBalance(balance) {
      return -1 * balance;
    },
    payCash() {
      this.paymentMethod = "Cash";
      this.pay();
    },
    payEC() {
      this.paymentMethod = "EC";
      this.pay();
    },
    pay() {
      if (this.newCredit <= 0) {
        this.$emit("close");
        this.$responseEventBus.$emit(
          "failureMessage",
          this.$t('messages.failure.negativePaymentNotAllowed')
        );
        return;
      }
      if (this.showPasswordForm && this.password !== "") {
        this.loginAndPay();
      }
      if (this.weekdayIsMonday()) {
        this.submitPayment();
      } else {
        this.showPasswordForm = true;
      }
    },
    submitPayment() {
      let AVHBS_payment = {
          User: this.user,
          Balance: this.newCredit,
          PaymentMethod: this.paymentMethod,
          IntentID: null,
        };
      this.$http
        .post("pay", AVHBS_payment)
        .then((response) => {
          AVHBS_payment.IntentID = response.data;
          this.$emit("close");
          this.handlePaymentIntent(AVHBS_payment);
        })
        .catch((response) => {
          //ToDo: internationalize this failure message
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    handlePaymentIntent(payment){
      // payment.IntentID
      let status = ""; 
      let pi_data = {};
      this.$http.post("confirmPaymentIntent", payment)
        .then((resp) => {
          pi_data = resp.data;
          status = pi_data.action.status;
          let message = `${this.$t('booking.payment.processing')}`// Standard is Processing
          this.$responseEventBus.$emit("close");
          this.$responseEventBus.$emit("processingMessage", message);
          if (status == "in_progress") {
            setTimeout(() => {this.handlePaymentIntent(payment)}, 1000);
          }
          if(status == "failed"){
            let message = `${this.$t('booking.payment.failed')} ${pi_data.action.failure_message}`; // Internationalize the Message
            this.$responseEventBus.$emit("failureMessage", message);
          }
          if(status == "succeeded"){
            //ToDo: remove currency Symbol and use internationalisation on payment
            this.$store.commit("getLastNBookEntries", 5);
            this.$store.commit("getUsers");
            this.$store.commit("selectUser", {});
            let message = `${this.displayUserName(this.user)} ${this.$t('messages.success.paid')} ${this.newCredit}€`;
            this.$responseEventBus.$emit("successMessage", message);
          }
        }).catch(() => {
          this.$responseEventBus.$emit("failureMessage", status);
          return;
        })
    },
    weekdayIsMonday() {
      var day = new Date().getDay();
      if (day === 1) {
        return true;
      }
      return true;
      // Return False to only allow Payment on Mondays.
      // Currently Payment is allowed on all Weekdays.
      // return false;
    },
    loginAndPay() {
      this.$http
        .post("login", this.password)
        .then(() => {
          this.submitPayment();
        })
        .catch(() => {
          this.validationError = `${this.$t('messages.failure.error')}. ${this.$t('messages.failure.wrongPassword')}`;
        });
    },
    cancel() {
      this.$emit("close");
    },
  },
  created() {
    this.$http
      .post("getUserDebts", this.user)
      .then((response) => {
        this.userBookEntries = response.body;
      })
      .catch(() => {
        this.$responseEventBus.$emit("failureMessage", this.$t('messages.failure.noDebts'));
      });
  },
};
</script>

<style lang="scss">
@import "../../assets/paymentModal.css";
</style>
