<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box">
          <UserTab/>
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box has-text-centered" style="width: 500px;">
          <Cart :user="selectedUser" :items="selectedItems"/>
          <img class="logo" src="~@/assets/avh_logo.png" width="250px">
        </div>
      </div>
      <div class="column">
        <div class="box">
          <ItemTab @selectedItem="getItem"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserTab from "./UserTab.vue";
import ItemTab from "./ItemTab.vue";
import Cart from "./Cart.vue";

export default {
  components: {
    UserTab,
    ItemTab,
    Cart
  },
  data: function() {
    return {
      selectedUser: {},
      selectedItems: [],
      stringCheckout: ""
    };
  },
  methods: {
    getUser: function(user) {
      this.selectedUser = user;
    },
    getItem: function(item) {
      this.selectedItems.push(item);
    },
    checkout: function() {
      this.$http.get("/checkout").then(response => {
        this.stringCheckout = response.body;
      });
      console.log("checkout");
    }
  }
};
</script>