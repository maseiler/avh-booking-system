import {defineStore} from 'pinia'
import {Product} from '../composables/product'
import { useProductVisibilityStore } from './ProductVisibilityStore'
import { useAccountStore } from './AccountStore'
import { useCategoryStore } from './CategoryStore'

export const useProductStore = defineStore('product', {
    state: () => {
        return {
            products: [] as Product[],
            selected: {} as Product
        }
    },
    actions: {
        getByCategory(categoryId: number, selectedAccountCategorys: number[], all?: boolean): Product[] {
            const currentCategoryProducts = this.products.filter((prod) => {
                // Hide products whose category is disabled (unless all=true, e.g. in settings)
                if (!all && !useCategoryStore().byId(prod.category)?.enabled) { return false }
                // ToDo Hide Products that are not available at this location
                if (categoryId == 0) {return true}
                return prod.category == categoryId
            });

            const visibleProducts = currentCategoryProducts.filter((prod) => {
                if (useAccountStore().selected.length == 0) {return true}
                // ToDo Show all Products for Accounts that have the corresponding Flag in their Account Options
                let visibilities = useProductVisibilityStore().byProductId(prod.id)
                let retVal = false;
                visibilities.forEach((visi) => {
                    if (selectedAccountCategorys.includes(visi.categoryId)) { 
                        retVal = true;
                        return
                    } 
                })
                return retVal;
            })

            // if(all) {
            //     // Gib alle Produkte aus, wenn die all-Flag gesetzt ist
            //     return visibleProducts.sort((a, b) => {
            //         return a.name.localeCompare(b.name)
            //     });
            // }

            // const nonGroupedProducts = visibleProducts.filter((prod) => prod.productGroup == 0);

            return visibleProducts.sort((a, b) => {
                return a.name.localeCompare(b.name);
            });
        },
        getBySearchAndCategory(searchString: string, categoryId: number, selectedAccountCategorys: number[], all?: boolean): Product[] {
            let search = searchString.toLowerCase();
            let byCategory = this.getByCategory(categoryId, selectedAccountCategorys, all);
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