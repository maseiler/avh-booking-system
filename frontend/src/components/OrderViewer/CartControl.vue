<template>
<div class="buttons">
  <button class="button is-warning is-inverted is-outlined" @click="$emit('cancelOrder')" title="discard cart and unselect account">
    <span class="icon"><icon :icon="['fas', 'trash']" /></span>
    <span>Cancel Order</span>
  </button>
  <button class="button is-success is-inverted is-outlined" @click="checkoutOrder">
    <span>Book now</span>
    <span class="icon"><icon :icon="['fas', 'beer']"/></span>
  </button>
</div>
</template>

<style scoped>
.buttons{
  margin-top:.5rem;
  justify-content: end;
}
</style>

<script lang="ts">

export default{


  data() {
    return {

    }
  },
  emits: {
    cancelOrder: null
  }
}
</script>