<template>
  <div>
    <br />
    <div class="tile is-ancestor is-narrow">
      <div class="tile is-parent is-vertical">
        <article class="tile is-child">
          <div class="box" style="height: 21rem; overflow: auto">
            <!-- Buttons -->
            <columns class="columns">
              <div class="column is-2">
                <button class="button is-link" @click="getBookEntries">
                  Get Book Entries
                </button>
              </div>
              <div class="column" align="left">
                <button
                  class="button is-link is-outlined"
                  @click="option = 'optLastN'"
                >
                  Last N
                </button>
                &nbsp;
                <button
                  class="button is-link is-outlined"
                  @click="option = 'optFromUser'"
                >
                  Of User
                </button>
                &nbsp;
                <button
                  class="button is-link is-outlined"
                  @click="option = 'optFromItem'"
                >
                  Of Item
                </button>
                &nbsp;
                <button
                  class="button is-link is-outlined"
                  @click="option = 'optUserPayments'"
                >
                  User Payments
                </button>
              </div>
              <div class="column is-1">
                <button
                  class="button is-warning is-outlined"
                  v-on:click="downloadCsv"
                >
                  Download
                </button>
              </div>
              <div class="column">
                &nbsp;
                <button class="button is-danger is-outlined" @click="undoEntry">
                  Undo Entry
                </button>
                &nbsp;
                <button
                  class="button is-danger is-outlined"
                  @click="deleteEntry"
                >
                  Delete Entry
                </button>
              </div>
            </columns>
            <!-- Input Fields -->
            <columns class="columns">
              <div class="column is-2">
                <div
                  v-if="
                    option === 'optFromUser' ||
                    option === 'optFromItem' ||
                    option === 'optUserPayments'
                  "
                >
                  <div class="field">
                    <div class="control">
                      <input
                        class="input is-info"
                        type="date"
                        placeholder="From"
                        v-model="from"
                      />
                    </div>
                    <br />
                    <div class="field">
                      <div class="control">
                        <input
                          class="input is-info"
                          type="date"
                          placeholder="To"
                          v-model="to"
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="column">
                <div class="field" v-if="option === 'optLastN'" align="left">
                  <div class="control">
                    <input
                      class="input is-info"
                      type="number"
                      placeholder="N"
                      v-model="n"
                      style="width: 20em"
                    />
                  </div>
                </div>
                <div
                  v-if="
                    option === 'optFromUser' || option === 'optUserPayments'
                  "
                >
                  <UserSearch />
                </div>
                <div v-if="option === 'optFromItem'">
                  <ItemSearch :mode="'single'" />
                </div>
              </div>
            </columns>
          </div>
        </article>
        <article class="tile is-child">
          <div
            class="box"
            style="height: calc(100vh - 25.25rem); overflow: auto"
          >
            <BookEntryList
              @selectEntry="selectEntry"
              :bookEntries="bookEntries"
            />
          </div>
        </article>
      </div>
    </div>
  </div>
</template>

<script>
import UserSearch from "./UserSearch.vue";
import ItemSearch from "./ItemSearch.vue";
import BookEntryList from "./BookEntryList.vue";

export default {
  components: {
    UserSearch,
    ItemSearch,
    BookEntryList,
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
      if (
        confirm(
          "The entry will be deleted and can not be restored. Prefer 'Undo Entry' for traceability. Use this button only in case of emergency.\n\nContinue?"
        )
      ) {
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
