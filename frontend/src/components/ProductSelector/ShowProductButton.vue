<template>
  <div class="accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">
    <div class="dictionary" v-for="(dict, key) of productsInOrder" :key="key">
      <span class="title is-1">{{ key }}</span>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <div v-for="product in dict">
          <button class="button is-fullwidth" title="select product" @click="cart$.addToCart(product)">
            {{ product.name }} {{ product.size }} {{ product.unit }}
            <span class="cartHint" v-if="cart$.productCartQuantity(product) != -1">{{ cart$.productCartQuantity(product) }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.accountList{
  height: var(--_height);
  overflow-y: scroll;
}
.dictionary{
  display:grid;
  grid-template-columns: 10% 90%;
  margin-bottom:1em;
  .title{
    justify-self: center;
    margin-bottom:0;
  }
}

.button:has(span){
  padding-right:1.8em;
}
.cartHint{
  position:absolute;
  right:.25em;
  top:50%;
  transform: translateY(-50%);
  font-size:.8em;
  background-color:rgba(255,255,255,.3);
  border-radius: 100vh;
  width:1.6em;
  aspect-ratio: 1;
}

@media screen and (min-width: 768px) {
  .button:has(span){
    padding: calc(var(--bulma-button-padding-vertical) - var(--bulma-button-border-width)) calc(var(--bulma-button-padding-horizontal) - var(--bulma-button-border-width));
  }
  .cartHint{
    display:none;
  }
}
</style>

<script lang="ts">
import { Product } from '../../composables/product';
import { useProductStore } from '../../store/ProductStore';
import type { PropType } from 'vue';
import { useResizeObserver } from '@vueuse/core';
import { useCartStore } from '../../store/CartStore';

export default {
  data(){
    return {
      product$: useProductStore(),
      height: 0,
      resizeElement: {} as HTMLElement,
      cart$: useCartStore(),
    }
  },
  props: {
    products: {
      type: Array as PropType<Product[]>
    }
  },
  computed: {
    productsInOrder() {
      var dict: {[key: string]: Product[]} = {};
      this.products?.forEach(prod => {
      var char = prod.name[0].toUpperCase();
      var charCode = char.charCodeAt(0);
      if (charCode >= 65 && charCode <= 90) { // A-Z
      } else if (charCode >= 48 && charCode <= 57) { // 0-9
        char = "#";
      } else {
        char = "?";
      }
      if (dict[char] === undefined) {
        dict[char] = [prod]
      } else {
        dict[char].push(prod);
      }
    })
    return dict;
    }
  },
  methods: {
    onResize(){
      let y = window.innerHeight;
      let _y = this.resizeElement.getBoundingClientRect().top;
      let dy = y - _y;
      this.height = dy -15 ;
    }
  },
  mounted() {
    this.resizeElement = this.$refs.resizeBox as HTMLElement
    // Watch for resizing
    useResizeObserver(this.resizeElement, () => {
      this.onResize();
    })
  },
}
</script>