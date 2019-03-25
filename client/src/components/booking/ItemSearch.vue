<template>
  <div>
    <!-- Search bar -->
    <div class="field">
      <div class="control has-icons-left">
        <input
          class="input"
          type="text"
          placeholder="Search items"
          v-model="search"
          v-on:keyup="searchItems"
        >
        <span class="icon is-small is-left">
          <font-awesome-icon icon="search"/>
        </span>
      </div>
    </div>
    <!-- Search results -->
    <div class="buttons" v-if="searchResults !== []">
      <button
        class="button"
        v-for="item in searchResults"
        :key="item"
        :class="[selectedItem === item ? 'is-link' : '']"
        @click="selectItem(item)"
      >{{ displayItem(item) }}</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    items: []
  },
  data: function() {
    return {
      search: "",
      searchResults: [],
      selectedItem: {}
    };
  },
  methods: {
    searchItems: function() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.items.filter(item =>
          item["Name"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    },
    displayItem: function(item) {
      if (item.Type == "alcoholic" || item.Type == "non-alcoholic") {
        return item.Name + " " + item.Size + " " + item.Unit;
      } else if (item.Type == "boat" || item.Type == "food") {
        return item.Name;
      } else {
        return "???";
      }
    },
    selectItem: function(item) {
      this.selectedItem = item;
      this.$emit("selectItem", item);
      this.$itemEventBus.$emit("sendToBus", item);
    },
    receiveFromEventBus(item) {
      this.selectedItem = item;
    }
  },
  created: function() {
    this.$itemEventBus.$on("sendToBus", this.receiveFromEventBus);
  }
};
</script>