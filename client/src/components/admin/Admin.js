import UserSettings from "./UserSettings.vue"
import ItemSettings from "./ItemSettings.vue"
import BookingSettings from "./BookingSettings.vue"
import OtherSettings from "./OtherSettings.vue"

export default {
  components: {
    UserSettings,
    ItemSettings,
    BookingSettings,
    OtherSettings
  },
  data: function () {
    return {
      users: [],
      items: [],
      lastBookings: [],
      showUserSettings: false,
      showItemSettings: false,
      showBookingSettings: false,
      showOtherSettings: false
    };
  },
  methods: {
    getUsers: function () {
      this.$http.get("/getUsers").then(response => {
        var temp = response.body
        this.users = [].concat.apply([], temp)
      });
    },
    getItems: function () {
      this.$http.get("/getUnreservedItems").then(response => {
        var temp = response.body
        this.items = [].concat.apply([], temp)
      });
      this.$http.get("/getReservedItems").then(response => {
        var temp = response.body
        this.items = this.items.concat(temp) //TODO
      });
    },
    getLastBookings: function () {
      //TODO magic number: 50
      this.$http.post("/getLastNBookings", 50).then(response => {
        var temp = response.body
        // this.lastBookings = [].concat.apply([], temp)
        this.lastBookings = this.lastBookings.concat(temp)
      });
    },
    showSetting(setting) {
      switch (setting) {
        case 'userSettings':
          this.showUserSettings = true
          this.showItemSettings = false
          this.showBookingSettings = false
          this.showOtherSettings = false
          break
        case 'itemSettings':
          this.showUserSettings = false
          this.showItemSettings = true
          this.showBookingSettings = false
          this.showOtherSettings = false
          break
        case 'bookingSettings':
          this.showUserSettings = false
          this.showItemSettings = false
          this.showBookingSettings = true
          this.showOtherSettings = false
          break
        case 'otherSettings':
          this.showUserSettings = false
          this.showItemSettings = false
          this.showBookingSettings = false
          this.showOtherSettings = true
          break
        default:
          break
      }
    }
  },
  created() {
    this.$nextTick(this.getUsers())
    this.$nextTick(this.getItems())
    this.$nextTick(this.getLastBookings())
  }
}