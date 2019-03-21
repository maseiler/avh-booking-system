import ItemTab from "./ItemTab.vue"
import ItemSearch from "./ItemSearch.vue"
import ItemList from "./ItemList.vue"
import Vue from "vue"

export default {
  components: {
    ItemTab,
    ItemSearch,
    ItemList
  },
  data: function () {
    return {
      allItems: [],
      selectedItem: {}
    };
  },
  methods: {
    getItems: function () {
      this.$http.get("/getUnreservedItems").then(response => {
        var temp = response.body;
        this.allItems = [].concat.apply([], temp)
        this.sortByName(this.allItems)
      });
    },
    sortByName: function (array) {
      array.sort(function (a, b) {
        var nameA = a["Name"].toLowerCase(), nameB = b["Name"].toLowerCase()
        var sizeA = a.Size.toLowerCase(), sizeB = b.Size.toLowerCase()
        if (nameA < nameB) return -1
        if (nameA > nameB) return 1
        if (sizeA < sizeB) return -1
        if (sizeA > sizeB) return 1
        return 0
      })
    },
    selectItem: function (item) {
      this.selectedItem = item
      this.$emit('selectedItem', item)
    }
  },
  created() {
    this.$nextTick(this.getItems())
  }
};

var ItemEventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $itemEventBus: {
    get: function () {
      return ItemEventBus;
    }
  }
});