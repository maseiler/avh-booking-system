<script lang="ts" setup>
import { Account } from '../../composables/account';

const  props = defineProps(['accounts', 'allowEdit'])

</script>

<template>
  <div class="selectedAccounts">
    <div class="tag is-primary is-medium" v-for="account in props.accounts">
      {{ account.getFullName() }}
      <button class="delete is-small" v-if="allowEdit" @click="$emit('unselect', account as Account)"></button>
    </div>
  </div>
</template>

<style scoped>
.selectedAccounts{
  margin-bottom:.75rem;
}
.tag{
  margin-bottom:.25em;
  margin-right:.25em;
}
</style>