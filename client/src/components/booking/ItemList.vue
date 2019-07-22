<template>
  <div>
    <div class="tabs">
      <ul>
        <li :class="[ activeTab === 'tab0' ? 'is-active' : '']">
          <a @click="activeTab='tab0'">
            <span class="icon is-small is-left">
              <font-awesome-icon icon="star" />
            </span>
            Favorites
          </a>
        </li>
        <li :class="[ activeTab === 'tab1' ? 'is-active' : '']">
          <a @click="activeTab='tab1'">
            <span class="icon is-small is-left">
              <font-awesome-icon icon="beer" />
            </span>
            Alcoholic
          </a>
        </li>
        <li :class="[ activeTab === 'tab2' ? 'is-active' : '']">
          <a @click="activeTab='tab2'">
            <span class="icon is-small is-left">
              <font-awesome-icon icon="glass-whiskey" />
            </span>
            Non-Alcoholic
          </a>
        </li>
        <li :class="[ activeTab === 'tab3' ? 'is-active' : '']">
          <a @click="activeTab='tab3'">
            <span class="icon is-small is-left">
              <font-awesome-icon icon="utensils" />
            </span>
            Food
          </a>
        </li>
        <li :class="[ activeTab === 'tab4' ? 'is-active' : '']">
          <a @click="activeTab='tab4'">
            <span class="icon is-small is-left">
              <font-awesome-icon icon="anchor" />
            </span>
            Boats
          </a>
        </li>
      </ul>
    </div>

    <div class="buttons" v-if="activeTab === 'tab0'">
      <button
        class="button"
        v-for="item in favoriteItems"
        :key="item"
        :class="[selectedItems.includes(item) ? 'is-link' : '']"
        @click="selectItem(item)"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab1'">
      <button
        class="button"
        v-for="item in itemsAlc"
        :key="item"
        :class="[selectedItems.includes(item) ? 'is-link' : '']"
        @click="selectItem(item)"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab2'">
      <button
        class="button"
        v-for="item in itemsNonAlc"
        :key="item"
        @click="selectItem(item)"
        :class="[selectedItems.includes(item) ? 'is-link' : '']"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab3'">
      <button
        class="button"
        v-for="item in itemsFood"
        :key="item"
        @click="selectItem(item)"
        :class="[selectedItems.includes(item) ? 'is-link' : '']"
      >{{ displayItem(item) }}</button>
    </div>

    <div class="buttons" v-if="activeTab === 'tab4'">
      <button
        class="button"
        v-for="item in itemsBoat"
        :key="item"
        @click="selectItem(item)"
        :class="[selectedItems.includes(item) ? 'is-link' : '']"
      >{{ displayItem(item) }}</button>
    </div>
  </div>
</template>

 <script>
export default {
  data() {
    return {
      favoriteItems: [],
      activeTab: "tab0"
    };
  },
  watch: {
    user() {
      this.getFavoriteItems();
    }
  },
  computed: {
    allItems() {
      return this.$store.state.items;
    },
    itemsAlc() {
      return this.$store.getters.itemsAlc;
    },
    itemsNonAlc() {
      return this.$store.getters.itemsNonAlc;
    },
    itemsFood() {
      return this.$store.getters.itemsFood;
    },
    itemsBoat() {
      return this.$store.getters.itemsBoat;
    },
    user() {
      return this.$store.state.selectedUser;
    },
    selectedItems() {
      return this.$store.state.selectedMultipleItems;
    }
  },
  methods: {
    selectItem(item) {
      this.$store.commit("selectMultipleItems", item);
    },
    getItem(id) {
      return this.allItems.find(i => {
        return i.ID == id;
      });
    },
    async getFavoriteItems() {
      var result = [];
      await this.$http
        .post("getFavoriteItemIDs", this.user)
        .then(response => {
          var favoriteItemIDs = [].concat.apply([], response.body);
          favoriteItemIDs.forEach(id => {
            var item = this.getItem(id);
            result = [].concat(result, item);
          });
        })
        .catch(response => {
          this.$responseEventBus.$emit(
            "failureMessage",
            "Couldn't get favorite items."
          );
        });
      this.favoriteItems = result;
    }
  }
};
</script>
