<template>
  <div>
    <div class="list is-hoverable">
      <a v-for="entry in bookings" :key="entry" class="list-item has-background-white">
        {{printDateTime(entry.TimeStamp)}}
        <br>
        {{displayUserName(getUserByID(entry.UserID))}} ~ {{entry.Amount}}x {{displayItem(getItemByID(entry.ItemID))}}
      </a>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    bookings() {
      return this.$store.getters.last5Bookings;
    }
  }
};
</script>