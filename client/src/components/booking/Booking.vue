<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box">
          <UserTab @selectedUser="getSelectedUser"/>
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box has-text-centered" style="width: 500px;">
          <div
            v-if="Object.keys(selectedUser).length !== 0 || Object.keys(selectedItems).length !== 0"
          >
            <Cart :user="selectedUser" :items="selectedItems"/>
            <hr>
            <UserInfo :user="selectedUser"></UserInfo>
          </div>
          <div v-else>
            <img class="logo" src="~@/assets/avh_logo.png" width="250px">
            <hr>
            <last5Bookings/>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="box">
          <ItemTab @selectedItems="getSelectedItems"/>
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
  data() {
    return {
      selectedUser: {},
      selectedItems: []
    };
  },
  methods: {
    getSelectedUser(user) {
      this.selectedUser = user;
    },
    getSelectedItems(items) {
      this.selectedItems = items;
    }
  }
};
</script>