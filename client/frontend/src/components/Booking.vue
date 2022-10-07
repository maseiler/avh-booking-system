<template>
  <div id="booking" class="booking">
    <div class="columns is-gapless">
      <div class="column image.png">
        <div class="box">
          <BookingUserCol />
        </div>
      </div>
      <div class="column is-one-fifth">
        <div class="box">
          <BookingMidCol />
        </div>
      </div>
      <div class="column is-two-fifth">
        <div class="box">
          <h1 class="title is-1">Item</h1>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import BookingUserCol from "./BookingUserCol.vue"
import BookingMidCol from "./BookingMidCol.vue"
import { useUserStore } from '../stores/userStore'

const userStore = useUserStore()

export default {
  components: {
    BookingUserCol,
    BookingMidCol
  },
  created(): void {
    userStore.fetchAllUsers
  }
}
</script>
