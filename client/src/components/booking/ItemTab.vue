<template>
  <div>
    <ItemSearch :mode="'multiple'" />
    <hr />
    <ItemList />
  </div>
</template>

<script>
import ItemSearch from "./ItemSearch.vue";
import ItemList from "./ItemList.vue";

export default {
  components: {
    ItemSearch,
    ItemList,
  },
};
</script>