<template>
  <div>
    <div class="list is-hoverable">
      <a v-for="entry in bookings" :key="entry" class="list-item has-background-white">
        {{printDateTime(entry.TimeStamp)}}
        <br />
        <div
          v-if="entry.ItemID === 0"
        >{{displayUserName(getUserByID(entry.UserID))}} ~ Payment: {{entry.TotalPrice * -1}}€</div>
        <div
          v-else
        >{{displayUserName(getUserByID(entry.UserID))}} ~ {{entry.Amount}}x {{displayItem(getItemByID(entry.ItemID))}}</div>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    bookings() {
      return this.$store.state.last5Bookings;
    }
  }
};
</script>