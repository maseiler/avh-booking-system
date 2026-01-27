<template>
  <h1 class="title">Edit Product: {{ product.name }}</h1>
  <div class="columns">
    <div class="column is-3">ID:</div>
    <div class="column">
      <input type="text" class="input" :value="product.id" disabled>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Name:</div>
    <div class="column">
      <input type="text" class="input" v-model="product.name">
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Size and Unit:</div>
    <div class="column is-1">
      <input type="number" class="input" v-model="product.size">
    </div>
    <div class="column is-3">
      <div class="select">
          <select v-model="product.unit">
            <option v-for="unit in unit$.all" :value="unit.id" class="has-icons-left">
              {{ unit.name }}
            </option>
        </select>
      </div>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Category:</div>
    <div class="column">
      <div class="control has-icons-left">
        <div class="select">
          <select v-model="product.category">
            <option v-for="category in category$.productCategories" :value="category.id" class="has-icons-left">
              {{ category.title }}
            </option>
          </select>
        </div>
        <div class="icon is-small is-left">
          <icon :icon="categoryIcon"/>
        </div>
      </div>
    </div>
  </div>

  <div class="columns" v-if="doneMounting">
    <div class="column is-3">Visibility:</div>
    <div class="column">
      TODO
      <!--
      <div class="buttons">
        <label class="checkbox button has-icons-right" v-for="category in category$.accountCategories">
          <input type="checkbox" :checked="product.visibility.includes(category.id)" @change="updateProductVisibility(category.id)"/>
          <span>{{ category.title }}</span>
          <div class="icon is-small is-right">
            <icon :icon="category.icon" />
          </div>
        </label>
      </div>
      -->
    </div>
  </div>
</template>

<script lang="ts">
import type { Product } from '../../composables/product';
import { useProductStore } from '../../store/ProductStore';
import { useCategoryStore } from '../../store/CategoryStore';
import {useUnitStore} from "@/store/UnitStore.ts";

export default {
  data() {
    return {
      product$: useProductStore(),
      category$: useCategoryStore(),
      unit$: useUnitStore(),
      product: {} as Product,
      doneMounting: false
    }
  },
  mounted() {
    this.product = this.product$.byId(parseInt(this.$route.params.productId.toString()))
    this.doneMounting = true
  },
  computed: {
    categoryIcon(){
      return this.category$.byId(this.product.category)?.icon
    }
  },
  methods: {
    updateProductVisibility(id: number){
      if (this.product.visibility.includes(id)){
        this.product.visibility = this.product.visibility.filter((visId) => visId != id)
        return
      }
      this.product.visibility.push(id)
      this.product.visibility.sort()
    }
  }
}
</script>