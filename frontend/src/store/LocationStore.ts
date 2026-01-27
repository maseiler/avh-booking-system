import {defineStore} from 'pinia'
import {Location} from '../composables/location'

export const useLocationStore = defineStore('location', {
    state: () => {
        return {
            locations: [] as Location[]
        }
    },
    getters: {
        all(): Location[] {
            return this.locations;
        }
    },
    actions: {
        byId(id: number | undefined): Location | undefined {
            return this.locations.find((location) => location.id == id)
        },
        patchLocations(newLocations: Location[]) {
            this.$patch(state => {
                newLocations.forEach(newLocation => {
                    const newLocationObj = new Location(newLocation);
                    const existing = state.locations.find(a => a.id === newLocationObj.id);
                    if (existing) {
                        Object.assign(existing, newLocationObj);
                    } else {
                        state.locations.push(newLocationObj);
                    }
                })
            })
        }
    }
})
