import ItemList from "./ItemList.vue"

export default {
  components: {
    ItemList
  },
  props: {
    allItems: []
  },
  data: function () {
    return {
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
    props: {
        allItems: []
    },
    data: function () {
        return {
            selectedItem: {},
            activeTab: 'tab0'
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
        displayItem: function (item) {
            if (item.Type == "alcoholic" || item.Type == "non-alcoholic") {
                return item.Name + " " + item.Size + " " + item.Unit
            } else if (item.Type == "boat" || item.Type == "food") {
                return item.Name
            } else {
                return "???"
            }
        },
        selectItem: function (item) {
            this.selectedItem = item
            this.$emit('selectItem', item)
        },
    }
  },
  methods: {
    displayItem: function (item) {
      if (item.Type == "alcoholic" || item.Type == "non-alcoholic") {
        return item.Name + " " + item.Size + " " + item.Unit
      } else if (item.Type == "boat" || item.Type == "food") {
        return item.Name
      } else {
        return "???"
      }
    },
    selectItem: function (item) {
      this.selectedItem = item
      this.$emit('selectItem', item)
      this.$itemEventBus.$emit('sendToBus', item);
    },
    receiveFromEventBus(item) {
      this.selectedItem = item
    }
  },
  created: function () {
    this.$itemEventBus.$on('sendToBus', this.receiveFromEventBus);
  }
};