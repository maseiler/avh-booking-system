import {defineStore} from 'pinia'
import {ProductGroup} from '../composables/productGroup'

export const useProductGroupStore = defineStore('productGroup', {
    state: () => {
        return {
            productGroups: [] as ProductGroup[]
        }
    },
    getters: {
        all(): ProductGroup[] {
            return this.productGroups;
        }
    },
    actions: {
        byId(id: number | undefined): ProductGroup | undefined {
            return this.productGroups.find((productGroup) => productGroup.id == id)
        },
        removeById(id: number) {
            this.$patch(state => {
                state.productGroups = state.productGroups.filter(g => g.id !== id)
            })
        },
        patchProductGroups(newGroups: ProductGroup[]) {
            this.$patch(state => {
                newGroups.forEach(newGroup => {
                    const newGroupObject = new ProductGroup(newGroup);
                    const existing = state.productGroups.find(a => a.id === newGroupObject.id);
                    if (existing) {
                        Object.assign(existing, newGroupObject);
                    } else {
                        state.productGroups.push(newGroupObject);
                    }
                })
            })
        }
    }
})
