<template>
  <div>
    <br />
    <div class="columns">
      <div class="column is-2">
        <br />
        <button class="button is-link is-fullwidth is-outlined" @click="getBookings">Get Bookings</button>
        <br />
        <button class="button is-link is-fullwidth" @click="option = 'optLastN'">Last N</button>
        <div class="box" v-if="option === 'optLastN'">
          <br />
          <div class="field">
            <div class="control">
              <input class="input is-info" type="number" placeholder="N" v-model="n" />
            </div>
          </div>
        </div>
        <br />

        <button class="button is-link is-fullwidth" @click="option = 'optFromUser'">Of User</button>
        <div class="box" v-if="option === 'optFromUser'">
          <br />
          <UserSearch />
          <p class="p has-text-grey-light">Time span can be left empty or in format YYY-MM-DD</p>
          <div class="columns">
            <div class="column">
              <div class="field">
                <div class="control">
                  <input class="input is-info" type="Text" placeholder="From" v-model="from" />
                </div>
              </div>
            </div>
            <div class="column">
              <div class="field">
                <div class="control">
                  <input class="input is-info" type="Text" placeholder="To" v-model="to" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <br />

        <button class="button is-link is-fullwidth" @click="option = 'optFromItem'">Of Item</button>
        <div class="box" v-if="option === 'optFromItem'">
          <br />
          <ItemSearch />
          <p class="p has-text-grey-light">Time span can be left empty or in format YYY-MM-DD</p>
          <div class="columns">
            <div class="column">
              <div class="field">
                <div class="control">
                  <input class="input is-info" type="Text" placeholder="From" v-model="from" />
                </div>
              </div>
            </div>
            <div class="column">
              <div class="field">
                <div class="control">
                  <input class="input is-info" type="Text" placeholder="To" v-model="to" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <br />

        <button
          class="button is-link is-fullwidth"
          @click="option = 'optUserPayments'"
        >User Payments</button>
        <div class="box" v-if="option === 'optUserPayments'">
          <br />
          <UserSearch />
          <p class="p has-text-grey-light">
            User can be empty.
            <br />Time span can be left empty or in format YYY-MM-DD
          </p>
          <div class="columns">
            <div class="column">
              <div class="field">
                <div class="control">
                  <input class="input is-info" type="Text" placeholder="From" v-model="from" />
                </div>
              </div>
            </div>
            <div class="column">
              <div class="field">
                <div class="control">
                  <input class="input is-info" type="Text" placeholder="To" v-model="to" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <br />
        <hr />
        <br />
        <button class="button is-danger is-outlined" @click="deleteEntry">Delete entry</button>
      </div>
      <div class="column">
        <LastBookings @selectEntry="selectEntry" :bookings="bookings" />
      </div>
    </div>
  </div>
</template>

<script>
import LastBookings from "./LastBookings.vue";
import UserSearch from "./../booking/UserSearch.vue";
import ItemSearch from "./../booking/ItemSearch.vue";

export default {
  components: {
    LastBookings,
    UserSearch,
    ItemSearch
  },
  data() {
    return {
      selectedEntry: {},
      bookings: [],
      option: "",
      n: 0,
      from: "",
      to: "",
      user: {},
      item: {}
    };
  },
  methods: {
    getBookings() {
      if (this.option === "optLastN") {
        if (this.n > 0) {
          this.$http.post("/getLastNBookings", this.n).then(response => {
            this.bookings = [].concat.apply([], response.body);
          });
        }
      } else if (this.option === "optFromUser") {
        var userFromTo = new Object();
        userFromTo.User = this.user;
        userFromTo.From = this.from;
        userFromTo.To = this.to;
        this.$http
          .post("getBookingsFromUserBetween", userFromTo)
          .then(response => {
            this.bookings = [].concat.apply([], response.body);
          })
          .catch(response => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't fetch bookings"
            );
          });
      } else if (this.option === "optFromItem") {
        var itemFromTo = new Object();
        itemFromTo.Item = this.item;
        itemFromTo.From = this.from;
        itemFromTo.To = this.to;
        this.$http
          .post("getBookingsFromItemBetween", itemFromTo)
          .then(response => {
            this.bookings = [].concat.apply([], response.body);
          })
          .catch(response => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't fetch bookings"
            );
          });
      } else if (this.option === "optUserPayments") {
        var userFromTo = new Object();
        userFromTo.User = this.user;
        userFromTo.From = this.from;
        userFromTo.To = this.to;
        this.$http
          .post("getPaymentsOfUser", userFromTo)
          .then(response => {
            this.bookings = [].concat.apply([], response.body);
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't fetch bookings"
            );
          });
      } else {
        this.$responseEventBus.$emit(
          "failureMessage",
          "Select an option first"
        );
      }
    },
    selectEntry(entry) {
      this.selectedEntry = entry;
    },
    deleteEntry() {
      if (Object.keys(this.selectedEntry).length !== 0) {
        this.$http
          .post("deleteBookEntry", this.selectedEntry)
          .then(function(response) {
            this.$store.commit("getLastNBookings", 50);
            this.$responseEventBus.$emit("successMessage", "Deleted bookEntry");
          })
          .catch(function(response) {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't delete item."
            );
          });
      } else {
        this.$responseEventBus.$emit(
          "failureMessage",
          "Select an entry first!"
        );
      }
    },
    selectUserFromBus(user) {
      this.user = user;
    },
    selectItemFromBus(items) {
      this.item = items[0];
    }
  },
  created() {
    this.$userEventBus.$on("selectUserToBus", this.selectUserFromBus);
    this.$itemEventBus.$on("selectItemsToBus", this.selectItemFromBus);
  }
};
</script>
