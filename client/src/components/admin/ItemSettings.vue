<template>
  <div>
    <br />
    <AddItemForm v-if="showAddItemForm" @close="showAddItemForm = false" />
    <ModifyItemForm
      v-if="showModifyItemForm"
      @close="showModifyItemForm = false"
    />
    <DeleteItemForm
      v-if="showDeleteItemForm"
      @close="showDeleteItemForm = false"
    />
    <div class="tile is-ancestor is-narrow">
      <div class="tile is-parent is-vertical">
        <article class="tile is-child">
          <div class="box field is-grouped">
            <button class="button is-link is-1" @click="showAddItemForm = true">
              {{$t('admin.itemSettings.add')}}
            </button>
            &nbsp;
            <button
              class="button is-link is-1"
              @click="showModifyItemForm = true"
            >
            {{$t('admin.itemSettings.modify')}}
            </button>
            &nbsp;
            <button
              class="button is-link is-1"
              @click="showDeleteItemForm = true"
            >
            {{$t('admin.itemSettings.deleteItemTitle')}}
            </button>
          </div>
        </article>
        <article class="tile is-child">
          <div class="box" style="height: 17rem; overflow: auto">
            <ItemSearch :mode="'single'" />
          </div>
        </article>
        <article class="tile is-child">
          <div
            class="box"
            style="height: calc(100vh - 26.25rem); overflow: auto"
          >
            <ItemTable />
          </div>
        </article>
      </div>
    </div>
  </div>
</template>

<script>
import AddItemForm from "../booking/AddItemForm.vue";
import ModifyItemForm from "./ModifyItemForm.vue";
import DeleteItemForm from "./DeleteItemForm.vue";
import ItemSearch from "./ItemSearch.vue";
import ItemTable from "./ItemTable.vue";

export default {
  components: {
    AddItemForm,
    ModifyItemForm,
    DeleteItemForm,
    ItemSearch,
    ItemTable,
  },
  data() {
    return {
      showAddItemForm: false,
      showModifyItemForm: false,
      showDeleteItemForm: false,
    };
  },
};
</script>