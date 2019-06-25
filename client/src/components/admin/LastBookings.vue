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
          <td>{{displayUserName(getUser(entry.UserID))}}</td>
          <td>{{displayItem(getItem(entry.ItemID))}}</td>
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
    printDateTime(dateTime) {
      var d = new Date(Date.parse(dateTime));
      return (
        ("0" + d.getDate()).slice(-2) +
        "." +
        ("0" + (d.getMonth() + 1)).slice(-2) +
        "." +
        d.getFullYear() +
        " " +
        ("0" + d.getHours()).slice(-2) +
        ":" +
        ("0" + d.getMinutes()).slice(-2) +
        ":" +
        ("0" + d.getSeconds()).slice(-2)
      );
    },
    getUser(id) {
      return this.users.find(u => {
        return u.UserID == id;
      });
    },
    getItem(id) {
      return this.items.find(i => {
        return i.ItemID == id;
      });
    },
    selectEntry(entry) {
      this.selectedEntry = entry;
      this.$emit("selectEntry", entry);
    }
  }
};
</script>