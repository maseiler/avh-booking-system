<template>
  <div>
    <button class="button">{{ displayItem(this.selectedItem) }}</button>
    <input
      class="input"
      type="text"
      placeholder="Category"
      v-model.lazy="newItemCategory.Category"
    />
    <button class="button is-link" @click="addItemCategory">Add</button>
  </div>
</template>

<script>
export default {
  computed: {
    selectedItem() {
      return this.$store.state.selectedSingleItem;
    },
  },
  data() {
    return {
      newItemCategory: {
        ItemID: null,
        Category: "",
      },
    };
  },
  methods: {
    addItemCategory() {
      this.newItemCategory.ItemID = this.selectedItem.ID;
      this.$http
        .post("addItemCategory", this.newItemCategory)
        .then(() => {
          let message = "".concat(
            "Added new Item Category: ",
            this.displayItem(this.selectedItem),
            " - ",
            this.newItemCategory.Category
          );
          this.$responseEventBus.$emit("successMessage", message);
          this.reset();
          this.$store.commit("getItemCategories");
        })
        .catch((response) => {
          this.$responseEventBus.$emit("failureMessage", response.data);
        });
    },
    reset() {
      this.$store.commit("selectSingleItem", {});
      this.newItemCategory.ItemID = null;
      this.newItemCategory.Category = "";
    },
  },
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
