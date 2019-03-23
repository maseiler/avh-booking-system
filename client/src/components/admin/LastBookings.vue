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
          <td>{{displayName(getUser(entry.UserID))}}</td>
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
  props: {
    lastBookings: [],
    users: [],
    items: []
  },
  data: function() {
    return {
      selectedEntry: {}
    };
  },
  methods: {
    printDateTime: function(dateTime) {
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
    getUser: function(id) {
      return this.users.find(u => {
        return u.UserID == id;
      });
    },
    displayName: function(user) {
      if (user.BierName != "") {
        return user.BierName;
      } else {
        if (user.LastName != "") {
          return user.FirstName + " " + user.LastName[0] + ".";
        }
        return user.FirstName;
      }
    },
    getItem: function(id) {
      return this.items.find(i => {
        return i.ItemID == id;
      });
    },
    displayItem: function(item) {
      if (item.Type == "boat" || item.Type == "food") {
        return item.Name;
      } else {
        return item.Name + " " + item.Size + item.Unit;
      }
    },
    selectEntry: function(entry) {
      this.selectedEntry = entry;
      this.$emit("selectEntry", entry);
    }
  }
};
</script>