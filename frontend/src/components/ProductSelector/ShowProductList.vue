<template>
  <div class="table-container">
    <table class="table is-striped is-hoverable">
      <tbody>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Size</th>
          <th>Unit</th>
          <th>Category</th>
          <th>Group</th>
          <th>Visibility</th>
          <th>Price</th>
          <th>Vat</th>
          <th v-show="hasEditProductRights">Actions</th>
        </tr>
        <tr :class="product$.selected == product ? 'is-primary' : ''" v-for="product in products" @click="product$.select(product)">
          <td>{{ product.id }}</td>
          <td>{{ product.name }}</td>
          <td>{{ product.size }}</td>
          <td>{{ product.getUnit()?.name }}</td>
          <td>
            <button class="tag">
              <span class="icon"><icon :icon="product.getCategory()?.icon" /></span>
              <span>{{ category$.byId(product.getCategory()?.id)?.title }}</span>
            </button>            
          </td>
          <td>{{ product.getGroup()?.name }}</td>
          <td>
            <!--
            <button v-for="categoryNumber in product.visibility" class="tag">
              <span class="icon"><icon :icon="category$.byId(categoryNumber)?.icon"/></span>
              <span>{{ category$.byId(categoryNumber)?.title }}</span>
            </button>

            TODO use ProductVisibility (TBD)
            -->
            TODO
          </td>
          <td>{{$n(product.price / 100, 'currency', 'de-DE')}}</td>
          <td>{{product.getVat()?.rate}}%</td>
          <td v-show="hasEditProductRights">
            <button class="tag">
              <router-link :to="{ name: 'ProductSettingsSingle', params: { productId: product.id } }">
                <span class="icon"><icon :icon="['fas', 'pen']"/></span>
              </router-link>
              </button>  
            <button class="tag">
              <span class="icon"><icon :icon="['fas', 'trash']"/></span>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.tag:not(:last-child){
  margin-right:.5em;
}
</style>

<script lang="ts">
import type { Product } from '../../composables/product';
import { useProductStore } from '../../store/ProductStore';
import { useCategoryStore } from '../../store/CategoryStore';
import type { PropType } from 'vue';

export default {
  data(){
    return {
      product$: useProductStore(),
      category$: useCategoryStore(),
      dev: false
    }
  },
  props: {
    products: {
      type: Array as PropType<Product[]>
    }
  },
  methods: {
    copyText(txt: string){
      navigator.clipboard.writeText(txt);
    }
  },
  computed: {
    hasEditProductRights(){
      if(this.dev){
        console.warn("User has elevated privileges to edit products because you are running this in development environment")
        return true;
      }
      //ToDo Check if currently loged in user is allowed to edit products
      return false
    }
  },
  mounted() {
    this.dev = import.meta.env.DEV;
  }
}
</script>

<style scoped>
  .has-copy-btn{
    position:relative;

    .icon {
      position:absolute;
      left:100%;
      z-index:10;
      padding:.8em;
      background-color:rgba(0,0,0,.3);
      border-radius:var(--bulma-radius);
      visibility:hidden;
      pointer-events: none;
      cursor:pointer;
    }
    &:hover .icon {
      visibility: visible;
      pointer-events: all;
    }
  }
</style>