import {defineStore} from 'pinia'
import {Vat} from '../composables/vat'

export const useVatStore = defineStore('vat', {
    state: () => {
        return {
            vats: [] as Vat[]
        }
    },
    getters: {
        all(): Vat[] {
            return this.vats;
        }
    },
    actions: {
        byId(id: number | undefined): Vat | undefined {
            return this.vats.find((vat) => vat.id == id)
        },
        getById(id: number | undefined): Vat | undefined {
            return this.vats.find((vat) => vat.id == id)
        },
        patchVats(newVats: Vat[]) {
            this.$patch(state => {
                newVats.forEach(newVat => {
                    const newVatObj = new Vat(newVat);
                    const existing = state.vats.find(a => a.id === newVatObj.id);
                    if (existing) {
                        Object.assign(existing, newVatObj);
                    } else {
                        state.vats.push(newVatObj);
                    }
                })
            })
        }
    }
})
