<template>
  <div class="order message is-info">

    <div class="message-header" @click="toggleOrderDetails">
      <span>
        <icon :class="showOrderDetails ? 'showDetails' : ''" class="order-icon details" :icon="['fas', 'arrow-up-short-wide']" />
        <icon :class="showOrderDetails ? 'showDetails' : ''" class="order-icon" :icon="['fas', 'arrow-down-short-wide']" />
        New Order
      </span>
      <button v-if="account$.selected.length > 0" class="delete" @click="cancelOrder" title="discard cart and unselect account"></button>
    </div>

    <div :class="showOrderDetails ? 'showDetails' : ''" class="order-ripped-teaser"> </div>
    <div :class="showOrderDetails ? 'showDetails' : ''" class="message-body fixed-grid has-3-cols">
      <div class="selectedAccounts">
        <div class="tag" v-for="account in account$.selected">
          {{ account.getFullName() }}
          <button class="delete is-small" @click="unselectAccount(account)"></button>
        </div>
      </div>

      <!-- Component Cart List -->
      <p v-if="account$.selected.length == 0">Please select an Accout first</p>
      <div class="cartList" v-if="account$.selected.length != 0">
        <div class="table-container">
          <table class="table is-striped">
            <thead><tr>
              <th>Quantity</th>
              <th>Product</th>
              <th class="has-text-right">Tax</th>
              <th class="has-text-right">Price</th>
              <th class="has-text-right">Amount</th>
            </tr></thead>
            <tbody>
              <CartProduct v-for="content in cart$.cartContents" :content="content"/>
            </tbody>
          </table>
        </div>
        <div class="dblhr"></div>
        <!-- Component CartSums -->
        <p class="cartSum">
          Summe:
          <span>{{ $n(cart$.getTotals[0] / 100, 'currency', 'de-DE') }}</span>
        </p>
        <p class="cartTax">
          Davon Steuer:
          <span>{{ $n(cart$.getTotals[1] / 100, 'currency', 'de-DE') }}</span>
        </p>

        <!-- Component OrderControls -->
        <div class="buttons">
          <button class="button is-warning is-inverted is-outlined" @click="cancelOrder">
            <span class="icon"><icon :icon="['fas', 'trash']" /></span>
            <span>Cancel Order</span>
          </button>
          <button class="button is-success is-inverted is-outlined" @click="checkoutOrder">
            <span>Book now</span>
            <span class="icon"><icon :icon="['fas', 'beer']"/></span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.selectedAccounts{
  margin-bottom:.5rem;
}
.buttons{
  margin-top:.5rem;
  justify-content: end;
}
.cartSum,
.cartTax{
  font-size:1.2rem;
  text-align: right;
  span{
    font-weight:600;
  }
}
.cartTax{
  font-size:.8rem;
}
.table-container{
  border-radius: var(--bulma-control-radius);
  margin-bottom:.5rem;
  .table{
    width:100%;
  }
}
.dblhr{
  --_height:6px;
  height:var(--_height);
  border-top:calc(var(--_height) / 3) solid var(--bulma-scheme-main);;
  border-bottom:calc(var(--_height) / 3) solid var(--bulma-scheme-main);;
}
.order{
  position:relative;
}
.order-icon{
  margin-right:.5em;
  &.details{
    display:none;
  }
}
.showDetails.order-icon{
  display:none;
  &.details{
    display:inline-block;
  }
}
.message-header{
  cursor:pointer;
}
.message-body{
  display:none;
  position:absolute;
  width:100%;
  z-index:20;
}
.order-ripped-teaser{
  --_bg-color:hsl(var(--bulma-message-h),var(--bulma-message-s),var(--bulma-message-background-l));
  background-color:var(--_bg-color);
  width:100%;
  height:0.5rem;
  position:relative;
}
.order-ripped-teaser::after {
  content: '';
  position: absolute;
  right: 0;
  left: -0%;
  top: 100%;
  z-index: 10;
  display: block;
  height: 20px;
  background-size: 20px 100%;
  background-image: linear-gradient(135deg, var(--_bg-color) 25%, transparent 25%), linear-gradient(225deg, var(--_bg-color) 25%, transparent 25%);
  background-position: 10% 0;
}
.showDetails{
  &.order-ripped-teaser{
    display: none;
  }
  &.message-body{
    display:block;
  }
}
@media screen and (min-width: 768px) {
  .order-icon,
  .showDetails.order-icon.details{
    display:none;
  }
  .message-header{
    cursor:auto;
    pointer-events:none;
  }
  .order-ripped-teaser{
    display:none;
  }
  .message-body{
    display:block;
  }
}
</style>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import type { Account } from '../../composables/account';
import { useCartStore } from '../../store/CartStore';
import CartProduct from './CartProduct.vue';

export default {
  data() {
    return {
      account$: useAccountStore(),
      showOrderDetails: false,
      cart$: useCartStore(),
    }
  },
  components: {
    CartProduct
  },
  methods: {
    unselectAccount(account: Account){
      this.account$.selectSubstract(account);
    },
    toggleOrderDetails(){
      this.showOrderDetails = !this.showOrderDetails;
    },
    cancelOrder(){
      this.account$.unselect();
      this.cart$.cartContents = [];
    },
    checkoutOrder(){
      alert("new Order received");
      this.cancelOrder();
    }
  }
}
</script>