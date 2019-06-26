<template>
  <div id="lastBookingNotification">
    <div class="notification is-link" v-for="entry in lastBookings" :key="entry">
      {{printDateTime(entry.TimeStamp)}} ł {{displayUserName(getUserByID(entry.UserID))}}
      <!-- <button class="button is-link">{{lastBookings[0]}}</button> -->
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    lastBookings() {
      return this.$store.getters.last3Bookings;
    }
  }
};
</script>

<style>
#lastBookingNotification {
  position: absolute;
  bottom: 0;
  left: 0;
}
</style>