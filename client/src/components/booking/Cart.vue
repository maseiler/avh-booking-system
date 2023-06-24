<template>
  <div>
    <table class="table is-fullwidth is-hoverable">
      <thead>
        <tr>
          <th>{{ $t("item.item") }}</th>
          <th>{{ $t("item.price") }}</th>
          <th>{{ $t("cart.amount") }}</th>
        </tr>
      </thead>
      <tfoot>
        <tr>
          <th></th>
          <th>{{ $n(sum, "currency", "de-DE") }}</th>
          <th>
            <button class="button is-success" @click="checkout">
              {{ $t("cart.checkout") }}
            </button>
          </th>
        </tr>
      </tfoot>
      <tbody>
        <tr></tr>
        <tr v-for="i in cart" :key="i">
          <td v-if="i.item.Type === 'food'">{{ i.item.Name }}</td>
          <td v-else>{{ i.item.Name }} {{ i.item.Size }}{{ i.item.Unit }}</td>
          <td>{{ $n(i.item.Price, "currency", "de-DE")}}</td>
          <td>
            <input
              class="input is-small has-text-black"
              type="text"
              style="width: 4em"
              v-model="i.amount"
              @change="updateCart(i.item)"
            />
            &nbsp;
            <button class="button is-small" @click="deleteItem(i)">
              <font-awesome-icon icon="trash" />
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  computed: {
    user() {
      return this.$store.state.selectedUser;
    },
    items() {
      return this.$store.state.selectedMultipleItems;
    },
  },
  watch: {
    items() {
      this.buildCart();
      this.updateSum();
    },
  },
  data() {
    return {
      cart: [],
      sum: 0,
      bookingError: "",
    };
  },
  methods: {
    buildCart() {
      var temp = [];
      this.items.forEach((item) => {
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
    contains(array, obj) {
      for (var i in array) {
        if (array[i].item.ID === obj.ID) {
          return true;
        }
      }
      return false;
    },
    increment(array, item) {
      for (var i in array) {
        if (array[i].item.ID === item.ID) {
          array[i].amount = array[i].amount + 1;
          return;
        }
      }
    },
    updateCart(item) {
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
    updateSum() {
      var temp = 0;
      for (var i in this.cart) {
        temp += this.cart[i].item.Price * this.cart[i].amount;
      }
      temp = temp.toFixed(2);
      this.sum = temp;
    },
    async checkout() {
      var packedCart = { cartItems: this.cart, user: this.user };
      await this.$http
        .post("checkout", packedCart)
        .then(() => {
          let message = `${this.$t("messages.success.checkoutFrom")}: ${this.displayUserName(this.user)}`;
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch((response) => {
          //ToDo: Failure Message internationalisation
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
      await this.emptyCart();
      await this.refresh();
      this.$root.$emit("set-item-tab-to-favorites", true);
      this.$root.$emit("set-user-tab-to-all", true);
    },
    deleteItem(item) {
      var index = this.cart.indexOf(item);
      if (index != -1) {
        this.cart[index].amount = 0;
        this.updateCart(item.item);
      }
    },
    emptyCart() {
      this.cart = [];
      this.sum = 0;
      this.$store.commit("selectUser", {});
      this.$store.state.selectedMultipleItems = [];
    },
    refresh() {
      this.$store.commit("getUsers");
      this.$store.commit("getItems");
      this.$store.commit("getLastNBookEntries", 5);
    },
  },
  created() {
    this.buildCart();
    this.updateSum();
  },
};
</script>
