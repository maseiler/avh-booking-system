<template>
  <div class="accountList" :style="`--_height:${height}px;`" ref="resizeBox" @mouseenter="onResize" @scroll="onResize">

    <div v-if="selectedGroup" class="group-overlay">
      <div class="group-overlay-header">
        <button class="button is-small" @click="selectedGroup = null">
          <span class="icon is-small"><icon :icon="['fas', 'arrow-left']" /></span>
          <span>Zurück</span>
        </button>
        <span class="group-overlay-title">{{ selectedGroup.name }}</span>
        <button class="delete is-medium" @click="selectedGroup = null" />
      </div>
      <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-align-content-flex-start is-gap-1">
        <Button
          v-for="product in productsForSelectedGroup"
          :key="product.id"
          @click="cart$.addToCart(product)"
          title="select product">
            {{ product.name }} {{ product.size }} {{ product.getUnit()?.name }}
            <span class="cartHint" v-if="cart$.productCartQuantity(product) != -1">{{ cart$.productCartQuantity(product) }}</span>
        </Button>
      </div>
    </div>

    <div class="dictionary" v-for="(dict, key) of processedProducts.dict" :key="key">
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
            @click="selectedGroup = product"
            title="select product group">
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
  position: relative;
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

.group-overlay {
  position: absolute;
  inset: 0;
  background-color: var(--bulma-scheme-main);
  overflow-y: auto;
  z-index: 10;
  padding: .5rem;
}
.group-overlay-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: .5rem;
  border-bottom: 1px solid var(--bulma-border);
  margin-bottom: .5rem;
}
.group-overlay-title {
  font-weight: 600;
  font-size: 1.1em;
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
      selectedGroup: null as ProductGroup | null,
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
    processedProducts() {
      const dict: {[key: string]: (Product | ProductGroup)[]} = {};
      const groupProductsMap: {[key: number]: Product[]} = {};

      this.products?.forEach(prod => {
        if (prod.productGroup != 0) {
          if (!groupProductsMap[prod.productGroup]) {
            groupProductsMap[prod.productGroup] = [];
          }
          groupProductsMap[prod.productGroup].push(prod);
        }

        let char = prod.name[0].toUpperCase();
        let charCode = char.charCodeAt(0);

        if (prod.productGroup != 0) {
          const pG = prod.getGroup();
          char = pG.name[0].toUpperCase();
          charCode = char.charCodeAt(0);
        }

        if (charCode >= 65 && charCode <= 90) { // A-Z
        } else if (charCode >= 48 && charCode <= 57) { // 0-9
          char = "#";
        } else {
          char = "?";
        }

        const toAdd: Product | ProductGroup = prod.productGroup != 0 ? prod.getGroup() : prod;

        if (dict[char] === undefined) {
          dict[char] = [toAdd];
        } else if (!dict[char].includes(toAdd)) {
          dict[char].push(toAdd);
        }
      });

      return { dict, groupProductsMap };
    },
    productsForSelectedGroup() {
      if (!this.selectedGroup) return [];
      return this.processedProducts.groupProductsMap[this.selectedGroup.id ?? -1] ?? [];
    }
  },
  watch: {
    products() {
      // this.selectedGroup = null;
    }
  },
  methods: {
    onResize(){
      let y = window.innerHeight;
      let _y = this.resizeElement.getBoundingClientRect().top;
      let dy = y - _y;
      this.height = dy - 15;
    }
  },
  mounted() {
    this.resizeElement = this.$refs.resizeBox as HTMLElement;
    useResizeObserver(this.resizeElement, () => {
      this.onResize();
    });
  },
}
</script>