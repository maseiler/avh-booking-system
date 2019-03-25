<template>
  <div>
    <h1>ITEM SETTINGS</h1>
    <div class="columns">
      <div class="column is-4">
        <div class="columns">
          <div class="column">
            <div class="field">
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  placeholder="Search item"
                  v-model="search"
                  v-on:keyup="searchItems"
                >
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="search"/>
                </span>
              </div>
            </div>
          </div>
        </div>
        <div class="buttons" v-if="searchResults != []">
          <button
            class="button"
            v-for="item in searchResults"
            :key="item"
            @click="selectedItem = item"
            :class="[selectedItem === item ? 'is-link' : '']"
          >{{ displayItem(item) }}</button>
        </div>
      </div>
      <div class="column is-narrow">
        <button class="button is-link" @click="showAddItemForm = true">Add Item</button>
        <button class="button is-link" @click="showModifyItemForm = true">Modify Item</button>
        <button class="button is-link" @click="showDeleteItemForm = true">Delete Item</button>
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
        <ItemInfo :item="selectedItem"/>
      </div>
    </div>
  </div>
</template>

<script>
import AddItemForm from "../booking/AddItemForm.vue";
import ModifyItemForm from "./ModifyItemForm.vue";
import DeleteItemForm from "./DeleteItemForm.vue";
import ItemInfo from "./ItemInfo.vue";

export default {
  components: {
    AddItemForm,
    ModifyItemForm,
    DeleteItemForm,
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
      search: "",
      searchResults: [],
      selectedItem: {}
    };
  },
  methods: {
    searchItems: function() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.items.filter(
          item =>
            item["Name"].toLowerCase().includes(tmpSearch) |
            item["Type"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    },
    displayItem: function(item) {
      if (item.Type == "boat" || item.Type == "food") {
        return item.Name;
      } else {
        return item.Name + " " + item.Size + item.Unit;
      }
    }
  }
};
</script>