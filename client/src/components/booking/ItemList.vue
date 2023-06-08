<template>
  <div>
    <tabs class="tabs">
      <ul>
        <li :class="[activeTab === 'tab0' ? 'is-active' : '']">
          <a @click="activeTab = 'tab0'">
            <icon class="icon is-small is-left">
              <font-awesome-icon icon="star" />
            </icon>
            {{ $t("generic.favorites") }}
          </a>
        </li>
        <li :class="[activeTab === 'tab1' ? 'is-active' : '']">
          <a @click="activeTab = 'tab1'">
            <icon class="icon is-small is-left">
              <font-awesome-icon icon="beer" />
            </icon>
            {{ $t("generic.alcoholic") }}
          </a>
        </li>
        <li :class="[activeTab === 'tab2' ? 'is-active' : '']">
          <a @click="activeTab = 'tab2'">
            <icon class="icon is-small is-left">
              <font-awesome-icon icon="glass-whiskey" />
            </icon>
            {{ $t("generic.nonAlcoholic") }}
          </a>
        </li>
        <li :class="[activeTab === 'tab3' ? 'is-active' : '']">
          <a @click="activeTab = 'tab3'">
            <icon class="icon is-small is-left">
              <font-awesome-icon icon="utensils" />
            </icon>
            {{ $t("generic.food") }}
          </a>
        </li>
      </ul>
    </tabs>
    <ItemDictionary v-if="activeTab === 'tab0'" :itemDict="favoriteItemDict" />
    <ItemDictionary v-if="activeTab === 'tab1'" :itemDict="itemsAlc" />
    <ItemDictionary v-if="activeTab === 'tab2'" :itemDict="itemsNonAlc" />
    <ItemDictionary v-if="activeTab === 'tab3'" :itemDict="itemsFood" /> 
  </div>
</template>

<script>
import helper from "../../helper.js";
import ItemDictionary from "./ItemDictionary.vue";

export default {
  components: {
    ItemDictionary,
  },
  data() {
    return {
      favoriteItemDict: {},
      activeTab: "tab0",
    };
  },
  watch: {
    user() {
      this.getFavoriteItemDict();
    },
  },
  created() {
    this.$root.$on("set-item-tab-to-favorites", () => {
      this.activeTab = "tab0";
    });
  },
  computed: {
    itemsAlc() {
      return helper.getItemsAsDict(this.$store.getters.itemsAlc);
    },
    itemsNonAlc() {
      return helper.getItemsAsDict(this.$store.getters.itemsNonAlc);
    },
    itemsFood() {
      return helper.getItemsAsDict(this.$store.getters.itemsFood);
    },
    user() {
      return this.$store.state.selectedUser;
    },
  },
  methods: {
    async getFavoriteItemDict() {
      var result = [];
      await this.$http
        .post("getFavoriteItemIDs", this.user)
        .then((response) => {
          var favoriteItemIDs = [].concat.apply([], response.body);
          favoriteItemIDs.forEach((id) => {
            var item = helper.getItemByID(this.$store.state.items, id);
            result = [].concat(result, item);
          });
        })
        .catch((e) => {
          console.log(e);
          this.$responseEventBus.$emit(
            "failureMessage",
            this.$t('messages.failure.noFavoriteItems')
          );
        });
      this.favoriteItemDict = helper.getFavoriteItemsAsDict(result);
    },
  },
};
</script>
