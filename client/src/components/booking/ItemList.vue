<template>
  <div>
    <p class="title">Items</p>
    <div class="tabs">
      <ul>
        <li :class="[ activeTab === 'tab0' ? 'is-active' : '']">
          <a @click="activeTab='tab0'">All</a>
        </li>
        <li :class="[ activeTab === 'tab1' ? 'is-active' : '']">
          <a @click="activeTab='tab1'">Alcoholic</a>
        </li>
        <li :class="[ activeTab === 'tab2' ? 'is-active' : '']">
          <a @click="activeTab='tab2'">Non-Alcoholic</a>
        </li>
        <li :class="[ activeTab === 'tab3' ? 'is-active' : '']">
          <a @click="activeTab='tab3'">Food</a>
        </li>
        <li :class="[ activeTab === 'tab4' ? 'is-active' : '']">
          <a @click="activeTab='tab4'">Boats</a>
        </li>
      </ul>
    </div>

    <div class="buttons" v-if="activeTab === 'tab0'">
      <button
        class="button"
        v-for="item in allItems"
        :key="item"
        :class="[selectedItem === item ? 'is-link' : '']"
        @click="selectedItem = item"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab1'">
      <button
        class="button"
        v-for="item in itemsAlc"
        :key="item"
        :class="[selectedItem === item ? 'is-link' : '']"
        @click="selectedItem = item"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab2'">
      <button
        class="button"
        v-for="item in itemsNonAlc"
        :key="item"
        @click="selectedItem = item"
        :class="[selectedItem === item ? 'is-link' : '']"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab3'">
      <button
        class="button"
        v-for="item in itemsFood"
        :key="item"
        @click="selectedItem = item"
        :class="[selectedItem === item ? 'is-link' : '']"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab4'">
      <button
        class="button"
        v-for="item in itemsBoat"
        :key="item"
        @click="selectedItem = item"
        :class="[selectedItem === item ? 'is-link' : '']"
      >{{ displayItem(item) }}</button>
    </div>
  </div>
</template>

 <script src="./ItemList.js"></script>