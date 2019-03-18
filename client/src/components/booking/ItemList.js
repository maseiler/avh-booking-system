import ItemList from "./ItemList.vue"

export default {
  components: {
    ItemList
  },
  data: function () {
    return {
      allItems: [],
      selectedItem: {},
      activeTab: 'tab0',
    };
  },
  computed: {
    itemsAlc: function () {
      return this.allItems.filter(function (item) {
        return item.Type == 'alcoholic'
      })
    },
    itemsNonAlc: function () {
      return this.allItems.filter(function (item) {
        return item.Type == 'non-alcoholic'
      })
    },
    itemsFood: function () {
      return this.allItems.filter(function (item) {
        return item.Type == 'food'
      })
    },
    itemsBoat: function () {
      return this.allItems.filter(function (item) {
        return item.Type == 'boat'
      })
    }
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
    displayItem: function (item) {
      if (item.Type == "alcoholic" || item.Type == "non-alcoholic") {
        return item.Name + " " + item.Size + " " + item.Unit
      } else if (item.Type == "boat" || item.Type == "food") {
        return item.Name
      } else {
        return "???"
      }
    }
  },
  created() {
    this.$nextTick(this.getItems())
  }
};