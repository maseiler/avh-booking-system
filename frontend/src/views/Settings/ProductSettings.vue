<template>
    <Buttons :addons="true">
    <Button
      icon-position="left"
      :fa-icon="['fas', 'plus']"
      @click="console.warn('Not yet implemented!')">
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

export default {
  data() {
    return {
      
    }
  },
  components: {
    ProductSelector,
    Buttons,
    Button
  }
}
</script>