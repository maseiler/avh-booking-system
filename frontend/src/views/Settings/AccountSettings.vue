<template>
  <h1 class="title">Account Settings</h1>
  <Buttons :addons="true">
    <Button
      icon-position="left"
      :fa-icon="['fas', 'plus']"
      @click="$router.push({ name: 'CategorySettingsAdd', params: { type: 'account' } })">
      Kategorie hinzufügen
    </Button>

    <Button
      icon-position="left"
      :fa-icon="['fas', 'user-plus']"
      @click="$router.push({name: 'AccountSettingsAdd'})">
      Account hinzufügen
    </Button>
  </Buttons>
  <AccountSelector show="list" :all="true"></AccountSelector>
</template>

<script lang="ts">
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