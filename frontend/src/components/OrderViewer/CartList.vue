<template>
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

  <CartSums />
  <CartControl @cancelOrder="$emit('cancelOrder')"/>
  <!-- Component OrderControls -->
  
</div>
</template>

<style scoped>
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
</style>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCartStore } from '../../store/CartStore';
import CartControl from './CartControl.vue';
import CartProduct from './CartProduct.vue';
import CartSums from './CartSums.vue';

export default {
  data() {
    return {
      account$: useAccountStore(),
      cart$: useCartStore(),
    }
  },
  components:{
    CartProduct,
    CartSums,
    CartControl
  },
  emits: {
    cancelOrder: null
  }
}
</script>