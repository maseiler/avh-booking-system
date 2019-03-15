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
        <AddItemForm v-if="showAddItemForm" @close="showAddItemForm = false"></AddItemForm>
        <ModifyItemForm
          :item="selectedItem"
          v-if="showModifyItemForm"
          @close="showModifyItemForm = false"
        ></ModifyItemForm>
        <DeleteItemForm
          :item="selectedItem"
          v-if="showDeleteItemForm"
          @close="showDeleteItemForm = false"
        ></DeleteItemForm>
        <ItemInfo :item="selectedItem"></ItemInfo>
      </div>
    </div>
  </div>
</template>

<script src="./ItemSettings.js"></script>