<template>
<p class="cartSum">
  Summe:
  <span>{{ $n(cart$.getTotals[0] / 100, 'currency', 'de-DE') }}</span>
</p>
<p class="cartTax">
  Davon Steuer:
  <span>{{ $n(cart$.getTotals[1] / 100, 'currency', 'de-DE') }}</span>
</p>
</template>

<style scoped>
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
</style>

<script lang="ts">
import { useCartStore } from '../../store/CartStore';

export default {
  data() {
    return {
      cart$: useCartStore()
    }
  },
}
</script>