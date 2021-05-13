<template>
  <div>
    <table class="table is-hoverable is-striped">
      <tbody>
        <tr v-for="icMap in icMaps" :key="icMap">
          <th>{{ icMap.Category }}</th>
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
    icMaps() {
      var icMaps = this.$store.state.itemCategoryMaps;
      var itemCategories = this.$store.state.itemCategories;
      var icMapList = [];
      var items;
      if (this.type === "alcoholic") {
        items = this.$store.getters.itemsNonAlc;
      } else if (this.type === "non-alcoholic") {
        //...
      }
      for (let i = 0; i < items.length; i++) {
        const item = items[i];
        for (let j = 0; j < icMaps.length; j++) {
          const icMap = icMaps[j];
          if (item.ID === icMap.ItemID) {
            let categoryName = this.getCategoryNameByID(
              itemCategories,
              icMap.CategoryID
            );
            if (this.categoryExists(icMapList, categoryName)) {
              icMapList = icMapList[categoryName].concat({ item }); //TODO
            } else {
              icMapList = icMapList.concat({ categoryName: [{ item }] });
            }
            // check if category exists in new list
            /*
            for (const icMap in icMapList) {
              if (!icMap.hasOwnProperty(categoryName)) {
                icMapList[categoryName].concat({ item });
                // TODO check if work
              } else {
                icMapList.concat({ categoryName: [{ item }] });
              }
            }
            */
          } else {
            /*
            if (!icMap.hasOwnProperty("Others")) {
              icMapList[categoryName].concat({ item });
            } else {
              icMapList.concat({ Others: [{ item }] });
            }
            */
          }
        }
      }
      console.log(icMapList);
      return icMapList;
    },
    selectedItem() {
      return this.$store.state.selectedSingleItem;
    },
  },
  methods: {
    getCategoryNameByID(itemCategories, id) {
      for (let i = 0; i < itemCategories.length; i++) {
        const ic = itemCategories[i];
        console.log(ic);
        if (ic.ID == id) {
          return ic.Name;
        }
      }
      return "";
    },
    categoryExists(icMapList, categoryName) {
      for (let i = 0; i < icMapList.length; i++) {
        const icMap = icMapList[i];
        if (icMap[categoryName] != null) {
          return true;
        }
      }
      return false;
    },
  },
};
</script>