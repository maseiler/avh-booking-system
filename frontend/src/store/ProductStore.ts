import {defineStore} from 'pinia'
import {Product} from '../composables/product'
import { useProductVisibilityStore } from './ProductVisibilityStore'
import { useAccountStore } from './AccountStore'

export const useProductStore = defineStore('product', {
    state: () => {
        return {
            products: [] as Product[],
            selected: {} as Product
        }
    },
    actions: {
        getByCategory(categoryId: number, selectedAccountCategorys: number[]): Product[] {
            const currentCategoryProducts = this.products.filter((prod) => {
                // ToDo Hide Products that are not available at this location
                if (categoryId == 0) {return true}
                return prod.category == categoryId 
            });

            const visibleProducts = currentCategoryProducts.filter((prod) => {
                if (useAccountStore().selected.length == 0) {return true}
                let visibilities = useProductVisibilityStore().byProductId(prod.id)
                let retVal = false;
                visibilities.forEach((visi) => {
                    if (selectedAccountCategorys.includes(visi.category)) { 
                        retVal = true;
                        return
                    } 
                })
                return retVal;
            })

            return visibleProducts.sort((a, b) => {
                return a.name.localeCompare(b.name);
            });
        },
        getBySearchAndCategory(searchString: string, categoryId: number, selectedAccountCategorys: number[]): Product[] {
            let search = searchString.toLowerCase();
            let byCategory = this.getByCategory(categoryId, selectedAccountCategorys);
            let searchResults = byCategory.filter((prod) => prod.name.toLowerCase().includes(search));
            return searchResults;
        },
        select(p: Product) {
            this.selected = p;
        },
        byId(id: number | undefined): Product | undefined{
            let foundProd = this.products.find((prod) => prod.id == id);
            return foundProd
        },
        patchProducts(newProducts: Product[]) {
            this.$patch(state => {
                newProducts.forEach(newProd => {
                    const newProdObject = new Product(newProd);
                    const existing = state.products.find(a => a.id === newProdObject.id);
                    if (existing) {
                        Object.assign(existing, newProdObject);
                    } else {
                        state.products.push(newProdObject);
                    }
                })
            })
        }
    }
})