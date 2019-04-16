<template>
  <div>
    <ItemSearch :items="allItems" @selectItems="selectItems"/>
    <hr>
    <ItemList :allItems="allItems" @selectItems="selectItems"/>
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
  data: function() {
    return {
      allItems: [],
      selectedItems: []
    };
  },
  methods: {
    getItems: function() {
      this.$http.get("/getUnreservedItems").then(response => {
        var temp = response.body;
        this.allItems = [].concat.apply([], temp);
        this.sortByName(this.allItems);
      });
    },
    sortByName: function(array) {
      array.sort(function(a, b) {
        var nameA = a["Name"].toLowerCase(),
          nameB = b["Name"].toLowerCase();
        var sizeA = a.Size.toLowerCase(),
          sizeB = b.Size.toLowerCase();
        if (nameA < nameB) return -1;
        if (nameA > nameB) return 1;
        if (sizeA < sizeB) return -1;
        if (sizeA > sizeB) return 1;
        return 0;
      });
    },
    selectItems: function(items) {
      this.selectedItems = items;
      this.$emit("selectedItems", this.selectedItems);
    }
  },
  created() {
    this.$nextTick(this.getItems());
  }
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