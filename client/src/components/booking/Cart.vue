<template>
  <div>
    <h3 class="subtitle is-3">{{bookingError}}</h3>
    <table class="table">
      <thead>
        <tr>
          <th>Item</th>
          <th>Price</th>
          <th>Amount</th>
        </tr>
      </thead>
      <tfoot>
        <tr>
          <th></th>
          <th>{{sum}} €</th>
          <th>
            <button class="button is-small" @click="checkout">Checkout</button>
          </th>
        </tr>
      </tfoot>
      <tbody>
        <tr v-for="i in cart" :key="i">
          <td>{{i.item.Name}} {{i.item.Size}}</td>
          <td>{{i.item.Price}} €</td>
          <td>
            <input class="input is-small" type="text" v-model="i.amount" @change="updateSum">
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  props: {
    user: {},
    items: []
  },
  watch: {
    items: function() {
      this.buildCart();
      this.updateSum();
    }
  },
  data: function() {
    return {
      cart: [],
      sum: 0,
      bookingError: ""
    };
  },
  methods: {
    buildCart: function() {
      var temp = [];
      this.items.forEach(item => {
        if (!this.contains(temp, item)) {
          var newItem = new Object();
          newItem.item = item;
          newItem.amount = 1;
          temp.push(newItem);
        } else {
          this.increment(temp, item);
        }
      });
      this.cart = temp;
    },
    contains: function(array, obj) {
      for (var i in array) {
        if (array[i].item.ItemID === obj.ItemID) {
          return true;
        }
      }
      return false;
    },
    increment: function(array, item) {
      for (var i in array) {
        if (array[i].item.ItemID === item.ItemID) {
          array[i].amount = array[i].amount + 1;
          return;
        }
      }
    },
    updateSum: function() {
      if (this.cart === []) {
        this.sum = 0;
      } else {
        var temp = 0;
        for (var i in this.cart) {
          temp += this.cart[i].item.Price * this.cart[i].amount;
        }
        this.sum = temp;
      }
    },
    checkout: function() {
      var packedCart = { cartItems: this.cart, user: this.user };
      this.$http
        .post("/checkout", packedCart)
        .then(function(response) {
          this.emptyCart();
        })
        .catch(function(response) {
          this.bookingError = response.data;
        });
    },
    emptyCart: function() {
      this.sum = 0;
      this.user = {};
      this.items = [];
      this.deselectUser();
      this.deselectItems();
    },
    deselectUser: function() {
      this.$userEventBus.$emit("deselectUserToBus");
    },
    deselectItems: function() {
      this.$itemEventBus.$emit("deselectItemsToBus");
    }
  }
};
</script>

