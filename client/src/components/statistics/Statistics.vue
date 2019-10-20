<template>
  <div>
    <div class="columns">
      <div class="column is-1 is-sidebar-menu">
        <div class="box" style="height:90vh;">
          <aside class="menu">
            <p class="menu-label">Statistics</p>
            <ul class="menu-list">
              <li>
                <a
                  @click="showSetting('showBookingsStats')"
                  :class="[ showBookingsStats ? 'is-active' : '']"
                >Bookings</a>
              </li>
              <li>
                <a
                  @click="showSetting('showFavoriteItemsStats')"
                  :class="[ showFavoriteItemsStats ? 'is-active' : '']"
                >Favorite Items</a>
              </li>
              <!-- <li>
                <a
                  @click="showSetting('showDrinkSharesStats')"
                  :class="[ showDrinkSharesStats ? 'is-active' : '']"
                >Drink Share</a>
              </li>
              <li>
                <a
                  @click="showSetting('showTypeStats')"
                  :class="[ showTypeStats ? 'is-active' : '']"
                >Type Shares</a>
              </li>-->
            </ul>
          </aside>
        </div>
      </div>
      <div class="column">
        <BookingStats v-if="showBookingsStats" />
        <FavoriteItemsStats v-if="showFavoriteItemsStats" />
        <!-- <DrinkSharesStats v-if="showDrinkSharesStats"/>
        <TypeStats v-if="showTypeStats"/>-->
      </div>
    </div>
  </div>
</template>

<script>
import BookingStats from "./BookingStats.vue";
import FavoriteItemsStats from "./FavoriteItemsStats.vue";

export default {
  components: {
    BookingStats,
    FavoriteItemsStats
  },
  data() {
    return {
      showBookingsStats: true,
      showFavoriteItemsStats: false,
      showDrinkSharesStats: false,
      showTypeStats: false
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
    }
  }
};
</script>