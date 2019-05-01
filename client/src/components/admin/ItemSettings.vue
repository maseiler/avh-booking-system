<template>
  <div>
    <br>
    <div class="columns">
      <div class="column is-1">
        <div class="buttons">
          <button class="button is-link is-fullwidth" @click="showAddItemForm = true">Add Item</button>
          <button class="button is-link is-fullwidth" @click="showModifyItemForm = true">Modify Item</button>
          <button class="button is-link is-fullwidth" @click="showDeleteItemForm = true">Delete Item</button>
        </div>
        <AddItemForm v-if="showAddItemForm" @close="showAddItemForm = false"/>
        <ModifyItemForm
          :item="selectedItem"
          v-if="showModifyItemForm"
          @close="showModifyItemForm = false"
        />
        <DeleteItemForm
          :item="selectedItem"
          v-if="showDeleteItemForm"
          @close="showDeleteItemForm = false"
        />
      </div>
      <div class="column is-4">
        <ItemSearch :items="items" @selectItems="selectItems"/>
      </div>
      <div class="column is-4">
        <ItemInfo :item="selectedItem"/>
      </div>
    </div>
  </div>
</template>

<script>
import AddItemForm from "../booking/AddItemForm.vue";
import ModifyItemForm from "./ModifyItemForm.vue";
import DeleteItemForm from "./DeleteItemForm.vue";
import ItemSearch from "../booking/ItemSearch.vue";
import ItemInfo from "./ItemInfo.vue";

export default {
  components: {
    AddItemForm,
    ModifyItemForm,
    DeleteItemForm,
    ItemSearch,
    ItemInfo
  },
  props: {
    items: []
  },
  data: function() {
    return {
      showAddItemForm: false,
      showModifyItemForm: false,
      showDeleteItemForm: false,
      selectedItem: []
    };
  },
  methods: {
    selectItems: function(items) {
      this.selectedItem = items[items.length - 1];
      var selectedItems = Array(1).fill(this.selectedItem);
      this.$itemEventBus.$emit("selectItemsToBus", selectedItems);
    }
  }
};
</script>