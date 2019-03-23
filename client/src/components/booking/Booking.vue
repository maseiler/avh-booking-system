<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box">
          <UserTab></UserTab>
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box has-text-centered" style="width: 500px;">
          <img class="logo" src="~@/assets/avh_logo.png" width="250px">
        </div>
      </div>
      <div class="column">
        <div class="box">
          <ItemTab @selectedItem="getItem"></ItemTab>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserTab from "./UserTab.vue";
import ItemTab from "./ItemTab.vue";

export default {
  components: {
    UserTab,
    ItemTab
  },
  data: function() {
    return {
      selectedUser: {},
      selectedItem: {},
      stringCheckout: ""
    };
  },
  methods: {
    getUser: function(user) {
      this.selectedUser = user;
    },
    getItem: function(item) {
      this.selectedItem = item;
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