<template>
  <div>
    <br />
    <div class="columns">
      <div class="column is-1">
        <div class="buttons">
          <button class="button is-link is-fullwidth" @click="showAddItemCategoryForm = true">Add</button>
          <!--<button class="button is-link is-fullwidth" @click="showModifyItemCategoryForm = true">Modify Item Category</button>-->
          <button class="button is-link is-fullwidth" @click="showDeleteItemCategoryForm = true">Delete</button>
        </div>
        <!--<AddItemCategoryMapForm v-if="showAddItemForm" @close="showAddItemCategoryForm = false" />-->
        <!--<ModifyItemForm v-if="showModifyItemForm" @close="showModifyItemCategoryForm = false" />-->
      </div>
      <div class="column is-6">
        <ItemInfo />
        <ItemSearch :mode="'single'" />
        <AddItemCategoryForm />
      </div>
      <div class="column">
        <div class="box" style="height:88vh; overflow:auto;">
          <ItemCategoryMapsTable />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
//import AddItemCategoryMapForm from "./AddItemCategoryMapForm.vue";
import ItemSearch from "../booking/ItemSearch.vue";
import ItemInfo from "./ItemInfo.vue";
import ItemCategoryMapsTable from "./ItemCategoryMapsTable.vue";

export default {
  components: {
    //AddItemCategoryMapForm,
    //ModifyItemForm,
    //DeleteItemCategoryForm,
    ItemCategoryMapsTable,
    ItemSearch,
    ItemInfo
  },
  data() {
    return {
      showAddItemCategoryForm: false,
      showModifyItemCategoryForm: false,
      showDeleteItemCategoryForm: false
    };
  }
};
</script>