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
    <button class="tag is-hoverable" :class="selectedCategory == 0 ? 'is-primary' : ''" @click="selectCategory(0)">{{ $t('all') }}</button>
    <button v-for="category in categories" class="tag is-hoverable" :class="selectedCategory == category.id ? 'is-primary' : ''" @click="selectCategory(category.id)">
      <span class="icon"><icon :icon="category.icon"/></span>
      <span>{{ category.title }}</span>
    </button>
  </div>
  <hr>

  <ShowProductButton v-if="show=='button'" :products="products" />
  <ShowProductList v-if="show=='list'" :products="products" />
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
import { useCategoryStore } from '../../store/CategoryStore';
import ShowProductButton from './ShowProductButton.vue';
import ShowProductList from './ShowProductList.vue';
import { useProductStore } from '../../store/ProductStore';
import { useAccountStore } from '../../store/AccountStore';
import { useProductGroupStore } from '../../store/ProductGroupStore.ts';

export default {
  components: {
    ShowProductButton,
    ShowProductList
  },
  props: {
    show: String,
    all: Boolean
  },
  data() {
    return {
      product$: useProductStore(),
      productGroup$: useProductGroupStore(),
      category$: useCategoryStore(),
      account$: useAccountStore(),
      selectedCategory: 0 as number,
      search: "" as string
    }
  },
  methods: {
    selectCategory(categoryId: number){
      this.selectedCategory = categoryId;
    },
  },
  computed: {
    products(){
      let selectedAccountCategorys = [...new Set(this.account$.selected.map(a => a.category))] as number[];
      // return this.product$.products;
      return this.product$.getBySearchAndCategory(this.search, this.selectedCategory, selectedAccountCategorys, this.all);
    },
    categories(){
      if (this.all) return this.category$.productCategories;
      return this.category$.productCategories.filter(cat => cat.enabled);
    }
  }
}
</script>