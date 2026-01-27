import {defineStore} from 'pinia'
import {Unit} from '../composables/unit'

export const useUnitStore = defineStore('unit', {
    state: () => {
        return {
            units: [] as Unit[]
        }
    },
    getters: {
        all(): Unit[] {
            return this.units;
        }
    },
    actions: {
        byId(id: number | undefined): Unit | undefined {
            return this.units.find((unit) => unit.id == id)
        },
        patchUnits(newUnits: Unit[]) {
            this.$patch(state => {
                newUnits.forEach(newUnit => {
                    const newUnitObj = new Unit(newUnit);
                    const existing = state.units.find(a => a.id === newUnitObj.id);
                    if (existing) {
                        Object.assign(existing, newUnitObj);
                    } else {
                        state.units.push(newUnitObj);
                    }
                })
            })
        }
    }
})
