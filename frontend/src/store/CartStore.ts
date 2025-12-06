import { defineStore } from 'pinia'
import { type CartContent } from '../composables/cartContent'
import { Product } from '../composables/product'
import { useAccountStore } from './AccountStore'

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
        let subTotal = cont.product.price * cont.quantity;
        let subTax = subTotal*(cont.product.tax / 100);
        total += subTotal;
        tax += subTax;
      })
      return [total, tax];
    },
    isOverdrawn(): Boolean{
      const account$ = useAccountStore();
      if (account$.selected.length > 1) {return false}
      return account$.selected[0].balance - this.getTotals[0] < account$.selected[0].maxDebt
    }
  },
  actions: {
    addToCart(product: Product){
      let alreadySelected = this.cartContents.filter((cont) => cont.product == product);
      if (alreadySelected.length > 0){
        alreadySelected[0].quantity ++;
        return;
      }
      this.cartContents.push({product: product, quantity: 1, productPrice: product.price} as CartContent);
    },
    removeFromCart(product: Product){
      this.cartContents = this.cartContents.filter((cont) => {
        return cont.product != product
      })
    },
    productCartQuantity(product: Product): number{
      let result = -1;
      this.cartContents.forEach((cont) => {
        if (cont.product != product){
          return;
        }
        result = cont.quantity;
      })
      return result;
    }
  }
})