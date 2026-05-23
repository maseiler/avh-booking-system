<template>
  <p class="control has-icons-left has-icons-right">
    <input v-model="search" type="text" class="input" :placeholder="$t('search')" />
    <span class="icon is-left">
      <icon :icon="['fas', 'search']" />
    </span>
    <span v-if="search != ''" class="icon is-right" @click="search = ''">
      <icon :icon="['fas', 'times']" />
    </span>
  </p>
  <!-- Category Filter -->
  <div class="tags is-flex is-justify-content-space-around category-filter">
    <button class="tag is-hoverable" :class="selectedCategory == 0 ? 'is-primary' : ''" @click="selectCategory(0)">All</button>
    <button v-for="category in categories" class="tag is-hoverable" :class="selectedCategory == category.id ? 'is-primary' : ''" @click="selectCategory(category.id)">
      <span class="icon"><icon :icon="category.icon"/></span>
      <span>{{ category.title }}</span>
    </button>
  </div>
  <hr>

  <ShowAccountButton v-if="show=='button'" :accounts="accounts" />
  <ShowAccountList v-if="show=='list'" :accounts="accounts" />
</template>

<style scoped>
.icon.is-right{
  pointer-events: all;
  cursor:pointer;
}
.category-filter{
  margin-bottom: 0;
  margin-top:.5rem;
}
hr{
  margin:.5rem;
}
</style>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCategoryStore } from '../../store/CategoryStore';
import ShowAccountButton from './ShowAccountButton.vue';
import ShowAccountList from './ShowAccountList.vue';

export default {
  components: {
    ShowAccountButton,
    ShowAccountList
  },
  props: {
    show: String,
    all: Boolean
  },
  data() {
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      selectedCategory: 0 as number,
      search: "" as string
    }
  },
  methods: {
    selectCategory(categoryId: number){
      this.selectedCategory = categoryId;
    }
  },
  computed: {
    accounts(){
      return this.account$.getBySearchAndCategory(this.search, this.selectedCategory, this.all);
    },
    categories(){
      if (this.all) return this.category$.accountCategories;
      return this.category$.accountCategories.filter(cat => cat.enabled);
    }
  }
}
</script>