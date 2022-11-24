<template>
  <div class="item-list">
    <columns
      class="columns"
      v-for="(items, key) in itemDict"
      :key="key"
      style="margin-bottom: 0cm"
    >
      <column class="column is-1">
        <h3 class="subtitle is-3 has-text-grey-light has-text-weight-bold">
          {{ key }}
        </h3>
      </column>
        <buttons class="buttons">
          <button
            class="button"
            v-for="item in items"
            :key="item"
            :class="[selectedItems.includes(item) ? 'is-link' : '']"
            @click="selectItem(item)"
          >
            {{ displayItem(item) }}
          </button>
        </buttons>
      </columns>
  </div>
</template>

<script>
export default {
  props: {
    itemDict: [],
  },
  computed: {
    selectedItems() {
      return this.$store.state.selectedMultipleItems;
    },
  },
};
</script>