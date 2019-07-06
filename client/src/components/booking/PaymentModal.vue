<template>
  <transition name="modal big">
    <div class="modal-mask big">
      <div class="modal-wrapper big">
        <div class="modal-container big">
          <div class="modal-header big">
            <h1 class="title is-4">Payment</h1>
            <hr />
          </div>

          <div class="modal-body big">
            <p class="subtitle is-4">
              <b>{{displayUserName(user)}}</b> has to pay
              <b>{{user.Balance}}€</b>.
            </p>
            <br />
            <p class="subtitle is-4">Please place the cash in the money box!</p>
          </div>

          <div class="modal-footer big">
            <div class="level">
              <div class="level-item has-text-centered">
                <button class="button is-link" @click="pay">Pay</button>
              </div>
              <div class="level-item has-text-centered">
                <button class="button" @click="cancel">Cancel</button>
              </div>
            </div>
            <hr />
            <p
              v-if="userBookings.LastPayment !== '0001-01-01T00:00:00Z'"
              class="subtitle is-6"
            >Last Payment: {{printDateTime(userBookings.LastPayment)}}</p>
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
  props: {
    user: {}
  },
  data() {
    return {
      userBookings: {}
    };
  },
  methods: {
    pay() {
      if (this.weekdayIsMonday()) {
        this.$http
          .post("pay", this.user)
          .then(function(response) {
            var message = "".concat(
              this.displayUserName(this.user),
              " payed ",
              this.user.Balance
            );
            this.$emit("close");
            this.$store.commit("getLast5Bookings");
            this.$store.commit("getUsers");
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch(function(response) {
            this.$responseEventBus.$emit("failureMessage", response.data);
          });
      } else {
        this.$emit("close");
        this.$responseEventBus.$emit(
          "failureMessage",
          "Payment only on Mondays possible!"
        );
      }
    },
    weekdayIsMonday() {
      return true;
      var day = new Date().getDay();
      if (day == 1) {
        return true;
      } else {
        //TODO: require password
        return false;
      }
    },
    cancel() {
      this.$emit("close");
    }
  },
  created() {
    this.$http
      .post("getUserDebts", this.user)
      .then(function(response) {
        this.userBookings = response.body;
      })
      .catch(function(response) {
        this.$responseEventBus.$emit("failureMessage", "Couldn't get debts.");
      });
  }
};
</script>

<style lang="scss">
@import "../../assets/modalBig.css";
</style>
