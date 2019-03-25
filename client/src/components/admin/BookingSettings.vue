<template>
  <div>
    <h1>BOOKING SETTINGS</h1>
    <div class="columns">
      <div class="column is-2">
        <button class="button is-link" @click="showModifyBookEntryForm = true">Modify Book Entry</button>
        <ModifyBookEntryForm v-if="showModifyBookEntryForm" :entry="selectedEntry"/>
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
import ModifyBookEntryForm from "./ModifyBookEntryForm.vue";

export default {
  components: {
    LastBookings,
    ModifyBookEntryForm
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
    }
  }
};
</script>