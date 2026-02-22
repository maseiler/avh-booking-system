import {defineStore} from 'pinia'
import { ProductVisibility} from '../composables/productVisibility'
import { useSocketStore } from './socketStore';

export const useProductVisibilityStore = defineStore('product_visibility', {
    state: () => {
        return {
            visibilities: [] as ProductVisibility[]
        }
    },
    getters: {
        all(): ProductVisibility[] {
            return this.visibilities;
        }
    },
    actions: {
        byId(id: number | undefined): ProductVisibility | undefined {
            return this.visibilities.find((vis) => vis.id == id)
        },
        byProductId(id: number){
            let foundVis = this.visibilities.filter((vis) => {
                return vis.product == id;
            })
            return foundVis;
        },
        categoryIsVisible(categoryId: number, productId: number): boolean{
            let returnValue = false;
            this.byProductId(productId).forEach((vis) => {
                if(vis.category == categoryId) {
                    returnValue = true;
                }
            })
            return returnValue;
        },
        toggleCategoryVisibility(categoryId: number, productId: number){
            let productVis = this.byProductId(productId);
            let categoryVis = productVis.filter((vis) => { return vis.category == categoryId});
            if(categoryVis.length > 0) {
                // this.visibilities = this.visibilities.filter((vis) => { return vis != categoryVis[0]});
                useSocketStore().removeVisibility(categoryVis[0].id);
                return
            }
            let newVis = {'category': categoryId, 'product': productId, 'location': 1} as ProductVisibility;
            useSocketStore().addVisibility(newVis);
        },
        patchVisibilities(newVisibilities: ProductVisibility[]) {
            this.$patch(state => {
                newVisibilities.forEach(newVis => {
                    const newVisObj = new ProductVisibility(newVis);
                    const existing = state.visibilities.find(a => a.id === newVisObj.id);
                    if (existing) {
                        Object.assign(existing, newVisObj);
                    } else {
                        state.visibilities.push(newVisObj);
                    }
                })
            })
        }
    }
})
