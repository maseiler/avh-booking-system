<template>
  <div>
    <p class="title">Items</p>
    <ItemSearch :items="allItems" @selectItem1="selectItem"></ItemSearch>
    <hr>
    <ItemList :allItems="allItems" @selectItem="selectItem"></ItemList>
  </div>
</template>

 <script src="./ItemTab.js"></script>