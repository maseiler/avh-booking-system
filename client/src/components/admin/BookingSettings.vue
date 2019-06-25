<template>
  <div>
    <div class="columns">
      <div class="column is-2">
        <button class="button is-link" @click="deleteEntry">Delete entry</button>
      </div>
      <div class="column">
        <LastBookings @selectEntry="selectEntry"/>
      </div>
    </div>
  </div>
</template>

<script>
import LastBookings from "./LastBookings.vue";

export default {
  components: {
    LastBookings
  },
  data() {
    return {
      selectedEntry: {},
      showModifyBookEntryForm: false
    };
  },
  methods: {
    selectEntry(entry) {
      this.selectedEntry = entry;
    },
    deleteEntry() {
      if (Object.keys(this.selectedEntry).length !== 0) {
        this.$http
          .post("/deleteBookEntry", this.selectedEntry)
          .then(function(response) {
            this.$store.commit("getLastNBookings", 50);
            this.$responseEventBus.$emit("successMessage", "Deleted bookEntry");
          })
          .catch(function(response) {
            this.$responseEventBus.$emit(
              "failureMessage",
              "Error: Couldn't delete item."
            );
          });
      } else {
        this.$responseEventBus.$emit(
          "failureMessage",
          "Select an entry first!"
        );
      }
    }
  }
};
</script>