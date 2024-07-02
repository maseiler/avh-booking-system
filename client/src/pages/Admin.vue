<template>
  <div>
    <div class="columns">
      <div class="column is-2 is-sidebar-menu">
        <div class="box max-height-scroll-y">
          <aside class="menu">
            <p class="menu-label">{{ $t("admin.menuLabel") }}</p>
            <ul class="menu-list">
              <li>
                <a
                  @click="showSetting('userSettings')"
                  :class="[showUserSettings ? 'is-active' : '']"
                  >{{ $t("admin.userSettings.menuLabel") }}</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('itemSettings')"
                  :class="[showItemSettings ? 'is-active' : '']"
                  >{{ $t("admin.itemSettings.menuLabel") }}</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('bookingSettings')"
                  :class="[showBookingSettings ? 'is-active' : '']"
                  >{{ $t("admin.bookingSettings.menuLabel") }}</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('otherSettings')"
                  :class="[showOtherSettings ? 'is-active' : '']"
                  >{{ $t("admin.otherSettings.menuLabel") }}</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('feedbackList')"
                  :class="[showFeedbackList ? 'is-active' : '']"
                  >{{ $t("admin.feedbackList") }}</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('reportSettings')"
                  :class="[showReportSettings ? 'is-active' : '']"
                  >{{ $t("admin.reportSettings.menuLabel") }}</a
                >
              </li>
              <li>
                <a
                  @click="showSetting('paymentSettings')"
                  :class="[showPaymentSettings ? 'is-active' : '']"
                  >{{ $t("admin.paymentSettings.menuLabel") }}</a
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
        <ReportSettings v-if="showReportSettings" />
        <PaymentSettings v-if="showPaymentSettings" />
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
import ReportSettings from "../components/admin/ReportSettings.vue";
import PaymentSettings from "../components/admin/PaymentSettings.vue";

export default {
  components: {
    UserSettings,
    ItemSettings,
    BookingSettings,
    OtherSettings,
    FeedbackList,
    ReportSettings,
    PaymentSettings
  },
  data() {
    return {
      showUserSettings: false,
      showItemSettings: false,
      showBookingSettings: false,
      showOtherSettings: false,
      showFeedbackList: false,
      showReportSettings: false,
      showPaymentSettings: false
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
          this.showReportSettings = false;
          this.showPaymentSettings = false;
          break;
        case "itemSettings":
          this.showUserSettings = false;
          this.showItemSettings = true;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          this.showReportSettings = false;
          this.showPaymentSettings = false;
          break;
        case "bookingSettings":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = true;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          this.showReportSettings = false;
          this.showPaymentSettings = false;
          break;
        case "otherSettings":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = true;
          this.showFeedbackList = false;
          this.showReportSettings = false;
          this.showPaymentSettings = false;
          break;
        case "feedbackList":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = true;
          this.showReportSettings = false;
          this.showPaymentSettings = false;
          break;
        case "reportSettings":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          this.showReportSettings = true;
          this.showPaymentSettings = false;
          break;
        case "paymentSettings":
          this.showUserSettings = false;
          this.showItemSettings = false;
          this.showBookingSettings = false;
          this.showOtherSettings = false;
          this.showFeedbackList = false;
          this.showReportSettings = false;
          this.showPaymentSettings = true;
          break;
        default:
      }
    },
  },
};
</script>