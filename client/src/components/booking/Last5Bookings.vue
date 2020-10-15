<template>
  <div>
    <div class="list is-hoverable">
      <a
        v-for="entry in bookings"
        :key="entry"
        class="list-item has-background-white"
      >
        {{ printDateTime(entry.TimeStamp) }}
        <br />
        {{ displayBooking(entry) }}
      </a>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    bookings() {
      return this.$store.state.last5Bookings;
    },
  },
  methods: {
    displayBooking(entry) {
      if (entry.ItemID === 0) {
        if (entry.Comment.startsWith("Undo")) {
          return (
            this.displayUserName(this.getUserByID(entry.UserID)) +
            " ~ Undid booking: " +
            entry.Amount +
            "€"
          );
        } else if (entry.Comment.startsWith("Payment")) {
          return (
            this.displayUserName(this.getUserByID(entry.UserID)) +
            " ~ Payment: " +
            entry.TotalPrice * -1 +
            "€"
          );
        }
        return this.displayUserName(this.getUserByID(entry.UserID)) + " ~ ???";
      } else {
        return (
          this.displayUserName(this.getUserByID(entry.UserID)) +
          " ~ " +
          entry.Amount +
          "x " +
          this.displayItem(this.getItemByID(entry.ItemID))
        );
      }
    },
  },
};
</script>