<template>
  <h1 class="title" v-if="doneMounting && isEdit">
    Edit Product:
    {{ product.name }}
  </h1>
  <h1 class="title" v-if="doneMounting && !isEdit">
    Neues Produkt erstellen
  </h1>

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

  <div class="columns">
    <div class="column is-3">Size and Unit:</div>
    <div class="column is-1">
      <input type="number" class="input no-controls" v-model="product.size">
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
    <div class="column is-3">
      Price:
    </div>
    <div class="column has-icons-left">
      <p class="control has-icons-left">
        <input type="number" class="input no-controls" v-model="product.price">
        <span class="icon is-small is-left">
            <icon :icon="['fas', 'usd']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">
      Vat:
    </div>
    <div class="column is-3">
      <div class="control has-icons-left">
        <div class="select">
          <select v-model="product.vat">
            <option v-for="vat in vat$.vats" :value="vat.id" class="has-icons-left">
              {{ vat.rate }}%
            </option>
          </select>
          <div class="icon is-small is-left">
            <icon :icon="['fas', 'university']"/>
          </div>
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

  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <Buttons>
        <Button :fa-icon="['fas', 'undo']" icon-position="left" @click="$router.go(-1)">
          Cancel
        </Button>

        <Button class="is-primary" @click="actionButtonClicked" :fa-icon="['fas', 'cloud-upload']" icon-position="right">
          {{ actionButton }}
        </Button>
      </Buttons>      
    </div>
  </div>
</template>

<script lang="ts">
import { Product } from '../../composables/product';
import { useProductStore } from '../../store/ProductStore';
import { useCategoryStore } from '../../store/CategoryStore';
import {useUnitStore} from "../../store/UnitStore.ts";
import Buttons from '../../composables/elements/Buttons.vue';
import Button from '../../composables/elements/Button.vue';
import { useSocketStore } from '../../store/socketStore';
import { useVatStore } from '../../store/VatStore.ts';

export default {
  data() {
    return {
      product$: useProductStore(),
      category$: useCategoryStore(),
      unit$: useUnitStore(),
      socket$: useSocketStore(),
      vat$: useVatStore(),
      product: {} as Product,
      doneMounting: false
    }
  },
  mounted() {
    if(this.isEdit){
      this.product = this.product$.byId(parseInt(this.$route.params.productId.toString()))?.copy();
    } else {
      let newProduct = {} as Product;
      this.product = new Product(newProduct);
    }
    this.doneMounting = true;
    console.log(this.product);
  },
  computed: {
    categoryIcon(){
      return this.category$.byId(this.product.category)?.icon
    },
    actionButton(){
      return this.isEdit ? "Speichern" : "Neu Erstellen";
    },
    isEdit(){
      return this.$route.params.productId?.toString().length > 0;
    },
  },
  methods: {
    // updateProductVisibility(id: number){
    //   if (this.productVisibility.includes(id)){
    //     this.product.visibility = this.product.visibility.filter((visId) => visId != id)
    //     return
    //   }
    //   this.product.visibility.push(id)
    //   this.product.visibility.sort()
    // },
    actionButtonClicked(){
      if(this.isEdit){
        // Update current User
        this.product$.byId(parseInt(this.$route.params.productId.toString()))?.update(this.product);
        this.$router.push({name:'ProductSettings'});
        return;
      }
      this.socket$.addProduct(this.product as Product);
      this.$router.push({name:'ProductSettings'});
    }
  },
  components: {
    Button,
    Buttons
  }
}
</script>

<style scoped>
.no-controls{
  -moz-appearance: textfield;
}
.no-controls::-webkit-outer-spin-button,
.no-controls::-webkit-inner-spin-button{
  -webkit-appearance: none;
  margin:0;
}
</style>