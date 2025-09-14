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
      <div class="grid">
        <div class="cell is-col-span-1">
          <div class="tag" v-for="account in account$.selected">
            {{ account.getFullName() }}
            <button class="delete is-small" @click="unselectAccount(account)"></button>
          </div>
        </div>
        <div class="cell is-col-span-2">
          <p v-if="account$.selected.length == 0">Please select an Accout first</p>
          <div v-if="account$.selected.length != 0">
            <CartProduct v-for="content in cart$.cartContents" :content="content"/>
            <!-- Component CartSums -->
            <p>Summe: {{ $n(cart$.getTotals[0] / 100, 'currency', 'de-DE') }} <br> Davon Steuer: {{ cart$.getTotals[1] }}</p>
            <!-- Component OrderControls -->
             <button class="button is-warning" @click="cancelOrder">Cancel Order</button>
             <button class="button is-success" @click="checkoutOrder">Book now</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
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