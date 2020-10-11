<template>
  <div>
    <table class="table is-hoverable is-striped">
      <tbody>
        <tr v-for="itemC in itemCs" :key="itemC">
          <th>{{ itemC.Category }}</th>
          <th>
            <button @click="mapCategoryItems">map</button>
            <!--
          <th>
            <div class="buttons" v-if="activeTab === 'tab1'">
              <button
                class="button"
                v-for="item in items"
                :key="item"
                :class="[selectedItems.includes(item) ? 'is-link' : '']"
              >
                {{ displayItem(item) }}
              </button>
            </div>
          </th>
          --></th>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  props: {
    type: String,
  },
  computed: {
    mapCategoryItems() {
      console.log("mapCategoryItems");
      var itemCategories = this.$store.state.itemCategories;
      var items = this.getFilteredItems(itemCategories);
      //var map = [];
      // List<{CategoryName, items}>
      return items;
    },
    itemCs() {
      return this.$store.state.itemCategories;
    },
    selectedItem() {
      return this.$store.state.selectedSingleItem;
    },
  },
  methods: {
    getFilteredItems(itemCategories) {
      console.log("getFilteredItemss");
      var itemList = [];
      var categoryList = [];
      var itemsOfType;
      if (this.type === "alcoholic") {
        itemsOfType = this.$store.getters.itemsAlc;
      } else if (this.type === "non-alcoholic") {
        //...
      }
      for (let i = 0; i < itemsOfType.length; i++) {
        const item = itemsOfType[i];
        for (let j = 0; j < itemCategories.length; j++) {
          const itemC = itemCategories[j];
          if (item.ID === itemC.ItemID) {
            itemList = itemList.concat(item);
            if (categoryList.indexOf(itemC.Category) < 0) {
              categoryList = categoryList.concat(itemC.Category);
            }
          }
        }
      }
      console.log(itemList);
      console.log(categoryList);
      return itemList;
    },
  },
};
</script>