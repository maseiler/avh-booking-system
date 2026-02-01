import { defineStore } from 'pinia'
import { type CartContent } from '../composables/cartContent'
import { Product } from '../composables/product'
import { useAccountStore } from './AccountStore'
import type { BookingTotals } from '../composables/booking'

export const useCartStore = defineStore('cart', {
  state: () => {
    return {
      cartContents: [] as CartContent[]
    }
  },
  getters:{
    getTotals(): BookingTotals{
      let total = 0;
      let tax = 0;
      this.cartContents.forEach((cont) => {
        let subTotal = cont.product.price * cont.quantity;
        let subTax = subTotal*(cont.product.vat.rate / 100);
        total += subTotal;
        tax += subTax;
      })
      return [total, tax] as BookingTotals;
    },
    isOverdrawn(): Boolean{
      const account$ = useAccountStore();
      if (account$.selected.length > 1) {return false}
      return account$.selected[0].balance - this.getTotals[0] < (account$.selected[0].maxDebt * -1)
    }
  },
  actions: {
    addToCart(product: Product){
      let alreadySelected = this.cartContents.filter((cont) => cont.product == product);
      if (alreadySelected.length > 0){
        alreadySelected[0].quantity ++;
        return;
      }
      let newCartContent = {product: product, quantity: 1, price: product.price, tax: product.vat.rate} as CartContent;
      this.cartContents.push(newCartContent);
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