<template>
<div class="buttons">
  <button class="button is-warning is-inverted is-outlined" @click="$emit('cancelOrder')" title="discard cart and unselect account">
    <span class="icon"><icon :icon="['fas', 'trash']" /></span>
    <span>Cancel Order</span>
  </button>

  <button class="button is-success is-inverted is-outlined" @click="checkoutOrder" v-if="!cart$.isOverdrawn">
    <span>Book now</span>
    <span class="icon"><icon :icon="['fas', 'beer']"/></span>
  </button>
  <button class="button is-inverted is-danger is-outlined" v-if="cart$.isOverdrawn">
    <span>Pay for this order</span>
    <span class="icon"><icon :icon="['fas', 'coins']" /></span>
  </button>
</div>
<div class="message is-danger hint" v-if="cart$.isOverdrawn">
  <div class="message-body">
    <icon :icon="['fas', 'warning']" />
    The Cart exceeds the allowance! <br>
    <router-link to="/payment">Top up the account</router-link> or pay for this order immediately with the button above.
  </div>
</div>
</template>

<style scoped>
.buttons{
  margin-top:.5rem;
  justify-content: space-between;
}
</style>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCartStore } from '../../store/CartStore';


export default{
  data() {
    return {
      account$: useAccountStore(),
      cart$: useCartStore()
    }
  },
  methods: {
    checkoutOrder(){
      console.warn("Not yet implemented");
      // ToDo: this
    }
  },
  emits: {
    cancelOrder: null
  }
}
</script>