<template>
  <div class="fixed-grid has-7-cols cartProduct">
    <div class="grid">
      <span class="productName cell is-col-span-3 is-row-start-1">{{ content.product.name }} ({{ content.product.size }})</span>
      <input v-model="content.amount" type="number" class="input cell is-col-start-4 is-col-span-1 is-row-start-1 productAmount"></input>
      <span class="productTax cell is-col-start-5 is-col-span-1 is-row-start-1">{{ content.product.tax }}%</span>
      <span class="productPrice cell is-col-start-6 is-col-span-2 is-row-start-1">{{ $n(content.product.price / 100, 'currency', 'de-DE') }}</span>
    </div>
  </div>
</template>

<style scoped>
.productTax,
.productPrice {
  justify-self: end;
}
.cartProduct{
  margin-bottom:0;
  .grid{
    height:3em;
    gap:0;
    align-items: center;
  }
}
.productAmount{
  -moz-appearance: textfield;
}
.productAmount::-webkit-outer-spin-button,
.productAmount::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>

<script lang="ts">
import { type CartContent } from '../../composables/cartContent';

export default{
  data() {
   return {

   }   
  },
  props:{
    content: {} as CartContent
  }
}
</script>