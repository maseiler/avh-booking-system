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
  },
  actions: {
    generateTestData(){
      if (this.categorys.length == 0) {
        this.categorys.push(...generateTestData());
      }
    },
    byId(id: number | undefined): Category | undefined {
      let foundCat = this.categorys.find((cat) => cat.id == id);
      return foundCat
    }
  }
})