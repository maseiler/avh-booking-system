<template>
  <div>
    <!-- Search bar -->
    <div class="field">
      <div class="control has-icons-left has-icons-right">
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
        <span v-if="search != ''"
              class="icon is-small is-right"
              @click="clearSearch"
              >
              <font-awesome-icon icon="times" />
            </span>
      </div>
    </div>
    <!-- Search results -->
    <div class="buttons" v-if="searchResults !== []">
      <button
        class="button"
        v-for="item in searchResults"
        :key="item"
        :class="[isSelected(item) ? 'is-link' : '']"
        @click="selectItem(item)"
      >
        {{ displayItem(item) }}
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    mode: String,
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
        this.searchResults = this.items.filter((item) => {
          if(item["Name"].toLowerCase().includes(tmpSearch) && item["Enabled"] == 1){
            return true
          }
          return false
        });
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
    clearSearch(){
      this.search = "";
      this.searchItems();
    },
  },
};
</script>