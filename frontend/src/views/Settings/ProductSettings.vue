<template>
  <h1 class="title">Product Settings</h1>
    <Buttons :addons="true">
    <Button
      icon-position="left"
      :fa-icon="['fas', 'plus']"
      @click="$router.push({ name: 'CategorySettingsAdd', params: { type: 'product' } })">
      Kategorie hinzufügen
    </Button>

    <Button
      icon-position="left"
      :fa-icon="['fas', 'cart-plus']"
      @click="$router.push({name: 'ProductSettingsAdd'})">
      Produkt hinzufügen
    </Button>
  </Buttons>
  <ProductSelector show="list"></ProductSelector>
</template>

<script lang="ts">
import ProductSelector from '../../components/ProductSelector/ProductSelector.vue';
import Button from '../../composables/elements/Button.vue';
import Buttons from '../../composables/elements/Buttons.vue';
import { useAccountStore } from '../../store/AccountStore';

export default {
  data() {
    return {
      
    }
  },
  components: {
    ProductSelector,
    Buttons,
    Button
  },
  mounted() {
      useAccountStore().selected = [];
  },
}
</script>