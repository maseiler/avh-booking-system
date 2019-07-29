<template>
  <div>
    <table class="table is-hoverable is-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Price</th>
          <th>Size</th>
          <th>Unit</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in items"
          :key="item"
          @click="selectItem(item)"
          :class="[selectedItem === item ? 'is-selected' : '']"
        >
          <th>{{item.ID}}</th>
          <td>{{item.Name}}</td>
          <td>{{item.Price}}</td>
          <td>{{item.Size}}</td>
          <td>{{item.Unit}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  computed: {
    items() {
      return this.$store.state.items.sort(function(a, b) {
        if (a.ID < b.ID) return -1;
        if (a.ID > b.ID) return 1;
        return 0;
      });
    },
    selectedItem() {
      return this.$store.state.selectedSingleItem;
    }
  },
  methods: {
    selectItem(item) {
      this.$store.commit("selectSingleItem", item);
    }
  }
  //TODO select
};
</script>
