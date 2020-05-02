<template>
  <div>
    <br />
    <div class="columns">
      <div class="column is-1">
        <div class="buttons">
          <button class="button is-link is-fullwidth" @click="showAddItemForm = true">Add Item</button>
          <button class="button is-link is-fullwidth" @click="showModifyItemForm = true">Modify Item</button>
          <button class="button is-link is-fullwidth" @click="showDeleteItemForm = true">Delete Item</button>
        </div>
        <AddItemForm v-if="showAddItemForm" @close="showAddItemForm = false" />
        <ModifyItemForm v-if="showModifyItemForm" @close="showModifyItemForm = false" />
        <DeleteItemForm v-if="showDeleteItemForm" @close="showDeleteItemForm = false" />
      </div>
      <div class="column is-4">
        <ItemInfo />
        <ItemSearch :mode="'single'" />
      </div>
      <div class="column">
        <div class="box" style="height:88vh; overflow:auto;">
          <ItemTable />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AddItemForm from "../booking/AddItemForm.vue";
import ModifyItemForm from "./ModifyItemForm.vue";
import DeleteItemForm from "./DeleteItemForm.vue";
import ItemSearch from "../booking/ItemSearch.vue";
import ItemTable from "./ItemTable.vue";
import ItemInfo from "./ItemInfo.vue";

export default {
  components: {
    AddItemForm,
    ModifyItemForm,
    DeleteItemForm,
    ItemSearch,
    ItemTable,
    ItemInfo
  },
  data() {
    return {
      showAddItemForm: false,
      showModifyItemForm: false,
      showDeleteItemForm: false
    };
  }
};
</script>