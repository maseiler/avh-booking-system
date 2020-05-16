<template>
  <div>
    <button class="button" @click="refreshPage">
      <span class="icon is-small">
        <font-awesome-icon icon="redo" />
      </span>
      <p>Refresh</p>
    </button>
    <br />
    <div
      v-if="Object.keys(this.$store.state.selectedUser).length !== 0 || Object.keys(this.$store.state.selectedMultipleItems).length !== 0"
    >
      <Cart />
      <hr />
      <UserInfo></UserInfo>
    </div>
    <div v-else>
      <img class="logo" src="~@/assets/avh_logo.png" width="250px" />
      <hr />
      <Last5Bookings />
    </div>
  </div>
</template>

<script>
import Cart from "./Cart.vue";
import UserInfo from "./UserInfo.vue";
import Last5Bookings from "./Last5Bookings.vue";

export default {
  components: {
    Cart,
    UserInfo,
    Last5Bookings
  },
  methods: {
    refreshPage() {
      this.$router.go();
    }
  }
};
</script>