<template>
  <div>
    <div class="columns">
      <div class="column is-2">
        <button class="button is-link" @click="deleteEntry">Delete entry</button>
      </div>
      <div class="column">
        <LastBookings
          :lastBookings="lastBookings"
          :users="users"
          :items="items"
          @selectEntry="selectEntry"
        />
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
  props: {
    lastBookings: [],
    users: [],
    items: []
  },
  data: function() {
    return {
      selectedEntry: {},
      showModifyBookEntryForm: false
    };
  },
  methods: {
    selectEntry: function(entry) {
      this.selectedEntry = entry;
    },
    deleteEntry: function() {
      if (Object.keys(this.selectedEntry).length !== 0) {
        this.$http
          .post("/deleteBookEntry", this.selectedEntry)
          .then(function(response) {
            this.$router.go();
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