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
          <Cart
            :user="selectedUser"
            :items="selectedItems"
            v-if="Object.keys(selectedUser).length !== 0 || Object.keys(selectedItems).length !== 0"
          />
          <img class="logo" src="~@/assets/avh_logo.png" width="250px" v-else>
          <UserInfo :user="selectedUser"></UserInfo>
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

export default {
  components: {
    UserTab,
    ItemTab,
    Cart,
    UserInfo
  },
  data: function() {
    return {
      selectedUser: {},
      selectedItems: []
    };
  },
  methods: {
    getSelectedUser: function(user) {
      this.selectedUser = user;
    },
    getSelectedItems: function(items) {
      this.selectedItems = items;
    }
  }
};
</script>