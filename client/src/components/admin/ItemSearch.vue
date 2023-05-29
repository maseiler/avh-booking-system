<template>
  <div class="columns">
    <!-- Search bar -->
    <div class="column is-5">
      <div class="field">
        <div class="control has-icons-left">
          <input
            class="input"
            type="text"
            v-bind:placeholder="$t('booking.search.items')"
            v-model="search"
            v-on:keyup="searchItems"
          />
          <span class="icon is-small is-left">
            <font-awesome-icon icon="search" />
          </span>
        </div>
      </div>
      <ItemInfo />
    </div>
    <!-- Search results -->
    <div class="column">
      <buttons class="buttons" v-if="searchResults !== []">
        <button
          class="button"
          v-for="item in searchResults"
          :key="item"
          :class="[isSelected(item) ? 'is-link' : '']"
          @click="selectItem(item)"
        >
          {{ displayItem(item) }}
        </button>
      </buttons>
    </div>
  </div>
</template>

<script>
import ItemInfo from "./ItemInfo.vue";

export default {
  props: {
    mode: String,
  },
  components: {
    ItemInfo,
  },
  computed: {
    items() {
      return this.$store.state.items;
    },
    selectedItems() {
      if (this.mode === "single") {
        return this.$store.state.selectedSingleItem;
      }
      return this.$store.state.selectedMultipleItems;
    },
  },
  data() {
    return {
      search: "",
      searchResults: [],
    };
  },
  methods: {
    searchItems() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.items.filter((item) =>
          item["Name"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    },
    selectItem(item) {
      if (this.mode === "single") {
        this.$store.commit("selectSingleItem", item);
      } else {
        this.$store.commit("selectMultipleItems", item);
      }
    },
    isSelected(item) {
      if (this.mode === "single") {
        if (this.selectedItems === item) {
          return true;
        }
        return false;
      }
      return this.selectedItems.includes(item);
    },
  },
};
</script>