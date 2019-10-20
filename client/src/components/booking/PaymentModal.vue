<template>
  <transition name="paymentModal">
    <div class="paymentModal-mask big">
      <div class="paymentModal-wrapper big">
        <div class="paymentModal-container big">
          <div class="paymentModal-header big">
            <h1 class="title is-4">Payment</h1>
            <hr />
          </div>

          <div class="paymentModal-body big">
            <p class="subtitle is-4">
              <b>{{displayUserName(user)}}</b> has to pay
              <b>{{user.Balance}}€</b>.
            </p>
            <br />
            <p class="subtitle is-4">Please place the cash in the money box!</p>
            <div class="level">
              <div class="level-item" style="text-align:right;">
                <div class="column is-4">
                  <button class="button is-link is-fullwidth" @click="pay">Pay</button>
                  <input class="input" type="text" placeholder="€" v-model.number="balancePart" />
                </div>
              </div>
              <div class="level-item">
                <div class="column" style="text-align:left;">
                  <button class="button" @click="cancel">Cancel</button>
                </div>
              </div>
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
              v-if="userBookings.LastPayment !== '0001-01-01T00:00:00Z'"
              class="subtitle is-6"
            >Last Payment: {{printDateTime(userBookings.LastPayment)}}</p>
          </div>

          <div class="paymentModal-footer big">
            <table class="table is-hoverable is-striped">
              <thead>
                <tr>
                  <th>Timestamp</th>
                  <th>Item</th>
                  <th>Amount</th>
                  <th>Total Price</th>
                  <th>Comment</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="entry in userBookings.Debts" :key="entry">
                  <td>{{printDateTime(entry.TimeStamp)}}</td>
                  <td>{{displayItem(getItemByID(entry.ItemID))}}</td>
                  <td>{{entry.Amount}}</td>
                  <td>{{entry.TotalPrice}}</td>
                  <td>{{entry.Comment}}</td>
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
    }
  },
  data() {
    return {
      userBookings: {},
      showPasswordForm: false,
      password: "",
      balancePart: 0
    };
  },
  methods: {
    pay() {
      if (this.balancePart <= 0) {
        this.$emit("close");
        this.$responseEventBus.$emit(
          "failureMessage",
          "Payments smaller or equal 0 are not allowed"
        );
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
        .post("pay", { User: this.user, DoubleValue: this.balancePart })
        .then(() => {
          var message = "".concat(
            this.displayUserName(this.user),
            " payed ",
            this.balancePart
          );
          this.$store.commit("getLast5Bookings");
          this.$store.commit("getUsers");
          this.$store.commit("selectUser", {});
          this.$emit("close");
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch(response => {
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
    }
  },
  created() {
    this.balancePart = this.user.Balance;

    this.$http
      .post("getUserDebts", this.user)
      .then(response => {
        this.userBookings = response.body;
      })
      .catch(() => {
        this.$responseEventBus.$emit("failureMessage", "Couldn't get debts.");
      });
  }
};
</script>

<style lang="scss">
@import "../../assets/paymentModal.css";
</style>
