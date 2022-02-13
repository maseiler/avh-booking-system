<template>
  <div>
    <table class="table is-hoverable is-striped is-fullwidth">
      <thead>
        <tr>
          <th>ID</th>
          <th>Timestamp</th>
          <th>User</th>
          <th>Item</th>
          <th>Amount</th>
          <th style="text-align: right">Price</th>
          <th>Comment</th>
          <th>Payment Method</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="entry in bookEntries"
          :key="entry"
          @click="selectEntry(entry)"
          :class="[selectedEntry === entry ? 'is-selected' : '']"
        >
          <th>{{ entry.ID }}</th>
          <td>{{ printDateTime(entry.TimeStamp) }}</td>
          <td>{{ displayUserName(getUserByID(entry.UserID)) }}</td>
          <td>{{ displayItem(getItemByID(items, entry.ItemID)) }}</td>
          <td>{{ entry.Amount }}</td>
          <td style="text-align: right">{{ entry.TotalPrice }} €</td>
          <td>{{ entry.Comment }}</td>
          <td>{{ entry.PaymentMethod }}</td>
        </tr>
      </tbody>
      <tfoot>
        <th></th>
        <th></th>
        <th></th>
        <th></th>
        <th></th>
        <th style="text-align: right">{{ sum }} €</th>
        <th></th>
        <th></th>
      </tfoot>
    </table>
  </div>
</template>

<script>
export default {
  props: {
    bookEntries: [],
  },
  computed: {
    users() {
      return this.$store.state.users;
    },
    items() {
      return this.$store.state.items;
    },
  },
  data() {
    return {
      selectedEntry: {},
      sum: 0,
    };
  },
  watch: {
    bookEntries() {
      this.updateSum();
    },
  },
  methods: {
    selectEntry(entry) {
      this.selectedEntry = entry;
      this.$emit("selectEntry", entry);
    },
    updateSum() {
      var temp = 0;
      this.bookEntries.forEach((entry) => {
        temp += entry.TotalPrice;
      });
      temp = temp.toFixed(2);
      this.sum = temp;
    },
  },
};
</script>