<template>
  <div>
    <table class="table is-hoverable is-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Time Stamp</th>
          <th>User</th>
          <th>Item</th>
          <th>Amount</th>
          <th>Total Price</th>
          <th>Comment</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="entry in lastBookings"
          :key="entry"
          @click="selectEntry(entry)"
          :class="[selectedEntry === entry ? 'is-selected' : '']"
        >
          <th>{{entry.BookEntryID}}</th>
          <td>{{printDateTime(entry.TimeStamp)}}</td>
          <td>{{displayUserName(getUserByID(entry.UserID))}}</td>
          <td>{{displayItem(getItemByID(entry.ItemID))}}</td>
          <td>{{entry.Amount}}</td>
          <td>{{entry.TotalPrice}}</td>
          <td>{{entry.Comment}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  computed: {
    lastBookings() {
      return this.$store.state.bookings;
    },
    users() {
      return this.$store.state.users;
    },
    items() {
      return this.$store.state.items;
    }
  },
  data() {
    return {
      selectedEntry: {}
    };
  },
  methods: {
    selectEntry(entry) {
      this.selectedEntry = entry;
      this.$emit("selectEntry", entry);
    }
  }
};
</script>