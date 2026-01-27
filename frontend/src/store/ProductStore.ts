import {defineStore} from 'pinia'
import {Product, createProduct} from '../composables/product'

export const useProductStore = defineStore('product', {
    state: () => {
        return {
            products: [] as Product[],
            selected: {} as Product
        }
    },
    actions: {
        getByCategory(categoryId: number, selectedAccountCategorys: number[]): Product[] {
            /*
            let visibleProducts = this.products.filter((prod) => {
                // Union of all Products / User Categorys
                // return prod.visibility.some(cat => selectedAccountCategorys.includes(cat))

                // Intersection
                let intersection = selectedAccountCategorys.filter(aCat => prod.visibility.includes(aCat));
                return JSON.stringify(intersection.sort()) == JSON.stringify(selectedAccountCategorys.sort());
            });

            const currentCategoryProducts = visibleProducts.filter((prod) => prod.category.id == categoryId);

            if (categoryId == 0) {
                return visibleProducts;
            }
            return currentCategoryProducts;
             */
            // TODO use class ProductVisibility (TBD)
            return this.products
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
        byId(id: number | undefined): Product | undefined {
            let foundProd = this.products.find((prod) => prod.id == id);
            return foundProd
        },
        patchProducts(newProducts: Product[]) {
            this.$patch(state => {
                newProducts.forEach(newProd => {
                    const newProdObject = createProduct(newProd);
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