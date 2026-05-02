import { defineStore } from "pinia";
import { type Setting } from "../composables/setting";


export const useSettingStore = defineStore('setting', {
  state: () => {
      return {
        settings: [] as Setting[],
        clientId: "",
      }
  },
  getters: {
    needsSetup(): boolean{
      return (!this.hasClientId) || false;
    },
    hasClientId(): boolean{
      return this.clientId != "";
    }
  },
  actions: {
    get(key: string){
      let foundSetting = this.settings.filter((set) => set.key == key)[0];
      return foundSetting == undefined || foundSetting.value == "" ? -1 : foundSetting;
    },
    set(key: string, value: string | ArrayBuffer | null){
      let newSetting = {
        key: key,
        value: value
      } as Setting;
      // This is just temporary and needs to be changed to send the new Setting to the WS
      // After that, the WS should notify the Client of the change and the Change will trigger the
      // patchSettings Function automatically.
      // The WS does not yet have this behaviour, that is why this workaround is placed in the frontend.
      let setArr = [] as Setting[];
      setArr.push(newSetting);
      this.patchSettings(setArr);
    },
    patchSettings(newSettings: Setting[]){
      this.$patch(state => {
        newSettings.forEach(newSet => {
          // const newSetObj = new Setting(newSet);
          const existing = state.settings.find(a => a.key === newSet.key);
          if(existing){
            Object.assign(existing, newSet);
          } else {
            state.settings.push(newSet);
          }
        })
      })
    }
  },
  persist:{
    storage: localStorage,
    pick: ['clientId'],
  }
})