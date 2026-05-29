<template>
  <div class="accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">
    <div class="dictionary" v-for="(dict, key) of productsInOrder" :key="key">
      <span class="title is-1">{{ key }}</span>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <template v-for="product in dict">
          <Button
          v-if="product.constructor.name == 'Product'"
          @click="cart$.addToCart(product)"
          title="select product">
            {{ product.name }} {{ product.size }} {{ product.getUnit()?.name }}
            <span class="cartHint" v-if="cart$.productCartQuantity(product) != -1">{{ cart$.productCartQuantity(product) }}</span>
          </Button>

          <Button
            v-if="product.constructor.name == 'ProductGroup'"
            title="select product">
              {{ product.name }}
          </Button>
        </template>
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
import { ProductGroup } from '../../composables/productGroup.ts';
import { useProductStore } from '../../store/ProductStore';
import type { PropType } from 'vue';
import { useResizeObserver } from '@vueuse/core';
import { useCartStore } from '../../store/CartStore';
import Button from '../../composables/elements/Button.vue';

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
  components: {
    Button
  },
  computed: {
    productsInOrder() {
      var dict: {[key: string]: Product[] | ProductGroup[]} = {};
      this.products?.forEach(prod => {
        var char = prod.name[0].toUpperCase();
        var charCode = char.charCodeAt(0);

        if(prod.productGroup != 0) { 
          // Use the Group Name instead if available
          let pG = prod.getGroup()
          char = pG.name[0].toUpperCase();
          charCode = char.charCodeAt(0);
        }

        if (charCode >= 65 && charCode <= 90) { // A-Z
        } else if (charCode >= 48 && charCode <= 57) { // 0-9
          char = "#";
        } else {
         char = "?";
        }

        let toAdd = prod;
        if(prod.productGroup != 0) {
          toAdd = prod.getGroup();
        }
        
        if (dict[char] === undefined) {
          dict[char] = [toAdd]
        } else if(!dict[char].includes(toAdd)) {
          dict[char].push(toAdd); 
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