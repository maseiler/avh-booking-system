<template>
  <Buttons :addons="true">
    <Button
      icon-position="left"
      :fa-icon="['fas', 'plus']"
      @click="$router.push({name: 'AccountSettingsAdd'})">
      Account hinzufügen
    </Button>
  </Buttons>
  <AccountSelector show="list" :all="true"></AccountSelector>
</template>

<script lasng="ts">
import AccountSelector from '../../components/AccountSelector/AccountSelector.vue';
import Button from '../../composables/elements/Button.vue';
import Buttons from '../../composables/elements/Buttons.vue';

export default {
  data() {
    return {

    }
  },
  components: {
    AccountSelector,
    Button,
    Buttons
  }
}
</script>