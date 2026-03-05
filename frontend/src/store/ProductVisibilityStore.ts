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
                return vis.productId == id;
            })
            return foundVis;
        },
        categoryIsVisible(categoryId: number, productId: number): boolean{
            let returnValue = false;
            this.byProductId(productId).forEach((vis) => {
                if(vis.categoryId == categoryId) {
                    returnValue = true;
                }
            })
            return returnValue;
        },
        toggleCategoryVisibility(categoryId: number, productId: number){
            let productVis = this.byProductId(productId);
            let categoryVis = productVis.filter((vis) => { return vis.categoryId == categoryId});
            if(categoryVis.length > 0) {
                // this.visibilities = this.visibilities.filter((vis) => { return vis != categoryVis[0]});
                useSocketStore().removeVisibility(categoryVis[0].id);
                return
            }
            let newVis = {'categoryId': categoryId, 'productId': productId, 'locationId': 1} as ProductVisibility;
            useSocketStore().addVisibility(newVis);
        },
        patchVisibilities(newVisibilities: ProductVisibility[]) {
            this.$patch(state => {
                state.visibilities = [];
                newVisibilities.forEach(newVis => {
                    const newVisObj = new ProductVisibility(newVis);
                    state.visibilities.push(newVisObj);
                })
            })
        }
    }
})
