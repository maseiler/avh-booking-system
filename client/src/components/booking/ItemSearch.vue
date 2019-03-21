<template>
  <div>
    <!-- Search bar -->
    <div class="field">
      <div class="control has-icons-left">
        <input
          class="input"
          type="text"
          placeholder="Search items"
          v-model="search"
          v-on:keyup="searchItems"
        >
        <span class="icon is-small is-left">
          <font-awesome-icon icon="search"/>
        </span>
      </div>
    </div>
    <!-- Search results -->
    <div class="buttons" v-if="searchResults !== []">
      <button
        class="button"
        v-for="item in searchResults.filter(item => item['Name'] != '')"
        :key="item"
        :class="[selectedItem === item ? 'is-link' : '']"
        @click="selectItem(item)"
      >{{ displayItem(item) }}</button>
    </div>
  </div>
</template>

<script src="./ItemSearch.js"></script>