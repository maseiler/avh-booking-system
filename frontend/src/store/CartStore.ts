import { defineStore } from 'pinia'
import { type CartContent } from '../composables/cartContent'
import { Product } from '../composables/product'

export const useCartStore = defineStore('cart', {
  state: () => {
    return {
      cartContents: [] as CartContent[]
    }
  },
  getters:{
    getTotals(){
      let total = 0;
      let tax = 0;
      this.cartContents.forEach((cont) => {
        let subTotal = cont.product.price * cont.amount;
        let subTax = subTotal*cont.product.tax;
        total += subTotal;
        tax += subTax;
      })
      return [total, tax];
    }
  },
  actions: {
    addToCart(product: Product){
      let alreadySelected = this.cartContents.filter((cont) => cont.product == product);
      if (alreadySelected.length > 0){
        alreadySelected[0].amount ++;
        return;
      }
      this.cartContents.push({product: product, amount: 1} as CartContent);
    },
    removeFromCart(product: Product){
      this.cartContents = this.cartContents.filter((cont) => {
        return cont.product != product
      })
    }
  }
})