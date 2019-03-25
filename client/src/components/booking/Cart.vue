<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Item</th>
          <th>Price</th>
          <th>Amount</th>
          <th></th>
        </tr>
      </thead>
      <tfoot>
        <tr>
          <th></th>
          <th>{{sum}} €</th>
          <th></th>
          <th>
            <button class="button is-small">Checkout</button>
          </th>
        </tr>
      </tfoot>
      <tbody>
        <tr v-for="i in cart" :key="i">
          <td>{{i.item.Name}} {{i.item.Size}}</td>
          <td>{{i.item.Price}} €</td>
          <td>{{i.amount}}</td>
          <td>
            <button class="button is-small" @click="increment(cart,i.item)">+</button>
            <button class="button is-small">-</button>
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
  data: function() {
    return {
      sum: 0
    };
  },
  computed: {
    cart: function() {
      return this.buildCart();
    }
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
          //   console.log(Number(newItem.item.Price));
          //   this.sum += Number(newItem.item.Price);
        } else {
          this.increment(temp, item);
        }
      });
      return temp;
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
      console.log(array);
      for (var i in array) {
        if (array[i].item.ItemID === item.ItemID) {
          array[i].amount = array[i].amount + 1;
          //   console.log(Number(array[i].item.Price));
          //   this.sum += Number(array[i].item.Price);
          console.log("inc");
          return;
        }
      }
    }
  }
};
</script>

