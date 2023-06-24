<template>
  <transition name="paymentModal">
    <div class="paymentModal-mask big">
      <div class="paymentModal-wrapper big">
        <div class="paymentModal-container big">
          <div class="paymentModal-header big">
            <h1 class="title is-4">{{$t('booking.payment.title')}}</h1>
            <hr />
          </div>

          <div class="paymentModal-body big">
            <p class="subtitle is-4">
              <b>{{ displayUserName(user) }}</b> {{$t('booking.payment.htp')}}
              <b>{{ user.Balance }}€</b>.
            </p>
            <br />
            <p class="subtitle is-4">{{$t('booking.payment.subtitle')}}</p>
            <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-3">
                <div class="field">
                  <p class="control has-icons-right">
                    <input
                      class="input"
                      type="text"
                      placeholder="€"
                      v-model.number="balancePart"
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
                <button class="button is-success is-fullwidth" @click="payCash">
                  {{$t('booking.payment.cash')}}
                </button>
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
                      placeholder="Password"
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
      balancePart: 0,
      paymentMethod: "",
    };
  },
  methods: {
    payCash() {
      this.paymentMethod = "Cash";
      this.pay();
    },
    payEC() {
      this.paymentMethod = "EC";
      this.pay();
    },
    pay() {
      if (this.balancePart <= 0) {
        this.$emit("close");
        this.$responseEventBus.$emit(
          "failureMessage",
          "Payments smaller or equal 0 are not allowed"
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
      this.$http
        .post("pay", {
          User: this.user,
          Balance: this.balancePart,
          PaymentMethod: this.paymentMethod,
        })
        .then(() => {
          var message = "".concat(
            this.displayUserName(this.user),
            " paid ",
            this.balancePart,
            " €"
          );
          this.$store.commit("getLastNBookEntries", 5);
          this.$store.commit("getUsers");
          this.$store.commit("selectUser", {});
          this.$emit("close");
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch((response) => {
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    weekdayIsMonday() {
      var day = new Date().getDay();
      if (day === 1) {
        return true;
      }
      return false;
    },
    loginAndPay() {
      this.$http
        .post("login", this.password)
        .then(() => {
          this.submitPayment();
        })
        .catch(() => {
          this.validationError = "Error. Wrong password?";
        });
    },
    cancel() {
      this.$emit("close");
    },
  },
  created() {
    this.balancePart = this.user.Balance;

    this.$http
      .post("getUserDebts", this.user)
      .then((response) => {
        this.userBookEntries = response.body;
      })
      .catch(() => {
        this.$responseEventBus.$emit("failureMessage", "Couldn't get debts.");
      });
  },
};
</script>

<style lang="scss">
@import "../../assets/paymentModal.css";
</style>
