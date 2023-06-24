<template>
  <div>
    <table class="table is-hoverable is-striped is-fullwidth">
      <thead>
        <tr>
          <th>{{$t('generic.id')}}</th>
          <th>{{$t('booking.payment.timestamp')}}</th>
          <th>{{$t('user.user')}}</th>
          <th>{{$t('item.item')}}</th>
          <th>{{$t('cart.amount')}}</th>
          <th style="text-align: right">{{$t('item.price')}}</th>
          <th>{{$t('generic.comment')}}</th>
          <th>{{$t('booking.payment.method')}}</th>
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
          <td style="text-align: right">{{ $n(entry.TotalPrice, "currency", "de-DE" ) }}</td>
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
        <th style="text-align: right">{{ $n(sum, "currency", "de-DE") }}</th>
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