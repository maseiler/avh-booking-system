import {defineStore} from 'pinia'
import {createProductVisibility, type ProductVisibility} from '../composables/productVisibility'

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
        patchVisibilities(newVisibilities: ProductVisibility[]) {
            this.$patch(state => {
                newVisibilities.forEach(newVis => {
                    const newVisObj = createProductVisibility(newVis);
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
