<template>
  <tr class="cartProduct">
    <td class="cell productQuantity">
      <div class="field has-addons" v-if="allowEdit">
        <p class="control">
          <button class="button" @click="cart$.removeFromCart(content.product)"><icon :icon="['fas', 'trash']" /></button>
        </p>
        <p class="control has-icons-left has-icons-right">
          <input v-model="content.quantity" type="number" class="input"></input>
          <span class="icon is-left" @click="reduceQuant()"><icon :icon="['fas', 'circle-minus']" /></span>
          <span class="icon is-right" @click="content.quantity ++"><icon :icon="['fas', 'circle-plus']" /></span>
        </p>
      </div>
      <div class="field input" v-if="!allowEdit">{{ content.quantity }}</div>
    </td>
    <td class="cell productName"><span>{{ content.product.name }} ({{ content.product.size }} {{ content.product.getUnit().name }})</span></td>
    <td class="cell productTax has-text-right"><span>{{ content.tax }}%</span></td>
    <td class="cell productPrice has-text-right"><span>{{ $n(content.price / 100, 'currency') }}</span></td>
    <td class="cell productAmount has-text-right"><span>{{ $n(content.price * content.quantity / 100, 'currency') }}</span></td>
  </tr>
</template>

<style scoped>
.cartProduct{
  margin-bottom:0;
  .grid{
    height:3em;
    gap:0;
    align-items: center;
  }
  td{
    vertical-align: middle;
  }
}

.productQuantity{
  .button{
    height:100%;
  }
  .icon{
    cursor:pointer;
    pointer-events: all;
  }
  input{
    text-align: center;
    -moz-appearance: textfield;
    max-width:15ch;
    min-width:12ch;

    &::-webkit-outer-spin-button,
    &::-webkit-inner-spin-button{
      -webkit-appearance: none;
      margin: 0;
    }
  }
}
</style>

<script lang="ts">
import { type CartContent } from '../../composables/cartContent';
import { useCartStore } from '../../store/CartStore';

export default{
  data() {
   return {
    cart$: useCartStore()
   }   
  },
  props:{
    content: {} as CartContent,
    allowEdit: Boolean,
  },
  methods: {
    reduceQuant(){
      this.content.quantity --;
      if(this.content.quantity == 0){
        this.cart$.removeFromCart(this.content.product);
      }
    },
  }
}
</script>