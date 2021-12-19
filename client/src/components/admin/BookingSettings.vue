<template>
  <div>
    <br />
    <div class="columns">
      <div class="column is-2">
        <br />
        <button
          class="button is-link is-fullwidth is-outlined"
          @click="getBookEntries"
        >
          Get Book Entries
        </button>
        <br />
        <button
          class="button is-link is-fullwidth"
          @click="option = 'optLastN'"
        >
          Last N
        </button>
        <div class="box" v-if="option === 'optLastN'">
          <br />
          <div class="field">
            <div class="control">
              <input
                class="input is-info"
                type="number"
                placeholder="N"
                v-model="n"
              />
            </div>
          </div>
        </div>
        <br />

        <button
          class="button is-link is-fullwidth"
          @click="option = 'optFromUser'"
        >
          Of User
        </button>
        <div class="box" v-if="option === 'optFromUser'">
          <br />
          <UserSearch />
          <p class="p has-text-grey-light">
            Time span can be left empty or in format YYYY-MM-DD
          </p>
          <div class="columns">
            <div class="column">
              <div class="field">
                <div class="control">
                  <input
                    class="input is-info"
                    type="Text"
                    placeholder="From"
                    v-model="from"
                  />
                </div>
              </div>
            </div>
            <div class="column">
              <div class="field">
                <div class="control">
                  <input
                    class="input is-info"
                    type="Text"
                    placeholder="To"
                    v-model="to"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
        <br />

        <button
          class="button is-link is-fullwidth"
          @click="option = 'optFromItem'"
        >
          Of Item
        </button>
        <div class="box" v-if="option === 'optFromItem'">
          <br />
          <ItemSearch :mode="'single'" />
          <p class="p has-text-grey-light">
            Time span can be left empty or in format YYYY-MM-DD
          </p>
          <div class="columns">
            <div class="column">
              <div class="field">
                <div class="control">
                  <input
                    class="input is-info"
                    type="Text"
                    placeholder="From"
                    v-model="from"
                  />
                </div>
              </div>
            </div>
            <div class="column">
              <div class="field">
                <div class="control">
                  <input
                    class="input is-info"
                    type="Text"
                    placeholder="To"
                    v-model="to"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
        <br />

        <button
          class="button is-link is-fullwidth"
          @click="option = 'optUserPayments'"
        >
          User Payments
        </button>
        <div class="box" v-if="option === 'optUserPayments'">
          <br />
          <UserSearch />
          <p class="p has-text-grey-light">
            User can be empty.
            <br />Time span can be left empty or in format YYYY-MM-DD
          </p>
          <div class="columns">
            <div class="column">
              <div class="field">
                <div class="control">
                  <input
                    class="input is-info"
                    type="Text"
                    placeholder="From"
                    v-model="from"
                  />
                </div>
              </div>
            </div>
            <div class="column">
              <div class="field">
                <div class="control">
                  <input
                    class="input is-info"
                    type="Text"
                    placeholder="To"
                    v-model="to"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
        <br />
        <button
          class="button is-link is-fullwidth is-outlined"
          v-on:click="downloadCsv"
        >
          Download
        </button>
        <br />
        <hr />
        <br />
        <button class="button is-danger is-outlined" @click="undoEntry">
          Undo Entry
        </button>
        <p class="p has-text-grey-light">
          Modifies user's balance accordingly and creates a new book entry.
        </p>
        <hr />
        <button class="button is-danger is-outlined" @click="deleteEntry">
          Delete Entry
        </button>
        <p class="p has-text-grey-light">
          Use deliberately!
          <br />The entry will be deleted and can not be restored.
        </p>
      </div>
      <div class="column">
        <div class="box" style="height: 88vh; overflow: auto">
          <BookEntryList @selectEntry="selectEntry" :bookEntries="bookEntries" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BookEntryList from "./BookEntryList.vue";
import UserSearch from "./../booking/UserSearch.vue";
import ItemSearch from "./../booking/ItemSearch.vue";

