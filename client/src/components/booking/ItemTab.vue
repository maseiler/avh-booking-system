<template>
  <div>
    <ItemSearch @selectItems="selectItems"/>
    <hr>
    <ItemList @selectItems="selectItems"/>
  </div>
</template>

 <script>
import ItemSearch from "./ItemSearch.vue";
import ItemList from "./ItemList.vue";
import Vue from "vue";

export default {
  components: {
    ItemSearch,
    ItemList
  },
  data() {
    return {
      selectedItems: []
    };
  },
  methods: {
    selectItems(items) {
      this.selectedItems = items;
      this.$emit("selectedItems", this.selectedItems);
    }
  },
};

var ItemEventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $itemEventBus: {
    get: function() {
      return ItemEventBus;
    }
  }
});
</script>