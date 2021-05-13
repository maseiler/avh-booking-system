<template>
  <div>
    <table class="table is-hoverable is-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Category</th>
          <th>Item</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="ic in icMaps"
          :key="ic"
          @click="selectItemCategory(ic)"
          :class="[selectedItemCategoryMap.ID === ic.ID ? 'is-selected' : '']"
        >
          <th>{{ ic.ID }}</th>
          <td>{{ getItemCategoryByID(ic.CategoryID).Name }}</td>
          <td>{{ displayItem(getItemByID(ic.ItemID)) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  computed: {
    icMaps() {
      //TODO sort
      return this.$store.state.itemCategoryMaps;
    },
    selectedItemCategoryMap() {
      return this.$store.state.selectedItemCategoryMap;
    },
  },
  methods: {
    selectItemCategory(icMap) {
      this.$store.commit("selectItemCategoryMap", icMap);
    },
  },
};
</script>