export default {
  components: {
    BookEntryList,
    UserSearch,
    ItemSearch,
  },
  computed: {
    user() {
      return this.$store.state.selectedUser;
    },
    item() {
      return this.$store.state.selectedSingleItem;
    },
  },
  data() {
    return {
      selectedEntry: {},
      bookEntries: [],
      option: "",
      n: 0,
      from: "",
      to: "",
    };
  },
  methods: {
    getBookEntries() {
      if (this.option === "optLastN") {
        if (this.n > 0) {
          this.$http
            .post("getLastNBookEntries", this.n)
            .then((response) => {
              this.bookEntries = [].concat.apply([], response.body);
            })
            .catch(() => {
              this.$responseEventBus.$emit(
                "failureMessage",
                "Couldn't fetch book entries."
              );
            });
        }
      } else if (this.option === "optFromUser") {
        var userFromTo = new Object();
        userFromTo.User = this.user;
        userFromTo.From = this.from;
        userFromTo.To = this.to;
        this.$http
          .post("getBookEntriesFromUserBetween", userFromTo)
          .then((response) => {
            this.bookEntries = [].concat.apply([], response.body);
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't fetch book entries."
            );
          });
      } else if (this.option === "optFromItem") {
        var itemFromTo = new Object();
        itemFromTo.Item = this.item;
        itemFromTo.From = this.from;
        itemFromTo.To = this.to;
        this.$http
          .post("getBookEntriesFromItemBetween", itemFromTo)
          .then((response) => {
            this.bookEntries = [].concat.apply([], response.body);
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't fetch book entries."
            );
          });
      } else if (this.option === "optUserPayments") {
        userFromTo = new Object();
        userFromTo.User = this.user;
        userFromTo.From = this.from;
        userFromTo.To = this.to;
        this.$http
          .post("getPaymentsOfUser", userFromTo)
          .then((response) => {
            this.bookEntries = [].concat.apply([], response.body);
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't fetch book entries."
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
          .then(() => {
            this.$store.commit("getLastNBookEntries", 5);
            this.$responseEventBus.$emit(
              "successMessage",
              "Deleted book entry"
            );
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't delete book entry."
            );
          });
      } else {
        this.$responseEventBus.$emit(
          "failureMessage",
          "Select an book entry first!"
        );
      }
    },
    undoEntry() {
      if (Object.keys(this.selectedEntry).length !== 0) {
        this.$http
          .post("undoBookEntry", this.selectedEntry)
          .then(() => {
            this.$store.commit("getLastNBookEntries", 5);
            this.$responseEventBus.$emit("successMessage", "Undid book entry");
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Couldn't undo book entry."
            );
          });
      } else {
        this.$responseEventBus.$emit(
          "failureMessage",
          "Select an book entry first!"
        );
      }
    },
    downloadCsv() {
      let csv = "ID,Time Stamp,User,Item,Amount,Price,Comment,Payment Method\n";
      this.bookEntries.forEach((entry) => {
        csv += entry.ID + ",";
        csv += this.printDateTime(entry.TimeStamp) + ",";
        csv +=
          this.displayUserName(this.getUserByID(entry.UserID)) +
          " (ID: " +
          entry.UserID +
          "),";
        csv +=
          this.displayItem(
            this.getItemByID(this.$store.state.items, entry.ItemID)
          ) +
          " (ID: " +
          entry.ItemID +
          "),";
        csv += entry.Amount + ",";
        csv += entry.TotalPrice + ",";
        csv += entry.Comment + ",";
        csv += entry.PaymentMethod + "\n";
      });
      let fileName =
        "book_entries_" +
        this.cleanTimeStamp(this.bookEntries[0].TimeStamp) +
        "_-_" +
        this.cleanTimeStamp(
          this.bookEntries[this.bookEntries.length - 1].TimeStamp
        ) +
        ".csv";

      let anchor = document.createElement("a");
      anchor.href = "data:text/csv;charset=utf-8," + encodeURIComponent(csv);
      anchor.target = "_blank";
      anchor.download = fileName;
      anchor.click();
    },
    cleanTimeStamp(timeStamp) {
      let t = timeStamp;
      t = t.replace("T", "_");
      t = t.replace(":", "-");
      t = t.replace(":", "-");
      t = t.replace("Z", "");
      return t;
    },
  },
};
</script>
