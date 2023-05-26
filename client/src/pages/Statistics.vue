<template>
  <div>
    <div class="columns">
      <div class="column is-1 is-sidebar-menu">
        <div class="box max-height-scroll-y">
          <aside class="menu">
            <p class="menu-label">{{ $t("statistics.menuLabel") }}</p>
            <ul class="menu-list">
              <li>
                <a
                  @click="showSetting('showBookingsStats')"
                  :class="[showBookingsStats ? 'is-active' : '']"
                  >{{ $t("statistics.showBookingsStats") }}</a>
              </li>
              <li>
                <a
                  @click="showSetting('showFavoriteItemsStats')"
                  :class="[showFavoriteItemsStats ? 'is-active' : '']"
                  >{{ $t("statistics.showFavoriteItemsStats") }}</a>
              </li>
            </ul>
          </aside>
        </div>
      </div>
      <div class="column">
        <BookingStats v-if="showBookingsStats" />
        <FavoriteItemsStats v-if="showFavoriteItemsStats" />
      </div>
    </div>
  </div>
</template>

<script>
import BookingStats from "../components/statistics/BookingStats.vue";
import FavoriteItemsStats from "../components/statistics/FavoriteItemsStats.vue";

export default {
  components: {
    BookingStats,
    FavoriteItemsStats,
  },
  data() {
    return {
      showBookingsStats: true,
      showFavoriteItemsStats: false,
      showDrinkSharesStats: false,
      showTypeStats: false,
    };
  },
  methods: {
    showSetting(setting) {
      switch (setting) {
        case "showBookingsStats":
          (this.showBookingsStats = true),
            (this.showFavoriteItemsStats = false),
            (this.showDrinkSharesStats = false),
            (this.showTypeStats = false);
          break;
        case "showFavoriteItemsStats":
          (this.showBookingsStats = false),
            (this.showFavoriteItemsStats = true),
            (this.showDrinkSharesStats = false),
            (this.showTypeStats = false);
          break;
        case "showDrinkSharesStats":
          (this.showBookingsStats = false),
            (this.showFavoriteItemsStats = false),
            (this.showDrinkSharesStats = true),
            (this.showTypeStats = false);
          break;
        case "showTypeStats":
          (this.showBookingsStats = false),
            (this.showFavoriteItemsStats = false),
            (this.showDrinkSharesStats = false),
            (this.showTypeStats = true);
          break;
        default:
          break;
      }
    },
  },
};
</script>