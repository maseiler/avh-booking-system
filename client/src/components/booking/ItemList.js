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
      receiveText: {}
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
      // this.$emit('selectItem', item)
      this.$eventBus.$emit('message', item);
    },
    onReceive(text) {
      this.selectedItem = text
      this.receiveText = text;
    },
  },
  created: function () {
    this.$eventBus.$on('message', this.onReceive);
  },
};