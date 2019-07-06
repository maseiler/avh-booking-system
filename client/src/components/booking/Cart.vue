<template>
  <div>
    <table class="table is-fullwidth is-hoverable">
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
            <button class="button" @click="checkout">Checkout</button>
          </th>
        </tr>
      </tfoot>
      <tbody>
        <tr></tr>
        <tr v-for="i in cart" :key="i">
          <td>{{i.item.Name}} {{i.item.Size}}</td>
          <td>{{i.item.Price}} €</td>
          <td>
            <input
              class="input is-small has-text-black"
              type="text"
              style="width:4em;"
              v-model="i.amount"
              @change="updateCart(i.item)"
            >
            &nbsp;
            <button class="button is-small" @click="deleteItem(i)">
              <font-awesome-icon icon="trash"/>
            </button>
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
        if (array[i].item.ID === obj.ID) {
          return true;
        }
      }
      return false;
    },
    increment: function(array, item) {
      for (var i in array) {
        if (array[i].item.ID === item.ID) {
          array[i].amount = array[i].amount + 1;
          return;
        }
      }
    },
    updateCart: function(item) {
      for (var i in this.cart) {
        if (this.cart[i].item === item) {
          var itemCount = 0;
          for (var j in this.items) {
            if (this.items[j] === item) {
              itemCount++;
            }
          }
          if (itemCount <= this.cart[i].amount) {
            while (itemCount < this.cart[i].amount) {
              this.items.push(item);
              itemCount++;
            }
          } else {
            while (itemCount > this.cart[i].amount) {
              var index = this.items.indexOf(item);
              if (index != -1) {
                this.items.splice(index, 1);
              }
              itemCount--;
            }
          }
          break;
        }
      }
    },
    updateSum: function() {
      var temp = 0;
      for (var i in this.cart) {
        temp += this.cart[i].item.Price * this.cart[i].amount;
      }
      temp = temp.toFixed(2);
      this.sum = temp;
    },
    checkout: async function() {
      var packedCart = { cartItems: this.cart, user: this.user };
      await this.$http
        .post("checkout", packedCart)
        .then(function(response) {
          var message = "".concat(
            "Checkout from ",
            this.displayUserName(this.user)
          );
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch(function(response) {
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
      await this.emptyCart();
      await this.$router.go();
    },
    deleteItem: function(item) {
      var index = this.cart.indexOf(item);
      if (index != -1) {
        this.cart[index].amount = 0;
        this.updateCart(item.item);
      }
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

