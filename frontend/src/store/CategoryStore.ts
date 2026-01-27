import {defineStore} from 'pinia'
import {Category, CategoryType, createCategory, generateTestData} from '../composables/category'

export const useCategoryStore = defineStore('category', {
    state: () => {
        return {
            categories: [] as Category[]
        }
    },
    getters: {
        accountCategories(state) {
            return state.categories.filter((cat) => cat.type == CategoryType.ACCOUNT)
        },
        productCategories(state) {
            return state.categories.filter((cat) => cat.type == CategoryType.PRODUCT)
        },
    },
    actions: {
        generateTestData() {
            if (this.categories.length == 0) {
                this.categories.push(...generateTestData());
            }
        },
        byId(id: number | undefined): Category | undefined {
            return this.categories.find((cat) => cat.id == id)
        },
        patchCategories(newCategories: Category[]) {
            // TODO? copy-pasta from AccountStore.ts
            this.$patch(state => {
                newCategories.forEach(newCat => {
                    const newCatObj = createCategory(newCat);
                    const existing = state.categories.find(a => a.id === newCatObj.id);
                    if (existing) {
                        Object.assign(existing, newCatObj);
                    } else {
                        state.categories.push(newCatObj);
                    }
                })
            })
        }
    }
})