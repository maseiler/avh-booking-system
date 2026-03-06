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
  margin-bottom:.75rem;
  border: 1px solid hsl(var(--bulma-scheme-h), var(--bulma-scheme-s), var(--bulma-border-l));
  overflow: hidden;
  .table{
    width:100%;
    background-color: transparent;
    margin-bottom:0;
  }
}
.dblhr{
  height: 2px;
  background: linear-gradient(
    to right,
    transparent,
    hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l)) 20%,
    hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l)) 80%,
    transparent
  );
  opacity: 0.5;
  margin-bottom: .75rem;
  border: none;
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