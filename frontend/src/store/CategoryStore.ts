import { defineStore } from 'pinia'
import { Category, generateTestData } from '../composables/category'
import { CategoryType } from '../composables/category'

export const useCategoryStore = defineStore('category', {
  state: () => {
    return {
      categorys: [] as Category[]
    }
  },
  getters:{
    accountCategorys(state){
      return state.categorys.filter((cat) => cat.type == CategoryType.ACCOUNT)
    },
    productCategorys(state){
      return state.categorys.filter((cat) => cat.type == CategoryType.PRODUCT)
    },
    byId(state){
      return (id: number) => state.categorys.find((cat) => cat.id == id)
    },
  },
  actions: {
    generateTestData(){
      if (this.categorys.length == 0) {
        this.categorys.push(...generateTestData());
      }
    }
  }
})