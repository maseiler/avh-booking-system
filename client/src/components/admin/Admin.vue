<template>
  <div>
    <div class="columns">
      <div class="column is-2">
        <div class="box">
          <aside class="menu">
            <p class="menu-label">Administration</p>
            <ul class="menu-list">
              <li>
                <a
                  @click="showSetting('userSettings')"
                  :class="[ showUserSettings ? 'is-active' : '']"
                >User Settings</a>
              </li>
              <li>
                <a
                  @click="showSetting('itemSettings')"
                  :class="[ showItemSettings ? 'is-active' : '']"
                >Item Settings</a>
              </li>
              <li>
                <a
                  @click="showSetting('bookingSettings')"
                  :class="[ showBookingSettings ? 'is-active' : '']"
                >Booking Settings</a>
              </li>
              <li>
                <a
                  @click="showSetting('otherSettings')"
                  :class="[ showOtherSettings ? 'is-active' : '']"
                >Other Settings</a>
              </li>
              <li>
                <a
                  @click="showSetting('feedbackList')"
                  :class="[ showFeedbackList ? 'is-active' : '']"
                >Feedback</a>
              </li>
            </ul>
          </aside>
        </div>
      </div>
      <div class="column">
        <UserSettings v-if="showUserSettings" :users="users"/>
        <ItemSettings v-if="showItemSettings" :items="items"/>
        <BookingSettings
          v-if="showBookingSettings"
          :lastBookings="lastBookings"
          :users="users"
          :items="items"
        />
        <OtherSettings v-if="showOtherSettings"/>
        <FeedbackList v-if="showFeedbackList" :feedbackList="feedbackList"/>
      </div>
    </div>
  </div>
</template>

<script>
import UserSettings from "./UserSettings.vue";
import ItemSettings from "./ItemSettings.vue";
import BookingSettings from "./BookingSettings.vue";
import OtherSettings from "./OtherSettings.vue";
import FeedbackList from "./FeedbackList.vue";

export default {
  components: {
    UserSettings,
    ItemSettings,
    BookingSettings,
    OtherSettings,
    FeedbackList
  },
  data: function() {
    return {
      users: [],
      items: [],
      lastBookings: [],
      feedbackList: [],
      showUserSettings: false,
      showItemSettings: false,
      showBookingSettings: false,
      showOtherSettings: false,
      showFeedbackList: false
    };
  },
  methods: {
    getUsers: function() {
      this.$http.get("/getUsers").then(response => {
        var temp = response.body;
        this.users = [].concat.apply([], temp);
      });
    },
    getItems: function() {
      this.$http.get("/getUnreservedItems").then(response => {
        var temp = response.body;
        this.items = [].concat.apply([], temp);
      });
      this.$http.get("/getReservedItems").then(response => {
        var temp = response.body;
        this.items = this.items.concat(temp); //TODO
      });
    },
    getLastBookings: function() {
      //TODO magic number: 50
      this.$http.post("/getLastNBookings", 50).then(response => {
        var temp = response.body;
        this.lastBookings = this.lastBookings.concat(temp);
      });
    },
    getFeedbackList: function() {
      this.$http.get("/getFeedback").then(response => {
        var temp = response.body;
        this.feedbackList = this.feedbackList.concat(temp);
      });
    },
    showSetting(setting) {
      switch (setting) {
        case "userSettings":
          this.showUserSettings = true;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          break;
        case "itemSettings":
          this.showUserSettings = false;
          this.showItemSettings = true;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          break;
        case "bookingSettings":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = true;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          break;
        case "otherSettings":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = true;
          this.showFeedbackList = false;
          break;
        case "feedbackList":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = true;
          break;
        default:
          break;
      }
    }
  },
  created() {
    this.$nextTick(this.getUsers());
    this.$nextTick(this.getItems());
    this.$nextTick(this.getLastBookings());
    this.$nextTick(this.getFeedbackList());
  }
};
</script>