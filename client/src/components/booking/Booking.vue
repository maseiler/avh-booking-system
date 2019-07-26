<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box">
          <UserTab />
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box has-text-centered" style="width: 500px;">
          <button class="button" @click="refreshPage">
            <span class="icon is-small is-left">
              <font-awesome-icon icon="redo" />
            </span>
            &nbsp;&nbsp;Refresh
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
      </div>
      <div class="column">
        <div class="box">
          <ItemTab />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserTab from "./UserTab.vue";
import ItemTab from "./ItemTab.vue";
import Cart from "./Cart.vue";
import UserInfo from "./UserInfo.vue";
import Last5Bookings from "./Last5Bookings.vue";

export default {
  components: {
    UserTab,
    ItemTab,
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