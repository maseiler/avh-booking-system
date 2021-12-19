<template>
  <div>
    <div class="columns">
      <div class="column is-2 is-sidebar-menu">
        <div class="box" style="height: 90vh">
          <aside class="menu">
            <p class="menu-label">Administration</p>
            <ul class="menu-list">
              <li>
                <a
                  @click="showSetting('userSettings')"
                  :class="[showUserSettings ? 'is-active' : '']"
                  >User Settings</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('itemSettings')"
                  :class="[showItemSettings ? 'is-active' : '']"
                  >Item Settings</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('bookingSettings')"
                  :class="[showBookingSettings ? 'is-active' : '']"
                  >Booking Settings</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('otherSettings')"
                  :class="[showOtherSettings ? 'is-active' : '']"
                  >Other Settings</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('feedbackList')"
                  :class="[showFeedbackList ? 'is-active' : '']"
                  >Feedback</a
                >
              </li>
            </ul>
          </aside>
        </div>
      </div>
      <div class="column">
        <UserSettings v-if="showUserSettings" />
        <ItemSettings v-if="showItemSettings" />
        <BookingSettings v-if="showBookingSettings" />
        <OtherSettings v-if="showOtherSettings" />
        <FeedbackList v-if="showFeedbackList" />
      </div>
    </div>
  </div>
</template>

<script>
import UserSettings from "../components/admin/UserSettings.vue";
import ItemSettings from "../components/admin/ItemSettings.vue";
import BookingSettings from "../components/admin/BookingSettings.vue";
import OtherSettings from "../components/admin/OtherSettings.vue";
import FeedbackList from "../components/admin/FeedbackList.vue";

export default {
  components: {
    UserSettings,
    ItemSettings,
    BookingSettings,
    OtherSettings,
    FeedbackList,
  },
  data() {
    return {
      showUserSettings: false,
      showItemSettings: false,
      showBookingSettings: false,
      showOtherSettings: false,
      showFeedbackList: false,
    };
  },
  methods: {
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
    },
  },
};
</script>