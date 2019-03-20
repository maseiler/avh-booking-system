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
    }
  },
  created() {
    this.$nextTick(this.getItems())
  }
};

var EventBus = new Vue();
Object.defineProperties(Vue.prototype, {
    $eventBus: {
        get: function () {
            return EventBus;
        }
    }
});

// event bus source: https://github.com/devjin0617/vuejs-eventbus-example