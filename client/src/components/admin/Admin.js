import UserSettings from "./UserSettings.vue"
import ItemSettings from "./ItemSettings.vue"
import OtherSettings from "./OtherSettings.vue"

export default {
  components: {
    UserSettings,
    ItemSettings,
    OtherSettings
  }, data: function () {
    return {
      showUserSettings: false,
      showItemSettings: false,
      showOtherSettings: false
    };
  },
  methods: {
    showSetting(setting) {
      switch (setting) {
        case 'userSetting':
          this.showUserSettings = true
          this.showItemSettings = false
          this.showOtherSettings = false
          break
        case 'itemSetting':
          this.showUserSettings = false
          this.showItemSettings = true
          this.showOtherSettings = false
          break
        case 'otherSetting':
          this.showUserSettings = false
          this.showItemSettings = false
          this.showOtherSettings = true
          break
        default:
          break
      }
    }
  }
}