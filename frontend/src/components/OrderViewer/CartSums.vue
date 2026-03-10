<template>
<p class="cartSum">
  Summe:
  <span>{{ $n(cart$.getTotals[0] / 100, 'currency') }}</span>
</p>
<p class="cartTax">
  Davon Steuer:
  <span>{{ $n(cart$.getTotals[1] / 100, 'currency') }}</span>
</p>
</template>

<style scoped>
.cartSum{
  font-size:1.3rem;
  text-align: right;
  color: hsl(var(--bulma-text-h), var(--bulma-text-s), var(--bulma-text-strong-l));
  span{
    font-weight:700;
    /* color: hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l)); */
  }
}
.cartTax{
  font-size:.8rem;
  text-align: right;
  /* opacity: 0.65; */
  span{
    font-weight:600;
  }
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